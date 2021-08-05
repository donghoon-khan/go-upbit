package model

import (
	"reflect"
	"testing"
)

func TestGetErrorResponseFromJson(t *testing.T) {
	cases := []struct {
		json     []byte
		expected *ErrorResponse
	}{
		{
			[]byte(`{"error":{"name":404,"message":"Code not found"}}`),
			&ErrorResponse{
				ErrorBody{
					StatusCode: 404,
					Message:    "Code not found",
				},
			},
		},
	}

	for _, c := range cases {
		actual, e := GetErrorResponseFromJson(c.json)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("GetErrorResponseFromJson(%s) == \n%#v\nexpected \n%#v\n",
				c.json, actual, c.expected)
		}
		if e != nil {
			t.Errorf("error : %s", e.Error())
		}
	}
}
