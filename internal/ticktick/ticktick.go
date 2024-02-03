package ticktick

import (
	"encoding/json"
	"family-dashboard/internal/config"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
)

type TickTick struct {
	cfg    config.Ticktick
	client resty.Client
}

// New creates a new TickTick client
func New(cfg config.Ticktick) *TickTick {
	return &TickTick{cfg: cfg, client: *resty.New()}
}

// Client returns the resty client
func (c *TickTick) Client() *http.Client {
	return c.client.GetClient()
}

// apiRequest returns a new resty request with the access token set
func (c *TickTick) request(params map[string]string) *resty.Request {
	return c.client.R().SetAuthToken(c.cfg.AccessToken).SetQueryParams(params)
}

// Projects returns all projects
func (c *TickTick) Projects() ([]Project, error) {
	req := c.request(nil)
	resp, err := req.Get(fmt.Sprintf(c.cfg.ApiUrl + "/project"))
	if err != nil {
		log.Panic(err)
	}
	var projects []Project
	err = json.Unmarshal(resp.Body(), &projects)
	return projects, err
}

// Task returns a single task
func (c *TickTick) Task(projectId, taskId string) (Task, error) {
	req := c.request(nil)
	resp, err := req.Get(fmt.Sprintf(c.cfg.ApiUrl+"/project/%s/task/%s", projectId, taskId))
	if err != nil {
		log.Panic(err)
	}
	var task Task
	err = json.Unmarshal(resp.Body(), &task)
	return task, err
}

func (c *TickTick) ProjectData(projectId string) (ProjectData, error) {
	req := c.request(nil)
	resp, err := req.Get(fmt.Sprintf(c.cfg.ApiUrl+"/project/%s/data", projectId))
	if err != nil {
		log.Panic(err)
	}
	var data ProjectData
	err = json.Unmarshal(resp.Body(), &data)
	return data, err
}
