package serializer

import (
	"konggo/model"
	"konggo/util"
)

// Example 样例序列化器
type Example struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"` //创建时间
}

// BuildExample 序列化样例
func BuildExample(example model.Example) Example {
	return Example{
		ID:        example.ID,
		Name:      example.Name,
		Age:       example.Age,
		CreatedAt: example.CreatedAt.Local().Format(util.TimeNormalFormat),
	}
}

// GetExampleList 样例列表序列化器
type GetExampleList struct {
	Total int       `json:"total"`
	List  []Example `json:"list"`
}

// BuildExample 序列化样例列表
func BuildExampleList(examples []model.Example, total int) GetExampleList {
	var res []Example
	for _, one := range examples {
		example := BuildExample(one)
		res = append(res, example)
	}
	return GetExampleList{
		Total: total,
		List:  res,
	}
}
