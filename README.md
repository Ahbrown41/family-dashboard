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

## Configuration
- Update Configuration file `config.yaml` with the appropriate settings.
```yaml
screen:
  output: epaper
ticktick:
  api_url: https://api.ticktick.com/open/v1
  focus_project: Your Focus Project ID
  access_token: The access token from the Tick Tick API
todoist:
  api_url: https://api.todoist.com/rest/v2
  access_token: The access token from the Todoist API
  project: Your Project ID
  labels:
    - Label ID
```
In addition you can use environment variables or a ".env" file to set the configuration.  The following environment variables are supported:
The environment variables take precedence over the `config.yaml` configuration file.
```shell
SCREEN.OUTPUT=png
TICKTICK.ACCESS_TOKEN=XXX
TICKTICK.FOCUS_PROJECT=XXX
TODOIST.ACCESS_TOKEN=XXX
TODOIST.PROJECT=XXX
TODOIST.LABELS=XXX
```

## Running the Application

- Run the application using the following to test it out
```shell
$ ./family-dashboard
```
- Setting up the application to run on boot on a schedule
```shell

```