package ticktick

type Project struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	SortOrder int64  `json:"sortOrder"`
	Kind      string `json:"kind"`
}

type Column struct {
	Id        string `json:"id"`
	ProjectId string `json:"projectId"`
	Name      string `json:"name"`
	SortOrder int    `json:"sortOrder"`
}

type ProjectData struct {
	Project `json:"project"`
	Tasks   []Task   `json:"tasks"`
	Columns []Column `json:"columns"`
}
