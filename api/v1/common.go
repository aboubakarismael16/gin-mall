package v1

import (
	"encoding/json"
	"gin-mall/serializer"
)

func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 400,
			Msg:    "JSON type not correct",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: 400,
		Msg:    "wrong parameter",
		Error:  err.Error(),
	}
}
