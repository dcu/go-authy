package authy

import(
    "testing"
    "authy"
)

func Test_RegisterUserWithInvalidData(t *testing.T) {
    api := authy.NewAuthyApi("bf12974d70818a08199d17d5e2bae630")
    api.ApiUrl = "http://sandbox-api.authy.com"

    user, err := api.RegisterUser("foo", "123", 1)

    if err == nil {
        t.Log("No comm error found")
    }

    if user.Valid == true {
        t.Error("User should not be valid.")
    }

    t.Log("Errors:", user.Errors)
    if user.Errors["email"] != "is invalid" {
        t.Error("Invalid error returned by server.")
    }
}

func Test_RegisterUserWithValidData(t *testing.T) {
    api := authy.NewAuthyApi("bf12974d70818a08199d17d5e2bae630")
    api.ApiUrl = "http://sandbox-api.authy.com"

    user, err := api.RegisterUser("foo@example.com", "432-123-2323", 1)

    if err != nil {
        t.Log("Comm error found:",err)
    }

    if user.Valid != true {
        t.Error("User should be valid.")
    }

    t.Log("Errors:", user.Errors)

    if user.Id == 0 {
        t.Error("User id should be set.")
    }
}


