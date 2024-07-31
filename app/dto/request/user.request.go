package request

import validation "github.com/go-ozzo/ozzo-validation"

type UserRegisterRequestInterface interface {
	Validate() error
}

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dto *UserRegisterRequest) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Username, validation.Length(4, 50), validation.Required),
		validation.Field(&dto.Password, validation.Length(6, 12), validation.Required),
	); err != nil {
		return err
	}
	return nil
}
