package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// FetchJiraAttachment downloads a specific attachment from JIRA
func FetchJiraAttachment(id string, name string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", getJiraURL(id, name), nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/octet-stream")
	req.SetBasicAuth("mama01", "BNxuRqLemq")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	//err = ioutil.WriteFile(name, body, 0644)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}

	return body, nil
}

func getJiraURL(id string, name string) string {

	return "https://edpconsult.atlassian.net/secure/attachment/" + id + "/" + url.PathEscape(name)

}
