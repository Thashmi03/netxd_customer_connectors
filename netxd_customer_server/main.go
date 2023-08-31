package main

import (
	"banking_with_grpc/netxd_customer_connectors/config"
	"banking_with_grpc/netxd_customer_connectors/constants"
	netxdcustomercontroller "banking_with_grpc/netxd_customer_connectors/netxd_customer_controller"
	netxddalservices "banking_with_grpc/netxd_dal/netxd_dal_services"
	"context"
	"fmt"
	"net"

	c "github.com/Thashmi03/netxd_customer"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func intiDatabase(client *mongo.Client){
	customerCollection:=config.GetCollection(client,"BankDatabase","Customer")
	netxdcustomercontroller.CustomerService=netxddalservices.InitCustomerService(customerCollection,context.Background())
}

func main(){
	mongoclient,err:=config.ConnectDatabase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	intiDatabase(mongoclient)
	lis,err:=net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s:=grpc.NewServer()
	c.RegisterCustomerServiceServer(s,&netxdcustomercontroller.RPCServer{})

	fmt.Println("sever listening on",constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}