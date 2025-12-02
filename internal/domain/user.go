package domain

type UserLogin struct {
	Username string
	Password string
}

type UserDetails struct {
	Age int
}

type User struct {
	Username       *string `json:"username"`
	HashedPassword *string `json:"password,omitempty"`
	Id             *int64  `json:"id,omitempty"`
	Age            *int    `json:"age,omitempty"`
}
