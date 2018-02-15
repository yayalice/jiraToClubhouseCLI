package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
)

// GetUserMap parses the userMap.json file and returns the content
func GetUserMap(mapFile string, token string) ([]UserMap, error) {
	jsonFile, err := os.Open(mapFile)
	if err != nil {
		return []UserMap{}, err
	}

	defer jsonFile.Close()
	JSONData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return []UserMap{}, err
	}

	// userMaps := []userMap
	var userMaps []UserMap
	err = json.Unmarshal(JSONData, &userMaps)
	if err != nil {
		return []UserMap{}, err
	}

	// We fetch the user ID from Clubhouse as it regularly changes
	userMaps, err = CHReadMembers(userMaps, token)
	if err != nil {
		return []UserMap{}, err
	}

	return userMaps, nil
}

// GetProjectMap parses the projectMap.json file and returns the content
func GetProjectMap(projectMapFile string) ([]ProjectMap, error) {
	jsonFile, err := os.Open(projectMapFile)
	if err != nil {
		return []ProjectMap{}, err
	}

	defer jsonFile.Close()
	JSONData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return []ProjectMap{}, err
	}

	// userMaps := []userMap
	var projectMaps []ProjectMap
	err = json.Unmarshal(JSONData, &projectMaps)
	if err != nil {
		return []ProjectMap{}, err
	}

	return projectMaps, nil
}

// ExportToJSON will import the XML and then export the data to the file specified.
func ExportToJSON(jiraFile string, userMaps []UserMap, projectMaps []ProjectMap, token string, exportFile string) error {
	export, err := GetDataFromXMLFile(jiraFile)
	if err != nil {
		return err
	}
	data, err := json.Marshal(export.GetDataForClubhouse(userMaps, projectMaps, token))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(exportFile, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// GetDataFromXMLFile will Unmarshal the XML file into the objects used by the application.
func GetDataFromXMLFile(jiraFile string) (JiraExport, error) {
	xmlFile, err := os.Open(jiraFile)
	if err != nil {
		return JiraExport{}, err
	}

	defer xmlFile.Close()
	XMLData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return JiraExport{}, err
	}

	jiraExport := JiraExport{}
	err = xml.Unmarshal(XMLData, &jiraExport)
	if err != nil {
		return JiraExport{}, err
	}

	return jiraExport, nil
}
