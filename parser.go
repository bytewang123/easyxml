package main

import (
	"fmt"
)

func Get(data []byte, keys ...string) (interface{}, error) {

	var i int
	dl := len(data)
	kn := len(keys)
	keylengths := make([]int, kn)
	for _, v := range keys {
		keylengths = append(keylengths, len(v))
	}

	hasXMLHeader, offset := validXMLHeader(data)
	if hasXMLHeader {
		i = offset
	} else {
		i = 0
	}

	var firstTime = true
	var level = 0
	var match bool

	for i < dl {
		if firstTime {
			if data[i] != '<' {
				return nil, InvalidXMLError
			}
			firstTime = false
		}

		match, currentNode, offset, err :=
			matchKey(i, data, level, keys, keylengths)

		if err != nil {
			return nil, err
		}
		if match {
			i = offset
			level++
			continue
		}
		offset = escapeCurrNode(i, data, currentNode)
		i = offset
	}
	return nil, nil
}

func validXMLHeader(data []byte) (bool, int) {
	var offset = 0
	dl := len(data)

	if dl < xmlBeginHalfCheck {
		return false, 0
	}
	xmlBeginHalf := data[:xmlBeginHalfCheck]

	if pass := xmlBeginHalf[0] == '<' &&
		xmlBeginHalf[1] == '?' &&
		xmlBeginHalf[2] == 'x' &&
		xmlBeginHalf[3] == 'm' &&
		xmlBeginHalf[4] == 'l'; !pass {
		return false, 0
	}

	i := xmlBeginHalfCheck
	for i < dl {
		switch data[i] {
		case '?':
			if data[i+1] == '>' {
				offset = i + 2
				return true, offset
			}
			return false, 0
		default:
			i++
		}
	}
	return false, 0
}
