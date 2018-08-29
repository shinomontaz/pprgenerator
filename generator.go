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
	stops         []*Stop
}

func (c *Config) Init() {
	// generate range of days
	equipment_num := (c.NumJobs / 10) + c.Rnd.Intn(c.NumJobs/10)
	tower_num := (equipment_num / 10) + c.Rnd.Intn(equipment_num/10)

	equipments := make([]*Equipment, 0, equipment_num)
	towers := make([]*Tower, 0, tower_num)

	for i := 0; i < tower_num; i++ {
		towers = append(towers, &Tower{
			Id:     i,
			Weight: 1000 * c.Rnd.Intn(c.NumJobs),
			Coords: LatLong{
				Lat:  180 * c.Rnd.Float64(),
				Long: 180 * c.Rnd.Float64(),
			},
		})
	}

	for i := 0; i < equipment_num; i++ {
		equipments = append(equipments, &Equipment{
			Id:       i,
			Name:     randomdata.SillyName(),
			Weight:   1000 * c.Rnd.Intn(c.NumJobs),
			FkTower:  towers[c.Rnd.Intn(tower_num)],
			Majority: c.Rnd.Float64(),
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
				Id:          i + 1,
				Name:        randomdata.SillyName(),
				Difficulty:  float64(c.Rnd.Intn(2)) + c.Rnd.Float64(),
				FkEquipment: equipments[c.Rnd.Intn(equipment_num)],
				IsBreaking:  randBool,
				Days:        days,
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

	randStops := c.Rnd.Perm(days_step)[:c.NumInterrupts]

	for i := 0; i < c.NumInterrupts; i++ {
		c.stops = append(c.stops, &Stop{
			Day:         c.Start.Add(time.Duration(randStops[i]*60*24) * time.Minute),
			Length:      c.Rnd.Float64() * 24,
			FkEquipment: equipments[c.Rnd.Intn(equipment_num)],
		})
	}
}

func (c *Config) Jobs() []*Job {
	return c.jobsList
}

func (c *Config) Interrupts() []*Interrupt {
	return c.interrupts
}

func (c *Config) Stops() []*Stop {
	return c.stops
}
