package user

type UserInformation struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserRegistrationPayload struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserLoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
