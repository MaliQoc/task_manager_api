package models

import "errors"

type Task struct {
	TaskID     string `json:"task_id"`
	TaskName   string `json:"task_name"`
	IsFinished bool   `json:"is_finished"`
}

var Tasks = []Task{
	{TaskID: "1", TaskName: "Install Operating System", IsFinished: false},
	{TaskID: "2", TaskName: "Update Software Packages", IsFinished: false},
	{TaskID: "3", TaskName: "Backup Important Files", IsFinished: false},
	{TaskID: "4", TaskName: "Run Virus Scan", IsFinished: false},
	{TaskID: "5", TaskName: "Configure Network Settings", IsFinished: false},
	{TaskID: "6", TaskName: "Set Up Firewall Rules", IsFinished: false},
	{TaskID: "7", TaskName: "Optimize System Performance", IsFinished: false},
	{TaskID: "8", TaskName: "Monitor System Logs", IsFinished: false},
}

func FindTaskByID(id string) (*Task, error) {
	for i, t := range Tasks {
		if t.TaskID == id {
			return &Tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}
