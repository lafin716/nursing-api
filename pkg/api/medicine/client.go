package medicine

type MedicineApiConfig struct {
	SummaryKey string
	AppearKey  string
}

type MedicineApi interface {
	Summary() SummaryApi
}

type medicineApi struct {
	config     *MedicineApiConfig
	summaryApi SummaryApi
}

func NewMedicineApi(config *MedicineApiConfig) MedicineApi {
	summary := NewSummaryApi(config.SummaryKey)
	return &medicineApi{
		config:     config,
		summaryApi: summary,
	}
}

func (m *medicineApi) Summary() SummaryApi {
	return m.summaryApi
}
