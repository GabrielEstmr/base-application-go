package main_gateways_ws_commonsresources

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

func (this *ControllerResponse) GetData() interface{} {
	return this.data
}

func (this *ControllerResponse) GetStatusCode() int {
	return this.statusCode
}
