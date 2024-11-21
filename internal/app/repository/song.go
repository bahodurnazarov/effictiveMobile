package repository

import (
	"errors"
	"github.com/bahodurnazarov/effictiveMobile/pkg/models"
	logger "github.com/bahodurnazarov/effictiveMobile/pkg/utils"
	"gorm.io/gorm"
	"log"
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
	song.DeletedAt = time.Now()
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

func (r *Repository) UpdateSong(songID string, newSong models.Song) (models.Song, error) {
	var song models.Song

	if err := r.Conn.First(&song, songID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.DebugLogger.Println(err)
			return models.Song{}, errors.New("song not found")
		}
		log.Println("Error fetching song:", err)
		logger.DebugLogger.Println(err)
		return models.Song{}, err
	}

	if newSong.SongName != "" {
		song.SongName = newSong.SongText
	}
	if newSong.Group != "" {
		song.Group = newSong.Group
	}
	if newSong.SongText != "" {
		song.SongText = newSong.SongText
	}
	if newSong.Link != "" {
		song.Link = newSong.Link
	}
	song.Deleted = newSong.Deleted

	if err := r.Conn.Save(&song).Error; err != nil {
		logger.DebugLogger.Println(err)
		return models.Song{}, err
	}
	return song, nil
}

func (r *Repository) GetAllSongs(limit int, offset int) ([]models.Song, error) {
	var songs []models.Song
	if err := r.Conn.Where("deleted = ?", false).Limit(limit).Offset(offset).Find(&songs).Error; err != nil {
		logger.DebugLogger.Println(err)
		return nil, err
	}
	return songs, nil
}
