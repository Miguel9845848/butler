package prereqs

import (
	"github.com/itchio/butler/redist"
	"github.com/itchio/ox"
	"github.com/pkg/errors"
)

func (pc *PrereqsContext) FilterPrereqs(names []string) ([]string, error) {
	consumer := pc.Consumer

	var result []string
	for _, name := range names {
		entry, err := pc.GetEntry(name)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		if entry == nil {
			consumer.Warnf("Prereq (%s) not found in registry, skipping...", name)
			continue
		}

		if !RedistHasPlatform(entry, pc.Host.Runtime.Platform) {
			consumer.Warnf("Prereq (%s) is not relevant on (%s), skipping...", name, pc.Host.Runtime.Platform)
			continue
		}
		result = append(result, name)
	}
	return result, nil
}

func RedistHasPlatform(redist *redist.RedistEntry, platform ox.Platform) bool {
	for _, p := range redist.Platforms {
		if p == string(platform) {
			return true
		}
	}
	return false
}
