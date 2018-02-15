package main

import "time"

// CHEpic is the object sent to the API to create an Epic. ClubHouseEpic is the return from the submission of this struct.
type CHEpic struct {
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	ExternalID  string    `json:"external_id"`
	Name        string    `json:"name"`
	id          int64     `json:"id"`
}

// CHFile is used in ClubHouseCreateStory for creating attachments
type CHFile struct {
	Uploader   string    `json:"uploader_id"`
	CreatedAt  time.Time `json:"created_at"`
	ExternalID string    `json:"external_id"`
	Name       string    `json:"name"`
	uploaded   bool
}

// CHFile is used to fetch attachments
type CHGETFile struct {
	Uploader   string    `json:"uploader_id"`
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
	epicLink    string
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

type CHGETStory struct {
	Comments      []CHComment `json:"comments"`
	CreatedAt     time.Time   `json:"created_at"`
	Description   string      `json:"description"`
	EpicID        int64       `json:"epic_id,omitempty"`
	Estimate      int64       `json:"estimate"`
	ExternalID    string      `json:"external_id"`
	Files         []CHGETFile `json:"files"`
	Labels        []CHLabel   `json:"labels"`
	Name          string      `json:"name"`
	OwnerIDs      []string    `json:"owner_ids"`
	ProjectID     int64       `json:"project_id"`
	RequestedBy   string      `json:"requested_by_id"`
	StoryType     string      `json:"story_type"`
	Tasks         []CHTask    `json:"tasks"`
	WorkflowState int64       `json:"workflow_state_id"`
}

type CHStoryForUpdate struct {
	FileIDs []int64 `json:"file_ids"`
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
	parent      string
}

// CHData is a container holding the data for submission of writing to a JSON file.
type CHData struct {
	Epics   []CHEpic  `json:"epics"`
	Stories []CHStory `json:"stories"`
}

// CHMember is a container holding the data for fetching user information from CH
type CHMember struct {
	ID      string    `json:"id"`
	Profile CHProfile `json:"profile"`
}

// CHProfile is a sub-part of the CHMember-container holding the data for fetching user information from CH
type CHProfile struct {
	Name string `json:"name"`
}
