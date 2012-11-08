package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type response struct {
	Results []struct {
		Created_at string
		Id_str     string
		From_user  string
		Text       string
		New        bool
		URL        string
	}
}

var responses [5]int
var updated_at time.Time
var started bool
var twitter_url = "http://search.twitter.com/search.json?q=OpenVE&rpp=5&result_type=recent"

func readTwitter() (response, error) {

	// Twitter Fetch
	tw := response{}
	resp, err := http.Get(twitter_url)
	if err != nil {
		println("Error retrieving from twitter: ", err)
		return tw, err
	}

	// Parsing the result body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	err = json.Unmarshal(body, &tw)
	if err != nil {
		println("Error parsing JSON: ", err)
		return tw, err
	}

	// Picking the new values
	if started {
		var created_at time.Time
		for i, v := range tw.Results {
			created_at, _ = time.Parse(time.RFC1123Z, v.Created_at)
			if updated_at.Sub(created_at) < 0 {
				tw.Results[i].New = true
			}
			tw.Results[i].URL = "https://twitter.com/" + v.From_user + "/status/" + v.Id_str
		}
	}

	// Saving the latest date
	updated_at, err = time.Parse(time.RFC1123Z, tw.Results[0].Created_at)
	if err != nil {
		println("Error parsing Time: ", err)
		return tw, err
	}

	if !started {
		started = true
	}

	return tw, err
}
