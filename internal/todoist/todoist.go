package todoist

import (
	"encoding/json"
	"family-dashboard/internal/config"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type Todoist struct {
	cfg    config.Todoist
	client resty.Client
}

// New returns a new Todoist client
func New(cfg config.Todoist) *Todoist {
	return &Todoist{cfg: cfg, client: *resty.New()}
}

// apiRequest returns a new resty request with the access token set
func (c *Todoist) request(params map[string]string) *resty.Request {
	return c.client.R().SetAuthToken(c.cfg.AccessToken).SetQueryParams(params)
}

// Client returns the resty client
func (c *Todoist) Client() *http.Client {
	return c.client.GetClient()
}

// Tasks returns all tasks
func (c *Todoist) Tasks(filter Task) ([]Task, error) {
	req := c.request(filter.query())
	url := fmt.Sprintf("%s/%s", c.cfg.ApiUrl, api_path_tasks)
	resp, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("invalid Response %d", resp.StatusCode())
	}
	var tasks []Task
	err = json.Unmarshal(resp.Body(), &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, err
}
