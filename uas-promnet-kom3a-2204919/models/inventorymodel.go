package models

import (
	"database/sql"
	"uas-promnet-kom3a-2204919/config"
	"uas-promnet-kom3a-2204919/entities"
)

type InventoryModel struct {
	conn *sql.DB
}

func NewInventoryModel() *InventoryModel {
	conn, err := config.DBConnectiom()
	if err!= nil {
		panic(err)
	}

	return &InventoryModel{
		conn: conn,
	}
}

func (i *InventoryModel) Create(inventory entities.Inventory) bool {
	result, err := i.conn.Exec(Insert into inventory (macaron, jumlah, harga_satuan, lokasi, deskripsi))
}