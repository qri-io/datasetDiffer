package datasetDiffer

import (
	"fmt"

	"github.com/qri-io/dataset"
)

// Diff contains a string description of a diff
type Diff string

// DiffList contains a slice of diffs in order of descending scope
type DiffList struct {
	diffs []Diff
}

// String returns the first (largest scope) change as a string
func (diffList DiffList) String() string {
	if len(diffList.diffs) > 0 {
		return string(diffList.diffs[0])
	}
	return ""
}

func diffStructure(a, b *dataset.Dataset) (*DiffList, error) {
	diffList := &DiffList{}
	diffDescription := Diff("")
	if len(a.Structure.Path().String()) > 1 && len(b.Structure.Path().String()) > 1 {
		if a.Structure.Path() != b.Structure.Path() {
			diffDescription = Diff("Structure Changed.")
			diffList.diffs = append(diffList.diffs, diffDescription)
		}
	} else {
		if len(a.Structure.Checksum) > 1 && len(b.Structure.Checksum) > 1 {
			if a.Structure.Checksum != b.Structure.Checksum {
				diffDescription = Diff("Structure Changed.")
				diffList.diffs = append(diffList.diffs, diffDescription)
			}
		} else {
			return nil, fmt.Errorf("error: structure path cannot be empty string")
		}
	}
	return diffList, nil
}

func diffTransform(a, b *dataset.Dataset) (*DiffList, error) {
	diffList := &DiffList{}
	diffDescription := Diff("")
	if len(a.Transform.Path().String()) > 1 && len(b.Transform.Path().String()) > 1 {
		if a.Transform.Path() != b.Transform.Path() {
			diffDescription = Diff("Transform Changed.")
			diffList.diffs = append(diffList.diffs, diffDescription)
		}
	}
	// else {
	// 	...
	// }
	return diffList, nil
}

func diffVisConfig(a, b *dataset.Dataset) (*DiffList, error) {
	diffList := &DiffList{}
	diffDescription := Diff("")
	if len(a.VisConfig.Path().String()) > 1 && len(b.VisConfig.Path().String()) > 1 {
		if a.VisConfig.Path() != b.VisConfig.Path() {
			diffDescription = Diff("VisConfig Changed.")
			diffList.diffs = append*(diffList.diffs, diffDescription)
		}
	} 
	// else {
	// ...
	// }
	return diffList, nil
}
//TODO: make work
func diffData(a, b *dataset.Dataset) (*DiffList, error) {
	temporarilyBlindToData := true // <-- REMOVE this
	diffList := &DiffList{}
	diffDescription := Diff("")
	if len(a.DataPath) > 1 && len(b.DataPath) > 1 {
		if a.DataPath != b.DataPath {
			diffDescription = Diff("Data Changed.")
			diffList.diffs = append(diffList.diffs, diffDescription)
		}
	} else {
		if !temporarilyBlindToData {
			return nil, fmt.Errorf("error: data path cannot be empty string")
		}
	}
	return diffList, nil
}

func diffMeta(a, b *dataset.Dataset) (*DiffList, error) {
	diffList := &DiffList{}
	diffDescription := Diff("")
	if a.Meta != nil && b.Meta != nil {
		if len(a.Meta.Path().String()) > 1 && len(b.Meta.Path().String()) > 1 {
			if a.Meta.Path() != b.Meta.Path() {
				diffDescription = Diff("Metadata Changed.")
				diffList.diffs = append(diffList.diffs, diffDescription)
			}
		} else {
			if a.Meta.Title != b.Meta.Title && a.Meta.Title != "" {
				diffDescription = Diff("Title Changed.")
				diffList.diffs = append(diffList.diffs, diffDescription)
			}
			if a.Meta.Description != b.Meta.Description && a.Meta.Description != "" {
				diffDescription = Diff("Description Changed.")
				diffList.diffs = append(diffList.diffs, diffDescription)
			}
		}
	}
	return diffList, nil
}

// DiffDatasets calculates diffs between two datasets and returns a
// dataset. Differences are checked in order of descending scope
// - dataset.Dataset.path
// - dataset.Dataset.Structure.path
// - dataset.Dataset.Data.path
// TODO: make diffs non-trivial
func DiffDatasets(a, b *dataset.Dataset) (*DiffList, error) {
	diffList := &DiffList{}
	// Compare Structure
	structureDiffList, err := diffStructure(a, b)
	if err != nil {
		return nil, err
	}
	if len(structureDiffList.diffs) > 0 {
		diffList.diffs = append(diffList.diffs, structureDiffList.diffs...)
	}
	// Compare Data
	dataDiffList, err := diffData(a, b)
	if err != nil {
		return nil, err
	}
	if len(dataDiffList.diffs) > 0 {
		diffList.diffs = append(diffList.diffs, dataDiffList.diffs...)
	}
	// Compare Metadata
	metaDiffList, err := diffMeta(a, b)
	if err != nil {
		return nil, err
	}
	if len(metaDiffList.diffs) > 0 {
		diffList.diffs = append(diffList.diffs, metaDiffList.diffs...)
	}
	// Compare Transform
	transformDiffList, err := diffTransform(a, b)
	if err != nil {
		return nil, err
	}
	if len(transformDiffList.diffs) > 0 {
		diffList.diffs = append(diffList.diffs, transformDiffList.diffs...)
	}
	// Compare VisConfig
	visConfigDiffList, err := diffVisConfig(a, b)
	if err != nil {
		return nil, err
	}
	if len(visConfigDiffList.diffs) > 0 {
		diffList.diffs = append(diffList.diffs, visConfigDiffList.diffs...)
	}
	return diffList, nil
}
