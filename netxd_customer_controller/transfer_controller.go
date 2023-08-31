package netxdcustomercontroller

import (
	"context"
	netxddalinterface "github.com/Thashmi03/netxd_dal/netxd_dal_interface"
	netxddalmodels "github.com/Thashmi03/netxd_dal/netxd_dal_models"
	t "github.com/Thashmi03/netxd_transfer"
)
type RPServer struct{
	t.UnimplementedCustomerServiceServer
}

var(
	TransferService netxddalinterface.Itransact
)

func(s *RPServer)Transfer(ctx context.Context,req * t.Details)(string,error){
	// dbTransfer:=&tmodel.Transaction{
	// 	Transaction_id: "00001",
	// 	From_account:   317,
	// 	To_account:     318,
	// 	Amount:         100,
	// }
	dbTransfer:=&netxddalmodels.Transaction{
		Transaction_id: "00001",
		From_account:   317,
		To_account:     318,
		Amount:         100,
	}
	_,err:=TransferService.Transfer(dbTransfer)
	
	if err != nil {
		return "no", err
	}
	return "yes",nil
	}
	

