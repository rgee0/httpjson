// Based on original work by Alex Ellis at http://blog.alexellis.io/golang-json-api-client/

// I cant claim this to be optimal in any way.  I Took the opportunity presented by Alex to experiment with nested JSON elements.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type jsonRes struct {
	Number  int    `json:"number"`
	Message string `json:"message"`
	People  []struct {
		Craft  string `json:"craft"`
		Person string `json:"name"`
	} `json:"people"`
}

func main() {

	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs - had to change this as it was timing out (a sign of how many ppl are trying the tutorial?)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	result := jsonRes{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// We ought to check whether there is any work to do, i.e. is anyone listed as being in space?
	if result.Number > 0 {

		//if there are, set some default words for the output
		isare := "are"
		ifplural := "people"

		//need to override these if there is only one person in space
		if result.Number == 1 {
			isare = "is"
			ifplural = "person"
		}
		//Print the into line
		fmt.Printf("Currently there %s %d %s in space. They are:\n", isare, result.Number, ifplural)

		//For each person in space print their name and the craft upon which they reside
		for i := range result.People {
			fmt.Println(result.People[i].Person + ", who is stationed on " + result.People[i].Craft)
		}

	} else {
		//Number is <= 0 so it looks as though nobody listed as being in space
		fmt.Println("Currently there are no people listed as being in space.")
	}
}
