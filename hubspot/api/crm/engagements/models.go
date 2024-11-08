package engagements

import "github.com/karman-digital/hubspot/hubspot/interfaces"

type Engagements struct {
	Notes  interfaces.Notes
	Tasks  interfaces.Tasks
	Calls  interfaces.Calls
	Emails interfaces.Emails
}
