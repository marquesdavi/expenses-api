package expenses

import "gorm.io/gorm"

type Expense struct {
    gorm.Model
    Description string  `json:"description"`
    Amount      float64 `json:"amount"`
}
