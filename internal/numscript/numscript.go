package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/Jeffail/gabs/v2"
	"github.com/spf13/cobra"

	"github.com/PagoPlus/numscriptex/internal/analysis"
	"github.com/PagoPlus/numscriptex/internal/interpreter"
	"github.com/PagoPlus/numscriptex/internal/parser"
	"github.com/PagoPlus/numscriptex/internal/utils"
)

type RunInputOpts struct {
	Script    string                       `json:"script"`
	Variables map[string]string            `json:"variables"`
	Meta      interpreter.AccountsMetadata `json:"metadata"`
	Balances  interpreter.Balances         `json:"balances"`
}

var Version string

func version() string {
	if Version == "" {
		return "dev"
	} else {
		return Version
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the app version",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(version())
	},
}

func SeverityToString(s analysis.Severity) string {
	switch s {
	case analysis.ErrorSeverity:
		return "error"
	case analysis.WarningSeverity:
		return "warning"
	case analysis.Information:
		return "info"
	case analysis.Hint:
		return "hint"
	default:
		return utils.NonExhaustiveMatchPanic[string](s)
	}
}

func check() {
	dat, err := io.ReadAll(os.Stdin)
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(1)
	}

	res := analysis.CheckSource(string(dat))
	sort.Slice(res.Diagnostics, func(i, j int) bool {
		p1 := res.Diagnostics[i].Range.Start
		p2 := res.Diagnostics[j].Range.Start

		return p2.GtEq(p1)
	})

	hasErrors := false
	jsonObj := gabs.New()
	for _, d := range res.Diagnostics {
		if d.Kind.Severity() == analysis.ErrorSeverity {
			hasErrors = true
		}

		logLevel := SeverityToString(d.Kind.Severity())

		buildCheckDetails(d, jsonObj, logLevel)
	}

	if hasErrors {
		jsonObj.Set(false, "valid")
	} else {
		jsonObj.Set(true, "valid")
	}

	fmt.Println(jsonObj.String())
}

func buildCheckDetails(diagnostic analysis.Diagnostic, jsonObj *gabs.Container, logLevel string) {
	logLevelKey := logLevel + "s"
	subJsonObj := gabs.New()

	subJsonObj.Set(diagnostic.Range.Start.Line, "line")
	subJsonObj.Set(diagnostic.Range.Start.Character, "character")
	subJsonObj.Set(logLevel, "level")
	subJsonObj.Set(diagnostic.Kind.Message(), logLevel)

	jsonObj.ArrayAppend(subJsonObj, logLevelKey)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check a numscript file",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		check()
	},
}

func run() {
	opt := RunInputOpts{
		Variables: make(map[string]string),
		Meta:      make(interpreter.AccountsMetadata),
		Balances:  make(interpreter.Balances),
	}

	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(1)
	}

	err = json.Unmarshal(bytes, &opt)
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(1)
	}

	parseResult := parser.Parse(opt.Script)
	if len(parseResult.Errors) != 0 {
		os.Stderr.Write([]byte(parser.ParseErrorsToString(parseResult.Errors, opt.Script)))
		os.Exit(1)
	}

	featureFlags := make(map[string]struct{})
	result, err := interpreter.RunProgram(
		context.Background(),
		parseResult.Value,
		opt.Variables,
		interpreter.StaticStore{
			Balances: opt.Balances,
			Meta:     opt.Meta,
		},
		featureFlags,
	)

	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(1)
	}

	out, err := json.Marshal(result)
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(1)
	}

	os.Stdout.Write(out)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Evaluate a numscript file",
	Long:  "Evaluate a numscript file, using the balances, the current metadata and the variables values as input.",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var rootCmd = &cobra.Command{
	Use:     "numscript",
	Short:   "Numscript cli",
	Long:    "Numscript cli",
	Version: version(),
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			os.Exit(1)
		}
	}()

	rootCmd.SetVersionTemplate(rootCmd.Version)

	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(versionCmd)

	rootCmd.Execute()
}
