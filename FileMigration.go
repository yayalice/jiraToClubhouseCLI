package main

import "fmt"

// UploadToClubhouse will import the XML, and upload it to Clubhouse
func MigrateFiles(jiraFile string, userMaps []userMap, token string, testMode bool) error {

	export, err := GetDataFromXMLFile(jiraFile)
	if err != nil {
		return err
	}

	jiraFileList := export.GetFileListFromXMLFile(userMaps)
	existingCHFiles, err := CHReadFileList(token)
	jiraFileList.RemoveDoubles(existingCHFiles)

	if !testMode {
		fmt.Println("Migrating files to Clubhouse...")
		err = jiraFileList.Migrate(token)
		if err != nil {
			return err
		}
	} else {
		for _, CHFile := range jiraFileList.CHFiles {
			fmt.Printf("Found file with Jira ID %v and name %v to be uploaded to CH\n", CHFile.ExternalID, CHFile.Name)
		}
	}

	return nil
}
