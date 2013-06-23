package authygo

import(
    "testing"
    "authygo"
)

func Test_RequestSms(t *testing.T) {
    api := authygo.NewAuthyApi("bf12974d70818a08199d17d5e2bae630")
    api.ApiUrl = "http://sandbox-api.authy.com"

    user, err := api.RegisterUser("foo@example.com", "432-123-2323", 1)
    verification, err := api.RequestSms(int(user.Id), true)

    if err != nil {
        t.Error("External error found", err)
    }

    if verification.Valid != true {
        t.Error("Verification should be valid.")
    }
}
