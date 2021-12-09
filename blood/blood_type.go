package blood

import (
	"errors"
	"sort"
)

type BloodMap map[BloodName]int

type BloodName int

const (
	AMinus BloodName = iota
	ABMinus
	BMinus
	OMinus
	APlus
	BPlus
	ABPlus
	OPlus
)

var (
	bank = &ArrBlood{}
)

type bloodType struct {
	Name    BloodName
	Accepts ArrBlood
}

func init() {

	var (
		OM  = bloodType{OMinus, []bloodType{}}
		AM  = bloodType{AMinus, []bloodType{OM}}
		BM  = bloodType{BMinus, []bloodType{OM}}
		ABM = bloodType{ABMinus, []bloodType{OM, AM, BM}}

		OP  = bloodType{OPlus, []bloodType{OM}}
		AP  = bloodType{APlus, []bloodType{AM, OM, OP}}
		BP  = bloodType{BPlus, []bloodType{OM, BM, OP}}
		ABP = bloodType{ABPlus, []bloodType{OM, AM, BM, ABM, OP, AP, BP}}
	)

	bank = &ArrBlood{OM, AM, BM, ABM, OP, AP, BP, ABP}

}

func GetAccepts(name BloodName) ([]BloodName, error) {
	for _, i := range *bank {
		if i.Name == name {
			var list []BloodName

			for _, j := range i.Accepts {
				list = append(list, j.Name)
			}

			return list, nil
		}
	}

	return nil, errors.New("blood type not found")
}

func (bt BloodMap) IsEmpty() bool {
	for _, i := range bt {
		if i != 0 {
			return false
		}
	}

	return true
}

func GetSorted() []bloodType {
	sort.Sort(bank)

	return *bank
}

func (ab ArrBlood) GetSorted() ArrBlood {
	sort.Sort(ab)

	return ab
}

type ArrBlood []bloodType

func (a ArrBlood) Len() int           { return len(a) }
func (a ArrBlood) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ArrBlood) Less(i, j int) bool { return len(a[i].Accepts) < len(a[j].Accepts) }
