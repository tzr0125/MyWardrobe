package models

// Item 只保存了一件物品的最基础信息
type Item struct {
	Uid uint64
	Type string
	Position string 
}

type Tag struct {
	Type string
	Value string
}