package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
)

// UploadAttachmentToCH will upload a file to CH and update it with JIRA-id as external_id
func UploadAttachmentToCH(externalID string, token string, file io.Reader, fileName string) (ClubHouseFile, error) {

	var updatedAttachment ClubHouseFile

	client := &http.Client{}
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("file", fileName)
	if err != nil {
		return updatedAttachment, err
	}
	if _, err = io.Copy(fw, file); err != nil {
		return updatedAttachment, err
	}

	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", GetURL("files", token), &b)
	if err != nil {
		return updatedAttachment, err
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	res, err := client.Do(req)
	if err != nil {
		return updatedAttachment, err
	}
	defer res.Body.Close()
	// Check the response
	if res.StatusCode > 299 {
		fmt.Println("response Status:", res.Status)
		fmt.Println("response Headers:", res.Header)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return updatedAttachment, err
	}

	newAttachments := []ClubHouseFile{}
	json.Unmarshal(body, &newAttachments)

	clubHouseID := newAttachments[0].ID
	updatedAttachment, err = updateAttachment(externalID, clubHouseID, token)
	if err != nil {
		return updatedAttachment, err
	}

	return updatedAttachment, nil
}

func updateAttachment(JiraID string, ClubHouseID int64, token string) (ClubHouseFile, error) {

	var updatedAttachment ClubHouseFile

	client := &http.Client{}

	var urlType = "files/" + strconv.FormatInt(ClubHouseID, 10)
	var chURL = GetURL(urlType, token)
	var jsonString = []byte(`{"external_id": "` + JiraID + `"}`)
	b := bytes.NewBuffer(jsonString)
	req, err := http.NewRequest("PUT", chURL, b)
	if err != nil {
		return updatedAttachment, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return updatedAttachment, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Println("response Status:", res.Status)
		fmt.Println("response Headers:", res.Header)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return updatedAttachment, err
	}

	json.Unmarshal(body, &updatedAttachment)

	return updatedAttachment, nil
}
