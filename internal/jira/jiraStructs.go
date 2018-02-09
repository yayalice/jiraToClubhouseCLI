package jira

import (
	"encoding/xml"
)

// Export is the container of Jira Items from the XML.
type Export struct {
	ElementName xml.Name `xml:"rss"`
	Items       []Item   `xml:"channel>item"`
}

type Assignee struct {
	Username string `xml:"username,attr"`
}

type Reporter struct {
	Username string `xml:"username,attr"`
}

// Item is the struct for a basic item imported from the XML
type Item struct {
	Assignee        Assignee     `xml:"assignee"`
	Attachments     []Attachment `xml:"attachments>attachment"`
	CreatedAtString string       `xml:"created"`
	Description     string       `xml:"description"`
	Key             string       `xml:"key"`
	Labels          []string     `xml:"labels>label"`
	Project         Project      `xml:"project"`
	Resolution      string       `xml:"resolution"`
	Reporter        Reporter     `xml:"reporter"`
	Status          string       `xml:"status"`
	Summary         string       `xml:"summary"`
	Title           string       `xml:"title"`
	Type            string       `xml:"type"`
	Parent          string       `xml:"parent"`

	Comments     []Comment     `xml:"comments>comment"`
	CustomFields []CustomField `xml:"customfields>customfield"`

	EpicLink string
}

// Attachment is the information for attacments
type Attachment struct {
	Author          string `xml:"author,attr"`
	CreatedAtString string `xml:"created,attr"`
	Name            string `xml:"name, attr"`
	ID              string `xml:"id,attr"`
}

// CustomField is the information for custom fields. Right now the only one used is the Epic Link
type CustomField struct {
	FieldName  string   `xml:"customfieldname"`
	FieldVales []string `xml:"customfieldvalues>customfieldvalue"`
}

// Comment is a comment from the imported XML
type Comment struct {
	Author          string `xml:"author,attr"`
	CreatedAtString string `xml:"created,attr"`
	Comment         string `xml:",chardata"`
	ID              string `xml:"id,attr"`
}

type Project struct {
	Key string `xml:"key,attr"`
}
