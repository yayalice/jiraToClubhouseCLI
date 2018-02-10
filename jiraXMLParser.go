package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
)

func GetUserMap(mapFile string) ([]userMap, error) {
	jsonFile, err := os.Open(mapFile)
	if err != nil {
		return []userMap{}, err
	}

	defer jsonFile.Close()
	JSONData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return []userMap{}, err
	}

	// userMaps := []userMap
	var userMaps []userMap
	err = json.Unmarshal(JSONData, &userMaps)
	if err != nil {
		return []userMap{}, err
	}

	return userMaps, nil
}

func GetProjectMap(projectMapFile string) ([]projectMap, error) {
	jsonFile, err := os.Open(projectMapFile)
	if err != nil {
		return []projectMap{}, err
	}

	defer jsonFile.Close()
	JSONData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return []projectMap{}, err
	}

	// userMaps := []userMap
	var projectMaps []projectMap
	err = json.Unmarshal(JSONData, &projectMaps)
	if err != nil {
		return []projectMap{}, err
	}

	return projectMaps, nil
}

// ExportToJSON will import the XML and then export the data to the file specified.
func ExportToJSON(jiraFile string, userMaps []userMap, projectMaps []projectMap, exportFile string) error {
	export, err := GetDataFromXMLFile(jiraFile)
	if err != nil {
		return err
	}
	data, err := json.Marshal(export.GetDataForClubhouse(userMaps, projectMaps))
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
