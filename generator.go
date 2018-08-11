package generator

import (
	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
)

func Generate(num int, start, end time.Time, rnd *rand.Rand) (jobs []*Job) {
	// generate range of days
	tower_num := (num / 10) + rnd.Intn(num/10)
	towers := make([]*Tower, 0, tower_num)
	for i := 0; i < tower_num; i++ {
		towers = append(towers, &Tower{
			Id:     i,
			Weight: 1000 * rnd.Intn(num),
			x:      1000 * rnd.Float64(),
			y:      1000 * rnd.Float64(),
		})
	}

	days_step := int(end.Sub(start).Seconds() / (60.0 * 60.0 * 24))

	for i := 0; i < num; i++ {
		rand_int := rnd.Intn(3)
		var rand_bool bool
		if rand_int > 1 {
			rand_bool = true
		}

		days_num := rnd.Intn(20)
		days := make([]time.Time, 0, days_num)
		curr_day := start.Add(time.Duration(rnd.Intn(days_step)*60*24) * time.Minute)

		for j := 0; j < days_num && curr_day.Before(end); j++ {
			days = append(days, curr_day)
			curr_day = curr_day.Add(time.Duration(60*24) * time.Minute)
		}
		jobs = append(jobs,
			&Job{
				Id:         i,
				Name:       randomdata.SillyName(),
				Difficulty: float64(rnd.Intn(2)) + rnd.Float64(),
				pTower:     towers[rnd.Intn(tower_num)],
				IsBreaking: rand_bool,
				Days:       days,
			},
		)
	}

	return jobs
}
