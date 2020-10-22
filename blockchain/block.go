package blockchain

import (
	//"DataCertPlatform/utils"
	//"bytes"
	"time"
)

/*
*定义区块结构体，用于表示区块
 */
type Block struct {
	Height     int64   //区块高度，第几个区块
	TimeStamp  int64   //
	PrevHash   []byte
	Data       []byte
	Hash       []byte
	Version    string
	Nonce      int64
}
/*
*创建一个新区块
 */
func NewBlock(height int64, prevHash []byte, data []byte,) Block {
	block := Block{
		Height: height,
		TimeStamp: time.Now().Unix(),
		PrevHash: prevHash,
		Data: data,
		Version: "0x01",
	}
	pow := NewPow(block)
	hash, nonce := pow.Run()
	block.Nonce = nonce
	block.Hash = hash

	/*heighBytes, _ := Int64ToByte(block.Height)
	timeStampByte, _ := Int64ToByte(block.TimeStamp)
	versionBytes, _ := StringToByte(block.Version)
	nonceBytes, _ := utils.Int64ToByte(block.Nonce)
	var blockBytes []byte
	//bytes.Jion拼接
	bytes.Join([][]byte{
		heighBytes,
		timeStampByte,
		block.PrevHash,
		block.Data,
		versionBytes,
		nonceBytes,
	},[]byte{})
	//调用Hash计算，对区块进行sha256哈希计算
	block.Hash = utils.SHA256HashBlock(blockBytes)
	*/
	return block

}
func CreateGenesisBlock() Block {
	 genesisBlock := NewBlock(0, []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,}, nil)
	return genesisBlock
}