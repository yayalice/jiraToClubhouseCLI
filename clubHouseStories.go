package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CHReadStoryList(CHProjectID int64, token string) ([]CHStorySlim, error) {

	client := &http.Client{}

	var urlType = "projects/" + strconv.FormatInt(CHProjectID, 10) + "/stories"
	var chURL = getURL(urlType, token)

	req, err := http.NewRequest("GET", chURL, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	stories := []CHStorySlim{}
	json.Unmarshal(body, &stories)

	return stories, nil

}
