package takehistory

import "nursing_api/pkg/ent"

func toModelList(entities []*ent.TakeHistory) []*TakeHistory {
	var models []*TakeHistory
	for _, entity := range entities {
		models = append(models, toModel(entity))
	}

	return models
}

func toModel(entity *ent.TakeHistory) *TakeHistory {
	if entity == nil {
		return nil
	}

	if entity.Edges.Timezone == nil {
		return nil
	}

	return &TakeHistory{
		ID:           entity.ID,
		TimezoneId:   entity.Edges.Timezone.ID,
		TimezoneName: entity.Edges.Timezone.TimezoneName,
		UserId:       entity.UserID,
		TakeDate:     entity.TakeDate,
		TakeStatus:   TakeStatus(entity.TakeStatus),
		Memo:         entity.Memo,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}

func toModelItemList(entities []*ent.TakeHistoryItem) []*TakeHistoryItem {
	var models []*TakeHistoryItem
	for _, entity := range entities {
		models = append(models, toModelItem(entity))
	}

	return models
}

func toModelItem(entity *ent.TakeHistoryItem) *TakeHistoryItem {
	return &TakeHistoryItem{
		ID:                 entity.ID,
		UserId:             entity.UserID,
		TakeHistoryId:      entity.TakeHistoryID,
		PrescriptionItemId: entity.PrescriptionItemID,
		TakeStatus:         entity.TakeStatus,
		PillName:           entity.Edges.PrescriptionItem.MedicineName,
		TakeAmount:         entity.TakeAmount,
		RemainAmount:       entity.RemainAmount,
		TotalAmount:        entity.TotalAmount,
		TakeUnit:           entity.TakeUnit,
		TakeDate:           entity.TakeDate,
		Memo:               entity.Memo,
		CreatedAt:          entity.CreatedAt,
		UpdatedAt:          entity.UpdatedAt,
	}
}
