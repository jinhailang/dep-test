package main

import (
	"fmt"
	"dep-test/internal/xx"
	"dep-test/internal/yy"
	"dep-test/vv"
	"github.com/jinhailang/dep-test/log"
)

func main(){
	fmt.Println("start.")
	xx.Fx()
	yy.Fx()
	vv.Fx()
    log.Fx()
	fmt.Println("end.")
}
