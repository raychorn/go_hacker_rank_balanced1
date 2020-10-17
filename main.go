package main

import (
	"fmt"
	"strconv"
	"strings"
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

func (n *node) asString() string {
	return fmt.Sprintf("{parent:%d, size:%d, visited:%t}", n.parent, n.size, n.visited)
}

func dumpNodesArray(nodes []node) string {
	output := "{"
	n := len(nodes)
	for i, node := range nodes {
		output += node.asString()
		if (i < n-1) {
			output += ", "
		}
	}
	output += "}"
	return output
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

func traverseToRoot(ptr int32, nodes []node) (string, int32) {
	branch := fmt.Sprintf("%d ", ptr)
	aNode := nodes[ptr]
	score := aNode.size
	for aNode.parent != -1 {
		fmt.Printf("node[%d] ? ", ptr)
		branch += fmt.Sprintf("%d ", ptr)
		score += aNode.size
		ptr = aNode.parent
		aNode = nodes[ptr]
	}
	if (aNode.parent == -1) {
		fmt.Printf("node[%d] ? ", ptr)
		branch += fmt.Sprintf("%d ", ptr)
		score += aNode.size
	}
	fmt.Printf("\nDEBUG: branch=%s, score=%d\n", branch, score)
	return branch, score
}

func getNodePtrsFromSubBranch(branch string) []int32 {
	toks := strings.Fields(branch)
	ptrs := make([]int32, len(toks))
	for i, tok := range toks {
		x, _ := strconv.ParseInt(tok, 10, 32)
		ptrs[i] = int32(x)
	}
	return ptrs
}

func generateTrees(nodes []node) map[string]int32 {
	trees := make(map[string]int32)
	fmt.Println("DEBUG: BEGIN")
	for i, node := range nodes {
		fmt.Printf("nodes[%d] -> %s\n", i, node.asString())
	}
	fmt.Println("DEBUG: END!!!")
	var aNode node
	branch := ""
	for j := len(nodes)-1; j >= 0; j-- {
		aNode = nodes[j]
		if (aNode.parent != -1) && (aNode.visited == false) {
			br, score := traverseToRoot(int32(j), nodes)
			ptrs := getNodePtrsFromSubBranch(br)
			fmt.Printf("br=%s, ptrs=%s\n", br, dumpArray(ptrs))
			for k := 0; k < len(ptrs)-1; k++ {
				nodes[ptrs[k]].visited = true
				fmt.Printf("nodes[%d](%d).visited=%t\n", ptrs[k], k, nodes[ptrs[k]].visited)
				//nodes[ptrs[k]] = nodes[ptrs[k]]
			}
			branch += br
			fmt.Printf("** branch: %s", branch)
			trees[branch] = score
			branch = ""
		}
	}
	return trees
}

func mostBalancedPartition1(parent []int32, filesSize []int32) (int32, string) {
    tree := make([]node, len(parent))
    for i := 0; i < len(parent); i++ {
        tree[i] = node{parent: parent[i], size: filesSize[i], visited: false}
	}
	output := "{"
	for _, t := range tree {
		output += t.asString() + ", "
	}
	output += "}"
	
	trees := generateTrees(tree)

	fmt.Printf("BEGIN: trees\n")
	for k,v := range trees {
		fmt.Printf("%s -> %d\n", k, v)
	}
	fmt.Printf("END!!!\n")
    return 0, output
}

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
					trees = append(trees, tree)
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
						fmt.Println("break")
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

	fmt.Printf("BEGIN: trees\n")
	for ii,vv := range trees {
		fmt.Printf("%d -> %s (%d)\n", ii, dumpTree(vv), scores[ii])
	}
	fmt.Printf("END!!!\n")
	fmt.Println()

	return 0, output
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
