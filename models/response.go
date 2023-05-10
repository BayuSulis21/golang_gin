package models

type SearchBarangJsonResponse struct {
	Code     string   `json:"code"`
	Message  string   `json:"message"`
	Kategori string   `json:"kategori"`
	ListData []Barang `json:"ListData"`
}
