package main

import (
	"fmt"
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
    parent int32;
    size   int32;
}

func (n *node) asString() string {
	return fmt.Sprintf("{parent:%d, size:%d }", n.parent, n.size)
}

func mostBalancedPartition(parent []int32, filesSize []int32) (int32, string) {
    tree := make([]node, len(parent))
    for i := 0; i < len(parent); i++ {
        tree[i] = node{parent: parent[i], size: filesSize[i]}
	}
	output := "{"
	for _, t := range tree {
		output += t.asString() + ", "
	}
	output += "}"
    
    return 0, output
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

func main() {
	var datasets = [][][]int32{
		{{-1, 0, 1, 2}, {1, 4, 3, 4}}, 
		{{-1, 0, 0, 0}, {10, 11, 10, 10}}, 
		{{-1, 0}, {20, 100}}, 
		{{-1, 0, 0, 0, 0, 3, 4, 6, 0, 3}, {298, 2187, 5054, 266, 1989, 6499, 5450, 2205, 5893, 8095}}, 
		{{-1, 0, 1, 2, 1, 0, 5, 2, 0, 0}, {8475, 6038, 8072, 7298, 5363, 9732, 3786, 5521, 8295, 6186}},
		{{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}, {8618, 5774, 7046, 459, 2279, 2894, 798, 2328, 1011, 2780}},
	}
	
	output := ""
	for _, val := range datasets {
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
	}
	fmt.Println(output)

}
