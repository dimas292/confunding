package handler

import (
	"confunding/helper"
	"confunding/transaction"
	"confunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionsHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionsHandler{
	return &transactionsHandler{service}
}

func (h *transactionsHandler) GetCampaignTransactin(c *gin.Context){
	
	var input transaction.GetCampaignTransactionsInput

	
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("failed to get transaction's", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser
	
	transactions, err := h.service.GetTransactionByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("failed to get transaction's", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	response := helper.APIResponse("campaign detail", http.StatusOK, "success", transaction.FormatCampaignsTransactions(transactions))
	c.JSON(http.StatusOK, response)
}