package cmd

import (
	"codeex/app/service"
	"log"
)

var client *service.Eth

func init() {
	var err error
	client, err = service.NewEth("https://rinkeby.infura.io/v3/abcab11733ea41109495495afe495d56", "0xdD430f0e762330cC2Af1d546A17E1495E8a6c228")
	if err != nil {
		log.Fatalln(err)
	}
}
