# Go Client for Authy API

A go library for using the Authy public API.

**This API is not stable yet. Use it at your own risk.**

## Installation

If you don't have `$GOPATH` configured please type the following commands:

```shell
$ export GOPATH=$HOME/GoCode
$ mkdir -p $GOPATH/src
$ echo 'export GOPATH=$HOME/GoCode' >> ~/.bashrc
```

If you already have `$GOPATH` configured then install the package:

```shell
$ go get github.com/dcu/go-authy
```

## Usage

To use this client you just need to import `go-authy` package and initialize it with your API KEY

```go
import(
    "github.com/dcu/go-authy"
)
authyAPI := authy.NewAuthyAPI("#your_api_key")
```

Now that you have an Authy API object you can start sending requests.

## Logger

By default, most operations and errors are logged to `stderr`. You can
access `authy.Logger` to replace the logger. Example:

```go
authy.Logger = log.New(...)
```

## Creating Users

__NOTE: User is matched based on cellphone and country code not e-mail.
A cellphone is uniquely associated with an authy_id.__

Creating users is very easy, you need to pass an email, a cellphone and a country code:

```go
user, err := authyAPI.RegisterUser("new_user@email.com", 44, "405-342-5699", url.Values{})
```

in this case `44` is the country code(UK), use `1` for USA or Canada.

You can easily see if the user was created by calling `user.Valid()`.
If request went right, you need to store the `authy id` in your database. Use `user.Id` to get this `id` in your database.

```go
if user.Valid() {
    # store userResponse.User.Id in your user database
}
```

If something goes wrong `user.Valid()` will return `false` and you can see the errors using the following code

```go
user.Errors
```

it returns a `map[string]string` explaining what went wrong with the request.


## Verifying Tokens


To verify users you need the user id and a token. The token you get from the user through your login form. 

```go
verification, err := authyAPI.VerifyToken(authy-id, "token-entered-by-the-user", url.Values{"ip":{"<user ip>"}})
```

Once again you can use `verification.Valid` to verify whether the token was valid or not.

```go
if verification.Valid() {
    # the user is valid
}
```

## Requesting SMS Tokens

To request a SMS token you only need the user id.

```go
sms, err := authyAPI.RequestSMS("authy-id", url.Values{})
```

As always, you can use `sms.Valid()` to verify if the token was sent. To be able to use this method you need to have activated the SMS plugin for your Authy App.

You should force this request to ensure the user will get a token even if it's using the Authy Mobile App.

## Requesting token via phone call

To request a token via Phone Call you only need the user id.

```go
phoneCall, err := authyAPI.RequestPhoneCall("authy-id", url.Values{"force":{"true"}})
```

As always, you can use `phoneCall.Valid()` to verify if the token was sent. To be able to use this method you need to have activated the Phone Call plugin for your Authy App.

You should force this request to ensure the user will get a token even if it's using the Authy App.

# OneTouch

## Sending an approval request

To send the push notification to a user use the method `SendApprovalRequest` which receives the Authy ID of the user, a message, the details to show to the user and any extra http param you want to send to the server.

```go
details := authy.Details{
    "Type":      "SSH Server",
    "Server IP": serverIP,
    "User IP":   clientIP,
    "User":      os.Getenv("USER"),
}
approvalRequest, err := authyAPI.SendApprovalRequest(authyID, "Log to your ssh server", details, url.Values{})
```

## Wait for approval request result

An easy way to get the response of the user is polling. The method `WaitForApprovalRequest` wraps all the polling code in just one method, use it as follows:

```go
status, err := authyAPI.WaitForApprovalRequest(approvalRequest.UUID, 45, url.Values{})
if status == authy.OneTouchStatusApproved {
    // the request was approved.
}
```


# Phone Verification

## Start a phone verification

To start a phone verification use the following code:
```go
verification, err := authyAPI.StartPhoneVerification(1, "555-555-5555",
authy.SMS)
```

## Check phone verification

To check a phone verification use the following code:
```go
verification, err := authyAPI.CheckPhoneVerification(1, "555-555-5555",
"000000")

if verification.Success {
}
```

## Contributing

Get the code:

```shell
$ go get -u github.com/dcu/go-authy
$ cd $GOPATH/src/github.com/dcu/go-authy
```

and start coding.

### Tests

To run the test just type:

```shell
make test
```

### More...

You can fine the full API documentation in the [official documentation](https://docs.authy.com) page.


