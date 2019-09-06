package remote_call

import (
	"io"
	"os"
	"time"
)

type Options struct {
	stdout         io.Writer
	stderr         io.Writer
	privateKeyFile string
	port           int
	user           string
	timeout        time.Duration
}

func newDefaultOptions() *Options {
	return &Options{
		stdout:         os.Stdout,
		stderr:         os.Stderr,
		privateKeyFile: "/root/.ssh/id_rsa",
		user:           "root",
		timeout:        3 * time.Second,
	}
}

func Stdout(w io.Writer) Option {
	return func(opts *Options) {
		opts.stdout = w
	}
}

func Stderr(w io.Writer) Option {
	return func(opts *Options) {
		opts.stderr = w
	}
}

func PrivateKeyFile(f string) Option {
	return func(opts *Options) {
		opts.privateKeyFile = f
	}
}

func User(u string) Option {
	return func(opts *Options) {
		opts.user = u
	}
}

func Timeout(d time.Duration) Option {
	return func(opts *Options) {
		opts.timeout = d
	}
}
