package game

//
//type PlayerMovementSystem struct {
//	playingSpace PlayingSpace
//	entities     []struct {
//		base                  *ecs.Entity
//		controllableComponent *SpeedControlComponent
//		position              *PositionedComponent
//	}
//}
//
//func NewPlayerMovementSystem(space PlayingSpace) *PlayerMovementSystem {
//
//	m := &PlayerMovementSystem{
//		playingSpace: space,
//		entities: []struct {
//			base                  *ecs.Entity
//			controllableComponent *SpeedControlComponent
//			position              *PositionedComponent
//		}{}}
//
//	return m
//}
//
//func (m *PlayerMovementSystem) Add(entity *ecs.Entity) {
//
//	m.entities = append(m.entities,
//		struct {
//			base                  *ecs.Entity
//			controllableComponent *SpeedControlComponent
//			position              *PositionedComponent
//		}{
//			base: entity,
//			controllableComponent: entity.Component(IsSpeedControllable).(*SpeedControlComponent),
//			position:              entity.Component(IsPositioned).(*PositionedComponent),
//		})
//}
//
//func (m *PlayerMovementSystem) Update(dt float32) {
//	for _, e := range m.entities {
//		moveAmount := dt * e.controllableComponent.Speed
//		if e.position.X+moveAmount < 0 {
//
//			e.lateralMoveComponent.Quad.Position = [4]float32{0, e.lateralMoveComponent.Quad.Position.Y(), e.lateralMoveComponent.Quad.Width(), e.lateralMoveComponent.Quad.Position.W()}
//			e.lateralMoveComponent.Speed = 0
//		} else if e.lateralMoveComponent.Quad.Position.Z()+moveAmount > m.width {
//			e.lateralMoveComponent.Quad.Position = [4]float32{m.width - e.lateralMoveComponent.Quad.Width(), e.lateralMoveComponent.Quad.Position.Y(), m.width, e.lateralMoveComponent.Quad.Position.W()}
//			e.lateralMoveComponent.Speed = 0
//		} else {
//			e.lateralMoveComponent.Quad.Position[0] += moveAmount
//			e.lateralMoveComponent.Quad.Position[2] += moveAmount
//		}
//	}
//}
//
//func (m *PlayerMovementSystem) Remove(_ ecs.BasicEntity) {}
