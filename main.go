package main

import (
	"DataCertPlatform/blockchain"
	_ "DataCertPlatform/routers"
	"fmt"
	"github.com/astaxie/beego"
	"DataCertPlatform/db_mysql"
)

func main() {
	//创建创世区块
	block0 := blockchain.CreateGenesisBlock()
	fmt.Println(block0)
	fmt.Printf("block0的哈希：%x",block0.Hash)
	return


	db_mysql.Connect()
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run( )
}

