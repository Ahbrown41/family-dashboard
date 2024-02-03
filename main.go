package main

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"family-dashboard/internal/config"
	"family-dashboard/internal/epaper"
	"family-dashboard/internal/google/calendar"
	"family-dashboard/internal/google/oauth"
	"family-dashboard/internal/todoist"
	"fmt"
	stripmd "github.com/writeas/go-strip-markdown"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

func main() {
	// Load Config
	cfg, err := config.LoadConfig("./config.yaml", "./.env")
	if err != nil {
		fmt.Printf("Error loading config: %s\n", err)
		return
	}

	// Get Hashes
	hashes := struct {
		CalendarHash, TodoHash string
	}{}
	fileBytes, err := os.ReadFile("./hash.yaml")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatalf("Unable to read hash file: %v", err)
	}
	if fileBytes != nil {
		err = yaml.Unmarshal(fileBytes, &hashes)
		if err != nil {
			log.Fatalf("Unable load hashes: %v", err)
		}
	}

	// Get Task Details
	tasksStr, err := TaskString(cfg)
	if err != nil {
		fmt.Printf("Error getting tasks: %s\n", err)
		return
	}
	apptStr, err := AppointmentString(cfg)
	if err != nil {
		fmt.Printf("Error getting appointments: %s\n", err)
		return
	}

	// Check/Write Hashes
	if hashes.CalendarHash == HashString(apptStr) && hashes.TodoHash == HashString(tasksStr) {
		fmt.Println("No changes")
		return
	} else {
		hashes.CalendarHash = HashString(apptStr)
		hashes.TodoHash = HashString(tasksStr)
		yfile, err := yaml.Marshal(&hashes)
		if err != nil {
			log.Fatalf("Unable to save hashes: %v", err)
		}
		err = os.WriteFile("./hash.yaml", yfile, 0644)
		if err != nil {
			log.Fatalf("Unable to save hashes: %v", err)
		}
	}

	// Draw Image
	err = WriteScreen(cfg.Screen, tasksStr, apptStr)
	if err != nil {
		fmt.Printf("Error writing screen: %s\n", err)
		return
	}
}

// WriteScreen writes the screen
func WriteScreen(cfg config.Screen, tasksStr string, apptStr string) error {
	// Setup Output
	paper := epaper.New()
	if cfg.Output == "screen" {
		paper.Init()
		paper.Clear()
	}

	paper.DrawCalendar(10, 10, time.Now())
	paper.DrawRectangle(10, 220, 629, 155) // Appointment Box
	paper.DrawRectangle(230, 10, 409, 210) // Task Box
	paper.DrawText(11, 240, 20, tasksStr)  // Left Bottom
	paper.DrawText(231, 28, 20, apptStr)   // Right Top

	// Write Output
	if cfg.Output == "screen" {
		paper.Render()
	} else {
		paper.SavePNG("out.png")
	}
	return nil
}

// HashString returns a sha1 hash of a string
func HashString(s string) string {
	hasher := sha1.New()
	_, err := hasher.Write([]byte(s))
	if err != nil {
		fmt.Printf("Error hashing tasks: %s\n", err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

// AppointmentString returns a string of appointments
func AppointmentString(cfg *config.Config) (string, error) {
	auth := oauth.New("./credentials.json", "./token.json")
	cal := calendar.New(auth)
	eventString := ""
	for i, evt := range cal.GetEvents() {
		eventString += fmt.Sprintf("- %s\n", evt.EventString())
		if i > 8 {
			break
		}
	}
	return eventString, nil
}

// TaskString returns a string of tasks
func TaskString(cfg *config.Config) (string, error) {
	var tasksStr string
	tdo := todoist.New(cfg.Todoist)
	tasks, err := tdo.Tasks(todoist.Task{Labels: cfg.Todoist.Labels, ProjectId: cfg.Todoist.Project})
	if err != nil {
		return "", err
	}
	for _, task := range tasks {
		due := task.Due.String
		if due, err := time.Parse(time.DateOnly, task.Due.Date); err == nil {
			if due.Before(time.Now()) {
				dueDate := fmt.Sprintf("(%s) ", task.Due.Date)
				tasksStr += fmt.Sprintf("- %s%s\n", dueDate, stripmd.Strip(task.Content))
			}
			continue
		}
		tasksStr += fmt.Sprintf("- %s%s\n", due, stripmd.Strip(task.Content))
	}
	return tasksStr, nil
}
