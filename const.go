package easyxml

import (
	"errors"
)

const (
	xmlBeginHalfCheck = 5
)

var (
	InvalidXMLError           = errors.New("invalid xml")
	KeyNotFoundError          = errors.New("key not found")
	KeysLevelNotMatchError    = errors.New("keys and levels not match")
	GetCurrNodeError          = errors.New("get curr node error")
	ContraryNodeNotFoundError = errors.New("contrary node not found")
)
