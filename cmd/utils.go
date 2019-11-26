package cmd

import (
	"time"
	"github.com/apognu/gocal"
)

func dateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func sameMonth(date1, date2 time.Time) bool {
	_, m1, _ := date1.Date()
	_, m2, _ := date2.Date()
	return m1 == m2
}

func nextMonth(date1, date2 time.Time) bool {
	_, m1, _ := date1.Date()
	_, m2, _ := date2.Date()
	return m1-1 == m2
}

func sameWeek(date1, date2 time.Time) bool {
	_, week1 := date1.Local().ISOWeek()
	_, week2 := date2.Local().ISOWeek()
	_, _, d1 := date1.Date()
	_, _, d2 := date2.Date()
	return week1 == week2 && d2 > d1
}

func nextWeek(date1, date2 time.Time) bool {
	_, week1 := date1.Local().ISOWeek()
	_, week2 := date2.Local().ISOWeek()
	return week2 == week1+1
}

func parseDay(events []gocal.Event) []gocal.Event {
	today := time.Now().Local()
	var todayEvents []gocal.Event
	for _, e := range events {
		if dateEqual(*e.Start, today) {
			todayEvents = append(todayEvents, e)
		}
	}
	return todayEvents
}

func parseTomorrow(events []gocal.Event) []gocal.Event {
	tomorrow := time.Now().UTC().Add(24 * time.Hour)
	var tomorrowEvents []gocal.Event
	for _, e := range events {
		if dateEqual(*e.Start, tomorrow) {
			tomorrowEvents = append(tomorrowEvents, e)
		}
	}
	return tomorrowEvents
}

func parseWeek(events []gocal.Event) []gocal.Event {
	today := time.Now().Local()
	var weekEvents []gocal.Event
	for _, e := range events {
		if sameWeek(today, *e.Start) {
			weekEvents = append(weekEvents, e)
		}
	}
	return weekEvents

}

func parseNextWeek(events []gocal.Event) []gocal.Event {
	today := time.Now().Local()
	var weekEvents []gocal.Event
	for _, e := range events {
		if nextWeek(today, *e.Start) {
			weekEvents = append(weekEvents, e)
		}
	}
	return weekEvents

}

func parseMonth(events []gocal.Event) []gocal.Event {
	today := time.Now().UTC()
	var monthEvents []gocal.Event
	for _, e := range events {
		if sameMonth(*e.Start, today) {
			monthEvents = append(monthEvents, e)
		}
	}
	return monthEvents
}

func parseNextMonth(events []gocal.Event) []gocal.Event {
	today := time.Now().UTC()
	var monthEvents []gocal.Event
	for _, e := range events {
		if nextMonth(*e.Start, today) {
			monthEvents = append(monthEvents, e)
		}
	}
	return monthEvents
}

func sameDay(date1, date2 time.Time) bool {
	_, _, d1 := date1.Date()
	_, _, d2 := date2.Date()
	return d1 == d2
}
