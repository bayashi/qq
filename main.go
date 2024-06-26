package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/adrg/xdg"
	"github.com/bayashi/colorpalette"
	"github.com/mattn/go-isatty"
	yaml "gopkg.in/yaml.v3"

	"github.com/bayashi/qq/dictionary"

	"github.com/bayashi/qq/resources/grpc"
	"github.com/bayashi/qq/resources/http"
	"github.com/bayashi/qq/resources/jsonrpc"
	"github.com/bayashi/qq/resources/mimetype"
)

type ExitCode int

const (
	EXIT_OK  ExitCode = 0
	EXIT_ERR ExitCode = 1
)

var funcExit = func(code ExitCode) {
	os.Exit(int(code))
}

type runner struct {
	isTTY   bool
	out     io.Writer
	err     io.Writer
	args    []string
	noColor bool // env:NO_COLOR
}

func main() {
	cli := &runner{
		isTTY:   isTTY(),
		out:     os.Stdout,
		err:     os.Stderr,
		args:    os.Args[1:],
		noColor: os.Getenv("NO_COLOR") != "",
	}
	exitCode, message := cli.run()

	if exitCode != EXIT_OK {
		cli.putErr(message)
	}
	funcExit(exitCode)
}

func (cli *runner) run() (ExitCode, string) {
	subcmd, arg := cli.parseArgs()
	err := cli.handle(subcmd, arg)
	if err != nil {
		return EXIT_ERR, fmt.Sprintf("Err: %s", err.Error())
	}

	return EXIT_OK, ""
}

func isTTY() bool {
	fd := os.Stdout.Fd()
	return isatty.IsTerminal(fd) || isatty.IsCygwinTerminal(fd)
}

func (cli *runner) handle(kind string, target string) error {
	var iMap map[int]dictionary.Element
	var sMap map[string]dictionary.Element
	typ := "int"
	asc := true
	switch kind {
	case "grpc":
		iMap = grpc.GetRef()
	case "http", "https":
		iMap = http.GetRef()
	case "jsonrpc":
		iMap = jsonrpc.GetRef()
		asc = false // desc
	case "mimetype", "mime", "contenttype", "ct":
		sMap = mimetype.GetRef()
		typ = "string"
	default:
		var err error
		typ, err = cli.loadCustomDictionary(kind, &iMap, &sMap)
		if err != nil {
			return err
		}
	}

	if target == "" {
		if typ == "int" {
			return cli.allIntMap(iMap, asc)
		} else {
			return cli.allStringMap(sMap)
		}
	}

	if typ == "int" {
		return cli.searchInt(iMap, target, asc)
	} else {
		return cli.searchString(sMap, target)
	}
}

func (cli *runner) allIntMap(m map[int]dictionary.Element, asc bool) error {
	keys := sortedIntKeys(m, asc)
	for _, k := range keys {
		dic := m[k]
		fmt.Fprintf(cli.out, "%2d %s\n", k, dic.Subject)
	}

	return nil
}

func (cli *runner) allStringMap(m map[string]dictionary.Element) error {
	keys := sortedStringKeys(m)
	for _, k := range keys {
		dic := m[k]
		fmt.Fprintf(cli.out, "%s %s\n", k, dic.Subject)
	}

	return nil
}

var HIGHLIGHTER = colorpalette.Get("bg_red")

func (cli *runner) replacer(re *regexp.Regexp, src string) string {
	if !cli.isTTY || cli.noColor {
		return src
	}
	return string(re.ReplaceAll([]byte(src), []byte(HIGHLIGHTER.Sprint("$1"))))
}

func (cli *runner) searchInt(m map[int]dictionary.Element, target string, asc bool) error {
	found := false
	targeti, err := strconv.Atoi(target)
	if err != nil {
		var re *regexp.Regexp
		keys := sortedIntKeys(m, asc)
		for _, k := range keys {
			dic := m[k]
			if strings.Contains(strings.ToLower(dic.Subject), strings.ToLower(target)) {
				if re == nil && cli.isTTY {
					re = regexp.MustCompile(fmt.Sprintf("(?i)(%s)", regexp.QuoteMeta(target)))
				}
				fmt.Fprintf(cli.out, "%2d %s\n", k, cli.replacer(re, dic.Subject))
				found = true
			}
		}
	} else if dic, isExists := m[targeti]; isExists {
		fmt.Fprintf(cli.out, "%d %s\n%s\n", targeti, dic.Subject, dic.Description)
		found = true
	} else {
		keys := sortedIntKeys(m, asc)
		for _, k := range keys {
			dic := m[k]
			kk := fmt.Sprintf("%d", k)
			if strings.Contains(kk, target) {
				kk = strings.ReplaceAll(kk, target, HIGHLIGHTER.Sprint(target))
				fmt.Fprintf(cli.out, "%2s %s\n", kk, dic.Subject)
				found = true
			}
		}
	}

	if found {
		return nil
	}

	return fmt.Errorf("not found `%s`", target)
}

func (cli *runner) searchString(m map[string]dictionary.Element, target string) error {
	found := false
	if dic, isExists := m[target]; isExists {
		fmt.Fprintf(cli.out, "%s %s\n%s\n", target, dic.Subject, dic.Description)
		found = true
	} else {
		var re *regexp.Regexp
		keys := sortedStringKeys(m)
		for _, k := range keys {
			dic := m[k]
			if strings.Contains(strings.ToLower(k), strings.ToLower(target)) ||
				strings.Contains(strings.ToLower(dic.Subject), strings.ToLower(target)) {
				if re == nil && cli.isTTY {
					re = regexp.MustCompile(fmt.Sprintf("(?i)(%s)", regexp.QuoteMeta(target)))
				}
				fmt.Fprintf(cli.out, "%s %s\n", cli.replacer(re, k), cli.replacer(re, dic.Subject))
				found = true
			}
		}
	}

	if found {
		return nil
	}

	return fmt.Errorf("not found `%s`", target)
}

func sortedIntKeys(m map[int]dictionary.Element, asc bool) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	if asc {
		sort.Ints(keys)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	}

	return keys
}

func sortedStringKeys(m map[string]dictionary.Element) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

var customDicPathFuc = func(kind string) string {
	return filepath.Join(xdg.ConfigHome, cmdName, fmt.Sprintf("%s.yaml", kind))
}

func (cli *runner) loadCustomDictionary(
	kind string,
	iMap *map[int]dictionary.Element,
	sMap *map[string]dictionary.Element,
) (string, error) {
	customDicPath := customDicPathFuc(kind)
	if _, err := os.Stat(customDicPath); err != nil {
		return "", fmt.Errorf("wrong sub command `%s`. Perhaps, you have custom YAML? It should be located on `%s`", kind, customDicPath)
	}

	bytes, err := os.ReadFile(customDicPath)
	if err != nil {
		return "", err
	}

	typ := "int"
	if bytes[2] == 0x69 { // 0x69:"i"
		err = yaml.Unmarshal(bytes, iMap)
	} else {
		err = yaml.Unmarshal(bytes, sMap)
		typ = "string"
	}
	if err != nil {
		return "", err
	}

	return typ, nil
}
