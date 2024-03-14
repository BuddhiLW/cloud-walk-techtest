package report

import (
	"log"
	"strconv"

	g "github.com/BuddhiLW/cloud-walk-techtest/game"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var ReportCmd = &Z.Cmd{
	Name:     `report`,
	Aliases:  []string{`data`},
	Usage:    `<number>`,
	Summary:  `*data*, or *report*, is the root of the quake-report binary-tree.`,
	NumArgs:  0,
	Commands: []*Z.Cmd{help.Cmd, RankCmd, StatisticsCmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		log.Print("JSON data for Match:")
		if len(args) == 0 {
			log.Print("Displaying all matches, in JSON format")
			g.GameOutput()
			return nil
		} else {
			log.Print("Match chosen: ", args[0])

			// Parse as int
			n, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}

			if n >= 1 {
				_ = g.GameOutput(n)
			} else {
				log.Fatal("Invalid Match number. Must be greater or equal to 1.")
			}
			return nil
		}
	},
}

var StatisticsCmd = &Z.Cmd{
	Name:     `statistics`,
	Commands: []*Z.Cmd{help.Cmd},
	Aliases:  []string{`stats`, `bydeath`},
	Summary:  `display statistics of a **match** by **death type**`,
	Call: func(x *Z.Cmd, args ...string) error {
		if len(args) == 0 {
			return x.UsageError()
		}
		// always use "log" and not "fmt" for errors and debugging
		log.Printf("Stats of match << %v >> (by death type)", args[0])

		// Parse as int
		n, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		if n >= 1 {
			_ = g.GameOutputStatistic(n)
		} else {
			log.Fatal("Invalid Match number. Must be greater or equal to 1.")
		}
		return nil
	},
}
