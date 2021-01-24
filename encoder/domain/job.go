package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewJob(output string, status string, video *Video) (*Job, error) {
	j := Job{
		OutputBucketPath: output,
		Status:           status,
		Video:            video,
	}
	j.prepare()
	if err := j.Validate(); err != nil {
		return nil, err
	}
	return &j, nil
}

type Job struct {
	ID               string    `valid:"uuid"`
	OutputBucketPath string    `valid:"notnull"`
	Status           string    `valid:"notnull"`
	Video            *Video    `valid:"-"`
	VideoID          string    `valid:"-"`
	Error            string    `valid:"-"`
	CreatedAt        time.Time `valid:"-"`
	UpdatedAt        time.Time `valid:"-"`
}

func (j *Job) Validate() error {
	if _, err := govalidator.ValidateStruct(j); err != nil {
		return err
	}
	return nil
}

func (j *Job) prepare() {
	j.ID = uuid.NewV4().String()
	j.CreatedAt = time.Now()
	j.UpdatedAt = time.Now()
}
