package main

import (
	"net/http"
	"uas-promnet-kom3a-2204919/controllers/inventorycontroller"
)

func main() {

	http.HandleFunc("/", inventorycontroller.Index)
	http.HandleFunc("/inventory", inventorycontroller.Index)
	http.HandleFunc("/inventory/create", inventorycontroller.Create)
	http.HandleFunc("/inventory/update", inventorycontroller.Update)
	http.HandleFunc("/inventory/delete", inventorycontroller.Delete)

	http.ListenAndServe(":8080", nil)

}