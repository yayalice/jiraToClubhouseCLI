package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// UploadAttachmentToCH will upload a file to CH and update it with JIRA-id as external_id
func UploadAttachmentToCH(externalID string, file []byte, token string) (int64, error) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", GetURL("files", token), bytes.NewBuffer(file))
	if err != nil {
		return 0, err
	}
	//req.Header.Set("Content-Type", "application/octet-stream")
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	newAttachment := ClubHouseFile{}
	json.Unmarshal(body, &newAttachment)

	return newAttachment.ID, nil

}
