package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (t *TasksService) CreateTaskWithAssociations(taskBody hubspotmodels.TaskPostBody, associations ...hubspotmodels.ObjectCreationAssociation) (hubspotmodels.Result, error) {
	var tasksResp hubspotmodels.Result
	for _, association := range associations {
		if taskBody.Associations == nil {
			taskBody.Associations = []hubspotmodels.ObjectCreationAssociation{}
		}
		taskBody.Associations = append(taskBody.Associations, association)
	}
	reqBody, err := json.Marshal(taskBody)
	if err != nil {
		return tasksResp, err
	}
	resp, err := t.SendRequest(http.MethodPost, "/crm/v3/objects/tasks", reqBody)
	if err != nil {
		return hubspotmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)

}
