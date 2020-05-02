# api-to-pubsub

Forward the response of an API call to a Pub/Sub topic

## Configure

- `URL` (Required): the URL to fetch
- `METHOD` (Optional, default to "GET"): the HTTP method to use
- `AUTH_TOKEN` (Optional, defaults to ""): Any `auth-token` header to add to the request
- `TOPIC` (Required): The ID of the PubSub topc to publish to (e.g. `my-topic`)
- `PROJECT_ID` (Optional, defaults to current project): The GCP project ID containing the Pub/Sub topic.