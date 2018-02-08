package clubHouseFiles

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"jiraToClubhouseCLI/internal/clubHouse"
	"mime/multipart"
	"net/http"
	"strconv"
)

func Create(file []byte, fileName string, externalID string, token string) (clubHouse.File, error) {

	client := &http.Client{}

	fr := bytes.NewReader(file)
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("file", fileName)
	if err != nil {
		return clubHouse.File{}, err
	}
	if _, err = io.Copy(fw, fr); err != nil {
		return clubHouse.File{}, err
	}

	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", getURL("files", token), &b)
	if err != nil {
		return clubHouse.File{}, err
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	res, err := client.Do(req)
	if err != nil {
		return clubHouse.File{}, err
	}
	defer res.Body.Close()

	// Check the response
	if res.StatusCode > 299 {
		fmt.Println("response Status:", res.Status)
		fmt.Println("response Headers:", res.Header)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return clubHouse.File{}, err
	}

	newAttachments := []clubHouse.File{}
	json.Unmarshal(body, &newAttachments)

	clubHouseID := newAttachments[0].ID
	clubHouseFile := clubHouse.File{ExternalID: externalID, ID: clubHouseID}

	// There is no need to update the file with a new external_id if the value is empty
	if externalID != "" {
		clubHouseFile, err = Update(clubHouseFile, token)
		if err != nil {
			return clubHouseFile, err
		}
	}

	return clubHouseFile, nil

}

func Read(clubHouseFileID int64, token string) (clubHouse.File, error) {

	client := &http.Client{}

	var urlType = "files/" + strconv.FormatInt(clubHouseFileID, 10)
	var chURL = getURL(urlType, token)
	req, err := http.NewRequest("GET", chURL, nil)
	if err != nil {
		return clubHouse.File{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return clubHouse.File{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Println("response Status:", res.Status)
		fmt.Println("response Headers:", res.Header)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return clubHouse.File{}, err
	}

	var clubHouseFile clubHouse.File
	json.Unmarshal(body, &clubHouseFile)

	return clubHouseFile, nil

}

func ReadList(token string) ([]clubHouse.File, error) {

	// CHAttachments := make(map[string]int)
	client := &http.Client{}

	req, err := http.NewRequest("GET", getURL("files", token), nil)

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
	files := []clubHouse.File{}
	json.Unmarshal(body, &files)

	return files, nil
}

func Update(clubHouseFile clubHouse.File, token string) (clubHouse.File, error) {

	client := &http.Client{}

	var urlType = "files/" + strconv.FormatInt(clubHouseFile.ID, 10)
	var chURL = getURL(urlType, token)
	var jsonString = []byte(`{"external_id": "` + clubHouseFile.ExternalID + `"}`)
	b := bytes.NewBuffer(jsonString)
	req, err := http.NewRequest("PUT", chURL, b)
	if err != nil {
		return clubHouse.File{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return clubHouse.File{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Println("response Status:", res.Status)
		fmt.Println("response Headers:", res.Header)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return clubHouse.File{}, err
	}

	var updatedClubHouseFile clubHouse.File
	json.Unmarshal(body, &updatedClubHouseFile)

	return updatedClubHouseFile, nil

}

func Delete(clubHouseFileID int64, token string) error {

	client := &http.Client{}

	var urlType = "files/" + strconv.FormatInt(clubHouseFileID, 10)
	var chURL = getURL(urlType, token)

	req, err := http.NewRequest("DELETE", chURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Println("response Status:", res.Status)
		fmt.Println("response Headers:", res.Header)
	}

	return nil

}

// GetURL will get the use the REST API v1 address, the resource provided and the API token to get the URL for transactions
func getURL(kind string, token string) string {
	return fmt.Sprintf("%s%s?token=%s", "https://api.clubhouse.io/api/v2/", kind, token)
}
