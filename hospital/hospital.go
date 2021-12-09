package hospital

import (
	b "github.com/CaptainObi/comding_comps/blood_bank/blood"
	p "github.com/CaptainObi/comding_comps/blood_bank/patient"
)

type Hospital struct {
	blood    b.BloodMap
	patients p.Patients
}

func (h Hospital) EnoughBlood() bool {
	needed_blood := h.patients.GetNeededBlood()
	available_blood := h.blood
	sorted_blood := b.GetSorted()

	for _, v := range sorted_blood {
		if needed_blood[v.Name] != 0 {
			new_accepts := append(v.Accepts, v)
			accepted := new_accepts.GetSorted()
			for i := accepted.Len() - 1; i >= 0; i-- {
				if available_blood[accepted[i].Name] != 0 {
					needed_blood[v.Name]--
					break
				}
			}
		}
	}

	return needed_blood.IsEmpty()
}

func CreateBlankHospital() Hospital {
	bloodMap := make(b.BloodMap)
	bloodMap[b.AMinus] = 0
	bloodMap[b.APlus] = 0
	bloodMap[b.BMinus] = 0
	bloodMap[b.BPlus] = 0
	bloodMap[b.ABMinus] = 0
	bloodMap[b.ABPlus] = 0
	bloodMap[b.OMinus] = 0
	bloodMap[b.OPlus] = 0

	return Hospital{patients: []p.Patient{}, blood: bloodMap}
}

func (h *Hospital) AddPatient(p p.Patient) {
	h.patients = append(h.patients, p)
}

func (h *Hospital) AddBlood(t b.BloodName, amt int) {
	h.blood[t] = amt
}
