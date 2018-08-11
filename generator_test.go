package generator

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	start := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2018, 6, 30, 0, 0, 0, 0, time.UTC)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	jobs := Generate(10, start, end, rnd)

	for _, job := range jobs {
		jsonStr, _ := json.Marshal(job)

		fmt.Printf("%s\n", jsonStr)

	}

	//	jsonStr, _ := json.Marshal(jobs)
}
