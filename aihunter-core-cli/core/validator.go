package core

import (
	"github.com/go-playground/validator/v10"
	"github.com/user/aihunter-x/aihunter-core-cli/configs/loader"
)

// Validate validates the configuration.
func Validate(cfg *config.Config) error {
	validate := validator.New()
	return validate.Struct(cfg)
}
