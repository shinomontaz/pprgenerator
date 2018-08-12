package generator

import (
	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
)

type Config struct {
	Start         time.Time
	End           time.Time
	Rnd           *rand.Rand
	NumJobs       int
	NumInterrupts int
	jobsList      []*Job
	interrupts    []*Interrupt
}

func (c *Config) Init() {
	// generate range of days
	tower_num := (c.NumJobs / 10) + c.Rnd.Intn(c.NumJobs/10)
	towers := make([]*Tower, 0, tower_num)
	for i := 0; i < tower_num; i++ {
		towers = append(towers, &Tower{
			Id:     i,
			Weight: 1000 * c.Rnd.Intn(c.NumJobs),
			x:      1000 * c.Rnd.Float64(),
			y:      1000 * c.Rnd.Float64(),
		})
	}

	days_step := int(c.End.Sub(c.Start).Seconds() / (60.0 * 60.0 * 24)) // количество дней в диапазоне

	for i := 0; i < c.NumJobs; i++ {
		var randBool bool
		if c.Rnd.Intn(3) > 1 {
			randBool = true
		}

		daysNum := c.Rnd.Intn(20)
		days := make([]time.Time, 0, daysNum)
		currDay := c.Start.Add(time.Duration(c.Rnd.Intn(days_step)*60*24) * time.Minute)

		for j := 0; j < daysNum && currDay.Before(c.End); j++ {
			days = append(days, currDay)
			currDay = currDay.Add(time.Duration(60*24) * time.Minute)
		}
		c.jobsList = append(c.jobsList,
			&Job{
				Id:         i + 1,
				Name:       randomdata.SillyName(),
				Difficulty: float64(c.Rnd.Intn(2)) + c.Rnd.Float64(),
				Tower:      towers[c.Rnd.Intn(tower_num)],
				IsBreaking: randBool,
				Days:       days,
			},
		)
	}

	randInterrupts := c.Rnd.Perm(days_step)[:c.NumInterrupts]

	for i := 0; i < c.NumInterrupts; i++ {
		c.interrupts = append(c.interrupts, &Interrupt{
			Day:    c.Start.Add(time.Duration(randInterrupts[i]*60*24) * time.Minute),
			Length: c.Rnd.Float64() * 24,
		})
	}
}

func (c *Config) Jobs() []*Job {
	return c.jobsList
}

func (c *Config) Interrupts() []*Interrupt {
	return c.interrupts
}
