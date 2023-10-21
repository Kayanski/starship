package transformer

import (
	"github.com/cosmology-tech/starship/pkg/types"
)

type Object interface {
	WriteToFile(dir string) error
	Validate() error
}

type Transformer interface {
	Transform(types.StarshipObject, types.ConvertOptions) (Object, error)
}