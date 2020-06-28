package incident

import (
	"time"

	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type Incident struct {
	Id              string            `json:"id"`
	ServiceId       string            `json:"serviceId"`
	TinyId          string            `json:"tinyId"`
	Message         string            `json:"message"`
	Status          string            `json:"status"`
	Tags            []string          `json:"tags"`
	CreatedAt       time.Time         `json:"createdAt"`
	UpdatedAt       time.Time         `json:"updatedAt"`
	Priority        Priority          `json:"priority"`
	OwnerTeam       string            `json:"ownerTeam"`
	Responders      []Responder       `json:"responders"`
	ExtraProperties map[string]string `json:"extraProperties"`
}

type RequestStatusResult struct {
	client.ResultMetadata
	Success       bool   `json:"success"`
	Action        string `json:"action"`
	ProcessedAt   string `json:"processedAt"`
	IntegrationId string `json:"integrationId"`
	IsSuccess     bool   `json:"isSuccess"`
	Status        string `json:"status"`
	IncidentId    string `json:"incidentId"`
}

type AsyncResult struct {
	client.ResultMetadata
	Result string `json:"result"`
}

type GetResult struct {
	client.ResultMetadata
	Incident
}

type ListResult struct {
	client.ResultMetadata
	Incidents []Incident `json:"data"`
	Paging    Paging     `json:"paging"`
}

type LogResult struct {
	Log       string    `json:"log"`
	Type      string    `json:"type"`
	Owner     string    `json:"owner"`
	CreatedAt time.Time `json:"createdAt"`
	Offset    string    `json:"offset"`
}

type ListLogsResult struct {
	client.ResultMetadata
	Logs   []LogResult `json:"data"`
	Paging Paging      `json:"paging"`
}

type NoteResult struct {
	Note      string    `json:"note"`
	Owner     string    `json:"owner"`
	CreatedAt time.Time `json:"createdAt"`
	Offset    string    `json:"offset"`
}

type ListNotesResult struct {
	client.ResultMetadata
	Notes  []NoteResult `json:"data"`
	Paging Paging       `json:"paging"`
}

type Actor struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type TimelineEntry struct {
	ID          string    `json:"id"`
	Group       string    `json:"group"`
	Type        string    `json:"type"`
	EventTime   time.Time `json:"eventTime"`
	Hidden      bool      `json:"hidden"`
	Actor       Actor     `json:"actor"`
	Description struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"description"`
	LastEdit struct {
		EditTime time.Time `json:"editTime"`
		Actor    Actor     `json:"actor"`
	} `json:"lastEdit"`
}

type ListTimelineResult struct {
	client.ResultMetadata
	Data struct {
		Entries    []TimelineEntry `json:"entries"`
		NextOffset string          `json:"nextOffset"`
	}
}

type Paging struct {
	Next  string `json:"next"`
	Prev  string `json:"prev"`
	First string `json:"first"`
	Last  string `json:"last"`
}
