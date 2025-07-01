package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/formancehq/numscript"
	"github.com/spf13/cobra"
)

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

type RunInputOpts struct {
	Script       string                     `json:"script"`
	Variables    map[string]string          `json:"variables"`
	Meta         numscript.AccountsMetadata `json:"metadata"`
	Balances     numscript.Balances         `json:"balances"`
	FeatureFlags map[string]struct{}        `json:"featureFlags"`
}

func run() {
	opt := RunInputOpts{
		Variables:    make(map[string]string),
		Meta:         make(numscript.AccountsMetadata),
		Balances:     make(numscript.Balances),
		FeatureFlags: make(map[string]struct{}),
	}

	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		panic(err)
	}

	err = json.Unmarshal(bytes, &opt)
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		panic(err)
	}

	parsedResult := numscript.Parse(opt.Script)
	errors := parsedResult.GetParsingErrors()

	if len(errors) != 0 {
		os.Stderr.Write([]byte(numscript.ParseErrorsToString(errors, opt.Script)))
		panic(fmt.Errorf("parsing errors"))
	}

	result, err := parsedResult.RunWithFeatureFlags(
		context.Background(),
		opt.Variables,
		numscript.StaticStore{
			Balances: opt.Balances,
			Meta:     opt.Meta,
		},
		opt.FeatureFlags,
	)

	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		panic(err)
	}

	out, err := json.Marshal(result)
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		panic(err)
	}

	os.Stdout.Write(out)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute a numscript",
	Long:  "Execute a numscript using the balances, the current metadata, and the variables values as input. Also accept feature flags to enable experimental features.",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var rootCmd = &cobra.Command{
	Use:     "numscript",
	Short:   "Numscript CLI",
	Long:    "Numscript CLI",
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

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.Execute()
}
