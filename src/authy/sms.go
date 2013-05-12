package authy

type SmsVerification struct {
    Message string `string:"message"`
    Valid bool
}

