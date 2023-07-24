package api

import (
	"github.com/PlatosCodes/momsrecipes/util"
	"github.com/go-playground/validator/v10"
)

var validCuisineType validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if cuisine_type, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCuisineType(cuisine_type)
	}
	return false
}
