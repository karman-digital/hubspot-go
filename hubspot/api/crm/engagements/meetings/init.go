package meetings

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

func NewMeetingsService(creds *credentials.Credentials) *MeetingsService {
	return &MeetingsService{
		Credentials: creds,
	}
}
