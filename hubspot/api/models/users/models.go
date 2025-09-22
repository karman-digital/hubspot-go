package usersmodels

type UserBody struct {
	Id               string   `json:"id,omitempty"`
	SuperAdmin       *bool    `json:"superAdmin,omitempty"`
	Email            string   `json:"email"`
	FirstName        string   `json:"firstName,omitempty"`
	LastName         string   `json:"lastName,omitempty"`
	PrimaryTeamId    string   `json:"primaryTeamId,omitempty"`
	SendWelcomeEmail bool     `json:"sendWelcomeEmail"`
	RoleId           string   `json:"roleId,omitempty"`
	SecondaryTeamIds []string `json:"secondaryTeamIds,omitempty"`
}
