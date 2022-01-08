package multiplication

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	MultiplicationService
}

func NewLoggingService(logger log.Logger, svdService MultiplicationService) MultiplicationService {
	return &loggingService{logger, svdService}
}

func (logService *loggingService) GetMultipliedMatrix(matrixHolder [][]float64) (_ [][]float64, err error) {
	defer func(begin time.Time) {
		logService.logger.Log(
			"Method", "GetSingularValueDecomposition",
			"Received At:", begin.Format(time.ANSIC),
			"Took", time.Since(begin),
			"Error", err,
		)
	}(time.Now())
	return logService.MultiplicationService.GetMultipliedMatrix(matrixHolder)
}
