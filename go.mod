module github.com/mymmsc/gox

go 1.16

// latest
//replace golang.org/x/sys v0.0.0-20190329044733-9eb1bfa1ce65 => github.com/golang/sys v0.0.0-20190329044733-9eb1bfa1ce65

//require golang.org/x/sys v0.0.0-20190329044733-9eb1bfa1ce65 // indirect

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1
	github.com/stretchr/testify v1.4.0
	golang.org/x/text v0.3.3
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
	rsc.io/qr v0.2.0
	xorm.io/builder v0.3.9 // indirect
	xorm.io/xorm v1.1.2
)
