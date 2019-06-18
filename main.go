package main

import (
	"log"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go"
)

func webhookHandler(writer http.ResponseWriter, request *http.Request) {
	event := cloudevents.NewEvent()
	event.SetID("ABC-123")
	event.SetType("io.jenkins.webhook.received")
	event.SetSource("http://localhost:8080/")
	event.SetData("{\"myData\": \"this\"}")
	event.SetDataContentType("application/json")
	log.Println("CloudEvent created")
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
