package main

import (
	"encoding/xml"
)

// JiraExport is the container of Jira Items from the XML.
type JiraExport struct {
	ElementName xml.Name   `xml:"rss"`
	Items       []JiraItem `xml:"channel>item"`
}

// JiraAssignee is the container of the username of an assignee from the xml
type JiraAssignee struct {
	Username string `xml:"username,attr"`
}

// JiraReporter is the container of the username of a reporter from the xml
type JiraReporter struct {
	Username string `xml:"username,attr"`
}

// JiraItem is the struct for a basic item imported from the XML
type JiraItem struct {
	Assignee        JiraAssignee     `xml:"assignee"`
	Attachments     []JiraAttachment `xml:"attachments>attachment"`
	CreatedAtString string           `xml:"created"`
	Description     string           `xml:"description"`
	Key             string           `xml:"key"`
	Labels          []string         `xml:"labels>label"`
	Project         JiraProject      `xml:"project"`
	Resolution      string           `xml:"resolution"`
	Reporter        JiraReporter     `xml:"reporter"`
	Status          string           `xml:"status"`
	Summary         string           `xml:"summary"`
	Title           string           `xml:"title"`
	Type            string           `xml:"type"`
	Parent          string           `xml:"parent"`

	Comments     []JiraComment     `xml:"comments>comment"`
	CustomFields []JiraCustomField `xml:"customfields>customfield"`

	EpicLink string
}

// JiraAttachment is the information for attacments
type JiraAttachment struct {
	Author          string `xml:"author,attr"`
	CreatedAtString string `xml:"created,attr"`
	Name            string `xml:"name, attr"`
	ID              string `xml:"id,attr"`
}

// JiraCustomField is the information for custom fields. Right now the only one used is the Epic Link
type JiraCustomField struct {
	FieldName  string   `xml:"customfieldname"`
	FieldVales []string `xml:"customfieldvalues>customfieldvalue"`
}

// JiraComment is a comment from the imported XML
type JiraComment struct {
	Author          string `xml:"author,attr"`
	CreatedAtString string `xml:"created,attr"`
	Comment         string `xml:",chardata"`
	ID              string `xml:"id,attr"`
}

type JiraProject struct {
	Key string `xml:"key,attr"`
}
