package redempay

import (
	"context"
	"log"
	"tapera/util/appcontext"

	mic "tapera.integrasi/grpc/client/mitraintegrasi/v1"
)

func (s *service) RedemptionPayment(ctx context.Context, parm *RedempPayData) (*RedemPayResponse, error) {
	logger := appcontext.Logger(ctx)
	var response *RedemPayResponse = new(RedemPayResponse)
	logger.Debug().Msg("Start Inquiry Confirm Subscribe")

	// create grpc client from manager
	miClient, err := s.miClientMgr.Client(ctx)
	if err != nil {
		return nil, err
	}
	// defer the close function in order to return back the grpc connection into grpc connection pool
	defer miClient.Close()

	var grpcData mic.RedemPayData
	grpcData = mic.RedemPayData{
		PaymentDate: parm.PayDate,
		ReferenceNo: parm.RefNo,
		NavDate:     parm.NavDate,
		Bankbiccode: parm.BankBicCode,
		Bankname:    parm.BankName,
		Accno:       parm.AccNo,
		Accname:     parm.AccName,
	}

	log.Println("Sending Data to gRPC Inquiry Confirm Subscribe ......")

	// inquiryRequest := &mic.PaymentPoolId{grpcData.Paymentpoolid}
	result, err := miClient.RedemptionPayment(ctx, &grpcData)
	log.Println("connected ")
	log.Println("a: ", result)
	log.Println("b: ", err)
	if err != nil {
		logger.Debug().Msgf("%v", err)
	}

	log.Println("response_code: ", result.ResponseCode)
	log.Println("response_description: ", result.ResponseDescription)
	log.Println("response_status: ", result.Status)

	response.ResponseCode = result.ResponseCode
	response.ResponseDescription = result.ResponseDescription
	response.ResponseStatus = result.Status
	response.ResponseRefNo = result.ReferenceNo

	return response, err
}
