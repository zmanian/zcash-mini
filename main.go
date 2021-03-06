package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/FiloSottile/zcash-mini/zcash"
)

var logo = `

             :88SX@888@@X8:
          %%Xt%tt%SSSSS:XXXt@@
        @S;;tt%%%t    ;;::XXXXSX
      .t:::;;%8888    88888tXXXX8;
     .%...:::8             8::XXX%;
     8888...:t888888X     8t;;::XX8
    %888888...:::;:8    :Xttt;;;::X@
    888888888...:St    8:%%tttt;;;:X
    88888888888S8    :%;ttt%%tttt;;X
    %888888888%t    8S:;;;tt%%%ttt;8
     8t8888888     S8888888Stt%%%t@
     .@tt888@              8;;ttt@;
      .8ttt8@SSSSS    SXXXX%:;;;X;
        X8ttt8888%    %88...::X8
          %8@tt88;8888%8888%8X
             :@888@XXX@888:

                    _       _
          _ __ ___ (_)_ __ (_)
         | '_ \ _ \| | '_ \| |
         | | | | | | | | | | |
         |_| |_| |_|_|_| |_|_|
`

var template = `%s

###############################################################
#
#  Here is your new z-address
#
#      %s
#
#  and here is the secret key
#
#      %s
#
#  KEEP IT SAFE, IT HAS NOT BEEN SAVED ANYWHERE
#
#  To spend, import it with
#
#      zcash-cli z_importkey KEY rescan
#
###############################################################

`

func main() {
	simpleMode := flag.Bool("simple", false, "output only address and key")
	flag.Parse()

	rawKey := zcash.GenerateKey()
	rawAddr, err := zcash.KeyToAddress(rawKey)
	if err != nil {
		log.Fatal(err)
	}

	key := zcash.Base58Encode(rawKey, zcash.ProdSpendingKey)
	addr := zcash.Base58Encode(rawAddr, zcash.ProdAddress)

	if *simpleMode {
		fmt.Printf("%s\n%s\n", addr, key)
	} else {
		fmt.Printf(template, logo, addr, key)
	}
}
