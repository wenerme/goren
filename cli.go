package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

type CloserFunc func() error

func (f CloserFunc) Close() error {
	return f()
}

func NewGorenApp() *cli.App {
	app := cli.NewApp()
	app.Name = "goren"
	app.Version = getVersion()
	app.UsageText = "goren [options]"
	app.Commands = []*cli.Command{
		{
			Name:        "render",
			Description: "Render template",
			Flags: []cli.Flag{
				&cli.PathFlag{Name: "in", Aliases: []string{"i"}, Usage: "Input file or directory", DefaultText: "-"},
				&cli.PathFlag{Name: "out", Aliases: []string{"o"}, Usage: "Output file or directory", DefaultText: "-"},
			},
			Action: func(ctx *cli.Context) error {
				proc := &Process{
					src:  ctx.Path("in"),
					dest: ctx.Path("out"),
				}
				fmt.Printf("Render %v -> %v\n", proc.src, proc.dest)
				render := NewRender()

				err := pipe(render, proc)
				return err
			},
		},
		{
			Name:        "watch",
			Description: "Rerender when changed",
			Action: func(context *cli.Context) error {
				fmt.Println("Watch")
				return nil
			},
		},
	}

	// step: the standard usage message isn't that helpful
	app.OnUsageError = func(context *cli.Context, err error, isSubcommand bool) error {
		fmt.Fprintf(os.Stderr, "[error] invalid options, %s\n", err)
		return err
	}

	return app
}

func pipe(render *Render, proc *Process) error {
	{
		if proc.src == "-" {
			proc.in = os.Stdin
		} else {
			in, err := os.Open(proc.src)
			if err != nil {
				return err
			}
			proc.in = in
		}

		if proc.dest == "-" {
			f := bufio.NewWriter(os.Stdout)
			proc.out = f
			proc.closers = append(proc.closers, CloserFunc(f.Flush))
		} else {
			f, err := os.Open(proc.dest)
			if err != nil {
				return err
			}
			proc.out = f
		}
	}
	{
		buf := new(strings.Builder)
		_, err := io.Copy(buf, proc.in)
		if err != nil {
			return err
		}
		proc.inputString = buf.String()
	}

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
