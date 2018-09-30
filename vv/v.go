package vv

import (
	"fmt"
	"dep-test/internal/xx"
	"github.com/jinhailang/dep-test/log"
	"github.com/pkg/errors"
)

func Fx(){
	fmt.Println("vv.v.Fx")
	xx.Fx()
	log.Fx()
}
