package models

import "github.com/go-playground/validator/v10"

type ItemRequest struct {
	Name     string `json:"name"`
	Price    int	`json:"price"`
	Quantity int	`json:"quantity"`
}

// ValidateStruct performs struct based validation
func (itemInput ItemRequest) ValidateStruct() []*ErrorResponse {
	// create a variable to store validation errors
	var errors []*ErrorResponse

	// create a new validator instance
	validate := validator.New()

	// validate the struct
	err := validate.Struct(itemInput)

	// if the validation is failed
	// insert the error message to errors variable
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.ErrorMessage = getErrorMessage(err)
			element.Field = err.Field()
			errors = append(errors, &element)			
			}
		}
	return errors
}

