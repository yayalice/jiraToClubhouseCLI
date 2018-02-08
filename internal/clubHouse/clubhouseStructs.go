package clubHouse

import "time"

// ClubHouseEpic is the data returned from the Clubhouse API when the Epic is created.
type Epic struct {
	ID int64 `json:"id"`
}

// ClubHouseCreateEpic is the object sent to the API to create an Epic. ClubHouseEpic is the return from the submission of this struct.
type CreateEpic struct {
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	ExternalID  string    `json:"external_id"`
	Name        string    `json:"name"`
	id          int64     `json:"id"`
}

type File struct {
	ID         int64  `json:"id"`
	ExternalID string `json:"external_id"`
}

// ClubhouseCreateAttachment is used in ClubHouseCreateStory for attachments
type CreateAttachment struct {
	Author     string    `json:"author_id"`
	CreatedAt  time.Time `json:"created_at"`
	ExternalID string    `json:"external_id"`
	Name       string    `json:"name"`
	Id         int64     `json:"id"`
}

// ClubHouseCreateComment is used in ClubHouseCreateStory for comments.
type CreateComment struct {
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Author    string    `json:"author_id"`
}

// ClubHouseCreateStory is the object sent to API to submit a Story, Tasks, & Comment
type CreateStory struct {
	Comments    []CreateComment `json:"comments"`
	CreatedAt   time.Time       `json:"created_at"`
	Description string          `json:"description"`
	Estimate    int64           `json:"estimate"`
	EpicID      int64           `json:"epic_id,omitempty"`
	EpicLink    string
	ExternalID  string        `json:"external_id"`
	Labels      []CreateLabel `json:"labels"`
	Name        string        `json:"name"`
	ProjectID   int64         `json:"project_id"`
	Tasks       []CreateTask  `json:"tasks"`
	StoryType   string        `json:"story_type"`

	OwnerIDs      []string `json:"owner_ids"`
	WorkflowState int64    `json:"workflow_state_id"`
	RequestedBy   string   `json:"requested_by_id"`
}

// ClubHouseCreateLabel is used to submit labels with stories, it looks like from the API that duplicates will not be created.
type CreateLabel struct {
	Name string `json:"name"`
}

// ClubHouseCreateTask is used for Tasks in stories.
type CreateTask struct {
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Complete    bool      `json:"complete"`
	Parent      string
}

// ClubHouseData is a container holding the data for submission of writing to a JSON file.
type Data struct {
	Epics   []CreateEpic  `json:"epics"`
	Stories []CreateStory `json:"stories"`
}
