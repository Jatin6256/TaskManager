package taskHandler

import (
	"encoding/json"
	"io/ioutil"
)

type Task struct {
	Txt      string
	Priority int
	// Position int
	Done bool
}

func SaveTasks(fileName string, tasks []Task) error {

	b, err := json.Marshal((tasks))

	if err != nil {
		return nil
	}

	err = ioutil.WriteFile(fileName, b, 0644)

	if err != nil {
		return err
	}

	// fmt.Println(string(b))
	return nil

}

func ReadFile(fileName string) ([]Task, error) {

	b, err := ioutil.ReadFile(fileName)

	if err != nil {
		return []Task{}, err
	}

	var tasks []Task

	err = json.Unmarshal(b, &tasks)

	if err != nil {
		return []Task{}, err
	}

	return tasks, nil
}

func (t *Task) SetPriority(pri int) {

	t.Priority = pri

}

func (t *Task) Prettier() string {

	if t.Priority == 1 {
		return "(1)"
	}

	if t.Priority == 3 {
		return "(3)"
	}

	return " "

}

// func (t *Task) Label() string {

// 	return strconv.Itoa(t.Position+1) + "."
// }

type ByPri []Task

func (e ByPri) Len() int {
	return len(e)
}

func (e ByPri) Less(i, j int) bool {
	if e[i].Done != e[j].Done {
		return !e[i].Done
	}
	return e[i].Priority <= e[j].Priority
}

func (e ByPri) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (t *Task) PrettyDone() string {
	if t.Done {
		return "X "
	} else {
		return ""
	}
}
