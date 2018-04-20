/* Yaml demo.
 *
 * Mainly to demonstrate the way to handle nested yaml type, in this case is the "tasks".
 * "tasks" here is an array of different types of element, so we define type for each
 * element(i.e. "Task1", "Task2",...).
 * Then in the "Config" struct definition, we define "Tasks" as []interface{}, so that
 * ensure it could unmarshal.
 * If we want to restore each task, we need to marshal it then unmarshal.
 */

package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

var fileContent string = `
hosts: all
gather_facts: no
remote_user: ubuntu
name: install latest nginx
tasks:
  - name: install the nginx key
    apt_key:
      url: http://nginx.org/keys/nginx_signing.key 
      state: present
    become: yes

  - name:  install aws cli
    command: pip3 install awscli
    become: yes`

type Task1 struct {
	Name   string `yaml:"name"`
	AptKey struct {
		Url   string `yaml:"url"`
		State string `yaml:"state"`
	} `yaml:"apt_key"`
}

type Task2 struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
	Become  string `yaml:"become"`
}

type Config struct {
	Hosts       string        `yaml:"hosts"`
	GatherFacts string        `yaml:"gather_facts"`
	RemoteUser  string        `yaml:"remote_user"`
	Name        string        `yaml:"name"`
	Tasks       []interface{} `yaml:"tasks"`
}

func main() {
	// Unmarshal the whole yaml file, but the Task(s) are in form of interface{}
	config := new(Config)
	err := yaml.Unmarshal([]byte(fileContent), config)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(config)

	// retrive Task1 via marshal-unmarshal approach
	out, err := yaml.Marshal(config.Tasks[0])
	if err != nil {
		log.Fatal(err)
	}
	task1 := new(Task1)
	err = yaml.Unmarshal(out, task1)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(task1)
}
