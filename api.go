package authy

import(
    "net/http"
    "net/url"
    "log"
    "strconv"
)

type Authy struct {
    ApiKey  string
    ApiUrl  string
}

func NewAuthyApi(apiKey string) *Authy {
    apiUrl := "https://api.authy.com"
    return &Authy{apiKey, apiUrl}
}

func (authy *Authy) RegisterUser(opts UserOpts) (*User, error) {
    log.Println("Creating user with", opts.Email, ",", opts.PhoneNumber, "and", opts.CountryCode)
    resp, err := http.PostForm(authy.ApiUrl+"/protected/json/users/new", url.Values{
        "user[cellphone]": {opts.PhoneNumber},
        "user[country_code]": {strconv.Itoa(opts.CountryCode)},
        "user[email]": {opts.Email},
        "api_key": {authy.ApiKey},
    })

    if err != nil {
        log.Fatal("Error while contacting the API:",err)
        return nil, err
    }

	userResponse, err := NewUser(resp)

    return userResponse, err
}

func (authy *Authy) VerifyToken(userId int, token string) (*TokenVerification, error) {
    resp, err := http.Get(authy.ApiUrl+"/protected/json/verify/"+url.QueryEscape(token)+"/"+url.QueryEscape(strconv.Itoa(userId))+"?api_key="+url.QueryEscape(authy.ApiKey) )
    defer resp.Body.Close()

    if err != nil {
        log.Fatal("Error while contacting the API:",err)
        return nil, err
    }

	tokenVerification, err := NewTokenVerification(resp)
    return tokenVerification, err
}

func (authy *Authy) RequestSms(userId int, force bool) (*SmsRequest, error) {
    resp, err := http.Get(authy.ApiUrl+"/protected/json/sms/"+url.QueryEscape(strconv.Itoa(userId))+"?api_key="+url.QueryEscape(authy.ApiKey)+"&force="+strconv.FormatBool(force) )

    defer resp.Body.Close()
    if err != nil {
        log.Fatal("Error while contacting the API:",err)
        return nil, err
    }

	smsVerification, err := NewSmsRequest(resp)
    return smsVerification, err
}

