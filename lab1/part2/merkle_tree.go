package main

import (
	"fmt"
)

// MerkleTree represents a merkle tree
type MerkleTree struct {
	RootNode *Node
	Leafs    []*Node
}

// Node represents a merkle tree node
type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Hash   []byte
}

const (
	leftNode = iota
	rightNode
)

// MerkleProof represents way to prove element inclusion on the merkle tree
type MerkleProof struct {
	proof [][]byte
	index []int64
}

// NewMerkleTree creates a new Merkle tree from a sequence of data
func NewMerkleTree(data [][]byte) *MerkleTree {
	// TODO(student)
	return nil
}

// NewMerkleNode creates a new Merkle tree node
func NewMerkleNode(left, right *Node, data []byte) *Node {
	// TODO(student)
	return &Node{}
}

// MerkleRootHash return the hash of the merkle root
func (mt *MerkleTree) MerkleRootHash() []byte {
	return mt.RootNode.Hash
}

// MakeMerkleProof returns a list of hashes and indexes required to
// reconstruct the merkle path of a given hash
//
// @param hash represents the hashed data (e.g. transaction ID) stored on
// the leaf node
// @return the merkle proof (list of intermediate hashes), a list of indexes
// indicating the node location in relation with its parent (using the
// constants: leftNode or rightNode), and a possible error.
func (mt *MerkleTree) MakeMerkleProof(hash []byte) ([][]byte, []int64, error) {
	// TODO(student)
	return [][]byte{}, []int64{}, fmt.Errorf("Node %x not found", hash)
}

// VerifyProof verifies that the correct root hash can be retrieved by
// recreating the merkle path for the given hash and merkle proof.
//
// @param rootHash is the hash of the current root of the merkle tree
// @param hash represents the hash of the data (e.g. transaction ID)
// to be verified
// @param mProof is the merkle proof that contains the list of intermediate
// hashes and their location on the tree required to reconstruct
// the merkle path.
func VerifyProof(rootHash []byte, hash []byte, mProof MerkleProof) bool {
	// TODO(student)
	return false
}
