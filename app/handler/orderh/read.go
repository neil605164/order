package orderh

import (
	"net/http"
	"order/app/global/helper"

	"github.com/gin-gonic/gin"
)

// OrderList 訂單清單
// @Description  訂單清單
// @Tags     order
// @Produce  json
// @Success      200  {object} structer.OrderList
// @Failure      400  {object} structer.APIResult "異常錯誤"
// @Router       /order [get]
func OrderList(c *gin.Context) {
	c.JSON(http.StatusOK, helper.Success(""))
}

// OrderDetail 訂單詳情
// @Description  訂單詳情
// @Tags     order
// @Produce  json
// @Param  id path int true "訂單流水號"
// @Success      200  {object} structer.OrderDetail
// @Failure      400  {object} structer.APIResult "異常錯誤"
// @Router       /order/{id} [get]
func OrderDetail(c *gin.Context) {
	c.JSON(http.StatusOK, helper.Success(""))
}
