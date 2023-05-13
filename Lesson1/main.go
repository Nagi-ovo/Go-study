package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type user struct {
	name     string
	password string
}

type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

func findUsers(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

// 函数
func add(a int, b int) string {
	sum := a + b
	result := fmt.Sprintf("a+b的值是%d", sum)
	return result
}

func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func add2(n int) {
	n += 2
}

func add2ptr(n *int) {
	*n += 2
}

func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Getenv("PATH"))
	fmt.Println(os.Setenv("AA", "BB"))

	n1, err := strconv.Atoi(("AAA"))
	fmt.Println(n1, err)
	n2, err := strconv.Atoi("123")
	fmt.Println(n2, err)

	// // 创建时间对象并输出
	// now := time.Now()
	// fmt.Println("当前时间：", now)

	// // 使用时间常量并输出
	// delay := 3 * time.Second
	// fmt.Printf("延迟 %s 后执行任务...\n", delay)

	// // 时间延迟
	// time.Sleep(delay)
	// fmt.Println("任务执行中...")

	// // 时间格式化
	// t1 := time.Date(2022, time.May, 3, 10, 34, 0, 0, time.Local)
	// fmt.Println("时间格式化：", t1.Format("2006-01-02 15:04:05"))

	// // 时间定时器
	// fmt.Println("启动定时器...")
	// timer := time.NewTimer(5 * time.Second)
	// <-timer.C
	// fmt.Println("定时器时间到...")

	dateString := "2022-07-28"
	fmt.Println(time.Parse("2006-01-02", dateString))

	timestamp := int64(1659729852)
	fmt.Println(time.Unix(timestamp, 0))
	now := time.Now()
	fmt.Println(now.Unix())

	t2 := time.Date(2022, time.August, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println(t2.Format("2006-01-02 15:04:05"))

	user_A := userInfo{Name: "wang", Age: 18, Hobby: []string{"Golang", "Typescript"}}
	buf, err := json.Marshal(user_A)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf)         // [123 34 78 97...]
	fmt.Println(string(buf)) // {"Name":"wang","age":18,"Hobby":["Golang","Typescript"]}

	buf, err = json.MarshalIndent(user_A, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
	/*{
		"Name": "wang",
		"age": 18,
		"Hobby": [
				"Golang",
				"Typescript"
		]
	}*/

	var user_B userInfo
	err = json.Unmarshal(buf, &user_B)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", user_B) // main.userInfo{Name:"wang", Age:18, Hobby:[]string{"Golang", "Typescript"}}

	str_a := "hellolL"
	fmt.Println(strings.Contains(str_a, "ll"))            // true
	fmt.Println(strings.Count(str_a, "l"))                // 4
	fmt.Println(strings.HasPrefix(str_a, "he"))           // 判断字符串 str_a 是否以 he 开头
	fmt.Println(strings.HasSuffix(str_a, "llo"))          // 判断字符串 str_a 是否以 llo 结尾
	fmt.Println(strings.Index(str_a, "ll"))               // 2
	fmt.Println(strings.Join([]string{"he", "llo"}, "-")) //  he-llo
	fmt.Println(strings.Repeat(str_a, 2))                 // hellollhelloll
	fmt.Println(strings.Replace(str_a, "l", "L", -1))     // heLLoLL   "-1"表示替换旧字符串次数，这里是替换所有
	fmt.Println(strings.Split("a-b--c", "-"))             // [a b  c]
	fmt.Println(strings.ToLower(str_a))                   // helloll
	fmt.Println(strings.ToTitle(str_a))                   // HELLOLL
	fmt.Println(len(str_a))                               // 7

	user_a := user{name: "wang", password: "1024"}
	user_b := user{"wang1", "1024"}
	user_c := user{name: "wang"}
	fmt.Println(user_a, user_b, user_c)
	user_c.password = "1024"
	d := user{} // var d user
	d.name = "wang"
	d.password = "1024"
	fmt.Println(d)
	fmt.Println(&user_a)

	//结构体方法
	user_test := user{name: "zhang", password: "4321"}
	fmt.Println("original password:", user_test.password)
	user_test.resetPassword("1234")
	fmt.Println("new password:", user_test.password)
	fmt.Println(user_test.checkPassword("1234"))

	struct_ptr_A := &user_a
	fmt.Println("name:", struct_ptr_A.name)
	fmt.Println("name:", user_b.name)
	struct_ptr_A.password = "1234"
	fmt.Println("修改后的A用户信息:", user_a)

	//错误处理
	u, err := findUsers([]user{{"wang", "1024"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name)

	if u, err := findUsers([]user{{"wang", "1024"}}, "li"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(u.name)
	}

	ptr := 5
	add2(ptr)
	fmt.Println(ptr)
	add2ptr(&ptr)
	fmt.Println(ptr)
	testMap := map[string]string{"hello": "Hello", "world": "World"}
	fmt.Println(exists(testMap, "hello"))
	fmt.Println(exists(testMap, "hi"))
	fmt.Println(exists(testMap, "world"))
	fmt.Println("Hello world")
	for j := 7; j < 9; j++ { //任何一段都可以省略
		fmt.Println(j)
	}
	for {
		fmt.Println("死循环")
		break
	}
	t := time.Now() // case可以使用任意变量类型，包括字符串和结构体
	switch {        //可以取代部分if else功能
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's afternoon")
	}
	var a [5]int //存放5个int的数组
	a[4] = 100
	fmt.Println(a[4], len(a))

	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s[1]) // b

	s = append(s, "d", "e")
	fmt.Println(s) // [a b c d e]

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c) // [a b c d e]

	fmt.Println(s[1:3]) // [b c]

	st := make([]string, 3)
	st[0] = "a"
	st[1] = "b"
	fmt.Println(st[1]) // b

	st = append(st, "c", "d")
	fmt.Println(st) // [a b  c d] //由于切片长度为3但只赋2个值，故空出一位
	fmt.Println(st[1:3])

	st[2] = "c"
	ct := make([]string, len(st))
	copy(ct, st)
	fmt.Println(ct) // [a b c d]

	//map
	m := make(map[string]int) //string:key类型  value类型：int
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m, len(m), m["two"], m["three"])

	//删除map的k-v对
	delete(m, "one")
	fmt.Println(m, len(m), m["two"], m["three"])

	//range
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println("index:", i, "num:", num) //index: 0 num: 2  index: 1 num: 3  index: 2 num: 4
	}
	fmt.Println(sum) // 9

	maps := map[string]string{"a": "A", "b": "B"}
	for k, v := range maps {
		fmt.Println(k, v)
	}
	for k := range maps {
		fmt.Println("key", k)
	}
	aa, bb := 1, 2
	fmt.Println(aa, bb)

	fmt.Println(add(1, 2))
}
