package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func echo(w http.ResponseWriter, req *http.Request){
	for name,value := range req.Header {
		fmt.Fprintf(w, "%s:%s \n",name,value[0])
	}
	body, err:= ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(w, string(body))
}

func notFound(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusNotFound)
	echo(w, req)
}
func serverError(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusInternalServerError)
	echo(w,req)
}

func invalidRequest(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusBadRequest)
	echo(w,req)
}
func main(){
	var port int
	flag.IntVar(&port, "port", 8666, "Port number")
	flag.Parse()
	log.Printf("Starting Echo server in port %d \n", port)
	http.HandleFunc("/", echo)
	http.HandleFunc("/404", notFound)
	http.HandleFunc("/500", serverError)
	http.HandleFunc("/400", invalidRequest)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port),nil)
	if err !=nil {
		log.Fatalln(err)
	}

}
