package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID       uuid.UUID `json:"engine_id"`
	Displacement   int64     `json:"displacement"`
	NoOfCyclinders int64     `json:"noofcylinders"`
	CarRange       int64     `json:"carrange"`
}

type EngineRequest struct {
	Displacement   int64 `json:"displacement"`
	NoOfCyclinders int64 `json:"noofcylinders"`
	CarRange       int64 `json:"carrange"`
}

func ValidateEngineRequest(engineReq EngineRequest) error {
	if err := validateDisplacement(engineReq.Displacement); err != nil {
		return err
	}
	if err := validateNoOfCylinders(engineReq.NoOfCyclinders); err != nil {
		return err
	}
	if err := validateCarRange(engineReq.CarRange); err != nil {
		return err
	}
	return nil
}

func validateDisplacement(displacement int64) error {
	if displacement <= 0 {
		return errors.New("engine displacement must be positive")
	}
	return nil
}

func validateNoOfCylinders(noOfCylinders int64) error {
	if noOfCylinders <= 0 {
		return errors.New("engine number of cylinders must be positive")
	}
	return nil
}

func validateCarRange(carRange int64) error {
	if carRange < 0 {
		return errors.New("engine car range must be positive")
	}
	return nil
}
