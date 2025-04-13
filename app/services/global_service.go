package services

import (
	"log"
	"time"
	"github.com/mfarhanz1/go-api-setoran-hafalan/app/types"
)

type GlobalService struct {}

func (gs *GlobalService) Introduce() (*types.GlobalServiceIntroduceType, error) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println("[ERROR] Error loading location:", err)
		return nil, err
	}

	t := time.Now().In(loc)
	timestamp := t.Format("Asia/Jakarta ~ 02/01/2006, 15:04:05 WIB")

	return &types.GlobalServiceIntroduceType{
		Response:    true,
		Message:     "Cihuy, Halow Semua 👋 ~ Selamat datang di API Setoran Hafalan! 🎉",
		Version:     "v2.1.4-alpha",
		Contributor: "https://github.com/MFarhanZ1/go-api-setoran-hafalan",
		Timestamp:   timestamp,
	}, nil
}

func NewGlobalService() *GlobalService {
	return &GlobalService{}
}