package main

import (
	"gitgub.com/diploma-mppr/backend_mppr/conf"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/auth/AuthHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/baseCriteria/BaseCriteriaHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/baseCriteria/BaseCriteriaRepository"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/baseCriteria/BaseCriteriaUseCase"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/borda/BordaHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/borda/BordaRepository"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/borda/BordaUseCase"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/data/DataHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/data/DataRepository"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/data/DataUseCase"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/metrics"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/middleware"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/nanson/NansonHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/nanson/NansonRepository"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/nanson/NansonUseCase"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pairComparisonCriteria/PairComparisonCriteriaHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pairComparisonCriteria/PairComparisonCriteriaRepository"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pairComparisonCriteria/PairComparisonCriteriaUseCase"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pareto/ParetoHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pareto/ParetoRepository"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pareto/ParetoUseCase"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pointScore/PointScoreHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pointScore/PointScoreRepository"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pointScore/PointScoreUseCase"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/weightedSum/WeightedSumHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/weightedSum/WeightedSumRepository"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/weightedSum/WeightedSumUseCase"
	"gitgub.com/diploma-mppr/backend_mppr/tools"
	"gitgub.com/diploma-mppr/backend_mppr/tools/authManager/jwtManager"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func main() {
	pgxManager, err := tools.NewPostgresqlX()
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating postgres agent"))
	}
	defer pgxManager.Close()

	RepositoryData := DataRepository.NewRepositoryData(pgxManager)
	UseCaseData := DataUseCase.NewUseCaseData(RepositoryData)
	HandlerData := DataHandler.NewHandlerData(UseCaseData)

	RepositoryPareto := ParetoRepository.NewRepositoryPareto(pgxManager)
	UseCasePareto := ParetoUseCase.NewUseCasePareto(RepositoryPareto)
	HandlerPareto := ParetoHandler.NewHandlerPareto(UseCasePareto)

	RepositoryBaseCriteria := BaseCriteriaRepository.NewRepositoryBaseCriteria(pgxManager)
	UseCaseBaseCriteria := BaseCriteriaUseCase.NewUseCaseBaseCriteria(RepositoryBaseCriteria)
	HandlerBaseCriteria := BaseCriteriaHandler.NewHandlerBaseCriteria(UseCaseBaseCriteria)

	RepositoryPointScore := PointScoreRepository.NewRepositoryPointScore(pgxManager)
	UseCasePointScore := PointScoreUseCase.NewUseCasePointScore(RepositoryPointScore)
	HandlerPointScore := PointScoreHandler.NewHandlerPointScore(UseCasePointScore)

	RepositoryPairComparisonCriteria := PairComparisonCriteriaRepository.NewRepositoryPairComparisonCriteria(pgxManager)
	UseCasePairComparisonCriteria := PairComparisonCriteriaUseCase.NewUseCasePairComparisonCriteria(RepositoryPairComparisonCriteria)
	HandlerPairComparisonCriteria := PairComparisonCriteriaHandler.NewHandlerPairComparisonCriteria(UseCasePairComparisonCriteria)

	RepositoryBorda := BordaRepository.NewRepositoryBorda(pgxManager)
	UseCaseBorda := BordaUseCase.NewUseCaseBorda(RepositoryBorda)
	HandlerBorda := BordaHandler.NewHandlerBorda(UseCaseBorda)

	RepositoryNanson := NansonRepository.NewRepositoryNanson(pgxManager)
	UseCaseNanson := NansonUseCase.NewUseCaseNanson(RepositoryNanson)
	HandlerNanson := NansonHandler.NewHandlerNanson(UseCaseNanson)

	RepositoryWeightedSum := WeightedSumRepository.NewRepositoryWeightedSum(pgxManager)
	UseCaseWeightedSum := WeightedSumUseCase.NewUseCaseWeightedSum(RepositoryWeightedSum)
	HandlerWeightedSum := WeightedSumHandler.NewHandlerWeightedSum(UseCaseWeightedSum)

	jwtManager := jwtManager.NewJwtManager()

	authHandler := AuthHandler.NewHandlerAuth(jwtManager)

	router := echo.New()

	m, err := metrics.CreateNewMetric("main")
	if err != nil {
		panic(err)
	}

	router.Use(m.CollectMetrics)

	serverRouting := conf.ServerHandlers{
		HandlerData:                   HandlerData,
		HandlerPareto:                 HandlerPareto,
		HandlerBaseCriteria:           HandlerBaseCriteria,
		HandlerPointScore:             HandlerPointScore,
		HandlerPairComparisonCriteria: HandlerPairComparisonCriteria,
		HandlerBorda:                  HandlerBorda,
		HandlerNanson:                 HandlerNanson,
		HandlerWeightedSum:            HandlerWeightedSum,
		AuthHandler:                   authHandler,
	}

	comonMw := middleware.NewCommonMiddleware(jwtManager)
	serverRouting.ConfigureRouting(router, &comonMw)

	httpServ := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	if err := router.StartServer(&httpServ); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
