package main

import "testing"

type GetCase struct {
	data []byte
	key  string
}

func TestValidXMLHeader(t *testing.T) {
	cases := map[string]GetCase{
		"easy_case": GetCase{
			data: []byte(`<?xml version="1.0" encoding="UTF-8"?><name>jack</name>`),
			key:  "name",
		},
	}
	for k, v := range cases {
		value, offset := validXMLHeader(v.data)
		t.Logf("case:%+v, value:%+v, offset:%+v", k, value, offset)
	}
}
