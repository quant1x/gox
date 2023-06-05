package num

import "math/big"

// Signed is a constraint that permits any signed integer type.
// If future releases of Go add new predeclared signed integer types,
// this constraint will be modified to include them.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
// If future releases of Go add new predeclared unsigned integer types,
// this constraint will be modified to include them.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is a constraint that permits any integer type.
// If future releases of Go add new predeclared integer types,
// this constraint will be modified to include them.
type Integer interface {
	Signed | Unsigned
}

// Float is a constraint that permits any floating-point type.
// If future releases of Go add new predeclared floating-point types,
// this constraint will be modified to include them.
type Float interface {
	~float32 | ~float64
}

// Complex is a constraint that permits any complex numeric type.
// If future releases of Go add new predeclared complex numeric types,
// this constraint will be modified to include them.
type Complex interface {
	~complex64 | ~complex128
}

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
type Ordered interface {
	Integer | Float | ~string
}

// NumberOfCPUBitsRelated The number of CPU bits is related
// Deprecated: 不推荐使用
type NumberOfCPUBitsRelated interface {
	~int | ~uint | ~uintptr
}

// /*nil, */ int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64 , bool, string
// ~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64 | ~int | ~uint | ~float32 | ~float64 | ~bool | ~string
// uintptr

// BaseType 基础类型
type BaseType interface {
	Integer | Float | ~string | ~bool
}

// GenericType Series支持的所有类型
// Deprecated: 不推荐使用
type GenericType interface {
	~bool | ~int32 | ~int64 | ~int | ~float32 | ~float64 | ~string
}

// StatType 可以统计的类型
// Deprecated: 不推荐使用
type StatType interface {
	~int32 | ~int64 | ~float32 | ~float64
}

type BigFloat = big.Float // 预留将来可能扩展float

// Deprecated: 不推荐使用
type Number8 interface {
	~int8 | ~uint8
}

// Deprecated: 不推荐使用
type Number16 interface {
	~int16 | ~uint16
}

// Deprecated: 不推荐使用
type Number32 interface {
	~int32 | ~uint32 | float32
}

// Deprecated: 不推荐使用
type Number64 interface {
	~int64 | ~uint64 | float64 | int | uint
}

// Deprecated: 已弃用
type MoveType interface {
	StatType | ~bool | ~string
}

// Number int和uint的长度取决于CPU是多少位
type Number interface {
	Integer | Float
}
