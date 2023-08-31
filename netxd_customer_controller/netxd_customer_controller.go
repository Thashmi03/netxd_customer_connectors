package netxdcustomercontroller

import (
	netxddalinterface "github.com/Thashmi03/netxd_dal/netxd_dal_interface"
	netxddalmodels "github.com/Thashmi03/netxd_dal/netxd_dal_models"
	"context"

	c "github.com/Thashmi03/netxd_customer"
)
type RPCServer struct{
	c.UnimplementedCustomerServiceServer
}

var(
	CustomerService netxddalinterface.ICustomer
)

func(s *RPCServer)CreateCustomer(ctx context.Context,req * c.Details)(*c.DetailResponse,error){
	dbProfile:=&netxddalmodels.Customer{
		CustomerId: req.CustomerId,
		Firstname:  req.Firstname,
		Lastname:   req.Lastname,
		BankId:     req.BankId,
		Balance:    req.Balance,
	}
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