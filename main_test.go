package main

import (
	"bytes"
	"testing"

	here "github.com/MakeNowJust/heredoc/v2"
	a "github.com/bayashi/actually"
)

var (
	stubCalled bool
	stubCode   ExitCode
)

func TestAllList(t *testing.T) {
	var o bytes.Buffer
	cli := &runner{
		out:  &o,
		args: []string{"grpc"},
	}
	ret, msg := cli.run()
	a.Got(ret).Expect(EXIT_OK).Same(t)
	a.Got(msg).Expect("").Same(t)

	expect := here.Doc(`
		 0 OK
		 1 Canceled
		 2 Unknown
		 3 InvalidArgument
		 4 DeadlineExceeded
		 5 NotFound
		 6 AlreadyExists
		 7 PermissionDenied
		 8 ResourceExhausted
		 9 FailedPrecondition
		10 Aborted
		11 OutOfRange
		12 Unimplemented
		13 Internal
		14 Unavailable
		15 DataLoss
		16 Unauthenticated
	`)
	a.Got(o.String()).Expect(expect).Same(t)
}

func TestDetaili(t *testing.T) {
	var o bytes.Buffer
	cli := &runner{
		out:  &o,
		args: []string{"grpc", "4"},
	}
	ret, msg := cli.run()
	a.Got(ret).Expect(EXIT_OK).Same(t)
	a.Got(msg).Expect("").Same(t)

	expect := here.Doc(`
		4 DeadlineExceeded
		The deadline expired before the operation could complete. For operations that change the state of the system, this error may be returned even if the operation has completed successfully.  For example, a successful response from a server could have been delayed long enough for the deadline to expire.
		HTTP Mapping: 504 Gateway Timeout
	`)
	a.Got(o.String()).Expect(expect).Same(t)
}

func TestDetails(t *testing.T) {
	var o bytes.Buffer
	cli := &runner{
		out:  &o,
		args: []string{"mimetype", "json"},
	}
	ret, msg := cli.run()
	a.Got(ret).Expect(EXIT_OK).Same(t)
	a.Got(msg).Expect("").Same(t)

	expect := here.Doc(`
		json application/json

	`)
	a.Got(o.String()).Expect(expect).Same(t)
}

func TestSearch(t *testing.T) {
	var o bytes.Buffer
	cli := &runner{
		out:  &o,
		args: []string{"grpc", "dead"},
	}
	ret, msg := cli.run()
	a.Got(ret).Expect(EXIT_OK).Same(t)
	a.Got(msg).Expect("").Same(t)

	expect := " 4 DeadlineExceeded\n"
	a.Got(o.String()).Expect(expect).Same(t)
}
