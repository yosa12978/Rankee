package repositories

import "github.com/yosa12978/Rankee/internal/domain/models"

type RankRepository interface {
	GetByUUID(uuid string)
	GetSubRanks(userUUID string)
	Search(query string)
	Create(rank models.Rank)
	Update(rank models.Rank)
	Delete(uuid string)
}
