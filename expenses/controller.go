package expenses

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type Controller struct {
    DB *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
    return &Controller{DB: db}
}

func (c *Controller) GetAllExpenses(ctx *gin.Context) {
    expenses, err := GetAllExpenses(c.DB)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, expenses)
}

func (c *Controller) GetExpenseByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    expense, err := GetExpenseByID(c.DB, uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
        return
    }

    ctx.JSON(http.StatusOK, expense)
}

func (c *Controller) CreateExpense(ctx *gin.Context) {
    var expense Expense
    if err := ctx.ShouldBindJSON(&expense); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := CreateExpense(c.DB, &expense); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, expense)
}

func (c *Controller) UpdateExpense(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var expense Expense
    if err := ctx.ShouldBindJSON(&expense); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    expense.ID = uint(id)
    if err := UpdateExpense(c.DB, &expense); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, expense)
}

func (c *Controller) DeleteExpense(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := DeleteExpense(c.DB, uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.Status(http.StatusNoContent)
}
