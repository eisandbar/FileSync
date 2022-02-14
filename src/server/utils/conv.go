package utils

func Atoi (s string) int64{
	res := int64(0)
	for _, c := range s {
		i := byte(c)
		if i < byte('0') || i > byte('9') {
			return int64(-1)
		}
		res = res * 10 + int64(i - byte('0'))
	}
	return res
}