package cgo

import (
	"log"
	"main/response"
	"main/vo"

	"github.com/gin-gonic/gin"
)

var (
	Maxsize int
)

type HashStruct struct {
	Count int
	Empty int
	Key   int
}

var HT []HashStruct

//hash
func Hash(key int) int {
	return key%Maxsize - 1
}

//初始化
func Inithash(c *gin.Context) {
	var requestHash vo.CreateHashPost
	if err := c.Bind(&requestHash); err != nil {
		log.Println(err.Error())
		response.Fail(c, nil, "创建初始化Hash请求失败")
		return
	}
	Maxsize = requestHash.Maxsize
	HashStruct := make([]HashStruct, Maxsize)
	HT = HashStruct
}

//增
func Inserthash(c *gin.Context) {
	var requestHash vo.RUDHashPost
	if err := c.Bind(&requestHash); err != nil {
		log.Println(err.Error())
		response.Fail(c, nil, "创建插入Hash请求失败")
		return
	}
	num := 0
	key := requestHash.Key
	addr := Hash(key)
	HT[addr].Count++
	for HT[addr].Empty == 0 {
		//处理冲突，线性再散列探查法
		addr++
		if addr == Maxsize+1 { //至表尾，从头开始
			addr = 1
		}
		if num == Maxsize { //整个表比较一遍还未找到
			break
		}
		num++ //比较次数加 1
	}
	if HT[addr].Empty == 1 {
		HT[addr].Empty = 0
		HT[addr].Key = key
		response.Success(c, gin.H{"conflict": num}, "插入(保存)成功")
		//存储记录
	} else {
		response.Fail(c, nil, "Hash表已满")
		//Hash 表已满
	}
}

//删
func Delethash(c *gin.Context) {
	var requestHash vo.RUDHashPost
	if err := c.Bind(&requestHash); err != nil {
		log.Println(err.Error())
		response.Fail(c, nil, "创建删除Hash请求失败")
		return
	}
	key := requestHash.Key
	pos := SearchHash(key)
	if pos != -1 { //找到记录
		HT[pos].Empty = 1
		HT[pos].Count--
		response.Success(c, nil, "删除成功")
	} else {
		response.Fail(c, nil, "删除失败")
	}

}

//单查
func SearchHash(key int) int { //找到返回记录所在位置，未找到返回-1

	A := 1 //查找次数
	addr := Hash(key)
	count := HT[addr].Count
	if count == 0 { //无记录
		return -1
	}

	for {
		if HT[addr].Empty == 1 { //记录空闲
			addr++
			if addr == Maxsize+1 { //至表尾，从头开始
				addr = 1
			}
			if A == Maxsize { //整个表比较一遍还未找到
				return -1
			}
			A++ //比较次数加 1
		} else if HT[addr].Empty == 0 { //记录非空闲
			if HT[addr].Key == key {
				return addr //找到位置
			} else {
				if Hash(HT[addr].Key) == Hash(key) {
					count--
					if count == 0 {
						return -1 //无记录
					} else {
						//取下一条记录
						addr++
						if addr == Maxsize+1 { //至表尾，从头开始
							addr = 1
						}
						if A == Maxsize { //整个表比较一遍还未找到
							return -1
						}
						A++ //比较次数加 1
					}
				} else {
					//取下一条记录
					addr++
					if addr == Maxsize+1 { //至表尾，从头开始
						addr = 1
					}
					if A == Maxsize { //整个表比较一遍还未找到
						return -1
					}
					A++ //比较次数加 1
				}
			}
		}
	}
}

func Searchhash(c *gin.Context) {
	var requestHash vo.RUDHashPost
	if err := c.Bind(&requestHash); err != nil {
		log.Println(err.Error())
		response.Fail(c, nil, "创建查找Hash请求失败")
		return
	}
	key := requestHash.Key
	pos := SearchHash(key)
	if pos != -1 { //找到记录
		response.Success(c, gin.H{"data": pos}, "查找成功")
	} else {
		response.Fail(c, nil, "查找失败")
	}
}

//全查
func Showhash(c *gin.Context) {
	if HT != nil {
		response.Success(c, gin.H{"data": HT}, "显示成功")
	} else {
		response.Fail(c, nil, "表为空")
	}
}
