package orderh

import (
	"net/http"
	"order/app/global/helper"

	"github.com/gin-gonic/gin"
)

// UpdateOrder 編輯訂單
// @Description  編輯訂單
// @Tags     order
// @Produce  json
// @Param id path int true "訂單流水號"
// @Param body body structer.CreateOrderReq true "新增訂單"
// @Success      200  {object} structer.UserList
// @Failure      400  {object} structer.APIResult "異常錯誤"
// @Router       /order/{id} [put]
func UpdateOrder(c *gin.Context) {
	c.JSON(http.StatusOK, helper.Success(""))
}
