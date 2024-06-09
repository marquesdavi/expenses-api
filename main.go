package main

import (
    "expenses-api/expenses"
    "github.com/gin-gonic/gin"
)

func main() {
    db := expenses.InitDB()

    router := gin.Default()
    controller := expenses.NewController(db)

    router.GET("/expenses", controller.GetAllExpenses)
    router.GET("/expenses/:id", controller.GetExpenseByID)
    router.POST("/expenses", controller.CreateExpense)
    router.PUT("/expenses/:id", controller.UpdateExpense)
    router.DELETE("/expenses/:id", controller.DeleteExpense)

    router.Run(":8080")
}
