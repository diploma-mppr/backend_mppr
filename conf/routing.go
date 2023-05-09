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
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"},                                   // "http://127.0.0.1:3000", "http://localhost:3001", "http://localhost:3000", "http://127.0.0.1:3001"
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderSetCookie}, // , echo.HeaderSetCookie
		AllowMethods:     []string{echo.GET, echo.POST},
		MaxAge:           86400,
	}))
	mwChain := []echo.MiddlewareFunc{
		mw.AuthMiddleware,
	}

	router.GET("metrics", echo.WrapHandler(promhttp.Handler()))

	router.GET("/api/get_all", sh.HandlerData.GetAll, mwChain...)
	//pareto
	router.GET("/api/get_pareto", sh.HandlerPareto.GetPareto, mwChain...)
	router.POST("/api/set_pareto", sh.HandlerPareto.SetPareto, mwChain...)
	router.POST("/api/update_pareto", sh.HandlerPareto.UpdatePareto, mwChain...)
	router.POST("/api/delete_pareto", sh.HandlerPareto.DeletePareto, mwChain...)
	//basicCriteria
	router.GET("/api/get_base_criteria", sh.HandlerBaseCriteria.GetBaseCriteria, mwChain...)
	router.POST("/api/set_base_criteria", sh.HandlerBaseCriteria.SetBaseCriteria, mwChain...)
	router.POST("/api/update_base_criteria", sh.HandlerBaseCriteria.UpdateBaseCriteria, mwChain...)
	router.POST("/api/delete_base_criteria", sh.HandlerBaseCriteria.DeleteBaseCriteria, mwChain...)
	//pointScore
	router.GET("/api/get_point_score", sh.HandlerPointScore.GetPointScore, mwChain...)
	router.POST("/api/set_point_score", sh.HandlerPointScore.SetPointScore, mwChain...)
	router.POST("/api/update_point_score", sh.HandlerPointScore.UpdatePointScore, mwChain...)
	router.POST("/api/delete_point_score", sh.HandlerPointScore.DeletePointScore, mwChain...)
	//pairComparisonCriteria
	router.GET("/api/get_pair_comparison_criteria", sh.HandlerPairComparisonCriteria.GetPairComparisonCriteria, mwChain...)
	router.POST("/api/set_pair_comparison_criteria", sh.HandlerPairComparisonCriteria.SetPairComparisonCriteria, mwChain...)
	router.POST("/api/update_pair_comparison_criteria", sh.HandlerPairComparisonCriteria.UpdatePairComparisonCriteria, mwChain...)
	router.POST("/api/delete_pair_comparison_criteria", sh.HandlerPairComparisonCriteria.DeletePairComparisonCriteria, mwChain...)
	//borda
	router.GET("/api/get_borda", sh.HandlerBorda.GetBorda, mwChain...)
	router.POST("/api/set_borda", sh.HandlerBorda.SetBorda, mwChain...)
	router.POST("/api/update_borda", sh.HandlerBorda.UpdateBorda, mwChain...)
	router.POST("/api/delete_borda", sh.HandlerBorda.DeleteBorda, mwChain...)
	//nanson
	router.GET("/api/get_nanson", sh.HandlerNanson.GetNanson, mwChain...)
	router.POST("/api/set_nanson", sh.HandlerNanson.SetNanson, mwChain...)
	router.POST("/api/update_nanson", sh.HandlerNanson.UpdateNanson, mwChain...)
	router.POST("/api/delete_nanson", sh.HandlerNanson.DeleteNanson, mwChain...)
	//weightedSum
	router.GET("/api/get_weighted_sum", sh.HandlerWeightedSum.GetWeightedSum, mwChain...)
	router.POST("/api/set_weighted_sum", sh.HandlerWeightedSum.SetWeightedSum, mwChain...)
	router.POST("/api/update_weighted_sum", sh.HandlerWeightedSum.UpdateWeightedSum, mwChain...)
	router.POST("/api/delete_weighted_sum", sh.HandlerWeightedSum.DeleteWeightedSum, mwChain...)
	// auth
	router.POST("/api/register", sh.AuthHandler.Register, mwChain...)
	router.POST("/api/login", sh.AuthHandler.Login, mwChain...)
	router.GET("/api/logout", sh.AuthHandler.Logout, mwChain...)
	router.GET("/api/get_user", sh.AuthHandler.GetUserById, mwChain...)
}
