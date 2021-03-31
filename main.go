package main

import (
	"context"
	"encoding/json"
	dapr "github.com/dapr/go-sdk/client"
	"io/ioutil"
	"net/http"
)

func main() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	http.HandleFunc("/publish", func(writer http.ResponseWriter, request *http.Request) {
		err = client.PublishEvent(context.Background(), "rabbitmq-pubsub", "test-topic", []byte("hello dapr"))
		respond(nil, err, writer)
	})

	http.HandleFunc("/invoke-service", func(writer http.ResponseWriter, request *http.Request) {
		out, err := client.InvokeMethodWithContent(context.Background(), "httpbin", "/headers", "GET", &dapr.DataContent{})
		respond(out, err, writer)
	})

	http.HandleFunc("/get-state", func(writer http.ResponseWriter, request *http.Request) {
		key := request.URL.Query().Get("key")
		stateItem, err := client.GetState(context.Background(), "my-redis", key)
		respond(stateItem.Value, err, writer)
	})

	http.HandleFunc("/save-state", func(writer http.ResponseWriter, request *http.Request) {
		req, _ := ioutil.ReadAll(request.Body)
		data := make(map[string]string)
		json.Unmarshal(req, &data)

		for k, v := range data {
			err := client.SaveState(context.Background(), "my-redis", k, []byte(v))
			respond(nil, err, writer)
		}
	})

	http.HandleFunc("/save-bulk-state", func(writer http.ResponseWriter, request *http.Request) {
		req, _ := ioutil.ReadAll(request.Body)
		data := make(map[string]string)
		json.Unmarshal(req, &data)

		var items []*dapr.SetStateItem
		for k, v := range data {
			item := &dapr.SetStateItem{
				Key: k,
				Value: []byte(v),
			}
			items = append(items, item)
		}

		err := client.SaveBulkState(context.Background(), "my-redis", items...)
		respond(nil, err, writer)
	})

	http.HandleFunc("/delete-state", func(writer http.ResponseWriter, request *http.Request) {
		err := client.DeleteState(context.Background(), "my-redis","my-key")
		respond(nil, err, writer)
	})

	http.ListenAndServe(":8080", nil)
}

func respond(data []byte, err error, writer http.ResponseWriter) {
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if data != nil {
		writer.Write(data)
	}
}
