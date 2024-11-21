package service

import (
	"errors"
	"github.com/bahodurnazarov/effictiveMobile/pkg/models"
	logger "github.com/bahodurnazarov/effictiveMobile/pkg/utils"
)

func (s *Service) CreateSong(requestBody models.Song) (models.Song, error) {
	if requestBody.SongName == "" {
		logger.DebugLogger.Println("song name is empty")
		return models.Song{}, errors.New("song name is empty")
	}

	song, err := s.Repository.CreateSong(requestBody)
	if err != nil {
		logger.DebugLogger.Println(err)
		return models.Song{}, err
	}
	return song, nil
}

func (s *Service) DeleteSong(songID string) error {
	return s.Repository.DeleteSong(songID)
}

func (s *Service) GetSongByID(songID string) (models.Song, error) {
	return s.Repository.GetSongByID(songID)
}

func (s *Service) UpdateSong(songID string, newSong models.Song) (models.Song, error) {
	return s.Repository.UpdateSong(songID, newSong)
}

func (s *Service) GetAllSongs(limit int, offset int) ([]models.Song, error) {
	return s.Repository.GetAllSongs(limit, offset)
}

func (s *Service) GetSongsWithFilter(requestBody models.RequestSong, limit int, offset int) ([]models.Song, error) {
	return s.Repository.GetSongsWithFilter(requestBody, limit, offset)
}
