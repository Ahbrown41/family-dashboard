package todoist

import (
	"strings"
	"time"
)

const api_path_tasks = "tasks"

// Task represents a task in Todoist
type Task struct {
	CreatorId    string    `json:"creator_id"`
	CreatedAt    time.Time `json:"created_at"`
	AssigneeId   string    `json:"assignee_id"`
	AssignerId   string    `json:"assigner_id"`
	CommentCount int       `json:"comment_count"`
	IsCompleted  bool      `json:"is_completed"`
	Content      string    `json:"content"`
	Description  string    `json:"description"`
	Due          struct {
		Date        string    `json:"date"`
		IsRecurring bool      `json:"is_recurring"`
		Datetime    time.Time `json:"datetime"`
		String      string    `json:"string"`
		Timezone    string    `json:"timezone"`
	} `json:"due"`
	Duration  interface{} `json:"duration"`
	Id        string      `json:"id"`
	Labels    []string    `json:"labels"`
	Order     int         `json:"order"`
	Priority  int         `json:"priority"`
	ProjectId string      `json:"project_id"`
	SectionId string      `json:"section_id"`
	ParentId  string      `json:"parent_id"`
	Url       string      `json:"url"`
}

func (t Task) query() map[string]string {
	params := make(map[string]string)
	if t.ProjectId != "" {
		params["project_id"] = t.ProjectId
	}
	if t.Labels != nil {
		params["label"] = strings.Join(t.Labels, ",")
	}
	return params
}
