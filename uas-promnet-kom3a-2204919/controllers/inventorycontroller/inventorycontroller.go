package inventorycontroller

import (
	"html/template"
	"net/http"
	"uas-promnet-kom3a-2204919/entities"
)

func Index(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("views/inventory/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)
}
func Create(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet{
		temp, err := template.ParseFiles("views/inventory/index.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)	
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var inventory entities.Inventory

		inventory.macaron = request.Form.Get("macaron")
		inventory.jumlah = request.Form.Get("jumlah")
		inventory.harga_satuan = request.Form.Get("harga_satuan")
		inventory.lokasi = request.Form.Get("lokasi")
		inventory.deskripsi = request.Form.Get("deskripsi")
	}
}
func Update(response http.ResponseWriter, request *http.Request) {

}
func Delete(response http.ResponseWriter, request *http.Request) {

}