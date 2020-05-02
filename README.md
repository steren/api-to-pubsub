# api-to-pubsub

Web service that fetches a given URL and forwards the response to a Google Cloud Pub/Sub topic.

## Prerequisites

* A Google Cloud Platform project
* A Cloud Pub/Sub topic in this project (create it with `gcloud pubsub topics create my-topic`)

## Configure

- `URL` (Required): the URL to fetch
- `METHOD` (Optional, default to "GET"): the HTTP method to use
- `AUTH_TOKEN` (Optional, defaults to ""): Any `auth-token` header to add to the request
- `TOPIC` (Required): The ID of the PubSub topc to publish to (e.g. `my-topic`)
- `PROJECT_ID` (Optional, defaults to current project): The GCP project ID containing the Pub/Sub topic.

## Deploy

This service can be deployed anywhere Go or containers can run.

Notably, it has been designed to be deployed as a private [Google Cloud Run](https://cloud.run) service, to be [invoked periodically by Google Cloud Scheduler](https://cloud.google.com/run/docs/triggering/using-scheduler) 

[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)