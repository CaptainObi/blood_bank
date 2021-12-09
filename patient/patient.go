package patient

import b "github.com/CaptainObi/comding_comps/blood_bank/blood"

type Patient struct {
	BType b.BloodName
}

type Patients []Patient

func (ps Patients) GetNeededBlood() b.BloodMap {

	bloodMap := make(b.BloodMap)

	for _, v := range ps {
		bloodMap[v.BType]++
	}

	return bloodMap

}
