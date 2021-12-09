package main

import (
	"fmt"

	"github.com/CaptainObi/comding_comps/blood_bank/blood"
	"github.com/CaptainObi/comding_comps/blood_bank/hospital"
	"github.com/CaptainObi/comding_comps/blood_bank/patient"
)

func main() {
	h := hospital.CreateBlankHospital()
	h.AddPatient(patient.Patient{BType: blood.ABMinus})
	h.AddPatient(patient.Patient{BType: blood.ABPlus})
	h.AddBlood(blood.OMinus, 1)
	h.AddBlood(blood.OPlus, 1)
	fmt.Println(h.EnoughBlood())
}
