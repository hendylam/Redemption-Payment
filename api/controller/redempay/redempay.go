package inquiry

import (
	"net/http"

	"tapera.integrasi/api/util/httpx"
	srv "tapera.integrasi/service/redempay"
)

func (c *Controller) RedemPay(w http.ResponseWriter, r *http.Request) {
	ext := httpx.New(w, r)
	var param srv.RedempPayData
	if err := ext.BindJSON(&param); err != nil {
		ext.JSONerr(http.StatusBadRequest, "Invalid request payload")
		return
	}
	var result *srv.RedemPayResponse
	result, _ = c.srv.RedemptionPayment(r.Context(), &param)

	ext.JSON(http.StatusOK, result)
}
