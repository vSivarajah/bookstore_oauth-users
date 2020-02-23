package users

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
