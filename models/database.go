package models

type Barang struct {
	Id_barang   string `json:"id_barang"  gorm:"column:id_barang"`
	Nama_barang string `json:"nama_barang"  gorm:"column:nama_barang"`
	Stock       string `json:"stock"  gorm:"column:stock"`
	Kategori    string `json:"kategori"  gorm:"column:kategori"`
	Satuan      string `json:"satuan"  gorm:"column:satuan"`
	Harga       string `json:"harga"  gorm:"column:harga"`
	Supplier    string `json:"supplier"  gorm:"column:supplier"`
}
