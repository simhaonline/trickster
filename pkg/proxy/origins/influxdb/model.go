/*
 * Copyright 2018 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package influxdb

import (
	"encoding/json"
	"time"

	"github.com/tricksterproxy/trickster/pkg/sort/times"
	"github.com/tricksterproxy/trickster/pkg/timeseries"

	"github.com/influxdata/influxdb/models"
)

// SeriesEnvelope represents a response object from the InfluxDB HTTP API
type SeriesEnvelope struct {
	Results      []Result              `json:"results"`
	Err          string                `json:"error,omitempty"`
	StepDuration time.Duration         `json:"step,omitempty"`
	ExtentList   timeseries.ExtentList `json:"extents,omitempty"`

	timestamps map[time.Time]bool // tracks unique timestamps in the matrix data
	tslist     times.Times
	isSorted   bool // tracks if the matrix data is currently sorted
	isCounted  bool // tracks if timestamps slice is up-to-date
}

// Result represents a Result returned from the InfluxDB HTTP API
type Result struct {
	StatementID int          `json:"statement_id"`
	Series      []models.Row `json:"series,omitempty"`
	Err         string       `json:"error,omitempty"`
}

// MarshalTimeseries converts a Timeseries into a JSON blob
func (c Client) MarshalTimeseries(ts timeseries.Timeseries) ([]byte, error) {
	// Marshal the Envelope back to a json object for Cache Storage
	return json.Marshal(ts)
}

// UnmarshalTimeseries converts a JSON blob into a Timeseries
func (c Client) UnmarshalTimeseries(data []byte) (timeseries.Timeseries, error) {
	se := &SeriesEnvelope{}
	err := json.Unmarshal(data, se)
	return se, err
}
