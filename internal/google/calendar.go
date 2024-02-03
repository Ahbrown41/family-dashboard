package main

import (
	"family-dashboard/internal/google/calendar"
	"family-dashboard/internal/google/oauth"
)

func main() {
	auth := oauth.New("../credentials.json", "../token.json")
	cal := calendar.New(auth)
	cal.PrintEvents()
}
