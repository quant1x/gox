module github.com/mymmsc/gox

go 1.16

// latest
//replace golang.org/x/sys v0.0.0-20190329044733-9eb1bfa1ce65 => github.com/golang/sys v0.0.0-20190329044733-9eb1bfa1ce65

//require golang.org/x/sys v0.0.0-20190329044733-9eb1bfa1ce65 // indirect

require (
	github.com/go-xorm/xorm v0.7.9
	github.com/handsomestWei/go-annotation v0.0.0-20201013025858-335644dfb2d9
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1
	github.com/satori/go.uuid v1.2.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	rsc.io/qr v0.2.0
)
