package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// CHReadStoryList is a db-operation in CH to get a list of all the stories for a given project
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

// CHUpdateStory is a db-operation in CH to update which files should be associated with a given story
func CHUpdateStory(clubHouseStorySlim CHStorySlim, fileIDs []int64, token string) (CHGETStory, error) {

	client := &http.Client{}

	// we must append only new files to the existing list of files
	var mergedFileIDs []int64
	mergedFileIDs = clubHouseStorySlim.FileIDs
	for _, fileID := range fileIDs {
		if !intInSlice(fileID, mergedFileIDs) {
			mergedFileIDs = append(mergedFileIDs, fileID)
		}
	}

	var urlType = "stories/" + strconv.FormatInt(clubHouseStorySlim.ID, 10)
	var chURL = getURL(urlType, token)
	requestBody := CHStoryForUpdate{FileIDs: fileIDs}
	var jsonString, err = json.Marshal(requestBody)
	// requestBodyString := string(jsonString)
	// fmt.Printf("Request body: %v", requestBodyString)
	if err != nil {
		return CHGETStory{}, err
	}
	b := bytes.NewBuffer(jsonString)
	req, err := http.NewRequest("PUT", chURL, b)
	if err != nil {
		return CHGETStory{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return CHGETStory{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Println("response Status:", res.Status)
		fmt.Println("response Headers:", res.Header)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return CHGETStory{}, err
	}

	var updatedClubHouseStory CHGETStory
	json.Unmarshal(body, &updatedClubHouseStory)

	return updatedClubHouseStory, nil

}

func splitToString(a []int64, sep string) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(b, sep)
}

func intInSlice(a int64, list []int64) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
