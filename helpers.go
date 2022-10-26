package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func fetchSWISData(url string) *[]byte {
	// push requests use PUT method
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
	}

	req.Header.Set("X-Auth-Token", "676c1618a631cffdsf5554xy545n4oo55q33ppvxcx555sa623a5aeea14e42ecfac7e77da8cfbcf4b69d6a3999828e9b0181ade")
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		log.Print(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
	}

	return &data
}
