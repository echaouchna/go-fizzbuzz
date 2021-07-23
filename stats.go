package main

type allStats struct {
	counts   map[FizzbuzzOptions]uint64
	maxHits  uint64
	mostUsed *FizzbuzzOptions
}

type MostHits struct {
	Hits           uint64           `json:"hits"`
	MostUsedParams *FizzbuzzOptions `json:"mostUsedParams"`
}

var (
	stats        = allStats{counts: make(map[FizzbuzzOptions]uint64), maxHits: 0, mostUsed: nil}
	statsChannel chan FizzbuzzOptions
)

func collectStats() {
	for {
		options := <-statsChannel
		stats.counts[options] += 1
		if stats.maxHits < stats.counts[options] {
			stats.maxHits = stats.counts[options]
			stats.mostUsed = &options
		}
	}
}

func getMostUsed() MostHits {
	mostUsed := MostHits{
		Hits:           stats.maxHits,
		MostUsedParams: stats.mostUsed,
	}
	return mostUsed
}
