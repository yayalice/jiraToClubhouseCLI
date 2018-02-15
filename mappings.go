package main

import "fmt"

// MapUser tries to map a given jira user with a clubhouse user
func MapUser(userMaps []UserMap, jiraUsername string) (CHID string, err error) {

	for _, u := range userMaps {
		if u.JiraUsername == jiraUsername {
			CHID = u.CHID
			return
		}
		if u.Default == true {
			CHID = u.CHID
		}
	}
	if CHID != "" {
		fmt.Printf("Unknown user %s. Will use default user: %s\n\n", jiraUsername, CHID)
	} else {
		err = fmt.Errorf("unknown user %s. No default user defined in userMap. Please define a default user and retry", jiraUsername)
		fmt.Println(err.Error())
	}

	return
}

// MapProject tries to map a given jira project with a clubhouse project
func MapProject(projectMaps []ProjectMap, jiraProjectKey string) (CHProjectID int64, err error) {
	//projectID := GetProjectInfo(projectMaps, jiraProjectKey)

	for _, u := range projectMaps {
		if u.JiraProjectKey == jiraProjectKey {
			CHProjectID = u.CHProjectID
		}
	}

	if CHProjectID == 0 {
		err = fmt.Errorf("[MapProject] JIRA project not found: %v", jiraProjectKey)
		fmt.Println(err.Error())
		CHProjectID = 299
	}

	return CHProjectID, err
}

// MapStory tries to map a given jira story with a clubhouse story
func MapStory(projectMaps []ProjectMap, jiraProjectKey string, jiraStoryKey string, token string) (CHStorySlim, error) {

	// get CHProjectID in some way
	clubHouseProjectID, err := MapProject(projectMaps, jiraProjectKey)
	if err != nil {
		return CHStorySlim{}, err
	}

	// fetch existing stories
	clubHouseStoryList, err := CHReadStoryList(clubHouseProjectID, token)
	if err != nil {
		return CHStorySlim{}, err
	}
	// loop through stories

	for _, clubHouseStorySlim := range clubHouseStoryList {
		if clubHouseStorySlim.ExternalID == jiraStoryKey {
			return clubHouseStorySlim, nil
		}
	}

	return CHStorySlim{}, fmt.Errorf("Could not find corresponding story for Jira story with key %v in Clubhouse for the project %v", jiraStoryKey, jiraProjectKey)
	// get ID from map

}

// GenerateMapForExistingCHFiles generates a mapping as map[jiraFileKey]clubhouseFileID
func GenerateMapForExistingCHFiles(existingCHFiles []CHGETFile) map[string]int64 {

	x := make(map[string]int64)
	for _, clubHouseFile := range existingCHFiles {
		x[clubHouseFile.ExternalID] = clubHouseFile.ID
	}
	return x

}

// GenerateMapForExistingCHStories generates a mapping as map[jiraStoryKey]clubhouseStoryID
func GenerateMapForExistingCHStories(existingCHStories []CHStorySlim) map[string]int64 {

	x := make(map[string]int64)
	for _, clubHouseStorySlim := range existingCHStories {
		x[clubHouseStorySlim.ExternalID] = clubHouseStorySlim.ID
	}
	return x

}

// GenerateMapForAttachmentMigrationList generates a mapping as map[jiraStoryKey]clubhouseFiles
func GenerateMapForAttachmentMigrationList(attachments []AttachmentGroup) map[string][]CHFile {
	x := make(map[string][]CHFile)
	for _, attachmentGroup := range attachments {
		x[attachmentGroup.JiraStoryKey] = attachmentGroup.CHFiles
	}
	return x
}
