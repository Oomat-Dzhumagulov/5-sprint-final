package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("invalid part length")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid steps: %w", err)
	}
	if steps <= 0 {
		return errors.New("invalid data")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}
	if duration <= 0 {
		return errors.New("invalid data")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("error: %w", err)
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
