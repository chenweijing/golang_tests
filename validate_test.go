package main

import "gopkg.in/go-playground/validator.v9"

type RegisterReq struct {
	Username       string `validate:"gt=0"`
	PasswordNew    string `validate:"gt=0"`
	PasswordRepeat string `validate:"eqfield=PasswordNew"`
	Email          string `validate:"email"`
}

validate := validator.New()

func validate(req RegisterReq) error {
	err := validator.Struct(req)

	if err != nil {
		return err
	}
}

func main() {

	var req = RegisterReq {
		Username :"Xargin",
		PasswordNew : "oh"
	}

}
