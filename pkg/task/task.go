package task

// Task is an interface that contains the spec for a task. It has
// one method Execute() that is used to run the task.
type Task interface {
	// Execute is used to execute the task. This method takes a map of
	// strings and interfaces as a parameter. It returns an map of strings
	// and interfaces as output and nil or nil and an error if one occurred.
	Execute(input map[string]interface{}) (map[string]interface{}, error)
}
