// Package heap provides immutable heap operations.  Operations such as
// Push and Pop will re-arrange the heap, but return a new slice or
// container for the heap.
package heap

func leftChild (node uint) uint {
    return 2*node + 1;
}

func rightChild (node uint) uint {
    return 2*node + 2;
}

func parent (node uint) uint {
    return (node - 1)/2
}

// Sift the start of the slice to the its proper position
// using comparisons for a max-heap
func siftDown (theArray []int, root uint) {
    for {
	lchild := leftChild(root)
	if lchild >= uint(len(theArray)) {
	    break
	}
	who2swap := root
	if theArray[who2swap] < theArray[lchild] {
	    who2swap = lchild
	}

	rchild := lchild + 1
	if rchild < uint(len(theArray)) && theArray[who2swap] < theArray[rchild] {
	    who2swap = rchild
	}

	if who2swap == root {
	    break;
	}
	
	theArray[who2swap],theArray[root] = theArray[root], theArray[who2swap]
	root = who2swap
    }
}

// Heapify re-arranges the elements of an []int slice to form a max-heap structure
func Heapify (theArray []int) {
    n := uint(len(theArray))
    if n < 2 {
	return
    }
    n--;
    for start := int(parent(n)) ; start >= 0; start-- {
	siftDown(theArray, uint(start))
    }
}

