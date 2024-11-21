package repository

import (
	"github.com/bahodurnazarov/effictiveMobile/pkg/models"
	logger "github.com/bahodurnazarov/effictiveMobile/pkg/utils"
)

func (r *Repository) CreateSong(requesBody models.Song) (models.Song, error) {
	// Save the new song to the database
	if err := r.Conn.Create(&requesBody).Error; err != nil {
		logger.DebugLogger.Println(err)
		return models.Song{}, err
	}

	return requesBody, nil
}
