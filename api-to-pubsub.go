package main

import (
	"fmt"
	"log"
	"net/http"
    "os"
    "errors"
    "io/ioutil"
)

const friendlyPackageName string = "API to Pub/Sub"

func fetchAndForward(url string, method string, authToken string) error {
    if(url == "") {
        return errors.New("No URL provided")
    }
    if(method == "") {
        method = "GET"
    }
    response := ""
    client := &http.Client{}
    req, _ := http.NewRequest(method, url, nil)
    if(authToken != "") {
        req.Header.Set("auth-token", authToken)
    }
    res, err := client.Do(req)
    if err == nil {
		defer res.Body.Close()
		responseBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		response = string(responseBody)
	}

    log.Print(response)
    return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
    fetchAndForward("https://api.co2signal.com/v1/latest?lon=-122.403171&lat=37.758218", "", "")
    fmt.Fprintf(w, "Done")
}

func main() {
	log.Printf("%s started.", friendlyPackageName)

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}