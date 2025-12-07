package bounds

import (
	"fmt"
	"strconv"
	"strings"
)

func Bounds(r string) (int, int, error) {
	bounds := strings.Split(r, "-")
	if len(bounds) != 2 {
		return 0, 0, fmt.Errorf("bounds: multiple hyphens %s", r)
	}

	lbound, err := strconv.Atoi(bounds[0])
	if err != nil {
		return 0, 0, fmt.Errorf("bounds: bad lower bound: %w", err)
	}

	rbound, err := strconv.Atoi(bounds[1])
	if err != nil {
		return 0, 0, fmt.Errorf("bounds: bad upper bound: %w", err)
	}

	return lbound, rbound, nil
}
