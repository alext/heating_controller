package controller

import (
	"fmt"

	"github.com/alext/heating-controller/thermostat"
	"github.com/alext/heating-controller/units"
)

// golang.org/x/tools/cmd/stringer
//go:generate stringer -type=Action

type Action uint8

const (
	Off Action = iota
	On
	SetTarget
	IncreaseTarget
	DecreaseTarget
)

func (a Action) MarshalText() ([]byte, error) {
	return []byte(a.String()), nil
}

func (a *Action) UnmarshalText(data []byte) error {
	switch string(data) {
	case "On":
		*a = On
	case "Off":
		*a = Off
	case "SetTarget":
		*a = SetTarget
	case "IncreaseTarget":
		*a = IncreaseTarget
	case "DecreaseTarget":
		*a = DecreaseTarget
	default:
		return fmt.Errorf("Unrecognised action value '%s'", data)
	}
	return nil
}

type ThermostatAction struct {
	Action Action            `json:"action"`
	Param  units.Temperature `json:"param"`
}
