package models

import "time"

type LeadsGen struct {
	LeadsID        int64     `json:"leads_id"`
	Source         string    `json:"source"`
	UsiaLeads      int64     `json:"usia_leads"`
	Level          int       `json:"level"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	AsssignedAt    time.Time `json:"assigned_at"`
	Status         string    `json:"status"`
	KodeWP         string    `json:"kode_wp"`
	NamaMitra      string    `json:"nama_mitra"`
	NamaWarung     string    `json:"nama_warung"`
	NoHP           string    `json:"no_hp"`
	Kota           string    `json:"kota"`
	Kecamatan      string    `json:"kecamatan"`
	Alamat         string    `json:"alamat"`
	Latlong        string    `json:"lat_long"`
	FotoWarung     string    `json:"foto_warung"`
	PatokanWarung  string    `json:"patokan_warung"`
	KategoriWarung string    `json:"kategori_warung"`
}
