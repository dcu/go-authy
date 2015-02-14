package authygo

import(
)

type UserOpts struct {
    Email string
    PhoneNumber string
    CountryCode int
}

type UserResponse struct {
    User *User `json:"user"`
    Errors map[string]string `json:"errors"`
    Message string `string:"message"`
    Valid bool
}

type User struct {
    Id int `json:"id"`
}

