package common

import "encoding/json"

func SwapToByJson(src, dst interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, dst)
	return err
}
