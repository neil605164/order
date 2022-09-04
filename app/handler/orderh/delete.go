package orderh

import (
	"net/http"
	"order/app/global/helper"

	"github.com/gin-gonic/gin"
)

// DeleteOrder 刪除訂單
// @Description  刪除訂單
// @Tags     order
// @Produce  json
// @Param  id path int true "訂單流水號"
// @Success      200  {object} structer.UserList
// @Failure      400  {object} structer.APIResult "異常錯誤"
// @Router       /order/{id} [delete]
func DeleteOrder(c *gin.Context) {
	c.JSON(http.StatusOK, helper.Success(""))
}
