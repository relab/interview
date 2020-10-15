# Creating a Blockchain

## Table of Contents

- [Lab 1: Creating a Blockchain](#lab-1-creating-a-blockchain)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Part 1](#part-1)
    - [Blocks](#blocks)
    - [Blockchain](#blockchain)
  - [Part 2](#part-2)
    - [Merkle Tree](#merkle-tree)
  - [Part 3](#part-3)
    - [Command line client](#command-line-client)
    - [Merkle tree benchmarks](#merkle-tree-benchmarks)
  - [Lab Approval](#lab-approval)

## Introduction

The main objective of this lab assignment is to build a simplified blockchain.
A blockchain is basically a distributed database of records.
What makes it unique is that it’s uses cryptographic hash functions to create a
tamper-proof mechanism of committed transactions through distributed consensus.
Most blockchains are permissionless, which means that they allow public membership of nodes,
often implemented on top of a peer-to-peer network, allowing a public distributed
database, i.e. everyone who uses the database has a full or partial copy of it.
A new record can be added only after the consensus between the other keepers of the database.
Also, it’s the blockchain that make crypto-currencies and smart contracts possible.

This lab consist of three parts. Each part will be explained in more detail in
their own sections.

1. **The chain of blocks:** Implement a chain of blocks as an ordered, back-linked list data structure.
   Use the provided skeleton code and unit tests.

2. **Efficient transactions and blocks verification:** Implement a efficient way
   to verify membership of certain transactions in a block using [Merkle Trees](https://en.wikipedia.org/wiki/Merkle_tree).
   Use the provided skeleton code and unit tests.

3. **Command line client and benchmarks**

For each part of the assignment you should copy your implementation of the previous part. But **do not copy the tests**, they can differ from each part, copy only your implementation. If you prefer, you can create a new branch for each part of the assignment.

## Part 1

### Blocks

In blockchain it’s the blocks that store valuable information.
For example, Bitcoin blocks store [transactions](https://en.bitcoin.it/wiki/Transaction), the essence of any crypto-currency.
Besides this, a block contains some technical information, like its version, current timestamp and the hash of the previous block.
In this assignment we will not implement the block as it’s described in current deployed blockchains or Bitcoin specifications, instead we’ll use a simplified version of it, which contains only significant information for learning purposes. Our block definition is defined in the `block.go` file and has the following structure:

```go
type Block struct {
  Timestamp     int64
  Transactions  []*Transaction
  PrevBlockHash []byte
  Hash          []byte
}

type Transaction struct {
  Data []byte
}
```

_Timestamp_ is the current timestamp (when the block is created), _Transactions_ is the actual valuable information containing in the block, _PrevBlockHash_ stores the hash of the previous block, and _Hash_ is the hash of the block.
In Bitcoin specification _Timestamp_, _PrevBlockHash_, and _Hash_ are block headers, which form a separate data structure, and _Transactions_ is a separate data structure (for now, our transaction is only a Two-dimensional slice of bytes contain the data to be stored). You can read more about how transactions are implemented [here](https://bitcoin.org/en/transactions-guide#introduction).

Each block is linked to the previous one using a hash function.
This linking of block hashes is what makes the blockchain secure.
The thing is that calculating a hash is a computationally difficult operation, it takes some time even on fast computers.
This is an intentional architectural design of blockchain systems, which makes adding new blocks difficult, thus preventing their modification after they’re added.

For now, you will just take block fields (i.e. headers), concatenate them, and calculate a SHA-256 hash on the concatenated combination. To do that, use the `SetHash` function. Feed the `PrevBlockHash`, `Transactions`, and `Timestamp` into the hash in this order.

To compute the SHA-256 checksum of the data you can use the [Sum256](https://golang.org/pkg/crypto/sha256/#Sum256) function from the go crypto package. The function receives as input parameters a byte slice, so you will need to convert each field of the header to bytes and then concatenate it. To convert the `timestamp`, you can use the function `IntToHex` in the file `utils.go`, this function uses the package [binary](https://golang.org/pkg/encoding/binary/) from the go standard library to write the binary representation of some data.

We also want all transactions in a block to be uniquely identified by a single hash.
To achieve this, you will get each transaction, concatenate them, and get a hash of the concatenated combination.
This hashing mechanism of providing unique representation of data will be given by the `HashTransactions` function, that will take all transactions of a block and return the hash of it to be included in the block _Hash_.

### Blockchain

Now let’s implement a blockchain.
In its essence blockchain is just a database with certain structure: it’s an ordered, back-linked list.
Which means that blocks are stored in the insertion order and that each block is linked to the previous one.
This structure allows to quickly get the latest block in a chain and to get a block by its hash.

In Golang this structure can be implemented by using an array and a map: the array would keep ordered hashes (arrays are ordered in Go), and the map would keep hash to block pairs (maps are unordered).
But for now, in your blockchain prototype you just need to use an array as shown below.

```go
type Blockchain struct {
  blocks []*Block
}
```

As every block need to be linked to the previous one, to add a new block we need an existing block, but there’re no blocks in the blockchain on the beginning.
So, in any blockchain there must be at least one block, and such block is the first in the chain and is called genesis block.

Your task is to implement all functions marked with `TODO(student)` in the file `blockchain.go`.
These functions are:

- `CreateBlockchain`: This function should create a new blockchain initializing a Genesis block with the
   hardcoded data `genesisCoinbaseData`.
   You can use the function `NewGenesisBlock` of the `block.go` to create the Genesis block.
- `AddBlock`: This function should get the previous block hash and add a new block linking it to the previous.
- `GetGenesisBlock`: This function should return the Genesis block.
- `CurrentBlock`: This function should return the last block.
- `GetBlock`: This function should return a block based on its hash.

## Part 2

### Merkle Tree

Until now we have used hashing as a mechanism for providing a unique representation of data, which give to us
an easy way to verify data integrity, i.e. if any of the transaction data in a block changes, the root hash will change, and tampering is identified.
We did that in the function `HashTransactions` in the `block.go` file, by getting each transaction in a block, concatenate them in a specific order and applied SHA256 to the concatenated combination.
Besides uniquely identifying all the transactions in a block with a single hash, for efficiency, we also want to be able to easily verify if some transaction is in the block without requiring to have all the block transactions.

[Merkle trees](https://xlinux.nist.gov/dads/HTML/MerkleTree.html) are used by [Bitcoin](https://bitcoin.org/bitcoin.pdf) to obtain transactions hash, which is then saved in block headers and is considered by the proof-of-work system.
The benefit of Merkle trees is that a node can verify membership of certain transaction without downloading the whole block, just using the transaction hash, the root hash of the merkle tree, and a set of intermediate hashes necessary to reconstruct the Merkle path for that transaction, which is know as merkle proof.
The Merkle path is simply the set of hashes from the transaction at the leaf node to the Merkle root.
A Merkle proof is a way of proving that a certain transaction is part of a merkle tree without requiring any of the other transactions to be exposed, just the hashes.
Each hash in the proof is the sibling of the hash in the path at the same level in the tree.

This optimization mechanism is crucial for the successful adoption of Bitcoin or any [permissionless blockchain](https://eprint.iacr.org/2017/375.pdf).
For example, the full Bitcoin database (i.e., blockchain) currently takes [more than 230 Gb of disk space](https://www.blockchain.com/charts/blocks-size).
Because of the decentralized nature of Bitcoin, every node in the network must be independent and self-sufficient, i.e. every node in the network must store a full copy of the blockchain.
With many people using Bitcoin, this rule becomes more difficult to follow: it’s not likely that everyone will run a full node.
However, fully-fledged participants of the network, must verify transactions and blocks.
Thus, nodes must interact with other nodes to download new blocks.

The above mechanism also enables SPV (Simple Payment Verification) in Bitcoin, allowing the creation of "light clients" that only store block headers (which includes the Merkle root) to perform transaction inclusion verifications.
Thus a light client doesn’t verify blocks and transactions, instead, it finds transactions in blocks (to verify payments) and maintain a connection with a full node to retrieve just necessary data.
This mechanism allows having multiple light nodes with running just one full node, but can also impose some centralization, since incentive less nodes to maintain the state consistency of the database.

A Merkle tree is built for each block, and it starts with leaves (the bottom of the tree), where a leaf is a transaction hash (Bitcoin uses double SHA256 hashing).
In a [Perfect Binary Merkle Tree](https://xlinux.nist.gov/dads/HTML/perfectBinaryTree.html), as shown in [Figure 1](#pmtree), every interior node has two children and all leaves have the same depth, but not every block contains an even number of transactions.
In case there is an odd number of transactions, the hash of the last transaction is duplicated (in the [Tree](https://github.com/bitcoin/bitcoin/blob/d0f81a96d9c158a9226dc946bdd61d48c4d42959/src/consensus/merkle.cpp#L8), not in the block!) to form a [Full Binary Merkle Tree](https://xlinux.nist.gov/dads//HTML/fullBinaryTree.html), in which every node has either 0 or 2 children.
This is shown in [Figure 2](#fmtree), where the nodes `23AF` and `5101` were duplicated during the process of build the tree.

![Perfect Binary Merkle Tree][pmtree]

Moving from the bottom up, leaves are grouped in pairs, their hashes are concatenated, and a new hash is obtained from the concatenated hashes.
The SHA256 hash is represented by the arrows in the figure.
The new hashes form new tree nodes.
This process is repeated until there’s just one node, which is called the root of the tree.
The root hash is then used as the unique representation of the transactions, is saved in block headers, and is used in the proof-of-work system.

Considering the example in [Figure 1](#pmtree).
The numbers inside the nodes represent the first 4 bytes of the hash of the transaction of that node.
Only leaf nodes store hash of real transactions, the internal nodes store the hash of its children.
The merkle path from the transaction `TX3` to the root hash `38C4` is shown by the _yellow nodes_ in [Figure 1](#pmtree).

The _blue nodes_ shows the set of the intermediate nodes (i.e, merkle proof) that can be used as proof to recreate the merkle path from the `TX3` to the root.
Thus, given `TX3`, the root hash `38C4` and the respective _blue nodes_: `D2B8`, `64B0` and `4A3B`, in this order and altogether with their respective orientations on the tree (i.e, left or right side), it's possible to show that `TX3` exists in the tree by hashing it with the intermediate nodes until reaching the same root.
The same logic can be applied for [Figure 2](#fmtree).

![Full Binary Merkle Tree][fmtree]

Thus, your task is to develop a Binary Merkle Tree by implementing all functions marked with `TODO(student)` in the `merkle_tree.go` file and change the function `HashTransactions` in the `block.go` to use it.
These functions are:

- `HashTransactions`: This function need to be changed in the `block.go` to take in consideration a merkle root instead of just the hash of all transactions.
- `NewMerkleTree`: This function should create a new Merkle tree from a sequence of data by using the `NewMerkleNode` function.
- `NewMerkleNode`: This function should create a node in the merkle tree, the node can be a leaf node, which store the hash of the data or a inner node, which is a hash of its children.
- `MerklePath`: This function should return a list of a nodes' hashes and indices (nodes' positions: left or right) required to reconstruct the inclusion proof of a given hash.
- `VerifyProof`: This function verifies the inclusion of a hash in the merkle tree by taking a hash and its merkle path and reconstructing the root hash of the merkle tree.

Remember to copy your implementation for the first part, but not the tests. The tests in `block_test.go`, for example, are different from the ones on the first part, since it now takes the merkle root hash into consideration.

For more information about the concept of Merkle Trees, and the [Bitcoin implementation](https://en.bitcoin.it/wiki/Protocol_documentation#Merkle_Trees) and its difference with the Ethereum implementation, please read [this](https://blog.ethereum.org/2015/11/15/merkling-in-ethereum/?source=post_page) article.

[pmtree]: perfect-merkle-tree.png "Figure 1"
[fmtree]: full-merkle-tree.png "Figure 2"

## Part 3

### Command line client

Until now we don't have any interface to interact with the blockchain.
Your task on this part is to create a command-line client application that will interact with your blockchain implementation.

You are free to choose any package of your preference to do this task, making a loop and reading the commands from the standard input, or using libraries like [promptui](https://github.com/manifoldco/promptui), it's up to you.

Independent of your choice, the command-line application should offer the following commands:

- createblockchain: Creates a blockchain initializing it with the genesis block.
- addtransaction: Adds a transaction data (e.g. an input string) to a buffer but does not create a block.
- addblock: Adds a block to the blockchain committing the transactions in the buffer.
- printchain: Prints all the blocks of the blockchain.

### Merkle tree benchmarks

Extend the merkle tree test file with two benchmarks.
Check this post for a [tutorial](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go) on writing benchmarks in Go.

- Write a benchmark function for the creation of Merkle trees.
  Benchmark the creation of a Merkle tree with 100, 1000 and 10000 data items.
  Each data item should be unique, but of constant size.
- Write a benchmark function for the validation of merkle proofs.
  Benchmark how long it takes to validate a Merkle proof from trees including 100, 1000 and 10000 data item.
