package calendar

import (
	"context"
	"family-dashboard/internal/google/oauth"
	"fmt"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"log"
	"time"
)

type Calendar struct {
	token *oauth.Token
	srv   *calendar.Service
}

type Event struct {
	Calendar string
	Event    calendar.Event
}

func New(token *oauth.Token) *Calendar {
	ctx := context.Background()
	client := token.GetClient()
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	return &Calendar{token: token, srv: srv}
}

func (e *Event) EventString() string {
	if e.Event.Start != nil {
		if e.Event.Start.DateTime == "" {
			return fmt.Sprintf("All Day - %s", e.Event.Summary)
		} else {
			t, _ := time.Parse(time.RFC3339, e.Event.Start.DateTime)
			return fmt.Sprintf("%s - %s", t.Format(time.Kitchen), e.Event.Summary)
		}
	}
	return fmt.Sprintf("%s", e.Event.Summary)
}

func (c *Calendar) GetEvents() []Event {
	today := time.Now()
	begin := time.Now().Format(time.RFC3339)
	end := time.Date(today.Year(), today.Month(), today.Day(), 23, 59, 59, 999, today.Location()).Format(time.RFC3339)
	log.Printf("Begin: %s End: %s\n", begin, end)
	cals, err := c.srv.CalendarList.List().
		ShowDeleted(false).
		ShowHidden(false).
		MinAccessRole("owner").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next Calendar list: %v", err)
	}
	evts := make([]Event, 0)
	for _, cal := range cals.Items {
		events, err := c.srv.Events.List(cal.Id).
			ShowDeleted(false).
			SingleEvents(true).
			TimeMin(begin).
			TimeMax(end).
			MaxResults(10).
			OrderBy("startTime").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve events for Calendar (%s): %v", cal.Id, err)
		}
		for _, evt := range events.Items {
			evts = append(evts, Event{Calendar: cal.Summary, Event: *evt})
		}
	}
	return evts
}

func (c *Calendar) PrintEvents() {
	events := c.GetEvents()
	fmt.Println("Upcoming events:")
	if len(events) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		var lastDate *calendar.EventDateTime
		calendar := ""
		for _, item := range events {
			if item.Calendar != calendar {
				fmt.Println(item.Calendar)
				calendar = item.Calendar
			}
			if item.Event.Start != lastDate {
				fmt.Println("\t" + getDateStr(item.Event.Start))
				lastDate = item.Event.Start
			}
			fmt.Printf("\t\t%s - %s (%s - %s)\n", item.Calendar, item.Event.Summary, getDateStr(item.Event.Start), getDateStr(item.Event.End))
		}
	}
}

func getDateStr(dateO *calendar.EventDateTime) string {
	date := dateO.DateTime
	if date == "" {
		date = dateO.Date
	}
	tt, err := time.Parse(time.DateOnly, date)
	if err != nil {
		tt, err = time.Parse(time.RFC3339, date)
		if err != nil {
			fmt.Println(err)
		}
		return tt.Format("2006-01-02")
	}
	return tt.Format("2006-01-02")
}
