package render

import (
	"github.com/go-gl/gl/all-core/gl"
)

type VertexBuffer struct {
	handle uint32
	count  int32
}

func NewVertexBuffer(values []float32) *VertexBuffer {

	var buffer uint32
	gl.GenBuffers(1, &buffer)

	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(values)*sizeOfFloat32, gl.Ptr(values), gl.DYNAMIC_DRAW)

	return &VertexBuffer{handle: buffer, count: int32(len(values))}
}

func (v *VertexBuffer) Update(values []float32) {
	v.Bind()
	gl.BufferData(gl.ARRAY_BUFFER, len(values)*sizeOfFloat32, gl.Ptr(values), gl.DYNAMIC_DRAW)
}

func (v *VertexBuffer) GetCount() int32 {
	return v.count
}

func (v *VertexBuffer) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, v.handle)
}

func (v *VertexBuffer) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

type VertexBufferElement struct {
	count int32
}

type VertexBufferLayout struct {
	elements []VertexBufferElement
}

func NewVertexBufferLayout() *VertexBufferLayout {
	return &VertexBufferLayout{
		elements: []VertexBufferElement{},
	}
}

func (l *VertexBufferLayout) getStride() int32 {
	size := int32(0)
	for _, e := range l.elements {
		size += e.getSize()
	}
	return size
}

func (l *VertexBufferLayout) AddLayoutFloats(floatCount int32) *VertexBufferLayout {
	l.elements = append(l.elements, VertexBufferElement{count: floatCount})
	return l
}

func (e *VertexBufferElement) getSize() int32 {
	return e.count * sizeOfFloat32
}

func (e *VertexBufferElement) getCount() int32 {
	return e.count
}
