package main

import (
	"bytes"
	"testing"

	a "github.com/bayashi/actually"
	c "github.com/fatih/color"
)

func init() {
	c.NoColor = false
}

func TestSearchIntWithColor(t *testing.T) {
	var o bytes.Buffer
	cli := &runner{
		isTTY: true,
		out:   &o,
		args:  []string{"grpc", "dead"},
	}
	ret, msg := cli.run()
	a.Got(ret).Expect(EXIT_OK).Same(t)
	a.Got(msg).Expect("").Same(t)

	expect := " 4 \x1b[41mDead\x1b[0mlineExceeded\n"
	a.Got(o.String()).Expect(expect).Same(t)
}

func TestSearchStringWithColor(t *testing.T) {
	var o bytes.Buffer
	cli := &runner{
		isTTY: true,
		out:   &o,
		args:  []string{"mimetype", "yao"},
	}
	ret, msg := cli.run()
	a.Got(ret).Expect(EXIT_OK).Same(t)
	a.Got(msg).Expect("").Same(t)

	expect := "yme application/vnd.\x1b[41myao\x1b[0mweme\n"
	a.Got(o.String()).Expect(expect).Same(t)
}
