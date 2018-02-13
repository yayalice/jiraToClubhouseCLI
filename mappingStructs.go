package main

type userMap struct {
	JiraUsername string
	CHID         string
	Default      bool
}

type projectMap struct {
	JiraProjectKey string
	CHProjectID    int64
}

type attachmentMap struct {
	JiraAttachmentKey string
	CHAttachmentID    int64
	JiraStoryKey      string
	CHStoryID         int64
}

type attachmentGroup struct {
	JiraStoryKey   string
	JiraProjectKey string
	CHFiles        []CHFile
}

type attachmentMigrationList struct {
	AttachmentGroups []attachmentGroup
}
