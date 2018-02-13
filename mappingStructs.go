package main

// UserMap is used to store association between Jira user and CH user
type UserMap struct {
	JiraUsername string
	CHID         string
	Default      bool
}

// ProjectMap is used to store association between Jira project and CH project
type ProjectMap struct {
	JiraProjectKey string
	CHProjectID    int64
}

// AttachmentGroup contains the list of files associated with a story
type AttachmentGroup struct {
	JiraStoryKey   string
	JiraProjectKey string
	CHFiles        []CHFile
}

// AttachmentMigrationList is a collection of list of files associated with each story
type AttachmentMigrationList struct {
	AttachmentGroups []AttachmentGroup
}
