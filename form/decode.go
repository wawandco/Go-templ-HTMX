package form

import (
	"time"

	"github.com/go-playground/form"
)

var Decoder = form.NewDecoder()

func init() {
	Decoder.RegisterCustomTypeFunc(TimeFormat, time.Time{})
}

func TimeFormat(vals []string) (interface{}, error) {
	return time.Parse("2006-01-02", vals[0])
}
