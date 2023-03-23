package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	api "github.ibm.com/mbg-agent/pkg/api"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get",
	Long:  `Get`,
	Run:   emptyRun,
}

// updateCmd represents the update command
var peerGetCmd = &cobra.Command{
	Use:   "peer",
	Short: "get peer list from the MBG",
	Long:  `get peer list from the MBG`,
	Run: func(cmd *cobra.Command, args []string) {
		mId, _ := cmd.Flags().GetString("myid")
		peerId, _ := cmd.Flags().GetString("id")
		m := api.Mbgctl{mId}
		if peerId == "" {
			pArr, err := m.GetPeers()
			if err != nil {
				fmt.Printf("Unable to get peers: %v", err)
				return
			}
			fmt.Printf("Peers :%+v\n", pArr)
		} else {
			m.GetPeer(peerId)
		}

	},
}

var serviceGetCmd = &cobra.Command{
	Use:   "service",
	Short: "get service list from the MBG",
	Long:  `get service list from the MBG`,
	Run: func(cmd *cobra.Command, args []string) {
		mId, _ := cmd.Flags().GetString("myid")
		serviceId, _ := cmd.Flags().GetString("id")
		serviceType, _ := cmd.Flags().GetString("type")
		m := api.Mbgctl{mId}
		i := 1
		if serviceId == "" {
			if serviceType == "local" {
				sArr, err := m.GetLocalServices()
				if err != nil {
					fmt.Printf("Unable to get local services: %v", err)
					return
				}
				fmt.Printf("Local services:\n")
				for _, s := range sArr {
					fmt.Printf("%d) Service ID: %s IP: %s Description: %s\n", i, s.Id, s.Ip, s.Description)
					i++
				}
			} else {
				sArr, err := m.GetRemoteServices()
				if err != nil {
					fmt.Printf("Unable to get remote services: %v\n", err)
					return
				}
				fmt.Printf("Remote Services:\n")
				for _, sA := range sArr {
					for _, s := range sA {
						fmt.Printf("%d) Service ID: %s Local IP: %s Remote MBGID: %s Description: %s\n", i, s.Id, s.Ip, s.MbgID, s.Description)
						i++
					}
				}
			}
		} else {
			if serviceType == "local" {
				s, err := m.GetLocalService(serviceId)
				if err != nil {
					fmt.Printf("Unable to get local service: %v\n", err)
					return
				}
				fmt.Printf("Local Service :%+v", s)
			} else {
				sArr, err := m.GetRemoteService(serviceId)
				if err != nil {
					fmt.Printf("Unable to get remote service: %v\n", err)
					return
				}
				for _, s := range sArr {
					fmt.Printf("Remote service ID: %s with Local IP: %s Remote MBGID %s Description %s \n", s.Id, s.Ip, s.MbgID, s.Description)
				}
			}
		}
	},
}

var policyGetCmd = &cobra.Command{
	Use:   "policy",
	Short: "Get policy list from the MBG",
	Long:  `Get policy list from the MBG (ACL and LB)`,
	Run: func(cmd *cobra.Command, args []string) {
		mId, _ := cmd.Flags().GetString("myid")
		m := api.Mbgctl{mId}

		rules, err := m.GetACLPolicies()
		if err != nil {
			fmt.Printf("Unable to get ACL Policies %v\n", err)
		} else {
			fmt.Printf("MBG ACL rules\n")
			for r, v := range rules {
				fmt.Printf("[Match]: %v [Action]: %v\n", r, v)
			}
		}

		lPolicies, err := m.GetLBPolicies()
		if err != nil {
			fmt.Printf("Unable to Get LB Policies %v\n", err)
		} else {
			fmt.Printf("MBG Load-balancing policies\n")
			for d, val := range lPolicies {
				for s, p := range val {
					fmt.Printf("ServiceSrc: %v ServiceDst: %v Policy: %v\n", s, d, p)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	// Get Peer
	getCmd.AddCommand(peerGetCmd)
	peerGetCmd.Flags().String("myid", "", "MBGCtl Id")
	peerGetCmd.Flags().String("id", "", "Peer id field")
	// Get Service
	getCmd.AddCommand(serviceGetCmd)
	serviceGetCmd.Flags().String("myid", "", "MBGCtl Id")
	serviceGetCmd.Flags().String("id", "", "service id field")
	serviceGetCmd.Flags().String("type", "remote", "service type : remote/local")
	// Get policy
	getCmd.AddCommand(policyGetCmd)
	policyGetCmd.Flags().String("myid", "", "MBGCtl Id")
}
