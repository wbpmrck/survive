package maps

//表示一个大小
type Size struct {
	width,height uint64
}
func(s *Size) getWidth() uint64{
	return s.width
}
func(s *Size) setWidth(val uint64)  {
	s.width = val
}
func(s *Size) getHeight() uint64{
	return s.height
}
func(s *Size) setHeight(val uint64)  {
	s.height= val
}