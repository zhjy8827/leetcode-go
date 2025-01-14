package main

import "fmt"

/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


解题思路：
//TODO

时间效率O()
空间效率O()
总结：
*/

//思路1-------------------------------------------

func convert(s string, numRows int) string {
	list := [][]int{}
	for i := 0; i < numRows; i++ {
		list = append(list, []int{})
	}
	for _, v := range s {

	}
}

//他人优秀代码----------------------------------
// 按行
func convert1(s string, numRows int) string {
	if s == "" || len(s) <= numRows || numRows < 2 {
		return s
	}
	result := make([]uint8, len(s))
	v := 0
	for i := 0; i < numRows; i++ { // i 为行数
		for k := i; k < len(s); k += (numRows - 1) * 2 { // k为要取的原字符串中字符的索引号
			result[v] = s[k]
			v++
			if i != 0 && i != numRows-1 {
				next := k + (numRows-i-1)*2
				if next < len(s) {
					result[v] = s[next]
					v++
				}
			}
		}

	}

	return string(result)
}

//主函数-------------------------------------------

func main() {
	s := "LEETCODEISHIRING"
	numRows := 4
	fmt.Println(convert(s, numRows))
}
