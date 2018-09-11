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

func InsertList(n *listNode, l *list) {

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

func NewTreeNode(c byte, frequency uint32, left *treeNode, right *treeNode) *treeNode {
	return &treeNode{c, frequency, left, right}
}

func NewListNode(treeN *treeNode) *listNode {
	return &listNode{treeN, nil}

}

func PopListMin(l *list) *treeNode {
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
			InsertList(NewListNode(NewTreeNode(byte(i), byteCount[i], nil, nil)), &l)
		}
	}

	for l.elements > 1 {
		leftNode := PopListMin(&l)
		rightNode := PopListMin(&l)

		sum := NewTreeNode('#', (leftNode.frequency)+(rightNode.frequency), leftNode, rightNode)

		InsertList(NewListNode(sum), &l)
	}

	return PopListMin(&l)
}
