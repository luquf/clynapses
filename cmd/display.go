package cmd

import (
	"github.com/apognu/gocal"
	"fmt"
	. "github.com/logrusorgru/aurora"
	"strings"
)

func displayEventToday(event gocal.Event) {
	d := strings.Split(event.Description, `\n`)
	fmt.Printf("Cours:\t%s\nDescription:\t%s\nIntervenant(s):\t%s\nJour:\t%s\nDébut:\t%s\nFin:\t%s\nSalle:\t%s\n\n",
		Brown(event.Summary),
		Brown(d[0]+" - "+d[1]+" - "+d[2]),
		Red(strings.Split(d[len(d)-2], " : ")[1]),
		Cyan(event.Start.Format("Monday 02/01/2006")),
		Cyan(event.Start.Local().Format("15:04")),
		Cyan(event.End.Local().Format("15:04")),
		Magenta(event.Location))
}

func displayEvent(event gocal.Event) {
	d := strings.Split(event.Description, `\n`)
	fmt.Printf("Cours:\t%s\nDescription:\t%s\nIntervenant(s):\t%s\nJour:\t%s\nDébut:\t%s\nFin:\t%s\nSalle:\t%s\n\n",
		Brown(event.Summary),
		Brown(d[0]+" - "+d[1]+" - "+d[2]),
		Red(strings.Split(d[len(d)-2], " : ")[1]),
		Cyan(event.Start.Format("Monday 02/01/2006")),
		Cyan(event.Start.Local().Format("15:04")),
		Cyan(event.End.Local().Format("15:04")),
		Magenta(event.Location))
}

func displayHeader(title string) {
	fmt.Printf("\n--------------- %s ---------------\n\n", title)
}

func displayHeaderDate(title string) {
	fmt.Printf("--------------- %s ---------------\n\n", title)
}



































































































