/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
	"github.ibm.com/mbg-agent/cmd/mbg/state"
	"google.golang.org/grpc"

	pb "github.ibm.com/mbg-agent/pkg/protocol"
)

/// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A start command set all parameter state of the Multi-cloud Border Gateway",
	Long: `A start command set all parameter state of the gateway-
			1) The MBG that the gateway is connected
			2) The IP of the gateway
			TBD now is done manually need to call some external `,
	Run: func(cmd *cobra.Command, args []string) {
		ip, _ := cmd.Flags().GetString("ip")
		id, _ := cmd.Flags().GetString("id")
		cport, _ := cmd.Flags().GetString("cport")
		exposePortRange, _ := cmd.Flags().GetString("exposePortRange")

		if ip == "" || id == "" || cport == "" {
			log.Println("Error: please insert all flag arguments for Mbg start command")
			os.Exit(1)
		}
		state.SetState(id, ip, cport, exposePortRange)
		startServer()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().String("id", "", "Multi-cloud Border Gateway id")
	startCmd.Flags().String("ip", "", "Multi-cloud Border Gateway ip")
	startCmd.Flags().String("cport", "", "Multi-cloud Border Gateway control port")
	startCmd.Flags().String("exposePortRange", "30000", " set the start port for exposing range ")

}

const (
	serverIp = ":50051"
)

/******* Commands **********/
//Expose
type ExposeServer struct {
	pb.UnimplementedExposeServer
}

func (s *ExposeServer) ExposeCmd(ctx context.Context, in *pb.ExposeRequest) (*pb.ExposeReply, error) {
	log.Printf("Received: %v", in.GetId())
	state.UpdateState()
	if in.GetDomain() == "Internal" {
		state.AddLocalService(in.GetId(), in.GetIp(), in.GetDomain())
		ExposeToMbg(in.GetId())
	} else { //Got the service from MBG so expose to local Gw
		state.AddRemoteService(in.GetId(), in.GetIp(), in.GetDomain(), in.GetMbgID())
		ExposeToGw(in.GetId())
	}
	return &pb.ExposeReply{Message: "Done"}, nil
}

//Hello
type HelloServer struct {
	pb.UnimplementedHelloServer
}

func (s *HelloServer) HelloCmd(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received Hello from MBG ip: %v", in.GetIp())
	state.UpdateState()
	state.AddMbgNbr(in.GetId(), in.GetIp(), in.GetCport())

	return &pb.HelloReply{Message: "MBG: " + state.GetMyIp() + " get hello message"}, nil
}

//Connect
type ConnectServer struct {
	pb.UnimplementedConnectServer
}

func (s *ConnectServer) ConnectCmd(ctx context.Context, in *pb.ConnectRequest) (*pb.ConnectReply, error) {
	log.Printf("Received Connect request from service: %v to service: %v", in.GetId(), in.GetIdDest())
	state.UpdateState()
	//svc := state.GetService(in.GetID())
	var listenPort, destIp string
	if state.IsServiceLocal(in.GetIdDest()) {
		destSvc := state.GetLocalService(in.GetIdDest())
		listenPort = destSvc.ListenPort
		destIp = destSvc.Service.Ip
	} else { //For Remtote service
		destSvc := state.GetRemoteService(in.GetIdDest())
		mbgIP := state.GetServiceMbgIp(destSvc.Service.Ip)
		SendConnectReq(in.GetId(), in.GetIdDest(), in.GetPolicy(), mbgIP)
		listenPort = destSvc.ListenPort
		destIp = destSvc.Service.Ip
	}

	go ConnectService(listenPort, destIp, in.GetPolicy())
	log.Printf("Send connect reply to service")
	return &pb.ConnectReply{Message: "Connecting the services"}, nil
}

/********************************** Server **********************************************************/
func startServer() {
	log.Printf("Gateway [%v] started", state.GetMyId())
	lis, err := net.Listen("tcp", serverIp)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterExposeServer(s, &ExposeServer{})
	pb.RegisterConnectServer(s, &ConnectServer{})
	pb.RegisterHelloServer(s, &HelloServer{})

	log.Printf("Control channel listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}