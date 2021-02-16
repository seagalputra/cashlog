package web

type GenericResponse struct {
	Message string      `json:"message"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data"`
}

func (g *GenericResponse) SuccessResponse(message string, data interface{}) GenericResponse {
	return GenericResponse{Message: message, Data: data}
}

func (g *GenericResponse) ErrorResponse(message string) GenericResponse {
	return GenericResponse{Message: message, Error: true}
}
