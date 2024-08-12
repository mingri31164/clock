package common

import (
	"go-admin/app/admin/models"
	"gorm.io/gorm"
	"log"
	"time"
)

func deleteOldRecords(db *gorm.DB) error {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	err := db.Where("date < ?", yesterday).Delete(&models.Durations{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func StartScheduledDeletion(db *gorm.DB) {
	go func() {
		for {
			now := time.Now()
			nextRun := now.Truncate(24 * time.Hour).Add(24 * time.Hour)
			time.Sleep(nextRun.Sub(now))

			err := deleteOldRecords(db)
			if err != nil {
				log.Printf("删除旧记录时出错: %v", err)
			} else {
				log.Println("成功删除旧记录")
			}
		}
	}()
}
