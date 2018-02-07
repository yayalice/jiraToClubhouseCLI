package main

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kennygrant/sanitize"
)

// JiraExport is the container of Jira Items from the XML.
type JiraExport struct {
	ElementName xml.Name   `xml:"rss"`
	Items       []JiraItem `xml:"channel>item"`
}

type JiraAssignee struct {
	Username string `xml:"username,attr"`
}

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

//JiraCustomField is the information for custom fields. Right now the only one used is the Epic Link
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

func GetUserInfo(userMaps []userMap, jiraUsername string) (CHID string) {

	defaultUser := ""

	for _, u := range userMaps {
		if u.JiraUsername == jiraUsername {
			return u.CHID
		}
		if u.Default == true {
			defaultUser = u.CHID
		}
	}
	if defaultUser != "" {
		fmt.Printf("Unknown user %s. Will use default user: %s\n\n", jiraUsername, defaultUser)
	} else {
		fmt.Printf("Unknown user %s. No default user defined in userMap. This story will not be created\n\n", jiraUsername)
	}

	return defaultUser
}

func GetProjectInfo(projectMaps []projectMap, jiraProjectKey string) (CHProjectID int) {

	for _, u := range projectMaps {
		// fmt.Printf("JiraProjectKey: %s | CHProjectID: %d\n\n", u.JiraProjectKey, u.CHProjectID)
		if u.JiraProjectKey == jiraProjectKey {
			return u.CHProjectID
		}
	}
	return 0
}

//GetDataForClubhouse will take the data from the XML and translate it into a format for sending to Clubhouse
func (je *JiraExport) GetDataForClubhouse(userMaps []userMap, projectMaps []projectMap) ClubHouseData {
	epics := []JiraItem{}
	tasks := []JiraItem{}
	stories := []JiraItem{}

	for _, item := range je.Items {
		switch item.Type {
		case "Epic":
			epics = append(epics, item)
			break
		case "Sub-task":
			tasks = append(tasks, item)
			break
		default:
			stories = append(stories, item)
			break
		}
	}

	chEpics := []ClubHouseCreateEpic{}

	for _, item := range epics {
		chEpics = append(chEpics, item.CreateEpic())
	}

	chTasks := []ClubHouseCreateTask{}
	chStories := []ClubHouseCreateStory{}

	for _, item := range tasks {
		chTasks = append(chTasks, item.CreateTask())
	}

	for _, item := range stories {
		chStories = append(chStories, item.CreateStory(userMaps, projectMaps))
	}

	// storyMap is used to link the JiraItem's key to its index in the chStories slice. This is then used to assign subtasks properly
	storyMap := make(map[string]int)
	for i, item := range chStories {
		storyMap[item.ExternalID] = i
	}

	for _, task := range chTasks {
		chStories[storyMap[task.parent]].Tasks = append(chStories[storyMap[task.parent]].Tasks, task)
	}

	return ClubHouseData{Epics: chEpics, Stories: chStories}
}

// CreateEpic returns a ClubHouseCreateEpic from the JiraItem
func (item *JiraItem) CreateEpic() ClubHouseCreateEpic {
	fmt.Printf("Epic Name: %s | Description: %s | Summary: %s\n\n", item.GetEpicName(), item.Description, item.Summary)

	return ClubHouseCreateEpic{Description: sanitize.HTML(item.Summary + "<br><br>" + item.Description), Name: sanitize.HTML(item.GetEpicName()), ExternalID: item.Key, CreatedAt: ParseJiraTimeStamp(item.CreatedAtString)}
}

// CreateTask returns a task if the item is a Jira Sub-task
func (item *JiraItem) CreateTask() ClubHouseCreateTask {
	return ClubHouseCreateTask{Description: sanitize.HTML(item.Summary), parent: item.Parent, Complete: false}
}

// CreateStory returns a ClubHouseCreateStory from the JiraItem
func (item *JiraItem) CreateStory(userMaps []userMap, projectMaps []projectMap) ClubHouseCreateStory {
	// fmt.Println("assignee: ", item.Assignee, "reporter: ", item.Reporter)
	// return ClubHouseCreateStory{}

	attachments := []ClubHouseCreateAttachment{}
	for _, attch := range item.Attachments {
		attachments = append(attachments, attch.CreateAttachment(userMaps))
	}

	comments := []ClubHouseCreateComment{}
	for _, c := range item.Comments {
		comments = append(comments, c.CreateComment(userMaps))
	}

	labels := []ClubHouseCreateLabel{}
	for _, label := range item.Labels {
		labels = append(labels, ClubHouseCreateLabel{Name: strings.ToLower(label)})
	}
	// Adding special label that indicates that it was imported from JIRA
	labels = append(labels, ClubHouseCreateLabel{Name: "JIRA"})

	// Adding Sprint as label
	sprintLabel := item.GetSprint()
	if sprintLabel != "" {
		labels = append(labels, ClubHouseCreateLabel{Name: sprintLabel})
	}

	// Overwrite supplied Project ID
	projectID := MapProject(projectMaps, item.Project.Key)
	// projectID, ownerID := GetUserInfo(userMaps, item.Assignee.Username)

	// Map JIRA assignee to Clubhouse owner(s)
	// Leave array empty if username is unknown
	// Must use "make" function to force empty array for correct JSON marshalling
	ownerID := MapUser(userMaps, item.Assignee.Username)
	var owners []string
	if ownerID != "" {
		// owners := []string{ownerID}
		owners = append(owners, ownerID)
	} else {
		owners = make([]string, 0)
	}

	// Map JIRA status to Clubhouse Workflow state
	// cases break automatically, no fallthrough by default
	var state int64 = 500000014
	switch item.Status {
	case "Open":
		// Open
		state = 500000003
	case "Done":
		// Done
		state = 500000002
	case "In Development":
		// In Development
		state = 500000004
	case "Waiting for Code Review":
		// Ready for Review
		state = 500000005
	case "In Code Review":
		// In Review
		state = 500000018
	case "Waiting for UX-Interaction Design Review":
		// Ready for Review
		state = 500000005
	case "In UX-Interaction Design Review":
		// In Review
		state = 500000018
	case "Waiting for UX-Design Review":
		// Ready for Review
		state = 500000005
	case "In UX-Design Review":
		// In Review
		state = 500000018
	case "Waiting for QA":
		// Ready for Test
		state = 500000017
	case "In QA":
		// In Test
		state = 500000019
	case "In QA Review":
		// In Test
		state = 500000019
	default:
		// Open
		state = 500000003
	}

	requestor := MapUser(userMaps, item.Reporter.Username)
	// _, requestor := GetUserInfo(userMaps, item.Reporter.Username)

	fmt.Printf("%s: JIRA Assignee: %s | Project: %d | Status: %s | Description: %s | Estimate: %d | Epic Link: %s | SprintTag: %s\n\n", item.Key, item.Assignee.Username, projectID, item.Status, item.GetDescription(), item.GetEstimate(), item.GetEpicLink(), item.GetSprint())

	return ClubHouseCreateStory{
		Comments:      comments,
		CreatedAt:     ParseJiraTimeStamp(item.CreatedAtString),
		Description:   item.GetDescription(),
		ExternalID:    item.Key,
		Labels:        labels,
		Name:          sanitize.HTML(item.Summary),
		ProjectID:     int64(projectID),
		StoryType:     item.GetClubhouseType(),
		epicLink:      item.GetEpicLink(),
		WorkflowState: state,
		OwnerIDs:      owners,
		RequestedBy:   requestor,
		Estimate:      item.GetEstimate(),
	}
}

func MapUser(userMaps []userMap, jiraUserName string) string {
	chUserID := GetUserInfo(userMaps, jiraUserName)

	if chUserID == "" {
		fmt.Println("[MapUser] JIRA user not found: ", jiraUserName)
		return ""
	}

	return chUserID
}

func MapProject(projectMaps []projectMap, jiraProjectKey string) int {
	projectID := GetProjectInfo(projectMaps, jiraProjectKey)

	if projectID == 0 {
		fmt.Println("[MapProject] JIRA project not found: ", jiraProjectKey)
		return 299
	}

	return projectID
}

func (attachment *JiraAttachment) CreateAttachment(userMaps []userMap) ClubHouseCreateAttachment {
	author := MapUser(userMaps, attachment.Author)

	return ClubHouseCreateAttachment{
		Author:     author,
		CreatedAt:  ParseJiraTimeStamp(attachment.CreatedAtString),
		ExternalID: attachment.ID,
		Name:       attachment.Name,
	}
}

// CreateComment takes the JiraItem's comment data and returns a ClubHouseCreateComment
func (comment *JiraComment) CreateComment(userMaps []userMap) ClubHouseCreateComment {
	commentText := sanitize.HTML(comment.Comment)
	if commentText == "\n" {
		commentText = "(empty)"
	}
	author := MapUser(userMaps, comment.Author)

	return ClubHouseCreateComment{
		Text:      commentText,
		CreatedAt: ParseJiraTimeStamp(comment.CreatedAtString),
		Author:    author,
	}
}

// GetEpicLink returns the Epic Link of a Jira Item.
func (item *JiraItem) GetEpicLink() string {
	for _, cf := range item.CustomFields {
		if cf.FieldName == "Epic Link" {
			return cf.FieldVales[0]
		}
	}
	return ""
}

// GetAcceptanceCriteria returns the acceptance criteria
func (item *JiraItem) GetAcceptanceCriteria() string {
	for _, cf := range item.CustomFields {
		if cf.FieldName == "Acceptance Criteria" {
			header := "<br>## Acceptance Criteria<br>"
			return header + cf.FieldVales[0]
		}
	}
	return ""
}

// GetEstimate returns the Story Points
func (item *JiraItem) GetEstimate() int64 {
	for _, cf := range item.CustomFields {
		if cf.FieldName == "Story Points" {
			storyPoint := cf.FieldVales[0]
			return ParseFloatStringToInt(storyPoint)
		}
	}
	return 0
}

// ParseFloatStringToInt parses a string containing a float into an uprounded int
func ParseFloatStringToInt(sFloat string) int64 {
	f, err := strconv.ParseFloat(sFloat, 64)
	if err == nil {
		i := int64(f + 0.5)
		return i
	}
	return 0
}

// GetDescription returns a concatenation of description and acceptance criteria
func (item *JiraItem) GetDescription() string {
	return sanitize.HTML(item.Description + item.GetAcceptanceCriteria())
}

// GetSprint returns a string to be used as tag for srint grouping
func (item *JiraItem) GetSprint() string {

	for _, cf := range item.CustomFields {
		if cf.FieldName == "Sprint" {
			sprint := cf.FieldVales[0]

			startPoint := strings.Index(sprint, "Sprint")
			if startPoint == -1 {
				startPoint = 0
			}
			sprintAfterNoise := sprint[startPoint:len(sprint)] + " " + item.Project.Key
			sprintAsTag := strings.ToLower(strings.Replace(sprintAfterNoise, " ", "_", -1))

			return sprintAsTag
		}
	}
	return ""

}

// GetEpicName returns the name of an epic stored in custom fields
func (item *JiraItem) GetEpicName() string {
	for _, cf := range item.CustomFields {
		if cf.FieldName == "Epic Name" {
			epicName := cf.FieldVales[0]
			return epicName
		}
	}
	return ""
}

// GetClubhouseType determines type based on if the Jira item is a bug or not.
func (item *JiraItem) GetClubhouseType() string {
	if item.Type == "Bug" {
		return "bug"
	}
	return "feature"
}

// ParseJiraTimeStamp parses the format in the XML using Go's magical timestamp.
func ParseJiraTimeStamp(dateString string) time.Time {
	format := "Mon, 2 Jan 2006 15:04:05 -0700"
	t, err := time.Parse(format, dateString)
	if err != nil {
		return time.Now()
	}
	return t
}
