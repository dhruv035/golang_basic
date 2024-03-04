package responses

import "basic_server/models"

type GasResponse struct {
	GasObjects []models.GasObject
}

type ErrorResponse struct {
	Msg string
}
