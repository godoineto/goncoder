package domain_test

import (
	"encoder/domain"
	"fmt"
	"os"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNotAnUUID(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "AAAAA"
	video.ResourceID = "BBBB"
	video.FilePath = "path.mp4"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}

func TestVideoValidateCorrectly(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	fmt.Fprintf(os.Stderr, "ERRO: %v\n", err)
	require.Nil(t, err)
}
