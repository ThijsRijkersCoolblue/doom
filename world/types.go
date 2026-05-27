package world

type Vertex struct {
	X float64
	Y float64
}

type Sector struct {
	FloorHeight   float64
	CeilingHeight float64
}

type Linedef struct {
	StartVertex int
	EndVertex   int
	FrontSector int
	BackSector  int
	TextureID   int
}
