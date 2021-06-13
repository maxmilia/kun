package day

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	TimezoneUTC()
	v := Now()
	name, offset := v.Zone()
	assert.Equal(t, "UTC", name)
	assert.Equal(t, 0, offset)

	name, offset = v.Timezone("America/New_York").Zone()
	assert.Equal(t, "EDT", name)
	assert.Equal(t, -4*60*60, offset)

	Timezone("America/New_York")
	v2 := Now()
	name, offset = v2.Zone()
	assert.Equal(t, "EDT", name)
	assert.Equal(t, -4*60*60, offset)

	name, offset = v2.Timezone("Beijing", 8*60*60).Zone()
	assert.Equal(t, "Beijing", name)
	assert.Equal(t, 8*60*60, offset)
}

func TestOf(t *testing.T) {

	TimezoneUTC()
	assert.Equal(t, 31, Of(Of("2019-12-31")).Day())
	assert.Equal(t, 31, Of(*Of("2019-12-31")).Day())
	assert.Equal(t, 31, Of(time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC)).Day())

	assert.Equal(t, 31, Of("2019-12-31").Day())
	assert.Equal(t, 30, Of("2019-12-31").Timezone("America/New_York").Day())
	assert.Equal(t, 31, Of("2019-12-31 08:31:56").Day())
	assert.Equal(t, 56, Of("2019-12-31 08:31:56").Second())

	Timezone("America/New_York")
	assert.Equal(t, 30, Of("2019-12-31").Day())
	assert.Equal(t, 31, Of("2019-12-31 08:31:56").Day())
	assert.Equal(t, 56, Of("2019-12-31 08:31:56").Second())

	assert.Panics(t, func() {
		fmt.Println(Of("error").Day())
	})
}

func TestLoad(t *testing.T) {
	TimezoneUTC()
	v := Now()
	assert.Equal(t, 31, v.Load("2019-12-31").Day())
	assert.Equal(t, 30, v.Load("2019-12-31").Timezone("America/New_York").Day())
	assert.Equal(t, 31, v.Load("2019-12-31 08:31:56").Day())
	assert.Equal(t, 56, v.Load("2019-12-31 08:31:56").Second())

	Timezone("America/New_York")
	assert.Equal(t, 30, v.Load("2019-12-31").Day())
	assert.Equal(t, 31, v.Load("2019-12-31 08:31:56").Day())
	assert.Equal(t, 56, v.Load("2019-12-31 08:31:56").Second())
}

func TestTimezones(t *testing.T) {
	zones := TimeZones()
	assert.True(t, len(zones) > 0)
}