package cmd

import (
	"fmt"
	"sort"
	"time"
	"github.com/urfave/cli"
	"clynapses/config"
	"clynapses/parser"
	. "github.com/logrusorgru/aurora"
)

func Action(ctx *cli.Context) error {
	resp, err := parser.Parse(config.GetIcsPath())
	if err != nil {
			return err
	}
	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Start.Before(*resp[j].Start)
	})
	switch {
	case ctx.Bool("day"):
		resp = parseDay(resp)
		if len(resp) == 0 {
			displayHeader("Aujourd'hui")
			fmt.Println(Red("Aucun cours\n"))
			break
		}
		displayHeader("Aujourd'hui")
		for _, event := range resp {
			displayEventToday(event)
		}
	case ctx.Bool("week"):
		resp = parseWeek(resp)
		if len(resp) == 0 {
			displayHeader("A venir cette semaine")
			fmt.Println(Red("Aucun cours\n"))
			break
		}
		last := time.Now().Local()
		displayHeader("A venir cette semaine")
		for _, event := range resp {
			if !sameDay(last, event.Start.Local()) {
				last = event.Start.Local()
				displayHeaderDate(last.Format("Monday 02/01/2006"))
			}
			displayEventToday(event)
		}
	case ctx.Bool("nextweek"):
		resp = parseNextWeek(resp)
		if len(resp) == 0 {
			displayHeader("La semaine prochaine")
			fmt.Println(Red("Aucun cours\n"))
			break
		}
		displayHeader("La semaine prochaine")
		last := time.Now().Local()
		for _, event := range resp {
			if !sameDay(last, event.Start.Local()) {
				last = event.Start.Local()
				displayHeaderDate(last.Format("Monday 02/01/2006"))
			}
			displayEvent(event)
		}
	case ctx.Bool("tomorrow"):
		resp = parseTomorrow(resp)
		if len(resp) == 0 {
			displayHeader("Demain")
			fmt.Println(Red("Aucun cours\n"))
			break
		}
		displayHeader("Demain")
		for _, event := range resp {
			displayEvent(event)
		}
	case ctx.Bool("month"):
		resp = parseMonth(resp)
		if len(resp) == 0 {
			displayHeader("Ce mois-ci")
			fmt.Println(Red("Aucun cours\n"))
			break
		}
		last := time.Now().Local()
		displayHeader("Ce mois-ci")
		for _, event := range resp {
			if !sameDay(last, event.Start.Local()) {
				last = event.Start.Local()
				displayHeaderDate(last.Format("Monday 02/01/2006"))
			}
			displayEvent(event)
		}
	case ctx.Bool("nextmonth"):
		resp = parseNextMonth(resp)
		if len(resp) == 0 {
			displayHeader("Le mois prochain")
			fmt.Println(Red("Aucun cours\n"))
			break
		}
		last := time.Now().Local()
		displayHeader("Le mois prochain")
		for _, event := range resp {
			if !sameDay(last, event.Start.Local()) {
				last = event.Start.Local()
				displayHeaderDate(last.Format("Monday 02/01/2006"))
			}
			displayEvent(event)
		}
	case ctx.Bool("all"):
		displayHeader("Tous les cours")
		if len(resp) == 0 {
			displayHeader("Tous les cours")
			fmt.Println(Red("Aucun cours\n"))
			break
		}
		last := time.Now().Local()
		for _, event := range resp {
			if !sameDay(last, event.Start.Local()) {
				last = event.Start.Local()
				displayHeaderDate(last.Format("Monday 02/01/2006"))
			}
			displayEvent(event)
		}
	default:
		resp = parseDay(resp)
		if len(resp) == 0 {
			displayHeader("Aujourd'hui")
			fmt.Println(Red("Aucun cours\n"))
			break
		}
		displayHeader("Aujourd'hui")
		for _, event := range resp {
			displayEventToday(event)
		}
	}
	return nil
}
