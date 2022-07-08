package models

type repository struct { 
        // We don't need a destination here as we will be using a standardised destination on the server
            Url string `json:"url" binding:"required"`
            Branch string `json:"branch" binding:"required"`
}

