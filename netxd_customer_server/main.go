package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Thashmi03/netxd_customer_connectors/config"
	"github.com/Thashmi03/netxd_customer_connectors/constants"
	netxdcustomercontroller "github.com/Thashmi03/netxd_customer_connectors/netxd_customer_controller"
	netxddalservices "github.com/Thashmi03/netxd_dal/netxd_dal_services"
	tcontroller "github.com/Thashmi03/netxd_customer_connectors/transfer_cotroller"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	c "github.com/Thashmi03/netxd_customer"
	
)



func intiDatabase(client *mongo.Client){
	customerCollection:=config.GetCollection(client,"BankDatabase","Customer")
	netxdcustomercontroller.CustomerService=netxddalservices.InitCustomerService(customerCollection,context.Background())
}
func intitransfer(client *mongo.Client){
	transferCollection:=config.GetCollection(client,"BankDatabase","Transfer")
	customerCollection:=config.GetCollection(client,"BankDatabase","Customer")
	netxdcustomercontroller.TransferService=netxddalservices.InitTransaction(customerCollection,transferCollection,context.Background(),client)
}

func main(){
	mongoclient,err:=config.ConnectDatabase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	intiDatabase(mongoclient)
	intitransfer(mongoclient)
	lis,err:=net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s:=grpc.NewServer()
	c.RegisterCustomerServiceServer(s,&netxdcustomercontroller.RPCServer{})
	c.RegisterCustomerServiceServer(s,&tcontroller.RPServer{})
	fmt.Println("sever listening on",constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}