package svd

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	SvdService
}

func NewLoggingService(logger log.Logger, svdService SvdService) SvdService {
	return &loggingService{logger, svdService}
}

func (logService *loggingService) GetSingularValueDecomposition(matrixHolder [][]float64) (_ [][]float64, err error) {
	defer func(begin time.Time) {
		logService.logger.Log(
			"Method", "GetSingularValueDecomposition",
			"Received At:", begin.Format(time.ANSIC),
			"Took", time.Since(begin),
			"Error", err,
		)
	}(time.Now())
	return logService.SvdService.GetSingularValueDecomposition(matrixHolder)
}
