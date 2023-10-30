package policytypes

type Direction int

const (
	Incoming Direction = iota
	Outgoing
)

type ConnectionRequest struct {
	SrcPeer string
	DstPeer string // ignored on an outgoing connection request

	SrcWorkloadAttrs WorkloadAttrs
	DstServiceName   string

	Direction Direction
}

type ConnectionResponse struct {
	Action  PolicyAction
	DstPeer string
}
