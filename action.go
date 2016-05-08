package tasker

import (
	"github.com/kr9ly/tasker/funcs"
)

func ExecAction(action Action, config *Config) {
	task := config.GetTask(action)
	fn := funcs.GetFunc(task.GetTaskType())
	options := map[string]interface{}{}
	for k,v := range config.GetOptionSet(task.GetOptionType()) {
		options[k] = v
	}
	for k,v := range task.GetOptionSet() {
		options[k] = v
	}
	for k,v := range action.GetOptionSet() {
		options[k] = v
	}
	fn(options)
}
