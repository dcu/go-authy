package authy

import(
    "testing"
    "authy"
)

func Test_VerifyTokenWithUnregisteredUser(t *testing.T) {
    api := authy.NewAuthyApi("bf12974d70818a08199d17d5e2bae630")
    api.ApiUrl = "http://sandbox-api.authy.com"

    verification, err := api.VerifyToken(0, "000000")

    if err == nil {
        t.Log("No comm error found")
    }

    if verification.Valid == true {
        t.Error("Verification should not be valid.")
    }

    if verification.Message != "User doesn't exist." {
        t.Error("Invalid test using a registered user.")
    }
}

func Test_VerifyTokenWithInvalidToken(t *testing.T) {
    api := authy.NewAuthyApi("bf12974d70818a08199d17d5e2bae630")
    api.ApiUrl = "http://sandbox-api.authy.com"

    user, err := api.RegisterUser("foo@example.com", "432-123-2323", 1)
    verification, err := api.VerifyToken(int(user.Id), "000000")

    if err == nil {
        t.Log("No comm error found")
    }

    if verification.Valid == true {
        t.Error("Verification should not be valid.")
    }

    if verification.Message == "User doesn't exist." {
        t.Error("Invalid test using an not registered user.")
    }
}

