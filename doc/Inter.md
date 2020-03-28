在自己的项目中集成
==================

准备
----

go get github.com/leeli73/goFileView

demo
----

package main

import(

"net/http"

"github.com/leeli73/goFileView/perview"

)

func index(w http.ResponseWriter, r \*http.Request) {

w.Write([]byte("I'm Index"))

}

func main(){

perview.Init("/perview/","no") //初始化

http.HandleFunc("/index",index)

http.HandleFunc("/perview/",perview.Handle) //绑定到preview的Handle

http.ListenAndServe(":80", nil)

}