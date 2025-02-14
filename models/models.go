package models

type Student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Test struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Question struct {
	Id       string `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"Answer"`
	TestId  string `json:"test_id"`
}
