package ingredients

import (
	"github.com/ajstarks/svgo"
	"github.com/google/uuid"
)

type ingredient struct {
	id uuid.UUID;
	name string;
	image svg.SVG;
	description string;
	alternates []uuid.UUID
}