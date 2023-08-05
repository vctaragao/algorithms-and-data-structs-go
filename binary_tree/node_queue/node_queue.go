package node_queue

import "github.com/vctaragao/algorithms-and-data-structs-go/queue"

type NodeQueue queue.Queue

func (q *NodeQueue) Pop() {
	q.Pop()
}
