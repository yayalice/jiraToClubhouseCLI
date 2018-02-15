package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CHReadMembers(userMaps []UserMap, token string) ([]UserMap, error) {

	client := &http.Client{}

	var chURL = getURL("members", token)
	req, err := http.NewRequest("GET", chURL, nil)
	if err != nil {
		return userMaps, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return userMaps, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Println("response Status:", res.Status)
		fmt.Println("response Headers:", res.Header)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return userMaps, err
	}

	var clubHouseMembers []CHMember
	json.Unmarshal(body, &clubHouseMembers)

	x := make(map[string]string)
	for _, member := range clubHouseMembers {
		x[member.Profile.Name] = member.ID
	}

	for i, userMap := range userMaps {
		userMaps[i].CHID = x[userMap.CHName]
		if userMaps[i].CHID == "" {
			return userMaps, fmt.Errorf("Could not find a user with name %v registered in Clubhouse", userMap.CHName)
		}
	}

	return userMaps, nil

}
