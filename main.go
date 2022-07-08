package main

import (
	"fmt"
        "os"
        "path"
	"net/http"
        "net/url"
        "github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
        "github.com/swaggo/files"
        swaggerFiles "github.com/swaggo/files"
        "github.com/swaggo/gin-swagger"
        "./docs"
)


func homeLink(c *gin.Context) {
        // send to swagger docs
        c.JSON(http.StatusOK, gin.H{"data": "Welcome Home!"})
}

func gitPull(c *gin.Context) {
        // check the filesystem.
        // if it exists just do a pull
        // otherwise do a clone
        //var json  struct - should be externalised as part of model
        json := struct { 
        // We don't need a destination here as we will be using a standardised destination on the server
            Url string `json:"url" binding:"required"`
            Branch string `json:"branch" binding:"required"`
        }{}

        if err := c.BindJSON(&json); err == nil {
           // IF NO ERROR IN BINDING
           c.JSON(http.StatusOK, gin.H{ 
             "url": json.Url,
             "branch": json.Branch,
           })
           // Perform GIT pull as a first try
           targetUrl, err := url.Parse(json.Url)
           r, err := git.PlainClone("/apps/" + path.Base(targetUrl.Path), false, &git.CloneOptions{
              URL:      json.Url,
              Progress: os.Stdout,
           })
           fmt.Printf("goober: %s\n", r) 
           if err != nil {
           //At this point, if there is something in the repository it means it has been cloned before
           //We should do a pull to update it rather than a clone.
             c.JSON(http.StatusOK, gin.H{
                   "error": err.Error(),
                   "message": json.Url})
             return
           } // end if error on Plainclone
        } else {
           fmt.Printf("stuff: %s\n", json.Url)
           // IF WE HIT HERE IT MEANS THERE IS AN ERROR IN BINDING
           c.JSON(http.StatusBadRequest, gin.H{
                "error": "VALIDATEERR-1",
                "message": err.Error(),
           })
        } // end if/else bindJSON

}

func newApp(c *gin.Context) {
        // Send new request to local unit that configures an app
        // need language type as a varible.

        c.JSON(http.StatusOK, gin.H{"data": "NEW APP"})
}


// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func main() {
	router := gin.Default()
        router.GET("/", homeLink)
        router.POST("/pull", gitPull)
        router.GET("/app", newApp)
        router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

        router.Run(":8080")
}
