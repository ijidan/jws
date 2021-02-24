package main

import "github.com/ijidan/jws/jws"

//入口函数
func main()  {
	go jws.StartWSServer()
	//阻塞
	ch:=make(chan struct{})
	<-ch
}
