package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'mostBalancedPartition' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY parent
 *  2. INTEGER_ARRAY files_size
 */

type node struct {
    parent int32
	size   int32
	visited bool
}

type nodes []node

func (n *node) asString() string {
	return fmt.Sprintf("{parent:%d, size:%d, visited:%t}", n.parent, n.size, n.visited)
}

func (s nodes) Len() int {
	return len(s)
}

func (s nodes) Less(i, j int) bool {
	return s[i].parent < s[j].parent
}

func (s nodes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func dumpArray(items []int32) string {
	output := "{"
	n := len(items)
	for i, val := range items {
		output += fmt.Sprintf("%d", val)
		if (i < n-1) {
			output += ", "
		}
	}
	output += "}"
	return output
}

func dumpTree(tree map[int32]node) string {
	output := "{"
	n := len(tree)
	i := 0
	for _,v := range tree {
		output += v.asString()
		if (i < n-1) {
			output += ", "
		}
		i++
	}
	output += "}"
	return output
}

func unpackTree(tree map[int32]node) nodes {
	theNodes := nodes{}
	for _,v := range tree {
		theNodes = append(theNodes, v)
	}
	return theNodes
}

/*
func sortTree(tree map[int32]node) []node {

}
*/

func isNodeInAnyTree(trees []map[int32]node, ptr int32) bool {
	hasPtr := false

	for _,tree := range trees {
		_, hasKey := tree[ptr]
		if (hasKey) {
			hasPtr = true
			break
		}
	}
	return hasPtr
}

func pluralize(val interface{}) string {
	plural := ""
	if (len(val) > 1) {
		plural = "s"
	}
	return plural
}

func mostBalancedPartition(parent []int32, filesSize []int32) (int32, string) {
	trees := make([]map[int32]node, 0)
	tree := make(map[int32]node)
	scores := make(map[int]int32)
    for ptr := int32(len(parent)-1); ptr > -1; ptr-- {
		inAny := isNodeInAnyTree(trees, ptr)
		if (inAny) {
			continue
		}
		fmt.Printf("ptr=%d is inAny=%t\n", ptr, inAny)
		_, hasKey := tree[ptr]
		fmt.Printf("current tree hasKey=%t\n", hasKey)
		if (!hasKey) {
			fmt.Println("BEGIN:")
			for iPtr := ptr; iPtr > 0; {
				fmt.Printf("iPtr=%d, parent[%d]=%d\n", iPtr, iPtr, parent[iPtr])
				tree[iPtr] = node{parent: parent[iPtr], size: filesSize[iPtr]}
				_, hasScore := scores[len(trees)]
				if (!hasScore) {
					scores[len(trees)] = 0
				}
				scores[len(trees)] += filesSize[iPtr]
				if (parent[iPtr] == -1) {
					fmt.Printf("(1) Store tree and break. iPtr=%d, parent[%d]=%d\n", iPtr, iPtr, parent[iPtr])
					trees = append(trees, tree)
					tree = make(map[int32]node)
					break
				} else {
					iPtr = parent[iPtr]
					if (iPtr == 0) || (iPtr == -1) {
						tree[iPtr] = node{parent: parent[iPtr], size: filesSize[iPtr]}
						_, hasScore = scores[len(trees)]
						if (!hasScore) {
							scores[len(trees)] = 0
						}
						scores[len(trees)] += filesSize[iPtr]
						trees = append(trees, tree)
						tree = make(map[int32]node)
						fmt.Println("(2) Store tree and break.")
						break
					}
				}
			}
			fmt.Println("END!!!")
		} else {
			fmt.Println("continue")
			continue
		}
	}
	output := "{"
	for k,v := range tree {
		output += fmt.Sprintf("%d -> %s, ", k, v.asString())
	}
	output += "}"

	if (false) {
		fmt.Printf("BEGIN: scores\n")
		for ii,vv := range scores {
			fmt.Printf("%d -> %d\n", ii, vv)
		}
		fmt.Printf("END!!!\n")
		fmt.Println()
	}

	theTrees := []nodes{}
	fmt.Printf("BEGIN: trees\n")
	for ii,vv := range trees {
		fmt.Printf("%d -> %s (%d)\n", ii, dumpTree(vv), scores[ii])
		theTrees = append(theTrees, unpackTree(vv))
	}
	fmt.Printf("END!!!\n")
	fmt.Println()

	response := 0
	if (len(trees) == 1) {
		fmt.Println("Only one tree.")
		response = int(scores[0])
	} else {
		// here we determine which edge to remove in the tree with the largest size to make the trees as equal as possible. Find the min difference.
		plural := pluralize(theTrees)
		fmt.Printf("Consider %d tree%s.", len(theTrees), plural)
		for i,someNodes := range theTrees {
			fmt.Println(someNodes)
			fmt.Printf("(1) #%d Is Sorted %t \n", i, sort.IsSorted(someNodes))
			sort.Sort(someNodes)
			fmt.Println(someNodes)
			fmt.Printf("(2) #%d Is Sorted %t \n", i, sort.IsSorted(someNodes))
			fmt.Println()
		}
	}

	return int32(response), output
}

func main() {
	var datasets = [][][]int32{
		{{-1, 0, 1, 2, -1}, {1, 4, 3, 4, 1}}, // two trees
		{{-1, 0, 1, 2}, {1, 4, 3, 4}},   // one tree
		{{-1, 0, 0, 0}, {10, 11, 10, 10}}, // three trees
		{{-1, 0}, {20, 100}}, 
		{{-1, 0, 0, 0, 0, 3, 4, 6, 0, 3}, {298, 2187, 5054, 266, 1989, 6499, 5450, 2205, 5893, 8095}}, 
		{{-1, 0, 1, 2, 1, 0, 5, 2, 0, 0}, {8475, 6038, 8072, 7298, 5363, 9732, 3786, 5521, 8295, 6186}},
		{{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}, {8618, 5774, 7046, 459, 2279, 2894, 798, 2328, 1011, 2780}},
	}
	
	output := ""
	for i, val := range datasets {
		output += "{"
		for _, val2 := range val {
			output += dumpArray(val2)
		}
		output += "},"
		var result int32
		result = -1
		sBuf := "ERROR"
		if (len(val) == 2) {
			result, sBuf = mostBalancedPartition(val[0], val[1])
		}
		output += fmt.Sprintf(" --> %d -> %s\n", result, sBuf)
		if (i == 0) {
			break
		}
	}
	fmt.Println(output)

}
