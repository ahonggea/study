package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("数据类型")
	fmt.Println("----------整数类型----------")

	// 有符号整数类型
	// 有符号整数类型二进制首位0/1表示正负数
	// 二进制 127 -> 01111111
	// 二进制 -128 ->  10000000
	// 127是怎么推算的
	// 01111111 -> 二进制 -> 十进制：
	//  1*2^6  +  1*2^5  +  1*2^4  +  1*2^3  +  1*2^2  +  1*2^1  +  1*2^0
	// =  64   +    32   +    16   +    8    +    4    +    2    +    1
	// =  127

	// int8 1字节 -2^7~2^7-1 -128~127
	var n1 int8 = 127
	fmt.Println(n1)
	// int16 2字节 -32768~32767
	var n2 int16 = 32767
	fmt.Println(n2)
	// int32 4字节 -2147483648~2147483647
	var n3 int32 = 2147483647
	fmt.Println(n3)
	// int64 8字节 -2^63~2^63-1

	// 无符号整数类型
	// 和有符号整数类型用前面的u区分

	fmt.Println("----------浮点数类型----------")
	var fnum1 float32 = 3.14
	var fnum2 float32 = -3.14
	fmt.Println(fnum1)
	fmt.Println(fnum2)
	// 浮点数可以用十进制表示形式、也可以用科学计数法表示形式 E大小写都可以
	var fnum3 float32 = 314e-2
	var fnum4 float32 = 314e+2
	fmt.Println(fnum3)
	fmt.Println(fnum4)
	// 浮点数可能会有精度的顺势，通常情况下对浮点数数据类型不确定使用float64
	// golang中默认的浮点数类型位：float64
	var fnum5 float32 = 256.0000000916
	var fnum6 float64 = 256.0000000916
	fmt.Println(fnum5)
	fmt.Println(fnum6)
	var fnum7 float64 = 314e+2
	fmt.Println(fnum7)
	fmt.Println("----------字符类型----------")
	var c1 byte = 'a' // 97
	fmt.Println(c1)
	var c2 byte = '6' // 54
	fmt.Println(c2)
	var c3 byte = 'c' // 99
	fmt.Println(c3)
	fmt.Println(c3 - c2)
	// 字符类型本质上就是一个整数，也可以参数与运算，输出字符的时候，会将对应的码值做一个输出
	// 字母、数字、标点等字符，底层是按照ASCII进行存储
	// 汉字字符，底层对应的是unicode码值
	// var c4 byte = "中"
	// fmt.Println(c4)
	fmt.Println("----------布尔类型----------")
	var flag1 bool = true
	var flag2 bool = false
	var flag3 bool = 5 < 9
	println(flag1, flag2, flag3)

	fmt.Println("----------字符串类型----------")
	// 定义一个字符串
	var str1 string = "hello golang"
	fmt.Println(str1)
	// 2字符串不可改变：指的是字符串一旦定义好，其中的字符的值不能改变，而不是不能整体重新赋值
	var str2 string = "abc"
	str2 = "def"
	// str2[0] = "t" //报错
	fmt.Println(str2)
	// 含有特殊符号可用`代替"
	var str3 = `test
	test2`
	fmt.Println(str3)
	// 可使用+拼接字符串
	var str4 = "abc" + "edf"
	str4 += "hij"
	fmt.Println(str4)
	// 字符串太长需要换行拼接时+保持在上面
	var str5 = "abcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabc" +
		"edfedfedfedfedfedfedfedfedfedfedf" +
		"edfedfedfedfedfedfedfedfedfedfedf" +
		"edfedfedfedfedfedfedfedfedfedfedf"
	fmt.Println(str5)

	fmt.Println("----------类型转换----------")
	var nn1 int = 100
	// var nn2 float32 = nn1 // 在这里自动转换不好使，比如显示转换
	var nn2 float32 = float32(nn1)
	fmt.Println(nn1)
	fmt.Println(nn2)

	fmt.Println("----------基本数据转为string[Sprintf]----------")
	var sn1 int = 50
	var sn2 float32 = 4.89
	var sn3 bool = false
	var sn4 byte = 'c'
	var ss1 string = fmt.Sprintf("%d", sn1)
	fmt.Printf("ss1对应的类型是: %T,ss1 = %q\n", ss1, ss1)
	var ss2 string = fmt.Sprintf("%f", sn2)
	fmt.Printf("ss2对应的类型是: %T,ss2 = %q\n", ss2, ss2)
	var ss3 string = fmt.Sprintf("%t", sn3)
	fmt.Printf("ss3对应的类型是: %T,ss3 = %q\n", ss3, ss3)
	var ss4 string = fmt.Sprintf("%c", sn4)
	fmt.Printf("ss4对应的类型是: %T,ss4 = %q\n", ss4, ss4)
	fmt.Println("----------基本数据转为string[strconv]----------")
	var sc1 int = 51
	var scs1 string = strconv.FormatInt(int64(sc1), 10)
	fmt.Printf("scs1对应的类型是: %T,scs1 = %q\n", scs1, scs1)
	var sc2 float64 = 51
	var scs2 string = strconv.FormatFloat(sc2, 'f', 9, 64)
	// 第二个参数：'f' (-ddd.dddd, no exponent), 第三个参数：9 保留小数点后面9位， 第四个参数表示这个小数是float64类型
	fmt.Printf("scs2对应的类型是: %T,scs2 = %q\n", scs2, scs2)
	var sc3 bool = true
	var scs3 string = strconv.FormatBool(sc3)
	fmt.Printf("scs3对应的类型是: %T,scs3 = %q\n", scs3, scs3)
	fmt.Println("----------string转为基本数据----------")
	// string -> bool
	var tc1 string = "true"
	var tt1 bool
	tt1, _ = strconv.ParseBool(tc1) // strconv.ParseBool返回两个值（value bool,err error）,因为在这里不关心error，直接用_忽略
	fmt.Printf("tt1对应的类型是: %T,tt1 = %v\n", tt1, tt1)
	// string -> int64
	var tc2 string = "15564651"
	var tt2 int64
	tt2, _ = strconv.ParseInt(tc2, 10, 64)
	fmt.Printf("tt2对应的类型是: %T,tt2 = %v\n", tt2, tt2)
	// string -> float64/32
	var tc3 string = "3.14"
	var tt3 float64
	tt3, _ = strconv.ParseFloat(tc3, 64)
	fmt.Printf("tt3对应的类型是: %T,tt3 = %v\n", tt3, tt3)
	// 注意： string向基本类型转换的时候，一定要确保string类型能够转成有效的数据类型，否则最后得到的结果就是按照对应类型的默认值输出

}
