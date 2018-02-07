package main

import "time"

// ClubHouseEpic is the data returned from the Clubhouse API when the Epic is created.
type ClubHouseEpic struct {
	ID int64 `json:"id"`
}

// ClubHouseCreateEpic is the object sent to the API to create an Epic. ClubHouseEpic is the return from the submission of this struct.
type ClubHouseCreateEpic struct {
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	ExternalID  string    `json:"external_id"`
	Name        string    `json:"name"`
	id          int64     `json:"id"`
}

type ClubHouseFile struct {
	ID         int64  `json:"id"`
	ExternalID string `json:"external_id"`
}

// ClubhouseCreateAttachment is used in ClubHouseCreateStory for attachments
type ClubHouseCreateAttachment struct {
	Author     string    `json:"author_id"`
	CreatedAt  time.Time `json:"created_at"`
	ExternalID string    `json:"external_id"`
	Name       string    `json:"name"`
	id         int64     `json:"id"`
}

// ClubHouseCreateComment is used in ClubHouseCreateStory for comments.
type ClubHouseCreateComment struct {
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Author    string    `json:"author_id"`
}

// ClubHouseCreateStory is the object sent to API to submit a Story, Tasks, & Comment
type ClubHouseCreateStory struct {
	Comments    []ClubHouseCreateComment `json:"comments"`
	CreatedAt   time.Time                `json:"created_at"`
	Description string                   `json:"description"`
	Estimate    int64                    `json:"estimate"`
	EpicID      int64                    `json:"epic_id,omitempty"`
	epicLink    string
	ExternalID  string                 `json:"external_id"`
	Labels      []ClubHouseCreateLabel `json:"labels"`
	Name        string                 `json:"name"`
	ProjectID   int64                  `json:"project_id"`
	Tasks       []ClubHouseCreateTask  `json:"tasks"`
	StoryType   string                 `json:"story_type"`

	OwnerIDs      []string `json:"owner_ids"`
	WorkflowState int64    `json:"workflow_state_id"`
	RequestedBy   string   `json:"requested_by_id"`
}

// ClubHouseCreateLabel is used to submit labels with stories, it looks like from the API that duplicates will not be created.
type ClubHouseCreateLabel struct {
	Name string `json:"name"`
}

// ClubHouseCreateTask is used for Tasks in stories.
type ClubHouseCreateTask struct {
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Complete    bool      `json:"complete"`
	parent      string
}

// ClubHouseData is a container holding the data for submission of writing to a JSON file.
type ClubHouseData struct {
	Epics   []ClubHouseCreateEpic  `json:"epics"`
	Stories []ClubHouseCreateStory `json:"stories"`
}
