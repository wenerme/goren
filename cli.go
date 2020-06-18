package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"strings"
)

func NewGorenApp() *cli.App {
	app := cli.NewApp()
	app.Name = "goren"
	app.Version = getVersion()
	app.UsageText = "goren [options]"

	// step: the standard usage message isn't that helpful
	app.OnUsageError = func(context *cli.Context, err error, isSubcommand bool) error {
		fmt.Fprintf(os.Stderr, "[error] invalid options, %s\n", err)
		return err
	}

	app.Action = func(cx *cli.Context) error {
		render := NewRender()
		f := bufio.NewWriter(os.Stdout)
		defer f.Flush()

		proc := &Process{in: os.Stdin, out: f}
		err := pipe(render, proc)
		return err
	}

	return app
}

func pipe(render *Render, proc *Process) error {

	buf := new(strings.Builder)
	_, err := io.Copy(buf, proc.in)
	if err != nil {
		return err
	}
	proc.inputString = buf.String()

	{
		tpl, err := render.tpl.Parse(proc.inputString)
		if err != nil {
			return err
		}
		proc.tpl = tpl
		proc.data = render.data
	}

	{
		buf := new(strings.Builder)
		err := proc.tpl.Execute(buf, proc.data)
		if err != nil {
			return err
		}
		proc.outputString = buf.String()
	}

	{
		_, err := io.Copy(proc.out, strings.NewReader(proc.outputString))
		if err != nil {
			return err
		}
	}

	return nil
}
