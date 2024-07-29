package hubspotmodels

type TaskPostBody struct {
	Properties   TaskProperties              `json:"properties"`
	Associations []ObjectCreationAssociation `json:"associations,omitempty"`
}

type TaskProperties struct {
	HsTimestamp     string `json:"hs_timestamp"`
	HsTaskBody      string `json:"hs_task_body"`
	HsTaskStatus    string `json:"hs_task_status"`
	HsTaskType      string `json:"hs_task_type"`
	HsTaskPriority  string `json:"hs_task_priority"`
	HsTaskSubject   string `json:"hs_task_subject"`
	HsTaskReminders string `json:"hs_task_reminders"`
}
