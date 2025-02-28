package models

type SpyCat struct {
	ID                uint    `json:"id,omitempty" gorm:"primaryKey"`
	Name              string  `json:"name" binding:"required" gorm:"size:255;not null"`
	YearsOfExperience float64 `json:"years_of_experience" binding:"required" gorm:"type:decimal(5,2);not null"`
	Breed             string  `json:"breed" binding:"required" gorm:"size:255;not null"`
	Salary            float64 `json:"salary" binding:"required" gorm:"type:decimal(10,2);not null"`

	Mission Mission `json:"mission" gorm:"foreignKey:CatID"`
}

type Mission struct {
	ID          uint  `gorm:"primaryKey"`
	CatID       *uint `json:"cat_id" gorm:"index"`
	IsCompleted bool  `gorm:"default:false"`

	Targets []Target `gorm:"foreignKey:MissionID"`
}

func (m *Mission) CheckTargetsUnique() bool {
	targetNames := make(map[string]bool)
	for _, target := range m.Targets {
		if _, exists := targetNames[target.Name]; exists {
			return false
		}
		targetNames[target.Name] = true
	}
	return true
}

type Target struct {
	ID          uint   `gorm:"primaryKey"`
	MissionID   uint   `gorm:"index"`
	Name        string `gorm:"size:255;not null"`
	Country     string `gorm:"size:255;not null"`
	Notes       string `gorm:"type:text"`
	IsCompleted bool   `gorm:"default:false"`

	Mission Mission `gorm:"foreignKey:MissionID" json:"-"`
}
