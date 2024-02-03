package ticktick

import (
	"family-dashboard/internal/config"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestProjects(t *testing.T) {
	tick := New(config.Ticktick{
		ApiUrl:      "https://api.ticktick.com/open/v1",
		AccessToken: "",
	})

	httpmock.ActivateNonDefault(tick.Client())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.ticktick.com/open/v1/project",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, []Project{
				{
					Id:   "1",
					Name: "My Project",
				},
			})
			return resp, err
		},
	)

	projs, err := tick.Projects()
	assert.NoError(t, err)
	assert.NotNil(t, projs)
	assert.NotEmpty(t, projs)
	assert.Equal(t, "1", projs[0].Id)
	assert.Equal(t, "My Project", projs[0].Name)
}

func TestProjectData(t *testing.T) {
	tick := New(config.Ticktick{
		ApiUrl:      "https://api.ticktick.com/open/v1",
		AccessToken: "",
	})

	httpmock.ActivateNonDefault(tick.Client())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.ticktick.com/open/v1/project/1/data",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, ProjectData{
				Project: Project{
					Id:   "1",
					Name: "My Project",
				},
				Tasks: []Task{
					{
						Id:    "1",
						Title: "My Task",
					},
				},
				Columns: nil,
			})
			return resp, err
		},
	)

	proj, err := tick.ProjectData("1")
	assert.NoError(t, err)
	assert.NotNil(t, proj)
	assert.NotEmpty(t, proj.Tasks)
	assert.Equal(t, "1", proj.Project.Id)
	assert.Equal(t, "My Project", proj.Project.Name)
	assert.Equal(t, "1", proj.Tasks[0].Id)
	assert.Equal(t, "My Task", proj.Tasks[0].Title)
}

func TestTask(t *testing.T) {
	tick := New(config.Ticktick{
		ApiUrl:      "https://api.ticktick.com/open/v1",
		AccessToken: "",
	})

	httpmock.ActivateNonDefault(tick.Client())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.ticktick.com/open/v1/project/1/task/1",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, Task{
				Id:    "1",
				Title: "My Task",
			})
			return resp, err
		},
	)

	proj, err := tick.Task("1", "1")
	assert.NoError(t, err)
	assert.NotNil(t, proj)
	assert.Equal(t, "1", proj.Id)
	assert.Equal(t, "My Task", proj.Title)
}
