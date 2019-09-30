package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//Define the configuration we should find in the configuration file.
type Configuration struct {
	AccessToken string
	DomainId    int
	RecordId    int
}

func main() {

	filename := "addns.json"
	fileLocation := ""
	//We should first look for the config file. Check to local directory, then check the users profile
	if _, err := os.Stat("./" + filename); err == nil {
		fileLocation = "./" + filename
	} else if _, err := os.Stat(os.Getenv("HOME") + "/" + filename); err == nil {
		fmt.Println("Using the Config from your Home directory")
		fileLocation = os.Getenv("HOME") + "/" + filename
	} else {
		log.Fatal("We could not find the config.json file. Please read the instructions.")
	}

	var configuration Configuration
	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	// Now update the linode record
	linodeUrl := fmt.Sprintf("https://api.linode.com/v4/domains/%v/records/%v", configuration.DomainId, configuration.RecordId)
	externalIP := GetOutboundIP()
	var jsonStr = []byte(fmt.Sprintf(`{"target":"%s"}`, externalIP))

	req, err := http.NewRequest(http.MethodPut, linodeUrl, bytes.NewBuffer(jsonStr))
	if err != nil {log.Fatal(err)}

	// Set a nice user-agent as we are nice people
	req.Header.Set("User-Agent", "ADDNS-updater")
	req.Header.Set("Authorization", "Bearer "+configuration.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	// Create a client and only allow 3 seconds for the request to complete
	addnsClient := http.Client{
		Timeout: time.Second * 3, // Maximum of 2 secs
	}

	//Send the query.
	res, getErr := addnsClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fmt.Printf("Query Result: %s\n", body)

}

// Get the external IP address from the site myexternalip.com
func GetOutboundIP() string {

	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {log.Fatal(err)}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}
	return string(body)

}