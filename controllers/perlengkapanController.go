package controllers

import (
    "encoding/json"
    "net/http"
    "api-asset2/config"
    "api-asset2/models"
    // "github.com/gorilla/mux"
)

func GetPerlengkapans(w http.ResponseWriter, r *http.Request) {
    rows, err := config.DB.Query("SELECT * FROM perlengkapans")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var perlengkapans []models.Perlengkapan
    for rows.Next() {
        var perlengkapan models.Perlengkapan
        if err := rows.Scan(&perlengkapan.ID, &perlengkapan.Nama,&perlengkapan.CreatedAt); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        perlengkapans = append(perlengkapans, perlengkapan)
    }
    json.NewEncoder(w).Encode(perlengkapans)
}

// func GetPermintaan(w http.ResponseWriter, r *http.Request) {
//     params := mux.Vars(r)
//     id, err := strconv.Atoi(params["id"])
//     if err != nil {
//         http.Error(w, "Invalid ID", http.StatusBadRequest)
//         return
//     }

//     var permintaan models.Permintaan
//     err = config.DB.QueryRow("SELECT id, event_id, tgl_permintaan, tenaga_it, dok_support, catatan, diminta_oleh, diserahkan_oleh, disetujui_oleh, created_at FROM perlengkapans WHERE id = ?", id).Scan(&permintaan.ID, &permintaan.EventID, &permintaan.TglPermintaan, &permintaan.TenagaIT, &permintaan.DokSupport, &permintaan.Catatan, &permintaan.DimintaOleh, &permintaan.DiserahkanOleh, &permintaan.DisetujuiOleh, &permintaan.CreatedAt)
//     if err == sql.ErrNoRows {
//         w.Header().Set("Content-Type", "application/json")
//         response := Response{Message: "ID tidak ditemukan"}
//         if err := json.NewEncoder(w).Encode(response); err != nil {
//             http.Error(w, err.Error(), http.StatusInternalServerError)
//         }
//         return
//     } else if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     json.NewEncoder(w).Encode(permintaan)
// }

// func CreatePermintaan(w http.ResponseWriter, r *http.Request) {
//     var permintaan models.Permintaan
//     if err := json.NewDecoder(r.Body).Decode(&permintaan); err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }

//     permintaan.CreatedAt = time.Now()
//     res, err := config.DB.Exec("INSERT INTO perlengkapans (event_id, tgl_permintaan, tenaga_it, dok_support, catatan, diminta_oleh, diserahkan_oleh, disetujui_oleh, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", permintaan.EventID, permintaan.TglPermintaan, permintaan.TenagaIT, permintaan.DokSupport, permintaan.Catatan, permintaan.DimintaOleh, permintaan.DiserahkanOleh, permintaan.DisetujuiOleh, permintaan.CreatedAt)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     id, err := res.LastInsertId()
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     permintaan.ID = int(id)
//     json.NewEncoder(w).Encode(permintaan)
// }

// func UpdatePermintaan(w http.ResponseWriter, r *http.Request) {
//     params := mux.Vars(r)
//     id, err := strconv.Atoi(params["id"])
//     if err != nil {
//         http.Error(w, "Invalid ID", http.StatusBadRequest)
//         return
//     }

//     var permintaan models.Permintaan
//     if err := json.NewDecoder(r.Body).Decode(&permintaan); err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }

//     var exists bool
//     err = config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM perlengkapans WHERE id = ?)", id).Scan(&exists)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     if !exists {
//         w.Header().Set("Content-Type", "application/json")
//         response := Response{Message: "ID tidak ditemukan"}
//         if err := json.NewEncoder(w).Encode(response); err != nil {
//             http.Error(w, err.Error(), http.StatusInternalServerError)
//         }
//         return
//     }

//     _, err = config.DB.Exec("UPDATE perlengkapans SET event_id = ?, tgl_permintaan = ?, tenaga_it = ?, dok_support = ?, catatan = ?, diminta_oleh = ?, diserahkan_oleh = ?, disetujui_oleh = ? WHERE id = ?", permintaan.EventID, permintaan.TglPermintaan, permintaan.TenagaIT, permintaan.DokSupport, permintaan.Catatan, permintaan.DimintaOleh, permintaan.DiserahkanOleh, permintaan.DisetujuiOleh, id)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     json.NewEncoder(w).Encode(permintaan)
// }

// func DeletePermintaan(w http.ResponseWriter, r *http.Request) {
//     params := mux.Vars(r)
//     id, err := strconv.Atoi(params["id"])
//     if err != nil {
//         http.Error(w, "Invalid ID", http.StatusBadRequest)
//         return
//     }

//     // Check if the ID exists
//     var exists bool
//     err = config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM perlengkapans WHERE id = ?)", id).Scan(&exists)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     if !exists {
//         w.Header().Set("Content-Type", "application/json")
//         response := Response{Message: "ID tidak ditemukan"}
//         if err := json.NewEncoder(w).Encode(response); err != nil {
//             http.Error(w, err.Error(), http.StatusInternalServerError)
//         }
//         return
//     }

//     _, err = config.DB.Exec("DELETE FROM perlengkapans WHERE id = ?", id)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     w.Header().Set("Content-Type", "application/json")
//     response := Response{Message: "Berhasil menghapus data"}
//     if err := json.NewEncoder(w).Encode(response); err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//     }
// }
