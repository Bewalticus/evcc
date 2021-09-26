package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decoratePhoenixEMEth(charger api.Charger, chargeTimer api.ChargeTimer, meter func() (float64, error), meterEnergy func() (float64, error), meterCurrent func() (float64, float64, float64, error)) api.Charger {
	switch {
	case meter == nil && meterCurrent == nil && meterEnergy == nil:
		return &struct {
			api.Charger
			api.ChargeTimer
		}{
			Charger:     charger,
			ChargeTimer: chargeTimer,
		}

	case meter != nil && meterCurrent == nil && meterEnergy == nil:
		return &struct {
			api.Charger
			api.ChargeTimer
			api.Meter
		}{
			Charger:     charger,
			ChargeTimer: chargeTimer,
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
		}

	case meter == nil && meterCurrent == nil && meterEnergy != nil:
		return &struct {
			api.Charger
			api.ChargeTimer
			api.MeterEnergy
		}{
			Charger:     charger,
			ChargeTimer: chargeTimer,
			MeterEnergy: &decoratePhoenixEMEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterCurrent == nil && meterEnergy != nil:
		return &struct {
			api.Charger
			api.ChargeTimer
			api.Meter
			api.MeterEnergy
		}{
			Charger:     charger,
			ChargeTimer: chargeTimer,
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEMEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter == nil && meterCurrent != nil && meterEnergy == nil:
		return &struct {
			api.Charger
			api.ChargeTimer
			api.MeterCurrent
		}{
			Charger:     charger,
			ChargeTimer: chargeTimer,
			MeterCurrent: &decoratePhoenixEMEthMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
		}

	case meter != nil && meterCurrent != nil && meterEnergy == nil:
		return &struct {
			api.Charger
			api.ChargeTimer
			api.Meter
			api.MeterCurrent
		}{
			Charger:     charger,
			ChargeTimer: chargeTimer,
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
			MeterCurrent: &decoratePhoenixEMEthMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
		}

	case meter == nil && meterCurrent != nil && meterEnergy != nil:
		return &struct {
			api.Charger
			api.ChargeTimer
			api.MeterCurrent
			api.MeterEnergy
		}{
			Charger:     charger,
			ChargeTimer: chargeTimer,
			MeterCurrent: &decoratePhoenixEMEthMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			MeterEnergy: &decoratePhoenixEMEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterCurrent != nil && meterEnergy != nil:
		return &struct {
			api.Charger
			api.ChargeTimer
			api.Meter
			api.MeterCurrent
			api.MeterEnergy
		}{
			Charger:     charger,
			ChargeTimer: chargeTimer,
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
			MeterCurrent: &decoratePhoenixEMEthMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			MeterEnergy: &decoratePhoenixEMEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}
	}

	return nil
}

type decoratePhoenixEMEthMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decoratePhoenixEMEthMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decoratePhoenixEMEthMeterCurrentImpl struct {
	meterCurrent func() (float64, float64, float64, error)
}

func (impl *decoratePhoenixEMEthMeterCurrentImpl) Currents() (float64, float64, float64, error) {
	return impl.meterCurrent()
}

type decoratePhoenixEMEthMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decoratePhoenixEMEthMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}
