package core

import (
	"github.com/amstee/easy-cut/services/appointment/src/vars"
	"github.com/amstee/easy-cut/src/es"
	"errors"
	"encoding/json"
	"github.com/amstee/easy-cut/src/logger"
	"fmt"
)

func CreateAppointment(appointment *vars.Appointment) error {
	resp, err := es.Index("appointment", "appointment", appointment); if err != nil {
		return err
	}
	appointment.Id = resp.Id
	return nil
}

func FindAppointment(appointment *vars.Appointment, appointmentId string) error {
	resp, err := es.GetById("appointment", "appointment", appointmentId); if err != nil {
		return err
	}
	if resp.Found {
		if appointment != nil {
			appointment.Id = resp.Id
			return json.Unmarshal(*resp.Source, appointment)
		}
		return nil
	}
	return errors.New("appointment not found")
}

func FindAppointments(extract vars.ExtractQuery) (*[]vars.Appointment, error) {
	var appointments []vars.Appointment
	query := extract.ConstructQuery()
	fmt.Println(query)
	result, err := es.Search("appointment", query, "", false, -1, 100)
	if err != nil {
		return nil, err
	}
	if result.Hits.TotalHits > 0 {
		for _, hit := range result.Hits.Hits {
			var appointment vars.Appointment
			appointment.Id = hit.Id
			err := json.Unmarshal(*(hit.Source), &appointment)
			if err != nil {
				logger.Error.Println("unable to unmarshal appointment if = ", hit.Id, ", raw data = ", hit)
			} else {
				appointments = append(appointments, appointment)
			}
		}
		return &appointments, nil
	}
	return &appointments, nil
}

func UpdateAppointment(rate vars.UpdateAppointment, appointmentId string) error {
	err := FindAppointment(nil, appointmentId); if err != nil {
		return errors.New("this appointment does not exist")
	}
	_, err = es.UpdateDoc("appointment", "appointment", appointmentId, rate); if err != nil {
		return err
	}
	return nil
}

func DeleteAppointment(appointmentId string, check bool) error {
	if check {
		err := FindAppointment(nil, appointmentId); if err != nil {
			return errors.New("this appointment does not exist")
		}
	}
	_, err := es.DeleteDoc("appointment", "appointment", appointmentId); if err != nil {
		return err
	}
	return nil
}