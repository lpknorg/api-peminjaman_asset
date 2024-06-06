package models

import "time"

type Permintaan struct {
    ID            int       `json:"id"`
    EventID       int       `json:"event_id"`
    TglPermintaan string    `json:"tgl_permintaan"`
    TenagaIT      int       `json:"tenaga_it"`
    DokSupport    int       `json:"dok_support"`
    Catatan       string    `json:"catatan"`
    DimintaOleh   string    `json:"diminta_oleh"`
    DiserahkanOleh string   `json:"diserahkan_oleh"`
    DisetujuiOleh string    `json:"disetujui_oleh"`
    CreatedAt     time.Time `json:"created_at"`
    PerlengkapanListId []int `json:"perlengkapan_list_id"`
    PerlengkapanListJumlah []int `json:"perlengkapan_list_jumlah"`
}
