package gmath

type ComparableType interface {
	NumberType | string
}

type NumberType interface {
	IntType | FloatType
}

type FloatType interface {
	float32 | float64
}

type IntType interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}
