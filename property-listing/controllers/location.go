// controllers/location.go
func (c *LocationController) FetchData() {
    apiKey, _ := beego.AppConfig.String("rapidapi_key")
    fetcher := services.NewDataFetcher(apiKey)
    
    err := fetcher.FetchAndStoreLocation("Dubai") // or any location
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "message": "Data fetched successfully",
        }
    }
    
    c.ServeJSON()
}