package domain

type Umkm struct {
	Date        string `json:"date"`
	NamaUmkm    string `json:"nama_umkm"`
	NamaBarang  string `json:"nama_brg"`
	Masuk       string `json:"masuk"`
	Keluar      string `json:"keluar"`
	HargaSatuan string `json:"hrg_satuan"`
	TotalHarga  string `json:"total_hrg"`
	Ket         string `json:"ket"`
}
