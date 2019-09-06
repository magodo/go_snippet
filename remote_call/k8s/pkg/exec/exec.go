package exec

import (
	"fmt"
	"io"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"os"
)

type Executor interface {
	ExecCommands(pod, container string, commands []string) error
	ExecScripts(pod, container string, scripts []string, args ...string) error
}

func NewExecutor(config *rest.Config, opts ...Option) Executor {
	options := newDefaultOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &executor{config, options}
}

type executor struct {
	*rest.Config
	*Options
}

func (e *executor) ExecCommands(pod, container string,  commands []string) error {
	client,err := kubernetes.NewForConfig(e.Config)
	if err != nil {
		return fmt.Errorf("failed to new client from config: %v", err)
	}
	req := client.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(pod).
		Namespace(e.Namespace).
		SubResource("exec")
	if container == "" {
	// If not specified container name, pick the 1st container in this pod
		Pod, err := client.CoreV1().Pods(e.Namespace).Get(pod, metav1.GetOptions{})
		if err != nil {
			return fmt.Errorf("failed to get pods: %v", err)
		}
		container = Pod.Spec.Containers[0].Name
	}
	req = req.Param("container", container)

	req.VersionedParams(&v1.PodExecOptions{
		Container: container,
		Command:   commands,
		Stdin:     e.Stdin != nil,
		Stdout:    e.Stdout != nil,
		Stderr:    e.Stderr != nil,
		TTY:       false,
	}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(e.Config, "POST", req.URL())
	if err != nil {
		return fmt.Errorf("failed to execute remote command: %v", err)
	}

	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  e.Stdin,
		Stdout: e.Stdout,
		Stderr: e.Stdout,
		Tty:    false,
	})
	return err
}

func (e *executor) ExecScripts(pod, container string,  scripts []string, args ...string) error {
	stdinR, stdinW := io.Pipe()
	defer stdinR.Close()
	e.Stdin = stdinR

	var openedScripts []io.Reader
	for _, script := range scripts {
		f, err := os.Open(script)
		if err != nil {
			for _, of := range openedScripts {
				if err := of.(io.Closer).Close(); err != nil {
					log.Printf("faield to close file: %v", err)
				}
			}
			return fmt.Errorf("failed to open script %s: %v", script, err)
		}
		openedScripts = append(openedScripts, f)
	}
	defer func() {
		for _, of := range openedScripts {
			if err := of.(io.Closer).Close(); err != nil {
				log.Printf("faield to close file: %v", err)
			}
		}
	}()

	mergedScript := io.MultiReader(openedScripts...)
	var errCmd error
	go func() {
		defer func() {
			defer stdinW.Close()
			if errCmd == nil {
				lastCmd := "main"
				for _, arg := range args {
					lastCmd += " \"" + arg + "\""
				}
				_, errCmd = stdinW.Write([]byte(lastCmd))
				if errCmd != nil {
					return
				}
			}
		}()
		firstCmd := "export REMOTE_MODE=on"
		_, errCmd = stdinW.Write([]byte(firstCmd))
		if errCmd != nil {
			return
		}
		_, errCmd = io.Copy(stdinW, mergedScript)
	}()

	return e.ExecCommands(pod, container, []string{"/bin/bash", "-c", `cat | /bin/bash`})
}
