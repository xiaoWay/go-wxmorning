package service

import (
	"encoding/hex"
	"fmt"
	"math/rand"
)

// 随机颜色函数设置
func randomcolor() string {
	// 生成一个16字节的随机种子
	seed := make([]byte, 16)
	rand.Read(seed)

	// 将随机种子解码为整数值作为随机数生成器的种子
	seedInt, err := hex.DecodeString(hex.EncodeToString(seed))
	if err != nil {
		panic(err)
	}
	source := rand.NewSource(int64(seedInt[0])<<56 | int64(seedInt[1])<<48 | int64(seedInt[2])<<40 |
		int64(seedInt[3])<<32 | int64(seedInt[4])<<24 | int64(seedInt[5])<<16 | int64(seedInt[6])<<8 |
		int64(seedInt[7]))

	// 使用随机数生成器生成三个0到255之间的随机数，并将它们格式化为十六进制颜色值
	random := rand.New(source)
	color := fmt.Sprintf("#%02X%02X%02X", random.Intn(256), random.Intn(256), random.Intn(256))

	return color
}
