package game

type TriangleIndexBufferGenerator struct {
	values []int32
}

func NewTriangleIndexBufferGenerator() *TriangleIndexBufferGenerator {
	return &TriangleIndexBufferGenerator{}
}

func (t *TriangleIndexBufferGenerator) Generate(count int) []int32 {

	if len(t.values) == 6*count {
		return t.values
	}

	t.values = []int32{}

	for i := int32(0); i < int32(count); i++ {
		t.values = append(t.values, i*4, i*4+1, i*4+2, i*4, i*4+3, i*4+2)
	}

	return t.values
}
