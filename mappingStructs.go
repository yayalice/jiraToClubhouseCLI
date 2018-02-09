package main

type userMap struct {
	JiraUsername string
	CHID         string
	Default      bool
}

type projectMap struct {
	JiraProjectKey string
	CHProjectID    int
}

type attachmentMap struct {
	JiraAttachmentKey string
	CHAttachmentID    int
}
