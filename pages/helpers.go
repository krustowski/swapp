package pages

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func fireSWIS(url string, method string) *[]byte {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Print(err)
	}

	req.Header.Set("X-Auth-Token", os.Getenv("SWAPI_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		log.Print(err)
		return nil
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
	}

	return &data
}

func putSWISData(url string) bool {
	data := fireSWIS(url, "PUT")
	if data == nil {
		return false
	}

	return true
}

func fetchSWISData(url string) *[]byte {
	data := fireSWIS(url, "GET")
	if data == nil {
		return nil
	}

	return data
}
