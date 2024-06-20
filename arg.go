package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
)

const (
	cmdName string = "qq"
)

var (
	version     = ""
	installFrom = "Source"

	supportSubCmds = []string{
		"grpc",
		"http",
		"https",
		"jsonrpc",
		"mimetype",
		"mime", // alias
		"contenttype",
		"ct", // alias
	}
)

func (cli *runner) hasHelpOption() bool {
	for _, arg := range cli.args {
		if arg == "-h" || arg == "--help" {
			return true
		}
	}

	return false
}

func (cli *runner) hasVersionOption() bool {
	for _, arg := range cli.args {
		if arg == "-v" || arg == "--version" {
			return true
		}
	}

	return false
}

func (cli *runner) parseArgs() (string, string) {
	if len(cli.args) == 0 {
		cli.putUsage()
		funcExit(EXIT_OK)
	} else if cli.hasHelpOption() {
		cli.putErr(fmt.Sprintf("Version %s", getVersion()))
		cli.putUsage()
		funcExit(EXIT_OK)
	} else if cli.hasVersionOption() {
		cli.putErr(versionDetails())
		funcExit(EXIT_OK)
	}

	return cli.getSubCmdAndArg()
}

func (cli *runner) getSubCmdAndArg() (string, string) {
	if len(cli.args) > 1 {
		return cli.args[0], cli.args[1]
	}

	return cli.args[0], ""
}

func versionDetails() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	compiler := runtime.Version()

	return fmt.Sprintf(
		"Version %s - %s.%s (compiled:%s, %s)",
		getVersion(),
		goos,
		goarch,
		compiler,
		installFrom,
	)
}

func getVersion() string {
	if version != "" {
		return version
	}
	i, ok := debug.ReadBuildInfo()
	if !ok {
		return "Unknown"
	}

	return i.Main.Version
}

func (cli *runner) putErr(message ...interface{}) {
	fmt.Fprintln(cli.err, message...)
}

func (cli *runner) putUsage() {
	cli.putErr(fmt.Sprintf("Usage: %s {%s} [REF]", cmdName, strings.Join(supportSubCmds, "|")))
}
