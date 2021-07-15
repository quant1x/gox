package main

import (
	"xorm.io/xorm"
	"github.com/mymmsc/gox/aspect/annotation/transaction"
	"github.com/mymmsc/gox/aspect/example"
)

// go build -gcflags=-l main.go
// main
func main() {
	scanPath := `/Users/wangfeng/projects/mymmsc/gox/aspect/example`
	transaction.NewTransactionManager(transaction.TransactionConfig{ScanPath: scanPath}).RegisterDao(new(example.ExampleDao))

	dao := new(example.ExampleDao)
	dao.Select()
	dao.Update(new(xorm.Session), "") // auto commit
	dao.Delete(new(xorm.Session)) // handle fail and auto rollback
}
