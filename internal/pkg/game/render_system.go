package game

//
//type Renderer interface {
//	Render()
//	Remove(id uuid.UUID)
//}
//
//type RenderSystem struct {
//	width          float32
//	height         float32
//	entities       []renderEntityHolder
//	circleRenderer *CircleRenderer
//	quadRenderer   *QuadRenderer
//	fontRenderer   *FontRenderer
//	renderers      []Renderer
//}
//
//type renderEntityHolder struct {
//	entity          *ecs.Entity
//	renderComponent *components.RenderComponent
//	position        *components.PositionedComponent
//}
//
//func NewRenderSystem(windowSize WindowSize) *RenderSystem {
//	proj := mgl32.Ortho(0, windowSize.Width, 0, windowSize.Height, -1.0, 1.0)
//	circleRenderer := NewCircleRenderer(proj)
//	quadRenderer := NewQuadRenderer(proj)
//	fontRenderer := NewFontRenderer(windowSize)
//
//	return &RenderSystem{
//		entities:       []renderEntityHolder{},
//		width:          windowSize.Width,
//		height:         windowSize.Height,
//		circleRenderer: circleRenderer,
//		quadRenderer:   quadRenderer,
//		fontRenderer:   fontRenderer,
//		renderers: []Renderer{
//			//circleRenderer,
//			quadRenderer,
//			//fontRenderer
//		},
//	}
//}
//
//func (r *RenderSystem) New(*ecs.World) {
//	render.UseDefaultBlending()
//	gl.ClearColor(colourDarkNavy.R, colourDarkNavy.G, colourDarkNavy.B, colourDarkNavy.A)
//}
//
//func (r *RenderSystem) Update(float32) {
//	gl.Clear(gl.COLOR_BUFFER_BIT)
//
//	for _, renderer := range r.renderers {
//		renderer.Render()
//	}
//}
//
//func (r *RenderSystem) Add(entity *ecs.Entity) {
//	renderComponent := entity.Component(components.IsRenderable).(*components.RenderComponent)
//	position := entity.Component(components.IsPositioned).(*components.PositionedComponent)
//
//	r.entities = append(r.entities, renderEntityHolder{
//		entity:          entity,
//		renderComponent: renderComponent,
//		position:        position,
//	})
//	if renderComponent.Circle != nil {
//		r.circleRenderer.Add(entity.ID(), renderComponent.Circle)
//	}
//	if renderComponent.Quad != nil {
//		r.quadRenderer.Add(entity.ID(), renderComponent.Quad, position)
//	}
//	if renderComponent.TextBox != nil {
//		r.fontRenderer.Add(entity.ID(), renderComponent.TextBox)
//	}
//}
//
//func (r *RenderSystem) Remove(basic *ecs.Entity) {
//	var del = -1
//	for index, e := range r.entities {
//		if e.entity.ID() == basic.ID() {
//			del = index
//			break
//		}
//	}
//	if del >= 0 {
//		r.entities = append(r.entities[:del], r.entities[del+1:]...)
//	}
//	for _, renderer := range r.renderers {
//		renderer.Remove(basic.ID())
//	}
//}
//
//func (r *RenderSystem) RequiredTypes() []interface{} {
//	return []interface{}{
//		components.IsRenderable,
//		components.IsPositioned,
//	}
//}
