package ch5

import "fmt"

// KeyType:表示键的类型。
// ValueType:表示键对应的值的类型。
// map[KeyType]ValueType

// map是引用类型

func init() {
	map1 := make(map[string]int)

	map1["age"] = 10
	map1["height"] = 178
	map1["weight"] = 140

	userInfo := map[string]string{
		"name": "go",
		"pwd":  "123",
	}

	fmt.Println(map1, userInfo)

	// 判断某个键是否存在
	value, ok := userInfo["name1"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("查无此人")
	}

	// 遍历

	for i := 0; i < len(userInfo); i++ {
		fmt.Println(i)

	}

	for key, v := range userInfo {
		fmt.Println(key, v)
	}

	// 切片类型map
	userInfos := map[string][]string{
		"name": {"js", "go"},
		"age":  {"1", "2"},
	}

	fmt.Println(userInfos)

	// 删除
	delete(userInfo, "age")
}
