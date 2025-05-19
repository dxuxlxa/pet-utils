package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates a struct using validator tags
func ValidateStruct(s any) error {
	if err := validate.Struct(s); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Error().Err(err).Msg("Invalid validation error")
			return err
		}

		// Return the first validation error
		for _, err := range err.(validator.ValidationErrors) {
			log.Debug().
				Str("field", err.Field()).
				Str("tag", err.Tag()).
				Str("param", err.Param()).
				Msg("Validation error")
			return err
		}
	}
	return nil
}
