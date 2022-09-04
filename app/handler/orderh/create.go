package orderh

import (
	"net/http"
	"order/app/global/helper"

	"github.com/gin-gonic/gin"
)

// CreateOrder 建立訂單
// @Description  建立訂單
// @Tags     order
// @Produce  json
// @Param body body structer.CreateOrderReq true "新增訂單"
// @Success      200  {object} structer.UserList
// @Failure      400  {object} structer.APIResult "異常錯誤"
// @Router       /order [post]
func CreateOrder(c *gin.Context) {
	c.JSON(http.StatusOK, helper.Success(""))
}
