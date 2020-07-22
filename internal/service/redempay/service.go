package redempay

import (
	"context"

	mic "tapera.integrasi/grpc/client/mitraintegrasi/v1"
)

type (
	Service interface {
		RedemptionPayment(ctx context.Context, parm *RedempPayData) (*RedemPayResponse, error)
	}

	service struct {
		miClientMgr mic.GrpcClientManager
	}

	RedempPayData struct {
		PayDate     string `json:"payment_date"`
		RefNo       string `json:"reference_no"`
		NavDate     string `json:"nav_date"`
		BankBicCode string `json:"redm_payment_bank_bic_code"`
		BankName    string `json:"redm_payment_bank_name"`
		AccNo       string `json:"redm_payment_acc_no"`
		AccName     string `json:"redm_payment_acc_name"`
	}

	RedemPayResponse struct {
		ResponseCode        string `json:"response_code"`
		ResponseDescription string `json:"response_description"`
		ResponseStatus      string `json:"response_status"`
		ResponseRefNo       string `json:"response_ref_no"`
	}
)

func NewService(miClientMgr mic.GrpcClientManager) Service {
	return &service{miClientMgr: miClientMgr}
}
