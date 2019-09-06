package exec

import (
	"io"
	apiv1 "k8s.io/api/core/v1"
	"os"
)

type Options struct {
	Namespace     string
	Stdin         io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func newDefaultOptions() *Options {
	return &Options{
		Namespace: apiv1.NamespaceDefault,
		Stdin: 		os.Stdin,
		Stdout: 	os.Stdout,
		Stderr: 	os.Stderr,
	}
}

type Option func(*Options)

func Namespace(ns string) Option {
	return func(opts *Options) {
		opts.Namespace = ns
	}
}

func Stdin(r io.Reader) Option {
	return func(opts *Options) {
		opts.Stdin = r
	}
}

func Stdout(w io.Writer) Option {
	return func(opts *Options) {
		opts.Stdout = w
	}
}

func Stderr(w io.Writer) Option {
	return func(opts *Options) {
		opts.Stderr = w
	}
}
