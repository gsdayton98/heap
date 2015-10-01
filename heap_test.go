// Test heap operations.
package heap

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// verifyHeap iterates through each element of theArray to verify it is the maximum element of its subtree
func verifyMaxHeap (theArray []int) bool {
    isHeap := true // innocent until proven otherwise
    for node := 0; 2*node+1 < len(theArray); node++ {
	if theArray[node] < theArray[2*node+1] {
	    isHeap = false
	    break
	}
	right := 2*node + 2
	if right < len(theArray) && theArray[node] < theArray[right] {
	    isHeap = false
	    break
	}
    }
    return isHeap
}

// TestHeapify exercises the Heapify function in delupes.
func TestHeapify (t *testing.T) {
    var testCases = []struct {
	    sample   []int
    }{
	    {[]int{3, 2, 1, 4}, },       // normal no duplicates
	    {[]int{1, 2, 3, 4}, },       // normal no duplicates
	    {[]int{}, },                 // empty input
	    {[]int{1}, },                // single element input
	    {[]int{2, 1}, },             // two element input no dups
	    {[]int{1, 2}, },             // two element input no dups
	    {[]int{2, 2}, },             // two element input all dups
	    {[]int{1, 2, 3}, },          // three element input no dups
	    {[]int{1, 3, 2}, },          // three element input no dups
	    {[]int{2, 1, 3}, },          // three element input no dups
	    {[]int{2, 3, 1}, },          // three element input no dups
	    {[]int{2, 1, 3}, },          // three element input no dups
	    {[]int{3, 1, 2}, },          // three element input no dups
	    {[]int{3, 2, 1, 4, 5, 6}, }, // even number no dups
	    {[]int{3, 3, 1, 4, 5, 6}, }, // even number 1 dup beginning
	    {[]int{3, 2, 1, 4, 6, 6}, }, // even number 1 dup end
	    {[]int{3, 1, 1, 1, 6, 6}, }, // even number 1 dup end
    }

    for _, testCase := range testCases {
	    Heapify(testCase.sample)

	    assert.True(t, verifyMaxHeap(testCase.sample), "Expected %v each element to be >= children", testCase)
    }
}

