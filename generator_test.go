package generator

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {

	conf := Config{
		Start:         time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		End:           time.Date(2018, 6, 30, 0, 0, 0, 0, time.UTC),
		Rnd:           rand.New(rand.NewSource(time.Now().UnixNano())),
		NumJobs:       100,
		NumInterrupts: 10,
	}

	conf.Init()
	jobs := conf.Jobs()
	interrupts := conf.Interrupts()

	for _, job := range jobs {
		jsonStr, _ := json.Marshal(job)
		fmt.Printf("%s\n", jsonStr)
	}

	jsonStr, _ := json.Marshal(interrupts)
	fmt.Printf("%s\n", jsonStr)
}
