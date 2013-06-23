# Go Client for Authy API

A go library for using the Authy public API.


## Installation

If you don't have `$GOPATH` configured please type the following commands:

    $ export GOPATH=$HOME/GoCode
    $ mkdir -p $GOPATH/src
    $ echo 'export GOPATH=$HOME/GoCode' >> ~/.bashrc


If you already have `$GOPATH` configured then install the package:

	$ go get github.com/dcu/authygo

## Usage

To use this client you just need to import `authygo` package and initialize it with your API KEY

    import "authygo"
    authy_api := authygo.NewAuthyApi('#your_api_key')

Now that you have an Authy API object you can start sending requests.


## Creating Users

__NOTE: User is matched based on cellphone and country code not e-mail.
A cellphone is uniquely associated with an authy_id.__  

Creating users is very easy, you need to pass an email, a cellphone and a country code:
   
    user, err := authy_api.RegisterUser('new_user@email.com', '405-342-5699', 57) #email, cellphone, area_code

in this case `57` is the country code(Colombia), use `1` for USA.

You can easily see if the user was created by calling `user.Valid`.
If request went right, you need to store the `authy id` in your database. Use `user.Id` to get this `id` in your database.

    if user.Valid {
        # store user.Id in your user database
    }

If something goes wrong `user.Valid` returns `false` and you can see the errors using the following code

    user.Errors

it returns a `map[string]string` explaining what went wrong with the request.


## Verifying Tokens


To verify users you need the user id and a token. The token you get from the user through your login form. 

    verification,err := authy_api.VerifyToken(authy-id, 'token-entered-by-the-user')

Once again you can use `verification.Valid` to verify whether the token was valid or not.

    if verification Valid {
        # the user is valid
    }


## Requesting SMS Tokens

To request a SMS token you only need the user id.

	sms,err := authy_api.RequestSms('authy-id', <true|false>) # authy id, force

As always, you can use `sms.Valid` to verify if the token was sent. To be able to use this method you need to have activated the SMS plugin for your Authy App.

You should force this request to ensure the user will get a token even if it's using the Authy Mobile App.

### More…

You can fine the full API documentation in the [official documentation](https://docs.authy.com) page.


## Tests

To run the test just type:

    make test



