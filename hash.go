package hash

import (
	"fmt"
)

/*
 * RS  从Robert Sedgwicks的 Algorithms in C一书中得到了。
 * 我(原文作者)已经添加了一些简单的优化的算法，以加快其散列过程。
 */
func RSHash(str string) uint64 {
	b := 378551
	a := 63689
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = hash*uint64(a) + uint64(str[i])
		a = a * b
	}
	fmt.Printf("RSHash %v\n", hash)
	return hash
}

/*
 *	Justin Sobel写的一个位操作的哈希函数。
 */

func JSHash(str string) uint64 {
	hash := uint64(1315423911)
	for i := 0; i < len(str); i++ {
		hash ^= ((hash << 5) + uint64(str[i]) + (hash >> 2))
	}
	fmt.Printf("JSHash %v\n", hash)
	return hash
}

//该散列算法是基于贝尔实验室的彼得J温伯格的的研究。在Compilers一书中（原则，技术和工具），建议采用这个算法的散列函数的哈希方法。

func PJWHash(str string) uint64 {
	BitsInUnsignedInt := (uint64)(4 * 8)
	ThreeQuarters := (uint64)((BitsInUnsignedInt * 3) / 4)
	OneEighth := (uint64)(BitsInUnsignedInt / 8)
	HighBits := (uint64)(0xFFFFFFFF) << (BitsInUnsignedInt - OneEighth)
	hash := uint64(0)
	test := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = (hash << OneEighth) + uint64(str[i])
		if test = hash & HighBits; test != 0 {
			hash = ((hash ^ (test >> ThreeQuarters)) & (^HighBits))
		}
	}
	fmt.Printf("PJWHash %v\n", hash)
	return hash
}

/*
这个算法来自Brian Kernighan 和 Dennis Ritchie的 The C Programming Language。
这是一个很简单的哈希算法,使用了一系列奇怪的数字,形式如31,3131,31...31,看上去和DJB
算法很相似。(参照我之前一篇博客，这个就是Java的字符串哈希函数)
*/
func BKDRHash(str string) uint64 {
	seed := uint64(131) // 31 131 1313 13131 131313 etc..
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + uint64(str[i])
	}
	fmt.Printf("BKDRHash %v\n", hash)
	return hash
}

//这个算法在开源的SDBM中使用，似乎对很多不同类型的数据都能得到不错的分布。

func SDBMHash(str string) uint64 {
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = uint64(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	fmt.Printf("SDBMHash %v\n", hash)
	return hash
}

//这个算法是Daniel J.Bernstein 教授发明的，是目前公布的最有效的哈希函数。

func DJBHash(str string) uint64 {
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = ((hash << 5) + hash) + uint64(str[i])
	}
	fmt.Printf("DJBHash %v\n", hash)
	return hash
}

// 由伟大的Knuth在《编程的艺术 第三卷》的第六章排序和搜索中给出。

func DEKHash(str string) uint64 {
	hash := uint64(len(str))
	for i := 0; i < len(str); i++ {
		hash = ((hash << 5) ^ (hash >> 27)) ^ uint64(str[i])
	}
	fmt.Printf("DEKHash %v\n", hash)
	return hash
}

/* AP【推荐使用********，其他几个在处理连续数据转换的字符串时结果也是连续的，但是这个是分散的】
 * 这是本文作者Arash Partow贡献的一个哈希函数，继承了上面以旋转以为和加操作
 */

func APHash(str string) uint64 {
	hash := uint64(0xAAAAAAAA)
	for i := 0; i < len(str); i++ {
		if (i & 1) == 0 {
			hash ^= ((hash << 7) ^ uint64(str[i])*(hash>>3))
		} else {
			hash ^= (^((hash << 11) + uint64(str[i]) ^ (hash >> 5)))
		}
	}
	fmt.Printf("APHash %v\n", hash)
	return hash
}
