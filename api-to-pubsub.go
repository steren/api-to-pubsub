package main

import (
    "context"
	"fmt"
	"log"
	"net/http"
    "os"
    "errors"
    "io/ioutil"

    "cloud.google.com/go/pubsub"
)

const friendlyPackageName string = "API to Pub/Sub"

func fetchAndPublish(url string, method string, authToken string, projectId string, topic string) error {
    response, err := fetch(url, method, authToken)
    if err != nil {
        return err
    }
    return publish(projectId, topic, response)
}

func fetch(url string, method string, authToken string) (string, error) {
    if(url == "") {
        return "", errors.New("No URL provided")
    }
    if(method == "") {
        method = "GET"
    }
    client := &http.Client{}
    req, _ := http.NewRequest(method, url, nil)
    if(authToken != "") {
        req.Header.Set("auth-token", authToken)
    }
    res, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
        return "", err
    }

    defer res.Body.Close()
    responseBody, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
        return "", err
    }

    return string(responseBody), nil
}

func getProjectIdFromMetadataServer() string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/project/project-id", nil)
	req.Header.Set("Metadata-Flavor", "Google")
	res, err := client.Do(req)
	if err == nil {
		defer res.Body.Close()
		responseBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		return string(responseBody)
    } else {
        return ""
    }
}

func publish(projectID, topicID, msg string) error {
    if(projectID == "") {
        return errors.New("project ID cannot be retrieved from instance metadata server, and no PROJECT_ID en var provided.")
    }

    if(topicID == "") {
        return errors.New("no Pub/Sub topic provided")
    }

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}
	log.Printf("Published a message; msg ID: %v\n", id)
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
    projectID := os.Getenv("PROJECT_ID")
    if(projectID == "") {
        projectID = getProjectIdFromMetadataServer()    
    }
    fetchAndPublish(os.Getenv("URL"), os.Getenv("METHOD"), os.Getenv("AUTH_TOKEN"), projectID, os.Getenv("TOPIC"))
    fmt.Fprintf(w, "Done")
}

func main() {
    log.Printf("%s started.", friendlyPackageName)
    
    url := os.Getenv("URL")
	if url == "" {
        log.Fatal("URL environment variable is not set")
        return
	}

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}