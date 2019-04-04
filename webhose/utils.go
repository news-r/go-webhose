package webhose

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const baseurl = "https://webhose.io"

func (client *NewsClient) constructURL() string {
	parameters := getParameters(client)
	url := baseurl + "/filterWebContent?" + parameters
	return (url)
}

func buildPageURL(next string) string {
	url := baseurl + next
	return (url)
}

func callAPI(url string) News {
	data := new(News)

	// call API
	cl := &http.Client{Timeout: 10 * time.Second}
	resp, err := cl.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close() // close response

	// Unmarshall
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Panic(err)
	}

	return (*data)
}
