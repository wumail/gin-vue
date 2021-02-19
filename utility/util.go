package utility

import (
	"math/rand"
	"time"
)

//RandomName 随机名字
func RandomName(n int) string {
	name := make([]byte, n)
	var randName = []byte("dasdasd1231dvooWEUQOFHoiweuqwdSCUPPuiioeq")
	rand.Seed(time.Now().Unix())
	for i := range name {
		name[i] = randName[rand.Intn(len(randName))]
	}
	return string(name)
}
