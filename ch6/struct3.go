package ch6

import (
	"encoding/json"
	"fmt"
)

// 序列化

type Part struct {
	ID    int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Name  string // json序列化是默认使用字段名作为key
	Count int
	// count int // 私有不能被json包访问
}

type Company struct {
	Title string
	Parts []*Part
}

func init() {
	c := &Company{
		Title: "ali",
		Parts: make([]*Part, 0, 200),
	}

	for i := 0; i < 10; i++ {
		part := &Part{
			Name:  "it",
			Count: 10,
			ID:    i,
		}
		c.Parts = append(c.Parts, part)
	}

	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}

	fmt.Println(string(data)) // {"Title":"ali","Parts":[{"id":0,"Name":"it","Count":10},{"id":1,"Name":"it","Count":10},{"id":2,"Name":"it","Count":10},{"id":3,"Name":"it","Count":10},{"id":4,"Name":"it","Count":10},{"id":5,"Name":"it","Count":10},{"id":6,"Name":"it","Count":10},{"id":7,"Name":"it","Count":10},{"id":8,"Name":"it","Count":10},{"id":9,"Name":"it","Count":10}]}

	// 反序列化
	str := `{"Title":"ali","Parts":[{"ID":0,"Name":"it","Count":10},{"ID":1,"Name":"it","Count":10},{"ID":2,"Name":"it","Count":10},{"ID":3,"Name":"it","Count":10},{"ID":4,"Name":"it","Count":10},{"ID":5,"Name":"it","Count":10},{"ID":6,"Name":"it","Count":10},{"ID":7,"Name":"it","Count":10},{"ID":8,"Name":"it","Count":10},{"ID":9,"Name":"it","Count":10}]}`
	c1 := &Company{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}
	fmt.Println(c1.Parts[1].ID)
}
