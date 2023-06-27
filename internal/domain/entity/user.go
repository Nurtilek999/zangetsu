package entity

type User struct {
	ID             int    `json:"ID"`
	RoleID         int    `json:"roleID"`
	FirstName      string `json:"firstName"`
	SecondName     string `json:"secondName"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RegisteredDate string `json:"registeredDate"`
	GmailBind      bool   `json:"gmailBind"`
}

type UserViewModel struct {
	//RoleID         int       `json:"roleID"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	//RegisteredDate time.Time `json:"registeredDate"`
}

type UserRegistrationModel struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type UserGroup struct {
	ID    int    `json:"ID"`
	Group string `json:"group"`
}
