package interfaces

import (
	"net/http"

	"github.com/shirocola/carbon-manage-be/internal/usecase"

	"github.com/labstack/echo/v4"
)

type CarbonFootprintHandler struct {
	usecase *usecase.CarbonFootprintUsecase
}

func NewCarbonFootprintHandler(u *usecase.CarbonFootprintUsecase) *CarbonFootprintHandler {
	return &CarbonFootprintHandler{usecase: u}
}

func (h *CarbonFootprintHandler) CreateCarbonFootprint(c echo.Context) error {
	var input struct {
		Source    string  `json:"source"`
		Emissions float64 `json:"emissions"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	err := h.usecase.CalculateAndSaveCarbonFootprint(input.Source, input.Emissions)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Carbon footprint saved successfully"})
}
