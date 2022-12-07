package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"webhook-queue/src/types"
	"webhook-queue/src/webhook_queue"

	"github.com/kardianos/osext"
)

func main() {
	//reading the executable path
	executablePath, _ := osext.ExecutableFolder()

	//setting the os path to the executable path
	os.Chdir(executablePath)

	// read the settngs.json file
	settings := readSettings()

	var queue []interface{}

	//handle the requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//check if the request method is post
		if r.Method == "POST" {
			//read the body of the request
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
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

	//start the server
	go http.ListenAndServe(":"+strconv.Itoa(settings.Port), nil)

	//show the user where the server is running
	fmt.Println("Webhook queue is running here: http://localhost:" + strconv.Itoa(settings.Port))

	//main loop
	for {
		//only execute if there are webhooks in the queue
		for index := range queue {
			//always only send the first webhook in the queue
			if index == 0 {
				//send the webhook
				webhook_queue.SendWebhook(queue[index].([]byte), settings.Webhook)
				//remove the webhook from the queue
				queue = queue[1:]

				fmt.Println("Webhooks in queue:", len(queue))
			}
		}
	}
}

func readSettings() types.Settings {
	settings, err := ioutil.ReadFile("settings.json")
	if err != nil {
		fmt.Println(err)
	}

	//load settings.json into the Settings struct
	var settingsJSON types.Settings

	//parse the settings.json file
	err = json.Unmarshal(settings, &settingsJSON)
	if err != nil {
		fmt.Println(err)
	}

	return settingsJSON
}
