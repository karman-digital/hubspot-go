package tasks

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewTasksService(creds *credentials.Credentials) *TasksService {
	return &TasksService{
		creds,
	}
}
