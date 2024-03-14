package report

import (
	"log"
	"strconv"

	g "github.com/BuddhiLW/cloud-walk-techtest/game"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var ReportCmd = &Z.Cmd{
	Name:     `qrep`,
	Aliases:  []string{`data`, `json`},
	Usage:    `qrep <number>`,
	Summary:  `*qrep* is a quake-report binary. It's goal is to output JSON data for the chosen match, arg should be a <number: integer>`,
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

var RankCmd = &Z.Cmd{
	Name:     `rank`,
	Commands: []*Z.Cmd{help.Cmd, JsonRankCmd, PrettyRankCmd},
	Aliases:  []string{`r`},
	Usage:    `<n: number>, or try: <rank help>`,
	Call: func(x *Z.Cmd, args ...string) error {
		if len(args) == 0 {
			return x.UsageError()
		}

		n, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Rank of match number << %v >>", args[0])
		log.Print("--------- (by kills) -----------")
		g.GameRanking("text", n)
		return nil
	},
}

var JsonRankCmd = &Z.Cmd{
	Name:     `text`,
	Commands: []*Z.Cmd{help.Cmd},
	Aliases:  []string{`j`, `json`},
	Usage:    `<n: number>, or try: <rank help>`,
	Summary:  `display json format of the ranking`,
	Call: func(x *Z.Cmd, args ...string) error {
		if len(args) == 0 {
			return x.UsageError()
		}

		n, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		g.GameRanking("json", n)
		return nil
	},
}

var PrettyRankCmd = &Z.Cmd{
	Name:     `pretty`,
	Commands: []*Z.Cmd{help.Cmd},
	Aliases:  []string{`pp`, `p`, `text`},
	Summary:  `display pretty print of the ranking (default)`,
	// Usage:    `<n: number>, or try: <rank help>`,
	Call: func(x *Z.Cmd, args ...string) error {
		if len(args) == 0 {
			return x.UsageError()
		}

		n, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		g.GameRanking("json", n)
		return nil
	},
}

var StatisticsCmd = &Z.Cmd{
	Name:     `statistics`,
	Commands: []*Z.Cmd{help.Cmd},
	Aliases:  []string{`stats`},
	Call: func(x *Z.Cmd, args ...string) error {
		if len(args) == 0 {
			return x.UsageError()
		}
		// always use "log" and not "fmt" for errors and debugging
		log.Printf("Stats of match << %v >> (by death type)", args[0])
		return nil
	},
}
