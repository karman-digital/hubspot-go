package transcripts

const ObjectTypeMeeting = "MEETING_EVENT"
const ObjectTypeCall = "CALL"

type Speaker struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}
type Utterance struct {
	ID              string  `json:"id,omitempty"`
	Speaker         Speaker `json:"speaker"`
	Text            string  `json:"text"`
	StartTimeMillis int64   `json:"startTimeMillis"`
	EndTimeMillis   int64   `json:"endTimeMillis"`
	LanguageCode    string  `json:"languageCode,omitempty"`
}
type CreateRequest struct {
	ObjectID   string      `json:"objectId"`
	ObjectType string      `json:"objectType,omitempty"`
	Utterances []Utterance `json:"transcriptUtterances"`
}
type CreateResponse struct {
	ID string `json:"id"`
}
type Transcript struct {
	ID         string      `json:"id"`
	ObjectID   string      `json:"objectId"`
	ObjectType string      `json:"objectType"`
	Source     string      `json:"transcriptSource"`
	CreatedAt  string      `json:"createdAt"`
	UpdatedAt  string      `json:"updatedAt"`
	Utterances []Utterance `json:"transcriptUtterances"`
}
type BestResponse struct {
	Status      string      `json:"status"`
	Result      *Transcript `json:"result,omitempty"`
	RequestedAt string      `json:"requestedAt"`
	CompletedAt string      `json:"completedAt,omitempty"`
}
