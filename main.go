package main

import (
	"ginblog/model"
	"ginblog/routers"
)

func main(){
	model.InitDb()

	routers.InitRouter()

}

