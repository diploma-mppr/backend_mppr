package conf

import (
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/auth/AuthHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/baseCriteria/BaseCriteriaHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/borda/BordaHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/data/DataHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/middleware"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/nanson/NansonHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pairComparisonCriteria/PairComparisonCriteriaHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pareto/ParetoHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/pointScore/PointScoreHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/app/weightedSum/WeightedSumHandler"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type ServerHandlers struct {
	HandlerData                   *DataHandler.HandlerData
	HandlerPareto                 *ParetoHandler.HandlerPareto
	HandlerBaseCriteria           *BaseCriteriaHandler.HandlerBaseCriteria
	HandlerPointScore             *PointScoreHandler.HandlerPointScore
	HandlerPairComparisonCriteria *PairComparisonCriteriaHandler.HandlerPairComparisonCriteria
	HandlerBorda                  *BordaHandler.HandlerBorda
	HandlerNanson                 *NansonHandler.HandlerNanson
	HandlerWeightedSum            *WeightedSumHandler.HandlerWeightedSum
	AuthHandler                   *AuthHandler.HandlerAuth
}

func (sh *ServerHandlers) ConfigureRouting(router *echo.Echo, mw *middleware.CommonMiddleware) {
	router.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"https://study-ai.ru/"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{echo.GET, echo.POST},
		MaxAge:           86400,
	}))
	mwChain := []echo.MiddlewareFunc{
		mw.AuthMiddleware,
	}

	router.GET("/api/get_all", sh.HandlerData.GetAll)
	//pareto
	router.GET("/api/get_pareto", sh.HandlerPareto.GetPareto)
	router.POST("/api/set_pareto", sh.HandlerPareto.SetPareto)
	router.POST("/api/update_pareto", sh.HandlerPareto.UpdatePareto)
	router.POST("/api/delete_pareto", sh.HandlerPareto.DeletePareto)
	//basicCriteria
	router.GET("/api/get_base_criteria", sh.HandlerBaseCriteria.GetBaseCriteria)
	router.POST("/api/set_base_criteria", sh.HandlerBaseCriteria.SetBaseCriteria)
	router.POST("/api/update_base_criteria", sh.HandlerBaseCriteria.UpdateBaseCriteria)
	router.POST("/api/delete_base_criteria", sh.HandlerBaseCriteria.DeleteBaseCriteria)
	//pointScore
	router.GET("/api/get_point_score", sh.HandlerPointScore.GetPointScore)
	router.POST("/api/set_point_score", sh.HandlerPointScore.SetPointScore)
	router.POST("/api/update_point_score", sh.HandlerPointScore.UpdatePointScore)
	router.POST("/api/delete_point_score", sh.HandlerPointScore.DeletePointScore)
	//pairComparisonCriteria
	router.GET("/api/get_pair_comparison_criteria", sh.HandlerPairComparisonCriteria.GetPairComparisonCriteria)
	router.POST("/api/set_pair_comparison_criteria", sh.HandlerPairComparisonCriteria.SetPairComparisonCriteria)
	router.POST("/api/update_pair_comparison_criteria", sh.HandlerPairComparisonCriteria.UpdatePairComparisonCriteria)
	router.POST("/api/delete_pair_comparison_criteria", sh.HandlerPairComparisonCriteria.DeletePairComparisonCriteria)
	//borda
	router.GET("/api/get_borda", sh.HandlerBorda.GetBorda)
	router.POST("/api/set_borda", sh.HandlerBorda.SetBorda)
	router.POST("/api/update_borda", sh.HandlerBorda.UpdateBorda)
	router.POST("/api/delete_borda", sh.HandlerBorda.DeleteBorda)
	//nanson
	router.GET("/api/get_nanson", sh.HandlerNanson.GetNanson)
	router.POST("/api/set_nanson", sh.HandlerNanson.SetNanson)
	router.POST("/api/update_nanson", sh.HandlerNanson.UpdateNanson)
	router.POST("/api/delete_nanson", sh.HandlerNanson.DeleteNanson)
	//weightedSum
	router.GET("/api/get_weighted_sum", sh.HandlerWeightedSum.GetWeightedSum)
	router.POST("/api/set_weighted_sum", sh.HandlerWeightedSum.SetWeightedSum)
	router.POST("/api/update_weighted_sum", sh.HandlerWeightedSum.UpdateWeightedSum)
	router.POST("/api/delete_weighted_sum", sh.HandlerWeightedSum.DeleteWeightedSum)
	// auth
	router.POST("/api/register", sh.AuthHandler.Register, mwChain...)
	router.POST("/api/login", sh.AuthHandler.Login, mwChain...)
	router.GET("/api/logout", sh.AuthHandler.Logout, mwChain...)
	router.GET("/api/get_user", sh.AuthHandler.GetUserById, mwChain...)
}
