package repository

import (
	"github.com/bahodurnazarov/effictiveMobile/pkg/models"
	logger "github.com/bahodurnazarov/effictiveMobile/pkg/utils"
	"time"
)

func (r *Repository) CreateSong(requesBody models.Song) (models.Song, error) {
	// Save the new song to the database
	if err := r.Conn.Create(&requesBody).Error; err != nil {
		logger.DebugLogger.Println(err)
		return models.Song{}, err
	}

	return requesBody, nil
}

func (r *Repository) DeleteSong(songID string) error {
	var song models.Song
	if err := r.Conn.First(&song, songID).Error; err != nil {
		logger.DebugLogger.Println(err)
		return err
	}

	song.Deleted = true
	song.UpdatedAt = time.Now()
	if err := r.Conn.Save(&song).Error; err != nil {
		logger.DebugLogger.Println(err)
		return err
	}
	return nil
}

func (r *Repository) GetSongByID(songID string) (models.Song, error) {
	var song models.Song
	if err := r.Conn.First(&song, songID).Error; err != nil {
		logger.DebugLogger.Println(err)
		return models.Song{}, err
	}
	return song, nil
}
