package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	baseurl         = "https://hotspot.ucn.dk"
	uri             = "/auth/action/ads"
	credentialsFile = "C:/Users/Nikola Velichkov/go/src/github.com/Federlizer/logoin/credentials.csv"
)

var values = make(map[string]string)

func main() {
	if connected() {
		fmt.Println("You're already connected. Exiting.")
		return
	}

	file, err := os.Open(credentialsFile)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		values[record[0]] = record[1]
	}

	err = login(baseurl+uri, values)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection successful")
}

func login(uri string, formValues map[string]string) error {
	form := url.Values{}
	for k, v := range formValues {
		form.Add(k, v)
	}

	response, err := http.PostForm(uri, form)
	if err != nil {
		return err
	}

	if response.StatusCode == http.StatusOK {
		return nil
	} else {
		return errors.New("Connection failed with status " + response.Status)
	}
}

func connected() bool {
	resp, err := http.Get("https://google.com")
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	return true
}
