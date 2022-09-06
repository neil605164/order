package orderh

import (
	"net/http"
	"order/app/business/orderb"
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"

	"github.com/gin-gonic/gin"
)

// DeleteOrder 刪除訂單
// @Description  刪除訂單
// @Tags     order
// @Produce  json
// @Param  id path int true "訂單流水號"
// @Success      200  {object} structer.APIResult
// @Failure      400  {object} structer.APIResult "異常錯誤"
// @Router       /order/{id} [delete]
func DeleteOrder(c *gin.Context) {
	// 取參數
	id := c.Param("id")

	// 檢查訂單狀態，只有尚未付款的才可以取消
	bus := orderb.Instance()
	order, apiErr := bus.GetOrderById(id)
	if apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	if order.Status != global.OrderUnpaid {
		apiErr := helper.ErrorHandle(global.WarnLog, errorcode.Code.CanNotDeleteOrder, "", order.Status)
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 刪除訂單(軟刪除)
	if apiErr = bus.DeleteOrderById(id); apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(""))
}
