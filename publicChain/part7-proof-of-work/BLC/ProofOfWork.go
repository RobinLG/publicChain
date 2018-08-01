package BLC

import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
)

// 比特币的工作量证明不是以下的算法与验证，此版用0来表示难度，可理解为简化版工作量证明

// 0000 0000 0000 0000 1001 0001 0000 ... 0001

// 256位Hash前面至少要有16个零
const targetBit  = 24

type ProofOfWork struct {
	Block *Block // 当前要验证的区块
	target *big.Int // 大数据存储,区块难度
}

//数据拼接，返回字节数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height)),
		},
		[]byte{},
	)

	return data
}

func (proofOfWork *ProofOfWork) run() ([]byte, int64) {

	// 1. 将Block的属性拼接成字节数组

	// 2. 生成Hash

	// 3. 判断Hash有效性，如果满足条件，跳出循环

	nonce := 0
	var hashInt big.Int // 存储我们新生成的Hash
	var hash [32]byte

	for {
		// 准备数据
		dataBytes := proofOfWork.prepareData(nonce)

		// 生成Hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)

		// 将Hash存储到hasInt
		hashInt.SetBytes(hash[:])

		// 判断hashInt是否小于Block里面的target
		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//   proofOfWork.target == x
		//	 hashInt == y
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}

		nonce = nonce + 1
	}

	return hash[:],int64(nonce)
}

// 创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {

	// 1. big.Int对象1
	// 2. targetBit = 2
	// 3. 0000 0001
	// 4. 8 - 2 = 6
	// 5. 0100 0000 64
	// 6. 0010 0000 找到的值小于等于32(0010 0000)即满足

	// 1.创建一个初始值为1的target
	target := big.NewInt(1)

	// 2.左移256 - target
	target = target.Lsh(target, 256 - targetBit)

	return &ProofOfWork{block, target}
}