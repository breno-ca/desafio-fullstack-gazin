package entity

import "github.com/go-playground/validator/v10"

// Singleton validator instance
var validate = validator.New()
