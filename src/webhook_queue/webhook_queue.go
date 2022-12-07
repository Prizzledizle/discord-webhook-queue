package webhook_queue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"webhook-queue/src/types"
)

func SendWebhook(webhook []byte, url string) {
	//make a new http client
	client := &http.Client{}

	//make a new http request
	req, err := http.NewRequest("POST", url, bytes.NewReader(webhook))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")

	//send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("send webhook, got status code", resp.StatusCode)

	//check if the status code is 429/ratelimited
	if resp.StatusCode == 429 {
		var ratelimitResponse types.RatelimitResponse

		//parse the response to json
		err = json.Unmarshal(body, &ratelimitResponse)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("got ratelimited, waiting", ratelimitResponse.Retry_after, "seconds")

		//wait the amount of seconds the api told us to wait
		time.Sleep(time.Duration(int(ratelimitResponse.Retry_after*1000)) * time.Millisecond)

		//send the webhook again
		SendWebhook(webhook, url)
	}
}
