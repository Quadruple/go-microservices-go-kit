package determinant

import (
	"go-microservices/matrix"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	DeterminantService
}

func NewLoggingService(logger log.Logger, detService DeterminantService) DeterminantService {
	return &loggingService{logger, detService}
}

func (logService *loggingService) GetDeterminant(matrixHolder *matrix.MatrixHolder) (_ float64, err error) {
	defer func(begin time.Time) {
		logService.logger.Log(
			"Method:", "GetDeterminant",
			"Received At:", begin.Format(time.ANSIC),
			"Took:", time.Since(begin),
			"Error:", err,
		)
	}(time.Now())
	return logService.DeterminantService.GetDeterminant(matrixHolder)
}
