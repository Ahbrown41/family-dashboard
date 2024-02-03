package ticktick

import (
	"time"
)

const dateFormat = "2006-01-02T15:04:05.000-0700"

type Task struct {
	Id            string   `json:"id"`
	IsAllDay      bool     `json:"isAllDay"`
	ProjectId     string   `json:"projectId"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	Desc          string   `json:"desc"`
	TimeZone      string   `json:"timeZone"`
	RepeatFlag    string   `json:"repeatFlag"`
	StartDate     string   `json:"startDate"`
	DueDate       string   `json:"dueDate"`
	Reminders     []string `json:"reminders"`
	Priority      int      `json:"priority"`
	Status        int      `json:"status"`
	CompletedTime string   `json:"completedTime"`
	SortOrder     int      `json:"sortOrder"`
	Items         []struct {
		Id            string `json:"id"`
		Status        int    `json:"status"`
		Title         string `json:"title"`
		SortOrder     int    `json:"sortOrder"`
		StartDate     string `json:"startDate"`
		IsAllDay      bool   `json:"isAllDay"`
		TimeZone      string `json:"timeZone"`
		CompletedTime string `json:"completedTime"`
	} `json:"items"`
}

func (c *Task) StartTime() *time.Time {
	if c.StartDate != "" {
		t, _ := time.Parse(dateFormat, c.StartDate)
		return &t
	}
	return nil
}

func (c *Task) DueTime() *time.Time {
	if c.DueDate != "" {
		t, _ := time.Parse(dateFormat, c.DueDate)
		return &t
	}
	return nil
}
