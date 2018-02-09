package main

import "fmt"

// MapUser tries to map a given jira user with a clubhouse user
func MapUser(userMaps []userMap, jiraUsername string) (CHID string, err error) {

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
func MapProject(projectMaps []projectMap, jiraProjectKey string) int {
	projectID := GetProjectInfo(projectMaps, jiraProjectKey)

	if projectID == 0 {
		fmt.Println("[MapProject] JIRA project not found: ", jiraProjectKey)
		return 299
	}

	return projectID
}
