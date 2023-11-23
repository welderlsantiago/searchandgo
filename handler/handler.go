package handler

import (
	"github.com/welderlsantiago/searchandgo/config"

	"gorm.io/gorm"
)

var {
	logger *config.Logger
	db *gorm.DB
}

func InitializeHandler() {
	logger = config.GetLogger("handler")
}