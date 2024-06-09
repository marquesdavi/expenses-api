package expenses

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func InitDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("expenses.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }

    db.AutoMigrate(&Expense{})
    return db
}

func GetAllExpenses(db *gorm.DB) ([]Expense, error) {
    var expenses []Expense
    result := db.Find(&expenses)
    return expenses, result.Error
}

func GetExpenseByID(db *gorm.DB, id uint) (Expense, error) {
    var expense Expense
    result := db.First(&expense, id)
    return expense, result.Error
}

func CreateExpense(db *gorm.DB, expense *Expense) error {
    result := db.Create(expense)
    return result.Error
}

func UpdateExpense(db *gorm.DB, expense *Expense) error {
    result := db.Save(expense)
    return result.Error
}

func DeleteExpense(db *gorm.DB, id uint) error {
    result := db.Delete(&Expense{}, id)
    return result.Error
}
