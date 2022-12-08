package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"webhook-queue/src/helper/cli"
	"webhook-queue/src/helper/log"
	"webhook-queue/src/types"
	"webhook-queue/src/webhook_queue"

	"github.com/kardianos/osext"
)

func main() {
	//reading the executable path
	executablePath, _ := osext.ExecutableFolder()

	//setting the os path to the executable path
	os.Chdir(executablePath)

	cli.ClearCLI()
	cli.PrintHeader()
	cli.RenameCLI([]types.QueueObject{})

	// read the settngs.json file
	settings := readSettings()

	for i := range settings.Webhooks {
		//start the webhook queue
		go webhook_queue.StartQueue(settings.Port, settings.Webhooks[i].Webhook, settings.Webhooks[i].Alias, i)
	}

	go http.ListenAndServe(":"+strconv.Itoa(settings.Port), nil)

	for {
		//main loop
	}
}

func readSettings() types.Settings {
	settings, err := ioutil.ReadFile("settings.json")
	if err != nil {
		log.Error(err.Error(), "")
	}

	//load settings.json into the Settings struct
	var settingsJSON types.Settings

	//parse the settings.json file
	err = json.Unmarshal(settings, &settingsJSON)
	if err != nil {
		log.Error(err.Error(), "")
	}

	return settingsJSON
}
