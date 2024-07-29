package engagements

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/api/crm/engagements/notes"
	"github.com/karman-digital/hubspot/hubspot/api/crm/engagements/tasks"
)

func NewEngagementService(creds *credentials.Credentials) Engagements {
	return Engagements{
		Notes: notes.NewNotesService(creds),
		Tasks: tasks.NewTasksService(creds),
	}
}
