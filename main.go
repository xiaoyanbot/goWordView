package main

import(
	"goFileViewLib/perview"
)

func main() {
	perview.Init("/perview/","0.0.0.0:8088")
	perview.StartServer()
}
