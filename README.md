# email-verifier

✉️ Fork of [email-verifier](https://github.com/AfterShip/email-verifier) with added functionality for handeling large input emails

## Features
- All existing features of parent project

## Install

Use `go get` to install this package.

```shell script
go get -u github.com/AfterShip/email-verifier
```

## Usage

### Basic usage

Use `VerifyEmails` method to verify an email addresses with different dimensions

```go
package main

import (
	"fmt"
	emailverifier "github.com/AfterShip/email-verifier"
)
var (
	verifier = emailverifier.NewVerifier()
)

func main() {
	email := "example@exampledomain.org"

	ret, err := verifier.Verify(email)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		return
	}
	if !ret.Syntax.Valid {
		fmt.Println("email address syntax is invalid")
		return
	}

	fmt.Println("email validation result", ret)
	/*
		results: {
			"email":"example@exampledomain.org",
			"disposable":false,
			"reachable":"unknown",
			"role_account":false,
			"free":false,
			"syntax":{
			"username":"example",
				"domain":"exampledomain.org",
				"valid":true
			},
			"has_mx_records":true,
			"smtp":null,
			"gravatar":null
		}],
		errors: null
	*/
}
```
## API 

We provide a simple **self-hosted** [API server](https://github.com/X0EF/email-verifier/tree/main/cmd/apiserver) script for reference.

The API interface is very simple. All you need to do is to send a GET request with the following URL.

The `email` parameter would be the target email you want to verify.

POST to `https://{your_host}/v1/bulk/verifications`
Body 
```
{emails: ["example@mail.com", "sample@email.com"]}
```
## License

This package is licensed under MIT license.
