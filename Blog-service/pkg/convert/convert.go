package convert

import "strconv"

type StrTo string

// 对返回的HTTP状态码和接口返回的响应结果进行判断
// 返回定义好的响应结果
func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s StrTo) MustInt() int{
	v, _ := s.Int()
	return v
}

func (s StrTo) UInt32()(uint32, error){
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s StrTo) MustUInt32() uint32{
	v, _ := s.UInt32()
	return v
}

