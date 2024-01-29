package photocontroller

import (
	"encoding/json"
	"net/http"

	"github.com/GibrannRafi/task-5-pbi-btpns-GibranRafiH-/app/helper"
	"github.com/GibrannRafi/task-5-pbi-btpns-GibranRafiH-/app/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	photos := &[]models.Photo{}

	// Ambil semua foto dari database
	err := models.DB.Find(&photos).Error

	// Periksa error
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Kirim response JSON berisi semua foto
	helper.ResponseJSON(w, http.StatusOK, photos)
}
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Photo := vars["id"]

	photo := &models.Photo{}

	if err := models.DB.First(&photo, Photo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Data tidak ditemukan", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	helper.ResponseJSON(w, http.StatusOK, photo)
}

func Create(w http.ResponseWriter, r *http.Request) {

	photo := &models.Photo{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&photo); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	if err := models.DB.Create(&photo).Error; err != nil {
		response := map[string]string{"message": "Gagal inser database"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	Photo := vars["id"]

	// Decode updated user data from request body
	photo := &models.Photo{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if models.DB.Model(&photo).Where("id = ?", Photo).Updates(&photo).RowsAffected == 0 {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Gagal MengUpdate Data", http.StatusNotFound)
			return
		}
		return
	}

	response := map[string]string{"message": "Berhasil Update Data"}
	helper.ResponseJSON(w, http.StatusOK, response)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	photo := vars["id"]

	// Konversi ID ke int64

	// Hapus foto berdasarkan ID
	if err := models.DB.Delete(&models.Photo{}, photo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Foto tidak ditemukan", http.StatusNotFound)
			return
		}
		http.Error(w, "Gagal menghapus foto", http.StatusInternalServerError)
		return
	}

	// Kirim response JSON bahwa data berhasil dihapus
	response := map[string]string{"message": "Data Berhasil Dihapus!"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
