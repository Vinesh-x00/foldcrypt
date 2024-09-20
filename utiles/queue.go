package utiles

type DirQueue []string

func (Q *DirQueue) Push(x string) {
	*Q = append(*Q, x)
}

func (Q *DirQueue) Pop() string {
	h := *Q
	var el string
	l := len(h)
	el, *Q = h[0], h[1:l]
	return el
}

func (Q *DirQueue) IsEmpty() bool {
	return len(*Q) == 0
}

func NewQueue() *DirQueue {
	return &DirQueue{}
}
