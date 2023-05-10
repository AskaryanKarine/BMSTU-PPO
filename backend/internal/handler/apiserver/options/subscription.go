package options

import (
	"DoramaSet/internal/handler/apiserver/middleware"
	"DoramaSet/internal/logic/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type subResponse struct {
	Id          int    `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Cost        int    `json:"cost,omitempty"`
	Duration    string `json:"duration,omitempty"`
}

func makeSubResponse(sub model.Subscription) subResponse {
	return subResponse{
		Id:          sub.Id,
		Description: sub.Description,
		Cost:        sub.Cost,
		Duration:    fmt.Sprintf("%s", sub.Duration),
	}
}

func (h *Handler) getAllSubs(c *gin.Context) {
	data, err := h.Services.GetAll()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response := make([]subResponse, len(data))
	for _, el := range data {
		response = append(response, makeSubResponse(el))
	}
	c.JSON(http.StatusOK, gin.H{"Data": response})
}

func (h *Handler) getInfoSub(c *gin.Context) {
	rowId := c.Param("id")
	id, err := strconv.Atoi(rowId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data, err := h.Services.GetInfo(id)
	if err != nil && fatalDB(err) {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": makeSubResponse(*data)})
}

func (h *Handler) subscribe(c *gin.Context) {
	rowId := c.Param("id")
	id, err := strconv.Atoi(rowId)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.SubscribeUser(token, id)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) unsubscribe(c *gin.Context) {
	token, err := middleware.GetUserToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = h.Services.UnsubscribeUser(token)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
