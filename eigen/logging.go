package eigen

import (
	"go-microservices/matrix"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	EigenValueService
}

func NewLoggingService(logger log.Logger, eigenService EigenValueService) EigenValueService {
	return &loggingService{logger, eigenService}
}

func (logService *loggingService) GetEigenValues(matrixHolder *matrix.MatrixHolder) (_ []float64, err error) {
	defer func(begin time.Time) {
		logService.logger.Log(
			"Method:", "GetEigenValues",
			"Received At:", begin.Format(time.ANSIC),
			"Took:", time.Since(begin),
			"Error:", err,
		)
	}(time.Now())
	return logService.EigenValueService.GetEigenValues(matrixHolder)
}
