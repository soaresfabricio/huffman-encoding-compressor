package compressor

type treeNode struct {
	c         byte
	frequency uint32
	left      *treeNode
	right     *treeNode
}

type listNode struct {
	n    *treeNode
	next *listNode
}

type list struct {
	head     *listNode
	elements int
}

func (l *list) InsertList(n *listNode) {

	if l.head == nil {
		l.head = n
	} else if n.n.frequency < l.head.n.frequency {
		n.next = l.head
		l.head = n
	} else {
		aux := (l.head.next)
		aux2 := (l.head)

		for aux != nil && (aux.n.frequency <= n.n.frequency) {
			aux2 = aux
			aux = aux2.next
		}
		aux2.next = n
		n.next = aux
	}
	(l.elements)++

}

func (l *list) PopListMin() *treeNode {
	aux := l.head
	aux2 := aux.n
	l.head = aux.next
	(l.elements)--
	return aux2
}

func BuildHuffmanTree(byteCount []uint32) *treeNode {

	l := list{nil, 0}

	for i := 0; i < 256; i++ {
		if byteCount[i] != 0 {
			l.InsertList(&listNode{&treeNode{byte(i), byteCount[i], nil, nil}, nil})
		}
	}

	for l.elements > 1 {
		leftNode := l.PopListMin()
		rightNode := l.PopListMin()

		sum := &treeNode{'#', (leftNode.frequency) + (rightNode.frequency), leftNode, rightNode}

		l.InsertList(&listNode{sum, nil})
	}

	return l.PopListMin()
}
