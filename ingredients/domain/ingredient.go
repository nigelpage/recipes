package ingredients

import (
	"github.com/ajstarks/svgo"
)

type ingredient struct {
	id int64;
	name string;
	image svg.SVG;
	description string;
	alternates []int64
}