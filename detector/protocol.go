// Copyright 2015 Eleme Inc. All rights reserved.

package detector

import (
	"strconv"
	"strings"

	"github.com/eleme/banshee/errors"
	"github.com/eleme/banshee/models"
)

// Function parseMetric parses protocol line string into Metric.
//   NAME  TIMESTAMP  VALUE \n
//   foo   1449481993 3.145 \n
//
func parseMetric(line string) (*models.Metric, error) {
	line = strings.TrimSpace(line)
	words := strings.Fields(line)
	if len(words) != 3 {
		return nil, errors.ErrProtocol
	}
	var err error
	m := &models.Metric{}
	m.Name = words[0]
	num, err := strconv.ParseUint(words[1], 10, 32)
	if err != nil {
		return nil, err
	}
	m.Stamp = uint32(num)
	m.Value, err = strconv.ParseFloat(words[2], 64)
	if err != nil {
		return nil, err
	}
	return m, nil
}
