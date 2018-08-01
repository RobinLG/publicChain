package BLC

import "math/big"

// 比特币的工作量证明不是以下的算法与验证，此版用0来表示难度，可理解为简化版工作量证明

// 0000 0000 0000 0000 1001 0001 0000 ... 0001

// 256位Hash前面至少要有16个零
const targetBit  = 16

type ProofOfWork struct {
	Block *Block // 当前要验证的区块
	target *big.Int // 大数据存储,区块难度
}

func (proofOfWork *ProofOfWork) run() ([]byte, int64) {
	return nil, 0
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

	return &ProofOfWork{block, 5}
}