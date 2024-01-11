package medicine_api

type MedicineApiConfig struct {
	SummaryKey string
	AppearKey  string
}

type MedicineApi interface {
	Summary() SummaryApi
	Appearance() AppearanceApi
}

type medicineApi struct {
	config        *MedicineApiConfig
	summaryApi    SummaryApi
	appearanceApi AppearanceApi
}

func NewMedicineApi(config *MedicineApiConfig) MedicineApi {
	summary := NewSummaryApi(config.SummaryKey)
	appearance := NewAppearanceApi(config.AppearKey)
	return &medicineApi{
		config:        config,
		summaryApi:    summary,
		appearanceApi: appearance,
	}
}

func (m *medicineApi) Summary() SummaryApi {
	return m.summaryApi
}

func (m *medicineApi) Appearance() AppearanceApi {
	return m.appearanceApi
}
