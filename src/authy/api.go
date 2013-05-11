package authy

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

func (authy *Authy) RegisterUser(email string, cellphone string, countryCode int) (User, error) {
    var user User
    var err error

    log.Println("Creating user with", email, ",", cellphone, "and", countryCode)
    resp, err := http.PostForm(authy.ApiUrl+"/protected/json/users/new", url.Values{
        "user[cellphone]": {cellphone},
        "user[country_code]": {strconv.Itoa(countryCode)},
        "user[email]": {email},
        "api_key": {authy.ApiKey},
    })

    if err != nil {
        log.Fatal("Error while contacting the API:",err)
        return user, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    user.Valid = (resp.StatusCode == 200)
    if err != nil {
        log.Fatal("Error reading from API:", err)
        return user, err
    }

    err = json.Unmarshal(body, &user)
    if err != nil {
        log.Fatal("Error parsing JSON:", err)
        return user, err
    }

    return user, err
}

func (authy *Authy) VerifyToken(userId int, token string) (TokenVerification, error) {
    var tokenVerification TokenVerification
    var err error

    resp, err := http.Get(authy.ApiUrl+"/protected/json/verify/"+url.QueryEscape(token)+"/"+url.QueryEscape(strconv.Itoa(userId))+"?api_key="+url.QueryEscape(authy.ApiKey)+"&force=true" )

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

