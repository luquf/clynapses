package parser

import (
	"net/http"
	"errors"
	"github.com/apognu/gocal"

)

var (
	hostUnreachableError = errors.New("Host unreachable")
	calendarParsingError = errors.New("Error parsing ICS calendar")
)


func Parse(target string) ([]gocal.Event, error) {
	resp, err := http.Get(target)
	if err != nil {
		return []gocal.Event{}, hostUnreachableError
	}
	defer resp.Body.Close()

	c := gocal.NewParser(resp.Body)
	err = c.Parse()
	if err != nil {
		return []gocal.Event{}, calendarParsingError
	}

	return c.Events, nil
}
















