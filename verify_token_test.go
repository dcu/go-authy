package authygo

import(
    "testing"
)

func Test_VerifyTokenWithUnregisteredUser(t *testing.T) {
    api := NewAuthyApi("bf12974d70818a08199d17d5e2bae630")
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
    api := NewAuthyApi("bf12974d70818a08199d17d5e2bae630")
    api.ApiUrl = "http://sandbox-api.authy.com"

    userResponse, err := api.RegisterUser(UserOpts{
		Email: "foo@example.com",
		PhoneNumber: "432-123-1111",
		CountryCode: 1,
	})
    verification, err := api.VerifyToken(userResponse.User.Id, "000000")

    if err != nil {
        t.Log("No comm error found")
    }

    if verification.Valid == true {
        t.Error("Verification should not be valid.")
    }

    if verification.Message == "User doesn't exist." {
        t.Error("Invalid test using an unregistered user.")
    }
}

