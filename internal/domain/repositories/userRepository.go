package repositories

import "github.com/yosa12978/Rankee/internal/domain/models"

type UserRepository interface {
	GetByUUID(uuid string)
	GetSubs(uuid string)
	Search(query string)
	Create(user models.User)
	Update(user models.User)
	Delete(uuid string)
}
