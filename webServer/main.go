package main

import (
	"net/http"
)

func main()  {
	http.HandleFunc("/",home)
	http.HandleFunc("/home",home)

	http.ListenAndServe(":8080",nil)
}

func home(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"./webServer/index.html")
}
