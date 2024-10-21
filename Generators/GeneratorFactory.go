package Generators

import (
	"fmt"
)

type GeneratorFactory struct{}

func NewGeneratorFactory() *GeneratorFactory {
	return &GeneratorFactory{}
}

func (f *GeneratorFactory) GetGenerator(generatorType GeneratorType) (Generator, error) {
	switch generatorType {
	case PRIM:
		return NewPrimsGenerator(75), nil
	case GrowingTree:
		return NewGrowingTreeGenerator(), nil
	default:
		return nil, fmt.Errorf("unknown generator type: %d", generatorType)
	}
}
