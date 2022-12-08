package webhook_queue

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"webhook-queue/src/helper/log"
	"webhook-queue/src/types"
)

func StartQueue(port int, webhook string, alias string) {
	var queue []interface{}

	//handle the requests
	http.HandleFunc("/"+alias, func(w http.ResponseWriter, r *http.Request) {
		//check if the request method is post
		if r.Method == "POST" {
			//read the body of the request
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Error("could not read body: "+err.Error(), alias)
			}

			//add the webhook to the queue
			queue = append(queue, body)

			//send a response
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Webhook added to queue"))
		} else {
			//send a response
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Webhook queue is running"))
		}
	})

	//show the user where the server is running
	log.Info("Webhook queue is running here: http://localhost:"+strconv.Itoa(port)+"/"+alias, alias)

	//main loop
	for {
		//only execute if there are webhooks in the queue
		for index := range queue {
			//always only send the first webhook in the queue
			if index == 0 {
				//send the webhook
				SendWebhook(queue[index].([]byte), webhook, alias)
				//remove the webhook from the queue
				queue = queue[1:]

				log.Info("Webhook in queue:"+strconv.Itoa(len(queue)), alias)
			}
		}
	}
}

func SendWebhook(webhook []byte, url string, alias string) {
	//make a new http client
	client := &http.Client{}

	//make a new http request
	req, err := http.NewRequest("POST", url, bytes.NewReader(webhook))
	if err != nil {
		log.Error("Error creating request"+err.Error(), alias)
	}
	req.Header.Set("Content-Type", "application/json")

	//send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Error sending request"+err.Error(), alias)
	}
	defer resp.Body.Close()

	//read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error reading response"+err.Error(), alias)
	}

	log.Info("Send webhook, got status code:"+strconv.Itoa(resp.StatusCode), alias)

	//check if the status code is 429/ratelimited
	if resp.StatusCode == 429 {
		var ratelimitResponse types.RatelimitResponse

		//parse the response to json
		err = json.Unmarshal(body, &ratelimitResponse)
		if err != nil {
			log.Error("Error parsing response"+err.Error(), alias)
		}

		log.Warning("Got ratelimited, waiting "+strconv.Itoa(int(ratelimitResponse.Retry_after))+" seconds", alias)

		//wait the amount of seconds the api told us to wait
		time.Sleep(time.Duration(int(ratelimitResponse.Retry_after*1000)) * time.Millisecond)

		//send the webhook again
		SendWebhook(webhook, url, alias)
	}
}
