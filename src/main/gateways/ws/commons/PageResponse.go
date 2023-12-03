package main_gateways_ws_commons

import (
	"encoding/json"
	"io"
)

type PageResponse struct {
	Content       any
	Page          int
	Size          int
	TotalElements int
	TotalPages    int
}

func (this *PageResponse) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(this)
}

func (this *PageResponse) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(this)
}
