package vars

// ============================================================================
// Named Types
// ============================================================================

type AppState int

type SlotType int

type Slot struct {
	Content    []byte
	Type       SlotType
	HasContent bool
}

// ============================================================================
// Constants
// ============================================================================

const (
	StateInit AppState = iota
	StateIdle
	StateCopying
	StatePasting
)

const (
	TypeText SlotType = iota
	TypeImage
)
