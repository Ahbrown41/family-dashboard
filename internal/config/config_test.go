package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := LoadConfig("./config_test.yaml", "./.env")
	assert.NoError(t, err)
	assert.Equal(t, "test_user", cfg.Ticktick.Username)
	assert.Equal(t, "test_password", cfg.Ticktick.Password)
	assert.Equal(t, "test_access_token", cfg.Ticktick.AccessToken)
	assert.Equal(t, "test_focus_project", cfg.Ticktick.FocusProject)
	assert.Equal(t, "test_api_url", cfg.Ticktick.ApiUrl)
	assert.Equal(t, "test_api_url", cfg.Todoist.ApiUrl)
	assert.Equal(t, "test_output", cfg.Screen.Output)
	assert.Equal(t, "test_access_token", cfg.Todoist.AccessToken)
	assert.Equal(t, "test_project", cfg.Todoist.Project)
	assert.Equal(t, []string{"test_label1", "test_label2"}, cfg.Todoist.Labels)
}
