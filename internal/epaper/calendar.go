package epaper

import (
	"fmt"
	"math"
	"time"
)

const (
	maxCols    = float64(7)
	daysOfWeek = "MTWTFSS"
)

// DrawCalendar Draws aa google object
func (p *Print) DrawCalendar(x float64, y float64, date time.Time) (float64, float64) {
	days := daysInMonth(date.Month(), date.Year())
	highlightDay := date.Day()

	// Write days of the week
	col := float64(0)
	row := float64(0)
	// Print days of the week
	for i := 0; i < 7; i++ {
		x1 := x + (30 * col)
		y1 := y + (30 * row)
		p.writeCalBox(x1, y1, string(daysOfWeek[i]), true)
		col++
	}

	// We want five weeks
	firstDay := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	started := false
	monthDay := 1
	for week := float64(1); week <= 5; week++ {
		y1 := y + (30 * week)
		// Iterate days i week
		for day := float64(0); day <= 6; day++ {
			x1 := x + (30 * day)
			if !started {
				fmt.Printf("%d %d\n", dayOfWeek(firstDay), day)
				if dayOfWeek(firstDay) == day {
					started = true
				}
			}
			if started {
				if monthDay == highlightDay {
					p.writeCalBox(x1, y1, fmt.Sprintf("%d", monthDay), true)
				} else {
					p.writeCalBox(x1, y1, fmt.Sprintf("%d", monthDay), false)
				}
				monthDay++
			}
			if monthDay > days {
				break
			}
		}

	}
	return math.Mod(float64(days), maxCols) * 30, 30 * maxCols
}

func (p *Print) writeCalBox(x float64, y float64, value string, fill bool) {
	p.ctx.SetRGB(0, 0, 0) // Black
	p.ctx.DrawRectangle(x, y, 30, 30)
	if fill {
		p.ctx.Fill()
		p.ctx.SetRGB(1, 1, 1) // White
	} else {
		p.ctx.Stroke()
	}
	p.ctx.DrawString(value, x+9, y+20)
}

// Gets day of week starting at 0 for Monday
func dayOfWeek(date time.Time) float64 {
	day := int(date.Weekday())
	if day == 0 {
		day = 7
	}
	return float64(day - 1)
}

// daysInMonth returns the number of days in a month
func daysInMonth(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
