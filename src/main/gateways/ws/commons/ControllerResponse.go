package main_gateways_ws_commons

type ControllerResponse struct {
	statusCode int
	data       interface{}
}

func NewControllerResponse(
	statusCode int,
	data interface{},
) *ControllerResponse {
	return &ControllerResponse{
		statusCode: statusCode,
		data:       data,
	}
}

func (c *ControllerResponse) GetData() interface{} {
	return c.data
}

func (c *ControllerResponse) GetStatusCode() int {
	return c.statusCode
}
