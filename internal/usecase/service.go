package usecase

import (
	"github.com/ivan12093/VK-DB-PROJECT/internal/models"
	"github.com/ivan12093/VK-DB-PROJECT/internal/repository"
	"github.com/ivan12093/VK-DB-PROJECT/internal/utils/errors"
)

type ServiceUsecase struct {
	serviceRepository repository.ServiceR
}

func NewServiceUsecase(serviceRepository repository.ServiceR) *ServiceUsecase {
	return &ServiceUsecase{serviceRepository: serviceRepository}
}

func (usecase *ServiceUsecase) Clear() (err error) {
	err = usecase.serviceRepository.Clear()
	if err != nil {
		err = errors.ServerInternal
		return
	}
	return
}

func (usecase *ServiceUsecase) Status() (status *models.ForumStatus, err error) {
	status, err = usecase.serviceRepository.Status()
	if err != nil {
		err = errors.ServerInternal
		return
	}
	return
}
