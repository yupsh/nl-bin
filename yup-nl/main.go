package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/nl"
)

const (
	flagBodyNumbering   = "body-numbering"
	flagHeaderNumbering = "header-numbering"
	flagFooterNumbering = "footer-numbering"
	flagNumberSeparator = "number-separator"
	flagNumberFormat    = "number-format"
	flagStartNumber     = "starting-line-number"
	flagIncrement       = "line-increment"
	flagNoRenumber      = "no-renumber"
)

func main() {
	app := &cli.App{
		Name:  "nl",
		Usage: "number lines of files",
		UsageText: `nl [OPTIONS] [FILE...]

   Write each FILE to standard output, with line numbers added.
   With no FILE, or when FILE is -, read standard input.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagBodyNumbering,
				Aliases: []string{"b"},
				Usage:   "use STYLE for numbering body lines (a=all, t=non-empty, n=none)",
				Value:   "t",
			},
		&cli.StringFlag{
			Name:  flagHeaderNumbering,
			Usage: "use STYLE for numbering header lines",
			Value: "n",
		},
			&cli.StringFlag{
				Name:    flagFooterNumbering,
				Aliases: []string{"f"},
				Usage:   "use STYLE for numbering footer lines",
				Value:   "n",
			},
			&cli.StringFlag{
				Name:    flagNumberSeparator,
				Aliases: []string{"s"},
				Usage:   "add STRING after (possible) line number",
				Value:   "\t",
			},
			&cli.StringFlag{
				Name:    flagNumberFormat,
				Aliases: []string{"n"},
				Usage:   "insert line numbers according to FORMAT (ln, rn, rz)",
				Value:   "rn",
			},
			&cli.IntFlag{
				Name:    flagStartNumber,
				Aliases: []string{"v"},
				Usage:   "first line number for each section",
				Value:   1,
			},
			&cli.IntFlag{
				Name:    flagIncrement,
				Aliases: []string{"i"},
				Usage:   "line number increment at each line",
				Value:   1,
			},
			&cli.BoolFlag{
				Name:    flagNoRenumber,
				Aliases: []string{"p"},
				Usage:   "do not reset line numbers at logical pages",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "nl: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments (or none for stdin)
	for i := 0; i < c.NArg(); i++ {
		params = append(params, yup.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.IsSet(flagBodyNumbering) {
		params = append(params, BodyNumbering(c.String(flagBodyNumbering)))
	}
	if c.IsSet(flagHeaderNumbering) {
		params = append(params, HeaderNumbering(c.String(flagHeaderNumbering)))
	}
	if c.IsSet(flagFooterNumbering) {
		params = append(params, FooterNumbering(c.String(flagFooterNumbering)))
	}
	if c.IsSet(flagNumberSeparator) {
		params = append(params, NumberSeparator(c.String(flagNumberSeparator)))
	}
	if c.IsSet(flagNumberFormat) {
		params = append(params, NumberFormat(c.String(flagNumberFormat)))
	}
	if c.IsSet(flagStartNumber) {
		params = append(params, StartNumber(c.Int(flagStartNumber)))
	}
	if c.IsSet(flagIncrement) {
		params = append(params, Increment(c.Int(flagIncrement)))
	}
	if c.Bool(flagNoRenumber) {
		params = append(params, NoRenumber)
	}

	// Create and execute the nl command
	cmd := Nl(params...)
	return yup.Run(cmd)
}
