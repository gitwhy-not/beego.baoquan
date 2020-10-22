package utils

import (
	"bytes"
	"encoding/binary"
)
//将一个int64转换为[]byte

func Int64ToByte(num int64) ([]byte, error){
	buff := new(bytes.Buffer)
	err := binary.Write(buff,binary.BigEndian, num)
	if  err != nil{
		return nil, err
	}
	return buff.Bytes(), nil
}

func StringToBytes(data string) []byte {
	return []byte(data)
}