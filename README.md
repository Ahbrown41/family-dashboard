# Family Dashboard

## Description
This is a family dashboard that allows family members to keep track 
of their schedules, chores, and other important information.  The 
dashboard leverages a Waveshare 7.5inch e-Paper HAT for Raspberry Pi.

## Features
- Display of a calendar of the current month with the current day highlighted.
- Integration with Google Calendar to display upcoming events.
- Integration with Todoist or TickTick to display upcoming tasks.

## Installation
- Clone the repository
- Build the application using
```shell
$ go build -o family-dashboard .
```
- Update Configuration file `config.yaml` with the appropriate settings.
```yaml
screen:
  output: epaper
ticktick:
  api_url: https://api.ticktick.com/open/v1
  focus_project: Your Focus Project ID
  username: Tick Tick Username
  password: Tick Tick Password
  access_token: The access token from the Tick Tick API
todoist:
  api_url: https://api.todoist.com/rest/v2
  access_token: Use .env file
  project: Use .env file
  labels:
    - Family
```
- Run the application using
```shell
$ ./family-dashboard
```