package motion

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"voter_backend/db"
)

type Motion struct {
	db.BaseModel
	Name        string   `binding:"required" json:"name"`
	Description string   `json:"description"`
	SessionID   int      `json:"sessionID"`
	Status      int8     `json:"status"`
	UserID      int      `json:"userID"`
	For         intArray `json:"for"     `
	Against     intArray `json:"against" `
	Abstain     intArray `json:"abstain"`
}

func (motion *Motion) AfterFind(tx *gorm.DB) (err error) {
	if motion.For == nil {
		motion.For = []int{}
	}
	if motion.Against == nil {
		motion.Against = []int{}
	}
	if motion.Abstain == nil {
		motion.Abstain = []int{}
	}
	return
}
func (motion *Motion) AfterCreate(tx *gorm.DB) (err error) {
	return motion.AfterFind(tx)
}

type intArray []int

func (p intArray) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *intArray) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &p)
}

//func (p intArray) Value() (driver.Value, error) {
//	s := fmt.Sprint(p)
//	return s[1 : len(s)-1], nil
//}

//func (p *intArray) Scan(value interface{}) error {
//	data := value.(string)
//	array := *p
//	if len(data) == 0 {
//		array = []int{}
//		return nil
//	}
//	stringArray := strings.Split(data, " ")
//	array = make([]int, len(stringArray))
//	var err error
//	for i := 0; i < len(stringArray); i++ {
//		array[i], err = strconv.Atoi(stringArray[i])
//	}
//	return err
//}
