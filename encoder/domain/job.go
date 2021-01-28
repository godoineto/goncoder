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
	ID               string    `json:"job_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OutputBucketPath string    `json:"output_bucket_path" valid:"notnull"`
	Status           string    `json:"status" valid:"notnull"`
	Video            *Video    `json:"video" valid:"-"`
	VideoID          string    `json:"-" valid:"-" gorm:"column:video_id;type:uuid;notnull"`
	Error            string    `valid:"-"`
	CreatedAt        time.Time `json:"created_at" valid:"-"`
	UpdatedAt        time.Time `json:"updated_at" valid:"-"`
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
