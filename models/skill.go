package models

type Skill struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string
    Level       string
    TargetHours int
    UserID      uint
    PracticeLogs []PracticeLog
}
