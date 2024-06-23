package db

import (
	"github.com/rs/zerolog/log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string `json:"Name"`
	CPF    int    `json:"CPF"`
	Email  string `json:"Email"`
	Age    int    `json:"Age"`
	Active bool   `json:"Active"`
}

type StudentHandler struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize SQLite: %s", err)
	}

	db.AutoMigrate(&Student{})

	return db
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student Student) error {

	if result := s.DB.Create(&student); result.Error != nil {
		log.Error().Msg("Failed to create student!")
		return result.Error
	}

	log.Info().Msg("Create Student!")
	return nil
}

func (s *StudentHandler) GetStudents() ([]Student, error) {

	students := []Student{}

	err := s.DB.Find(&students).Error
	return students, err
}
