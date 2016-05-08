package funcs

type TaskFunc func(option map[string]interface{})

func GetFunc(taskType string) TaskFunc {
	switch taskType {
	case "ssh":
		return ssh
	}
	return nil
}
