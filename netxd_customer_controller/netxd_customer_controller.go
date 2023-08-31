package netxdcustomercontroller

import (
	netxddalinterface "github.com/Thashmi03/netxd_dal/netxd_dal_interface"
	netxddalmodels "github.com/Thashmi03/netxd_dal/netxd_dal_models"
	"context"
	// tmodel "github.com/Thashmi03/transfer_model"
	tinterface "github.com/Thashmi03/transfer_interface"
	c "github.com/Thashmi03/netxd_customer"
	// t "github.com/Thashmi03/netxd_transfer"
)
type RPCServer struct{
	c.UnimplementedCustomerServiceServer
}

var(
	CustomerService netxddalinterface.ICustomer
	TransferService tinterface.Itransact
)

func(s *RPCServer)CreateCustomer(ctx context.Context,req * c.Details)(*c.DetailResponse,error){
	dbProfile:=&netxddalmodels.Customer{
		CustomerId: req.CustomerId,
		Firstname:  req.Firstname,
		Lastname:   req.Lastname,
		BankId:     req.BankId,
		Balance:    req.Balance,
	}
	// dbTransfer:=&tmodel.Transaction{
	// 	Transaction_id: "00001",
	// 	From_account:   317,
	// 	To_account:     318,
	// 	Amount:         100,
	// }
	res,err:=CustomerService.CreateCustomer(dbProfile)
	
	if err != nil {
		return nil, err
	}else {
		responseProfile := &c.DetailResponse{
			CustomerId: res.CustomerId,
			CreatedAt: res.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		return responseProfile, nil
	}

}