package repositories

import (
	"encoder/domain"
	"fmt"

	"github.com/jinzhu/gorm"
)

//JobRepository is Job interface
type JobRepository interface {
	Insert(jog *domain.Job) (*domain.Job, error)
  Find(id string) (*domain.Job, error)
  Update(job *domain.Job) (*domain.Job, error)
}


//JobRepositoryDb is Job struct
type JobRepositoryDb struct {
	Db *gorm.DB
}

//NewJobRepository is new Job repository
func NewJobRepository(db *gorm.DB) *JobRepositoryDb {
	return &JobRepositoryDb{Db: db}
}

//Insert is create job method
func (repo JobRepositoryDb) Insert(job *domain.Job) (*domain.Job, error) {

	err := repo.Db.Create(job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}

//Find is search job by id
func (repo JobRepositoryDb) Find(id string) (*domain.Job, error) {

	var job domain.Job
	repo.Db.Preload("Video").First(&job, "id = ?", id)

	if job.ID == "" {
		return nil, fmt.Errorf("job does not exist")
	}

	return &job, nil

}

//Update is create job method
func (repo JobRepositoryDb) Update(job *domain.Job) (*domain.Job, error) {

  err := repo.Db.Save(&job).Error

  if err != nil {
		return nil, err
	}

  return job, nil

}
