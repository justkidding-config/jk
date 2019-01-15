package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/jkcfg/jk/pkg/deferred"
	"github.com/jkcfg/jk/pkg/resolve"
	"github.com/jkcfg/jk/pkg/std"

	v8 "github.com/ry/v8worker2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute a jk program",
	Args:  runArgs,
	Run:   run,
}

type paramKind int

const (
	paramKindFromFile paramKind = iota
	paramKindBoolean
	paramKindNumber
	paramKindString
	paramKindObject
)

type paramsOption struct {
	params *std.Params
	kind   paramKind
}

func (p *paramsOption) String() string {
	return ""
}

func (p *paramsOption) setFromFile(s string) error {
	f, err := os.Open(s)
	if err != nil {
		return err
	}

	params, err := std.NewParamsFromJSON(f)
	if err != nil {
		return fmt.Errorf("%s isn't valid JSON: %v", s, err)
	}

	p.params.Merge(params)

	f.Close()
	return nil
}

func (p *paramsOption) setFromCommandline(s string) error {
	parts := strings.Split(s, "=")
	if len(parts) != 2 {
		return errors.New("input parameters are of the form name=value")
	}
	path := parts[0]
	v := parts[1]

	switch p.kind {
	case paramKindBoolean:
		b, err := strconv.ParseBool(v)
		if err != nil {
			return fmt.Errorf("could not parse '%s' as a boolean", v)
		}
		p.params.SetBool(path, b)
	case paramKindNumber:
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return fmt.Errorf("could not parse '%s' as a float64", v)
		}
		p.params.SetNumber(path, n)
	case paramKindString:
		p.params.SetString(path, v)
	case paramKindObject:
		o, err := std.NewParamsFromJSON(strings.NewReader(v))
		if err != nil {
			return fmt.Errorf("could not parse JSON '%s': %v", v, err)
		}
		params := std.NewParams()
		params.SetObject(path, o)
		p.params.Merge(params)
	}

	return nil
}

func (p *paramsOption) Set(s string) error {
	if p.kind == paramKindFromFile {
		return p.setFromFile(s)
	}
	return p.setFromCommandline(s)
}

func (p *paramsOption) Type() string {
	if p.kind == paramKindFromFile {
		return "filename"
	}
	return "name=value"
}

var runOptions struct {
	verbose         bool
	outputDirectory string
	parameters      std.Params
}

func parameters(kind paramKind) pflag.Value {
	return &paramsOption{
		params: &runOptions.parameters,
		kind:   kind,
	}
}

func init() {
	runOptions.parameters = std.NewParams()
	runCmd.PersistentFlags().BoolVarP(&runOptions.verbose, "verbose", "v", false, "verbose output")
	runCmd.PersistentFlags().StringVarP(&runOptions.outputDirectory, "output-directory", "o", "", "where to output generated files")
	runCmd.PersistentFlags().VarP(parameters(paramKindFromFile), "parameters", "p", "load parameters from a JSON file")
	runCmd.PersistentFlags().Var(parameters(paramKindBoolean), "pb", "boolean input parameter")
	runCmd.PersistentFlags().Var(parameters(paramKindNumber), "pn", "number input parameter")
	runCmd.PersistentFlags().Var(parameters(paramKindString), "ps", "string input parameter")
	runCmd.PersistentFlags().Var(parameters(paramKindObject), "po", "object input parameter")
	jk.AddCommand(runCmd)
}

func runArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("run requires an input script")
	}
	return nil
}

type exec struct {
	worker *v8.Worker
}

func (e *exec) onMessageReceived(msg []byte) []byte {
	return std.Execute(msg, e.worker, std.ExecuteOptions{
		Verbose:         runOptions.verbose,
		Parameters:      runOptions.parameters,
		OutputDirectory: runOptions.outputDirectory,
	})
}

func run(cmd *cobra.Command, args []string) {
	engine := &exec{}
	worker := v8.New(engine.onMessageReceived)
	engine.worker = worker
	filename := args[0]
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	resolver := resolve.NewResolver(worker, path.Dir(filename),
		&resolve.StaticImporter{Specifier: "std", Source: std.Module()},
		&resolve.StaticImporter{Specifier: "@jkcfg/std", Source: std.Module()},
		&resolve.FileImporter{},
		&resolve.NodeModulesImporter{},
	)
	if err := worker.LoadModule(path.Base(filename), string(input), resolver.ResolveModule); err != nil {
		log.Fatal(err)
	}
	deferred.Wait() // TODO(michael): hide this in std?
}
