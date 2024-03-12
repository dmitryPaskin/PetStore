package entities

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	UserStatus int    `json:"userStatus"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
