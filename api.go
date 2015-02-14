package authygo

import(
    "net/http"
    "net/url"
    "encoding/json"
    "io/ioutil"
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

func (authy *Authy) RegisterUser(opts UserOpts) (*UserResponse, error) {
    userResponse := &UserResponse{}
    var err error

    log.Println("Creating user with", opts.Email, ",", opts.PhoneNumber, "and", opts.CountryCode)
    resp, err := http.PostForm(authy.ApiUrl+"/protected/json/users/new", url.Values{
        "user[cellphone]": {opts.PhoneNumber},
        "user[country_code]": {strconv.Itoa(opts.CountryCode)},
        "user[email]": {opts.Email},
        "api_key": {authy.ApiKey},
    })

    if err != nil {
        log.Fatal("Error while contacting the API:",err)
        return userResponse, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    userResponse.Valid = (resp.StatusCode == 200)
    if err != nil {
        log.Fatal("Error reading from API:", err)
        return userResponse, err
    }

    err = json.Unmarshal(body, userResponse)
    if err != nil {
        log.Fatal("Error parsing JSON:", err)
        return userResponse, err
    }

    return userResponse, err
}

func (authy *Authy) VerifyToken(userId int, token string) (TokenVerification, error) {
    var tokenVerification TokenVerification
    var err error

    resp, err := http.Get(authy.ApiUrl+"/protected/json/verify/"+url.QueryEscape(token)+"/"+url.QueryEscape(strconv.Itoa(userId))+"?api_key="+url.QueryEscape(authy.ApiKey) )

    if err != nil {
        log.Fatal("Error while contacting the API:",err)
        return tokenVerification, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    tokenVerification.Valid = (resp.StatusCode == 200)
    if err != nil {
        log.Fatal("Error reading from API:", err)
        return tokenVerification, err
    }

    err = json.Unmarshal(body, &tokenVerification)
    if err != nil {
        log.Fatal("Error parsing JSON:", err)
        return tokenVerification, err
    }

    return tokenVerification, err
}

func (authy *Authy) RequestSms(userId int, force bool) (SmsVerification, error) {
    var smsVerification SmsVerification
    var err error

    resp, err := http.Get(authy.ApiUrl+"/protected/json/sms/"+url.QueryEscape(strconv.Itoa(userId))+"?api_key="+url.QueryEscape(authy.ApiKey)+"&force="+strconv.FormatBool(force) )

    if err != nil {
        log.Fatal("Error while contacting the API:",err)
        return smsVerification, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    smsVerification.Valid = (resp.StatusCode == 200)
    if err != nil {
        log.Fatal("Error reading from API:", err)
        return smsVerification, err
    }

    err = json.Unmarshal(body, &smsVerification)
    if err != nil {
        log.Fatal("Error parsing JSON:", err)
        return smsVerification, err
    }

    return smsVerification, err
}

