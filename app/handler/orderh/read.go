package orderh

import (
	"net/http"
	"order/app/business/orderb"
	"order/app/global/helper"

	"github.com/gin-gonic/gin"
)

// OrderList 訂單清單
// @Description  訂單清單
// @Tags     order
// @Produce  json
// @Success      200  {object} structer.OrderListResp
// @Failure      400  {object} structer.APIResult "異常錯誤"
// @Router       /order [get]
func OrderList(c *gin.Context) {

	// 取訂單清單
	bus := orderb.Instance()
	orders, apiErr := bus.GetOrderList()
	if apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(orders))
}

// OrderDetail 訂單詳情
// @Description  訂單詳情
// @Tags     order
// @Produce  json
// @Param  id path int true "訂單流水號"
// @Success      200  {object} structer.OrderDetailResp
// @Failure      400  {object} structer.APIResult "異常錯誤"
// @Router       /order/{id} [get]
func OrderDetail(c *gin.Context) {
	// 取參數
	id := c.Param("id")

	// 取訂單清單
	bus := orderb.Instance()
	order, apiErr := bus.GetOrderById(id)
	if apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(order))
}
