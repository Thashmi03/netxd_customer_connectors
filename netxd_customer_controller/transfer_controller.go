package netxdcustomercontroller

import (
	"context"
	tmodel "github.com/Thashmi03/transfer_model"
	tinterface "github.com/Thashmi03/transfer_interface"
	
	t "github.com/Thashmi03/netxd_transfer"
)
type RPCServer struct{
	t.UnimplementedCustomerServiceServer
}

var(
	TransferService tinterface.Itransact
)

func(s *RPCServer)Transfer(ctx context.Context,req * t.Details)(*t.DetailResponse,error){
	// dbTransfer:=&tmodel.Transaction{
	// 	Transaction_id: "00001",
	// 	From_account:   317,
	// 	To_account:     318,
	// 	Amount:         100,
	// }
	dbTransfer:=&tmodel.Transaction{
		Transaction_id: "00001",
		From_account:   317,
		To_account:     318,
		Amount:         100,
	}
	res,err:=TransferService.Transfer(dbTransfer)
	
	if err != nil {
		return nil, err
	}else {
		responseProfile := &t.DetailResponse{
			Transaction_id: res.CustomerId,
		}
		return responseProfile, nil
	}

}