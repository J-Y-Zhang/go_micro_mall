package common

import "encoding/json"

//通过json tag实现结构体快速赋值
func SwapToByJson(src, dst interface{}) error {
    data, err := json.Marshal(src)
    if err != nil {
        return err
    }
    err = json.Unmarshal(data, dst)
    return err
}