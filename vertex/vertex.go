package syrup

import (
	"waffle"
)

type SyrupVertex struct {
	waffle.VertexBase
}

func (v *SyrupVertex) SendMessageToAllNeighbors(m waffle.Msg) {
	for _, e := range v.OutEdges() {
		v.SendMessageTo(e.Target(), m)
	}
}

type StreamVertex struct {
	waffle.VertexBase
}

func (v *StreamVertex) Compute(msg []waffle.Msg) {
}