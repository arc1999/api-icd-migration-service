package transformer

import (
	"api-icd-migration-service/model"
)

func Transform(micds []model.ICDMongo) []model.ICD {
	var icds []model.ICD
	for _, micd := range micds {
		var icd model.ICD
		icd.DateUpdated = micd.DateUpdated
		icd.Code = micd.Code
		icd.CreatedByName = micd.CreatedByName
		icd.DateCreated = micd.DateCreated
		icd.UpdatedBy = micd.UpdatedBy
		icd.UpdatedByName = micd.UpdatedByName
		icd.DiseaseChiefComplaint = micd.DiseaseChiefComplaint
		icd.IcdCode = micd.IcdCode
		icd.ID=micd.ID
		var ct string
		for _, terms := range micd.CommonTerms {
			ct = ct + terms	+","
		}
		if len(ct)>0{
			ct = ct[:len(ct) - 1]
		}
		icd.CommonTerms = ct
		icds = append(icds, icd)
	}
	return icds
}
