package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

type users struct {
	users map[string]user `json:"users"`
}

type user struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Active   bool   `json:"active"`
}

func fetchRemoteStream(url string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Fatal("not a HTTP/200 status code on socket list fetch --- got " + strconv.Itoa(resp.StatusCode))
	}

	body := resp.Body

	return body, nil
}
