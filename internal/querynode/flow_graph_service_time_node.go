package querynode

import (
	"log"
)

type serviceTimeNode struct {
	BaseNode
	replica collectionReplica
}

func (stNode *serviceTimeNode) Name() string {
	return "stNode"
}

func (stNode *serviceTimeNode) Operate(in []*Msg) []*Msg {
	//fmt.Println("Do serviceTimeNode operation")

	if len(in) != 1 {
		log.Println("Invalid operate message input in serviceTimeNode, input length = ", len(in))
		// TODO: add error handling
	}

	serviceTimeMsg, ok := (*in[0]).(*serviceTimeMsg)
	if !ok {
		log.Println("type assertion failed for serviceTimeMsg")
		// TODO: add error handling
	}

	// update service time
	stNode.replica.getTSafe().set(serviceTimeMsg.timeRange.timestampMax)
	//fmt.Println("update tSafe to:", getPhysicalTime(serviceTimeMsg.timeRange.timestampMax))
	return nil
}

func newServiceTimeNode(replica collectionReplica) *serviceTimeNode {
	maxQueueLength := Params.FlowGraphMaxQueueLength
	maxParallelism := Params.FlowGraphMaxParallelism

	baseNode := BaseNode{}
	baseNode.SetMaxQueueLength(maxQueueLength)
	baseNode.SetMaxParallelism(maxParallelism)

	return &serviceTimeNode{
		BaseNode: baseNode,
		replica:  replica,
	}
}
