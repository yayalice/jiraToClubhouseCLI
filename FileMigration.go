package main

import "fmt"

// MigrateFiles will import the XML, upload files to Clubhouse and link them to the appropriate story
func MigrateFiles(jiraFile string, userMaps []userMap, projectMaps []projectMap, token string, testMode bool) error {

	export, err := GetDataFromXMLFile(jiraFile)
	if err != nil {
		return err
	}

	jiraFileList := export.GetFileListFromXMLFile(userMaps)
	existingCHFiles, err := CHReadFileList(token)
	jiraFileList.RemoveDoubles(existingCHFiles)

	if !testMode {
		fmt.Println("Migrating files to Clubhouse...")
		err = jiraFileList.Migrate(projectMaps, token)
		if err != nil {
			return err
		}
	} else {
		for _, group := range jiraFileList.AttachmentGroups {
			for _, CHFile := range group.CHFiles {
				fmt.Printf("Found file File with Jira ID %v and name %v to be uploaded to CH\n", CHFile.ExternalID, CHFile.Name)
			}
		}
	}

	return nil
}
