package models

import "time"

type PracticeLog struct {
    ID        uint      `gorm:"primaryKey"`
    SkillID   uint
    Date      time.Time
    Duration  int
    Notes     string
}
