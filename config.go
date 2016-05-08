package tasker

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	OptionSets OptionSets `yaml:"options"`
	NodeSets   NodeSets   `yaml:"nodes"`
	Tasks      Tasks      `yaml:"tasks"`
	Triggers   Triggers   `yaml:"triggers"`
}

type OptionSets map[string]OptionSet
type OptionSet map[string]interface{}

type NodeSets map[string]NodeSet
type NodeSet []string

type Tasks map[string]Task
type Task map[string]OptionSet

type Triggers []Trigger
type Trigger struct {
	Time  string
	Actions Actions
}

type Actions []Action
type Action map[string]OptionSet

func (t Task) GetTaskType() string {
	for k,_ := range t {
		return k
	}
	return ""
}

func (t Task) GetOptionType() string {
	return t.GetOptionSet()["option"].(string)
}

func (t Task) GetOptionSet() OptionSet {
	for _,v := range t {
		return v
	}
	return nil
}

func (a Action) GetActionType() string {
	for k,_ := range a {
		return k
	}
	return ""
}

func (a Action) GetOptionSet() OptionSet {
	for _,v := range a {
		return v
	}
	return nil
}

func (c *Config) GetOptionSet(name string) OptionSet {
	return c.OptionSets[name]
}

func (c *Config) GetTask(a Action) Task {
	for k,v := range c.Tasks {
		if k == a.GetActionType() {
			return v
		}
	}
	return nil
}

func Load(in []byte) *Config {
	c := Config{}
	err := yaml.Unmarshal(in, &c)
	if err != nil {
		panic(err)
	}
	return &c
}