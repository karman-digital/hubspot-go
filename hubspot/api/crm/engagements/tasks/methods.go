package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"

	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	associationsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/associations"
	taskmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/tasks"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (t *TasksService) CreateTaskWithAssociations(taskBody taskmodels.TaskPostBody, associations ...associationsmodels.ObjectCreationAssociation) (crmmodels.Result, error) {
	var tasksResp crmmodels.Result
	for _, association := range associations {
		if taskBody.Associations == nil {
			taskBody.Associations = []associationsmodels.ObjectCreationAssociation{}
		}
		taskBody.Associations = append(taskBody.Associations, association)
	}
	reqBody, err := json.Marshal(taskBody)
	if err != nil {
		return tasksResp, err
	}
	resp, err := t.SendRequest(http.MethodPost, "/crm/v3/objects/tasks", reqBody)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)

}
