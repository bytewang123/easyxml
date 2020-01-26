package easyxml

import (
	"bytes"
	"fmt"
)

func matchKey(
	start int,
	data []byte,
	level int, keys []string, keylengths []int) (
	bool,
	[]byte,
	int,
	error) {
	var key string
	var keyLen int
	dl := len(data)
	if level >= len(keys) {
		return false, nil, -1, KeysLevelNotMatchError
	}
	var i = start
	for i < dl {
		if data[i] == '<' {
			i++
			continue
		}
		key = keys[level]
		keyLen = keylengths[level]
		getKey := data[i : i+keyLen]
		if match := equalStr(&getKey, key); match {
			return true, getKey, i + keyLen + 1, nil
		}
		currNode, offset := getCurrNode(data, start)
		if currNode != nil {
			return false, currNode, offset, nil
		}
		return false, nil, -1, GetCurrNodeError
	}
}

func getCurrNode(data []byte, start int) ([]byte, int) {
	i := start
	if data[i] != '<' {
		return nil, -1
	}
	for i < len(data) {
		if data[i] != '>' {
			i++
			continue
		}
		return data[start:i], i
	}
	return nil, -1
}

func escapeCurrNode(start int, data []byte, currNode []byte) (int, error) {
	i := start
	for i < len(data) {
		switch data[i] {
		case '<':
			switch data[i+1] {
			case `/`:
				contrary := data[i+1 : i+1+len(currNode)]
				if bytes.Compare(contrary, currNode) == 0 {
					return i + 1 + len(currNode), nil
				}
			default:
				i++
				continue
			}
		default:
			i++
			continue
		}
	}
	return -1, ContraryNodeNotFoundError
}
