package main

import (
	"encoding/json"
	"testing"
)

func toString(s MostHits) string {
	v, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(v)
}

func TestStats(t *testing.T) {
	statsChannel = make(chan FizzbuzzOptions)
	defer close(statsChannel)
	go collectStats()
	options1 := FizzbuzzOptions{
		3,
		5,
		"fizz",
		"buzz",
	}
	options2 := FizzbuzzOptions{
		4,
		5,
		"fizz",
		"buzz",
	}
	statsChannel <- options2
	statsChannel <- options1
	statsChannel <- options1
	expected := MostHits{
		Hits:           2,
		MostUsedParams: &options1,
	}
	mostUsed := getMostUsed()
	if mostUsed.Hits != expected.Hits ||
		mostUsed.MostUsedParams.Int1 != expected.MostUsedParams.Int1 ||
		mostUsed.MostUsedParams.Int2 != expected.MostUsedParams.Int2 ||
		mostUsed.MostUsedParams.Str1 != expected.MostUsedParams.Str1 ||
		mostUsed.MostUsedParams.Str2 != expected.MostUsedParams.Str2 {
		t.Errorf("Expected %s, got %s", toString(expected), toString(mostUsed))
	}
}
