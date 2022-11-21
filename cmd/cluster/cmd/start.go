/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.ibm.com/mbg-agent/cmd/cluster/state"

	"context"
	"log"
	"net"

	pb "github.ibm.com/mbg-agent/pkg/protocol"
	"google.golang.org/grpc"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A start command set all parameter state of the cluster",
	Long: `A start command set all parameter state of the cluster-
			1) The MBG that the cluster is connected
			2) The IP of the cluster
			TBD now is done manually need to call some external `,
	Run: func(cmd *cobra.Command, args []string) {
		ip, _ := cmd.Flags().GetString("ip")
		id, _ := cmd.Flags().GetString("id")
		mbgIP, _ := cmd.Flags().GetString("mbgIP")
		cport, _ := cmd.Flags().GetString("cport")

		state.SetState(ip, id, mbgIP, cport)
		startServer()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().String("id", "", "Cluster Id")
	startCmd.Flags().String("ip", "", "Cluster IP")
	startCmd.Flags().String("mbgIP", "", "IP address of the MBG connected to the Cluster")
	startCmd.Flags().String("cport", "", "Cluster control port")
}

/******* Commands **********/
//Expose
func (s *ExposeServer) ExposeCmd(ctx context.Context, in *pb.ExposeRequest) (*pb.ExposeReply, error) {
	log.Printf("Received: %v", in.GetId())
	state.UpdateState()
	state.AddService(in.GetId(), in.GetIp(), in.GetDomain())
	return &pb.ExposeReply{Message: "Done"}, nil
}

type ExposeServer struct {
	pb.UnimplementedExposeServer
}

//Connect
type ConnectServer struct {
	pb.UnimplementedConnectServer
}

func (s *ConnectServer) connectCmd(ctx context.Context, in *pb.ConnectRequest) (*pb.ConnectReply, error) {
	log.Printf("Received Connect request from service: %v to service: %v", in.GetId(), in.GetIdDest())
	state.UpdateState()
	svc := state.GetService(in.GetId())
	destSvc := state.GetService(in.GetIdDest())

	connectClient(svc.Service.Ip, destSvc.Service.Ip)

	return &pb.ConnectReply{Message: "Connecting the services"}, nil
}

/********************************** Server **********************************************************/
func startServer() {
	log.Printf("Cluster [%v] started", state.GetId())
	clusterCPort := state.GetIP() + ":" + state.GetCport()
	lis, err := net.Listen("tcp", clusterCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExposeServer(s, &ExposeServer{})
	pb.RegisterConnectServer(s, &ConnectServer{})
	log.Printf("Control channel listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}