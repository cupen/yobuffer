package types

func SizeOf(_type string) int {
	switch _type {
	case "int8":
		return 1
	case "int16":
		return 2
	case "int32":
		return 4
	case "int64":
		return 8
	case "uint8":
		return 1
	case "uint16":
		return 2
	case "uint32":
		return 4
	case "uint64":
		return 8
	case "bool":
		return 1
	case "float32":
		return 4
	case "float64":
		return 8
	case "string":
		return 0
	default:
		return -1
	}
}
