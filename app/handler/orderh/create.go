package orderh

import (
	"net/http"
	"order/app/business/orderb"
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/global/structer"

	"github.com/gin-gonic/gin"
)

// CreateOrder 建立訂單
// @Description  建立訂單
// @Tags     order
// @Produce  json
// @Param body body structer.CreateOrderReq true "新增訂單"
// @Success      200  {object} structer.APIResult
// @Failure      400  {object} structer.APIResult "異常錯誤"
// @Router       /order [post]
func CreateOrder(c *gin.Context) {
	// 取參數
	raw := structer.CreateOrderReq{}
	if err := c.ShouldBind(&raw); err != nil {
		apiErr := helper.ErrorHandle(global.WarnLog, errorcode.Code.BindParamError, err, raw)
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 驗證參數
	if err := helper.ValidateStruct(raw); err != nil {
		apiErr := helper.ErrorHandle(global.WarnLog, errorcode.Code.ValidateParamError, err.Error(), raw)
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 寫入 DB
	bus := orderb.Instance()
	apiErr := bus.CreateOrder(raw)
	if apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(raw))
}
