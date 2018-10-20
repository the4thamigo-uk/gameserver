package memorystore

import (
	"encoding/base64"
	"encoding/json"
)

func toJSONBase64(obj interface{}) (string, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	s := base64.StdEncoding.EncodeToString(b)
	return s, nil
}

func fromJSONBase64(s string, obj interface{}) error {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, obj)
}
