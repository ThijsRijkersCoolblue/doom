package level

import (
	"doom/world"
	"fmt"
	"strconv"
	"strings"
)

func parseVertexLine(line string, level *Level) error {
	parts := strings.Fields(line)
	if len(parts) < 2 {
		return fmt.Errorf("invalid vertex line: %s", line)
	}

	start := 0
	if len(parts) >= 3 {
		start = 1
	}

	x, err := strconv.ParseFloat(parts[start], 64)
	if err != nil {
		return fmt.Errorf("invalid vertex x: %s", line)
	}
	y, err := strconv.ParseFloat(parts[start+1], 64)
	if err != nil {
		return fmt.Errorf("invalid vertex y: %s", line)
	}

	level.Vertices = append(level.Vertices, world.Vertex{X: x, Y: y})
	return nil
}

func parseSectorLine(line string, level *Level) error {
	parts := strings.Fields(line)
	if len(parts) < 2 {
		return fmt.Errorf("invalid sector line: %s", line)
	}

	start := 0
	if len(parts) >= 3 {
		start = 1
	}

	floorHeight, err := strconv.ParseFloat(parts[start], 64)
	if err != nil {
		return fmt.Errorf("invalid sector floor height: %s", line)
	}
	ceilingHeight, err := strconv.ParseFloat(parts[start+1], 64)
	if err != nil {
		return fmt.Errorf("invalid sector ceiling height: %s", line)
	}

	level.Sectors = append(level.Sectors, world.Sector{
		FloorHeight:   floorHeight,
		CeilingHeight: ceilingHeight,
	})
	return nil
}

func parseLinedefLine(line string, level *Level) error {
	parts := strings.Fields(line)
	if len(parts) < 5 {
		return fmt.Errorf("invalid linedef line: %s", line)
	}

	start := 0
	if len(parts) >= 6 {
		start = 1
	}

	startVertex, err := strconv.Atoi(parts[start])
	if err != nil {
		return fmt.Errorf("invalid linedef start vertex: %s", line)
	}
	endVertex, err := strconv.Atoi(parts[start+1])
	if err != nil {
		return fmt.Errorf("invalid linedef end vertex: %s", line)
	}
	frontSector, err := strconv.Atoi(parts[start+2])
	if err != nil {
		return fmt.Errorf("invalid linedef sector: %s", line)
	}
	backSector, err := strconv.Atoi(parts[start+3])
	if err != nil {
		return fmt.Errorf("invalid linedef back sector: %s", line)
	}
	textureID, err := strconv.Atoi(parts[start+4])
	if err != nil {
		return fmt.Errorf("invalid linedef texture id: %s", line)
	}

	level.Linedefs = append(level.Linedefs, world.Linedef{
		StartVertex: startVertex,
		EndVertex:   endVertex,
		FrontSector: frontSector,
		BackSector:  backSector,
		TextureID:   textureID,
	})
	return nil
}
