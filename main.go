package main

import (
	"encoding/json"
	"fmt"
	"github.com/fahrettinbyrm/antalya-me/Order"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Router
	//-Sipariş Ekleme
	//-Sipariş Test Etme
	// 1.Adımda inMemory(Proje çi bir değişkene kaydedilecek.)
	// 2.Adımda postgreSQL'e kayıt edilecek
	//SiparisDeneme()

	r := mux.NewRouter()
	r.HandleFunc("/siparis/ver", SiparisVer).Methods("POST")
	r.HandleFunc("/siparisler", TumSiparisler).Methods("GET")
	fmt.Println("Server 9097'da ayağa kalktı...")
	http.ListenAndServe(":9097", r)

}

func SiparisVer(w http.ResponseWriter, r *http.Request) {
	//Dışarıdan JSON formatında gelen veri okuma mekanizması.

	//Dışarıdan gelen veriyi bu nesnede tutuyoruz.
	var requestBody struct {
		Description string `json:"description"`
		IsUser      bool   `json:"is_user"`
	}

	//Dışarıdan gelen veri nesne içersine yazmak için tuttuk.
	json.NewDecoder(r.Body).Decode(&requestBody)
	if requestBody.Description == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Açıklama Boş Bırakılamaz..!"))
		return
	}
	if !requestBody.IsUser {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Sipariş geçebilmek için önce login olmalısınız..!"))
		return
	}
	//Dışarıdan gelen veriyi oluşturduğumuz nesneye yazdırdık.
	siparis := Order.NewSiparis(requestBody.Description)
	//Dışarıdan gelen veriyi listeledik.
	json.NewEncoder(w).Encode(siparis)
}

func TumSiparisler(w http.ResponseWriter, r *http.Request) {
	siparisler := []Order.Siparis{}

	//Dizi olarak çoklama?
	for _, siparis := range Order.Siparisler {
		fmt.Println(siparis.Code)
		if !siparis.IsDelivered {
			siparisler = append(siparisler, *siparis)

		}

		json.NewEncoder(w).Encode(siparisler)

	}

}

func SiparisDeneme() {
	siparis1 := Order.NewSiparis("Siparis 1")
	siparis2 := Order.NewSiparis("Siparis 2")
	siparis3 := Order.NewSiparis("Siparis 3")
	siparis4 := Order.NewSiparis("Siparis 4")

	siparis1.IsDelivered = true
	siparis2.IsDelivered = true
	siparis3.IsDelivered = false
	siparis4.IsDelivered = true

	for _, v := range Order.Siparisler {
		fmt.Println("Sipariş içeriği: ", v)
	}

}
