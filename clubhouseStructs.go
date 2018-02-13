package main

import "time"

// CHEpic is the object sent to the API to create an Epic. ClubHouseEpic is the return from the submission of this struct.
type CHEpic struct {
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	ExternalID  string    `json:"external_id"`
	Name        string    `json:"name"`
	ID          int64     `json:"id"`
}

// CHFile is used in ClubHouseCreateStory for attachments
type CHFile struct {
	Author     string    `json:"author_id"`
	CreatedAt  time.Time `json:"created_at"`
	ExternalID string    `json:"external_id"`
	Name       string    `json:"name"`
	ID         int64     `json:"id"`
}

// CHComment is used in ClubHouseCreateStory for comments.
type CHComment struct {
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Author    string    `json:"author_id"`
}

// CHStory is the object sent to API to submit a Story, Tasks, & Comment
type CHStory struct {
	Comments    []CHComment `json:"comments"`
	CreatedAt   time.Time   `json:"created_at"`
	Description string      `json:"description"`
	Estimate    int64       `json:"estimate"`
	EpicID      int64       `json:"epic_id,omitempty"`
	EpicLink    string
	ExternalID  string    `json:"external_id"`
	FileIDs     []int64   `json:"file_ids"`
	Labels      []CHLabel `json:"labels"`
	Name        string    `json:"name"`
	ProjectID   int64     `json:"project_id"`
	Tasks       []CHTask  `json:"tasks"`
	StoryType   string    `json:"story_type"`

	OwnerIDs      []string `json:"owner_ids"`
	WorkflowState int64    `json:"workflow_state_id"`
	RequestedBy   string   `json:"requested_by_id"`
}

// CHStorySlim is the object fetched from the API to get a list of stories
type CHStorySlim struct {
	ID         int64   `json:"id"`
	ExternalID string  `json:"external_id"`
	Name       string  `json:"name"`
	FileIDs    []int64 `json:"file_ids"`
}

// CHLabel is used to submit labels with stories, it looks like from the API that duplicates will not be created.
type CHLabel struct {
	Name string `json:"name"`
}

// CHTask is used for Tasks in stories.
type CHTask struct {
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Complete    bool      `json:"complete"`
	Parent      string
}

// CHData is a container holding the data for submission of writing to a JSON file.
type CHData struct {
	Epics   []CHEpic  `json:"epics"`
	Stories []CHStory `json:"stories"`
}
