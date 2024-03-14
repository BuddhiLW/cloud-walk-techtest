package report

import (
	"log"
	"strconv"

	g "github.com/BuddhiLW/cloud-walk-techtest/game"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

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
	Name:     `json`,
	Commands: []*Z.Cmd{help.Cmd},
	Aliases:  []string{`j`},
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
