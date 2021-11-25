package main

import (
	"fmt"
)

// TODO: Cleaning consts into group variables ()
// Sets Default values which dynamicly changes as Struct properties
const Capacity_Def int = 12
const Length_Def int = 12

// Sets Default values which are immutable
const LoadFactorMax_Def float32 = 0.7
const LoadFactorMin_Def float32 = 0.25

type Node struct {
	Key   int         // sets Key which will be stored and hashed in array
	Value interface{} // Any value TODO: normalize the value into asci
	Next  *Node       // Linked list for chaining
}

// Constructor for creating Node
func NewNode(k int, v interface{}) *Node {
	return &Node{Key: k, Value: v}
}

type HashTable struct {
	Capacity     int     // Sets the capacity of array for calculate hashFunc
	CurrentNodes int     // Stores actual count of nodes
	Nodes        []*Node // Array of Nodes
}

// Constructor for creating Hash Table
// making empty slice with initilize default values of consts at package level
func NewHashTable() *HashTable {
	slice := make([]*Node, Length_Def, Capacity_Def)
	return &HashTable{Capacity_Def, 0, slice}
}

// Accept any value and key as a integer
func (h *HashTable) Insert(v interface{}, k int) {
	hashed := h.HashFunction(k)                         // hashed any key into our universe of possible keys -> default is 12
	node := NewNode(k, v)                               // create node with to store unhashed key and value
	if foundNode := h.Nodes[hashed]; foundNode != nil { // if is already stored any node with the same hashed value TODO: Colission is when two nodes has same hashed value but not same key hashFunc(ki) == hashFunc(kj) && k != j
	HERE:
		for foundNode != nil { //Colission -> solve -> by chaining with linked list structure
			if foundNode.Next == nil {
				foundNode.Next = node
				break HERE
			} else {
				foundNode = foundNode.Next
			}
		}
	} else {
		h.CurrentNodes += 1       // increase counter of nodes
		h.Nodes[hashed] = node    // actual store the node
		h.CheckLoadFactorUpdate() // always check if we should shrink or expand (Table-Doubling) the array
	}
}

// Delete by given key
func (h *HashTable) Delete(k int) {
	hashed := h.HashFunction(k) // hashed the key to find the specific node
	if isFound := h.Nodes[hashed]; isFound != nil {
		h.Nodes[hashed] = nil     // sets the node struct as a nil
		h.CurrentNodes -= 1       // decrease counter
		h.CheckLoadFactorUpdate() // always check if we should shrink or expand (Table-Doubling) the array
	}
}

func (h *HashTable) CheckLoadFactorUpdate() {
	// Table Doubling
	loadFactor := h.CalcLoadFactor()
	fmt.Println("[x] LoadFactor [x] ==>", loadFactor)
	if loadFactor >= LoadFactorMax_Def {
		h.Capacity = h.Capacity * 2
		newListNodes := make([]*Node, h.Capacity, h.Capacity)
		for _, v := range h.Nodes {
			if v != nil {
				hashed := h.HashFunction(v.Key)
				newListNodes[hashed] = v
			}
		}
		h.Nodes = newListNodes
	} else if loadFactor <= LoadFactorMin_Def && h.Capacity != Capacity_Def {
		h.Capacity = h.Capacity / 2
		newListNodes := make([]*Node, h.Capacity)
		for _, v := range h.Nodes {
			if v != nil {
				hashed := h.HashFunction((*v).Key)
				newListNodes[hashed] = v
			}
		}
		h.Nodes = newListNodes
	}
}

func (h *HashTable) HashFunction(k int) int {
	return (k ^ 2) % h.Capacity
}

func (h *HashTable) CalcLoadFactor() float32 {
	return float32(h.CurrentNodes) / float32(len(h.Nodes))
}

func main() {
	hashTable := NewHashTable()
	hashTable.Display(hashTable)
}

func (h *HashTable) Display(hashTable *HashTable) {
	for k, v := range hashTable.Nodes {
		fmt.Println("[x] HASHED KEY [x] ==>", k)
		fmt.Println("[x]    VALUE   [x] ==>", v)
	}
}
