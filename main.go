package main

import (
	_ "github.com/huydeerpets/socialloot/conf/inits"
	_ "github.com/huydeerpets/socialloot/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
