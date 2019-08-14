package presenter

import (
	"bytes"
	"encoding/json"
)

type Presenter interface {
	Output(v interface{}) (string, error)
}

type presenter struct {
}

func New() Presenter {
	return &presenter{}
}

func (p *presenter) Output(v interface{}) (string, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return "", err
	}
	return buf.String(), nil
}
