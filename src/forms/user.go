package forms

type UserCredential struct {
	// This should be improved by using custom validator
	Username string `json:"username" binding:"required,alphanum,lowercase,min=3,max=15" db:"username"`
	Password string `json:"password" binding:"required,min=6,max=20" db:"password"`
}

type UserInfo struct {
	Name     string `json:"name" binding:"required"`
	Birthday string `json:"birthday,omitempty"`
	Gender   string `json:"gender,omitempty"`
	PhotoURL string `json:"photo_url,omitempty"`
}
