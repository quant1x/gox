package main

import (
	"github.com/mymmsc/gox/annotation/example"
	"github.com/mymmsc/gox/annotation/transaction"
	"xorm.io/xorm"
)

// go build -gcflags=-l main.go
// main
func main() {
	scanPath := `/Users/wangfeng/projects/mymmsc/gox/annotation/example`
	transaction.NewTransactionManager(transaction.TransactionConfig{ScanPath: scanPath}).RegisterDao(new(example.ExampleDao))

	dao := new(example.ExampleDao)
	dao.Select()
	dao.Update(new(xorm.Session), "") // auto commit
	dao.Delete(new(xorm.Session))     // handle fail and auto rollback
}
