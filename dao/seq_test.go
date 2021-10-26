package dao

import (
	"fmt"
	"git-knowledge/db"
	"testing"
)

func InitSeqDao() SeqDao {
	resource, err := db.NewResource("127.0.0.1", "27017", "app", "root", "root123")
	if err != nil {
		panic(err)
	}
	return NewSeqDao(resource)
}

func TestSeqDaoImpl_GenUserId(t *testing.T) {
	dao := InitSeqDao()
	id, err := dao.GenUserId()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
