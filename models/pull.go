package models

// @BasePath /api/v1
// HealthCheck godoc
// @Summary Pull a github repository down.
// @Description Pull a github repository down.
// @Tags root
// @Accept application/json
// @Produce json
// @Param   branch body string true "Branch Name"
// @Success 200 {object} map[string]interface{}
// @Router /pull [post]
type pull struct { 
        // We don't need a destination here as we will be using a standardised destination on the server
	Url string `json:"url" binding:"required"`
	Branch string `json:"branch" binding:"required"`
}

