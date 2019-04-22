module github.com/mymmsc/gox

go 1.11.4

// latest
//replace golang.org/x/sys v0.0.0-20190329044733-9eb1bfa1ce65 => github.com/golang/sys v0.0.0-20190329044733-9eb1bfa1ce65

//require golang.org/x/sys v0.0.0-20190329044733-9eb1bfa1ce65 // indirect

require (
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1
	golang.org/x/text v0.3.0
	rsc.io/qr v0.2.0
)
