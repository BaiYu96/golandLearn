// 这个是在testgodoc.go文件的内容
// package Name: test
//  describe:    程序的入口
//
// ---testgodoc.go end---
package test

import (
	"fmt"
)

// 说明: Person 用于定义个人信息的参数
//	Name		：人名
//	Phone		：电话
//	Address		：住址
//	Email		：邮箱
type Person struct {
	Name    string
	Phone   string
	Address string
	Email   string
}
// 定义了Baiyu这个变量名
const Baiyu  = "baiyu"
// 定义了一组基本信息
const (
	tel = 1101191201234  //你会发现这个变量看不到
	//你会发现tel这个变量看不到
	Email = "xxx@163.com"
	Address = "米花町2段21号"
)
// 说明: 定义了Computer
//  CPU		: CPU牌子
//  Mouse		: 鼠标
//  ProductKey	: 公钥
//  DeviceName	: DeviceName是设备名称
type Computer struct {
	CPU        string
	Mouse      string
	ProductKey string
	DeviceName string
}

// 这一行注释时不会显示的，因为没有紧跟着关键字

// 功能: 测试函数
//  参  数:
//     	param1	: 说明参数值
//     	inter	: 运算的数字
//  返回值:
//     	k 	: 返回值
func Test(param1 string, inter int) (k string) {
	fmt.Println("测试函数")
	return param1
}

// 这个smallTest你会发现他是没办法在go doc里面出现的
// 因为函数名的首字符没有大写，根据go语言的可见性原则，这个函数对外不可见的
func smallTest(param1 string, inter int) (k string) {
	fmt.Println("测试函数small")
	return param1
}
// 功能: 函数H的注释
//  参  数:
//     	h1	: 说明参数值
//     	h2	: 运算的数字
//  返回值:
//     	reValue 	: 返回值

//你会发现函数H的功能注释没有显示，这是因为功能函数注释与关键字func之间空了一行
func H(h1 string, h2 string) (reValue string) {
	fmt.Println(h1, h2)
	return h1
}
// BUG(Z): 这个注释会生成在文档最后后面，同时因为紧跟着fun Z，所以在上面的func列表里面也有显示
// 功能: 生成Z签名
func Z() () {
	fmt.Println()
}
// BUG(who): 因为前面有BUG(who)这个关键字，所以这句注释就算没有紧跟关键字不会被隐藏掉

// BUG(): 这一句的BUG(who)关键字书写不正确，不显示

// BUG6: 这一句的BUG(who)关键字书写不正确，不显示

// BUG: 这一句的BUG(who)关键字书写不正确，不显示

// BUG(6): BUG(6):格式正确，所以这句注释就算没有紧跟关键字不会被隐藏掉，前面的BUG():、BUG6:、BUG:都是不正确的BUG(who)命名
