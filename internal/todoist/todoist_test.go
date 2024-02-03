package todoist

import (
	"family-dashboard/internal/config"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestClient_GetTasks(t *testing.T) {
	tok := New(config.Todoist{
		ApiUrl:      "https://example.com/api",
		AccessToken: "",
		Project:     "",
		Labels:      nil,
	})

	httpmock.ActivateNonDefault(tok.Client())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://example.com/api/tasks?label=Family",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, []Task{
				{
					Id:           "1",
					CreatorId:    "1",
					CreatedAt:    time.Now(),
					AssigneeId:   "1",
					AssignerId:   "",
					CommentCount: 0,
					IsCompleted:  false,
					Content:      "",
					Description:  "My Task",
					Duration:     nil,
					Labels:       []string{"Family"},
				},
			})
			return resp, err
		},
	)

	tasks, err := tok.Tasks(Task{Labels: []string{"Family"}})
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
	assert.NotEmpty(t, tasks)
	assert.Equal(t, "1", tasks[0].Id)
	assert.Equal(t, "My Task", tasks[0].Description)
}
