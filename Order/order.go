package Order

import "github.com/google/uuid"

// Map içerisinde pointerler adresleriyle aranmalıdır.
var Siparisler map[string]*Siparis // key:string değer - value:Siparis Değerleri

type Siparis struct {
	Code         string `json:"code"`         //Dışarıdan okuma
	Dependencies string `json:"dependencies"` //Dışarıdan okuma
	IsDelivered  bool   `json:"is_delivered"` //Dışarıdan okuma

}

// Kütüphaneye erişim sağlandığında bellekte boş bir map string oluşturulur.
// Ekleme işlemi öncesi bellekte mutlaka yer açılmalıdır.
func init() {
	Siparisler = map[string]*Siparis{}
}

func NewSiparis(dependencies string) *Siparis {
	siparis := &Siparis{
		Code:         uuid.New().String(),
		Dependencies: dependencies,
		IsDelivered:  false,
	}
	Siparisler[siparis.Code] = siparis
	return siparis
}
