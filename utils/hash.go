package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

/*
 *对于一个字符串数据进行MD5哈希计算
 */
func MD5HashString(data string) string {
   md5Hash := md5.New()
   md5Hash.Write([]byte(data))
   bytes := md5Hash.Sum(nil)
	return hex.EncodeToString(bytes)

}
/*
 *io : input output
 */
func MD5HashReader(reader io.Reader) (string,error) {
	md5Hash := md5.New()
	readerBytes, err := ioutil.ReadAll(reader)
	if err != nil{
		return "", err
	}
	md5Hash.Write(readerBytes)
	hashBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}
/*
 *读取io流中的数据，并对数据进行HASH计算返回SHA256 hash的值
 */
func SHA256HashReader(reader io.Reader) (string,error){
	sha256Hash := sha256.New()
	readerBytes, err := ioutil.ReadAll(reader)
	if err != nil{
		return "", err
	}
	sha256Hash.Write(readerBytes)
	hashBytes := sha256Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}
//对数据区块进行SHA256哈希计算
func SHA256HashBlock(bs []byte) []byte{
	//2，将转换后的[]byte字节切片输入Write方法
	sha256Hash := sha256.New()
	sha256Hash.Write(bs)
	hash := sha256Hash.Sum(nil)
	return hash
}