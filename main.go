package main

import (
	"fmt"
	"github.com/fahrettinbyrm/antalya-me/Handle"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Router
	//-Sipariş Ekleme
	//-Siparişleri listeleme

	r := mux.NewRouter()
	r.HandleFunc("/siparis/ver", Handle.SiparisVer).Methods("POST")
	r.HandleFunc("/siparisler", Handle.TumSiparisler).Methods("GET")
	fmt.Println("Server 9097'da ayağa kalktı...")
	http.ListenAndServe(":9097", r)

}
