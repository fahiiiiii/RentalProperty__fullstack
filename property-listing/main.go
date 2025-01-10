
//!main.go

package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    
    "property-listing/conf"
    "property-listing/controllers"

    // "property-listing/routers"
    "github.com/beego/beego/v2/server/web"

    // "io/ioutil"
    
)

func main() {
    
    // Initialize database
    conf.InitDB()
    
    // Create a new controller instance
    controller := controllers.NewBookingController()

    web.Router("/", controller, "get:Index")
    web.Router("/v1/property/list", controller, "get:ListProperties")
    web.Router("/v1/property/details/:propertyId", controller, "get:GetPropertyDetails")

    // Serve static files
    web.SetStaticPath("/static", "static")
    web.Run()

    

    
    fmt.Println("Processing cities...")
    err := controller.ProcessAllCities()
    if err != nil {
        log.Fatalf("Failed to process cities: %v", err)
    }
    
    // Process properties for all cities
    fmt.Println("Processing properties...")
    err  = controller.ProcessAllProperties()
    if err != nil {
        log.Fatalf("Failed to process properties: %v", err)
    }
    
    // Save data to database
    fmt.Println("Saving data to database...")
    err = controller.SaveToDatabase()
    if err != nil {
        log.Fatalf("Failed to save to database: %v", err)
    }
    
    //!save rental property data
    // Save rental properties to database
    fmt.Println("Saving rental properties to database...")
    err = controller.SaveRentalPropertiesToDatabase() // Call the new function here
    if err != nil {
        log.Fatalf("Failed to save rental properties to database: %v", err)
    }



     //! Finally, process hotel details
     fmt.Println("Processing Hoteldetails...")
     err = controller.ProcessAllHotelDetails()
     if err != nil {
         log.Fatalf("Error processing hotel details: %v", err)
     }
     
     
     //! Process hotel descriptions
     fmt.Println("Processing hotel descriptions...")
     err = controller.ProcessAllHotelDescriptions() // Call to process hotel descriptions
     if err != nil {
         log.Fatalf("Error processing hotel descriptions: %v", err)
    }
    

    //!for fetching images
    err = controller.ProcessAllHotelImages()
    if err != nil {
        log.Fatalf("Failed to process hotel images: %v", err)
    }
    //!for rating review
    err = controller.ProcessAllHotelRatingsAndReviews()
    if err != nil {
        log.Fatalf("Failed to process hotel ratings and reviews: %v", err)
    }
    
    //!for stroring  desc,images and rating review
    
    fmt.Println("Saving rental properties to database...")
    controller.ProcessPropertyDetails() // Call the new function here
    if err != nil {
        log.Fatalf("Failed to save property details to database: %v", err)
    }
    
    
    // Get the final summary
    summary := controller.GetSummary()

    // Convert to JSON
    summaryJSON, err := json.MarshalIndent(summary, "", "  ")
    if err != nil {
        log.Fatalf("Failed to create JSON summary: %v", err)
    }
    
    // Save JSON to file
    err = os.WriteFile("city_summary.json", summaryJSON, 0644)
    if err != nil {
        log.Printf("Failed to save JSON summary: %v", err)
        } else {
        fmt.Println("JSON Summary saved to city_summary.json")
    }

    
    
    // //!for fetching againgst created api
    web.Run()
    
    
   
}




















// ----------------------------

// package main

// import (
//     "encoding/json"
//     "fmt"
//     "io/ioutil"
//     "log"
//     "net/http"
// )

// // HotelData represents the main structure for hotel information
// type HotelData struct {
//     Data struct {
//         HotelID              int    `json:"hotel_id"`
//         HotelName            string `json:"hotel_name"`
//         AccommodationTypeName string `json:"accommodation_type_name"`
//         Rooms                map[string]struct {
//             Description string `json:"description"`
//             PrivateBathroomCount int `json:"private_bathroom_count"`
//         } `json:"rooms"`
//         FacilitiesBlock struct {
//             Facilities []struct {
//                 Name string `json:"name"`
//             } `json:"facilities"`
//         } `json:"facilities_block"`
//     } `json:"data"`
// }

// // ExtractedHotelInfo represents the final structure we want to output
// type ExtractedHotelInfo struct {
//     HotelID    int      `json:"hotel_id"`
//     PropertyName string  `json:"property_name"`
//     Type        string  `json:"type"`
//     Bathroom    int     `json:"bathroom"`
//     Amenities   []string `json:"amenities"`
// }

// func main() {
//     // API configuration
//     url := "https://booking-com18.p.rapidapi.com/stays/detail"
//     apiKey := "3426928f2amsh2e4d4b598c36eefp10e72ajsn894381081bec"
//     apiHost := "booking-com18.p.rapidapi.com"

//     // Create HTTP client and request
//     client := &http.Client{}
//     req, err := http.NewRequest("GET", url, nil)
//     if err != nil {
//         log.Fatalf("Error creating request: %v", err)
//     }

//     // Add query parameters
//     q := req.URL.Query()
//     q.Add("hotelId", "56166")
//     q.Add("checkinDate", "2025-01-09")
//     q.Add("checkoutDate", "2025-01-23")
//     q.Add("units", "metric")
//     req.URL.RawQuery = q.Encode()

//     // Add headers
//     req.Header.Add("x-rapidapi-key", apiKey)
//     req.Header.Add("x-rapidapi-host", apiHost)

//     // Make the request
//     resp, err := client.Do(req)
//     if err != nil {
//         log.Fatalf("Error making request: %v", err)
//     }
//     defer resp.Body.Close()

//     // Read response body
//     body, err := ioutil.ReadAll(resp.Body)
//     if err != nil {
//         log.Fatalf("Error reading response: %v", err)
//     }

//     // Parse JSON response
//     var hotelData HotelData
//     if err := json.Unmarshal(body, &hotelData); err != nil {
//         log.Fatalf("Error parsing JSON: %v", err)
//     }

//     // Extract required information
//     extractedInfo := ExtractedHotelInfo{
//         HotelID:     hotelData.Data.HotelID,
//         PropertyName: hotelData.Data.HotelName,
//         Type:        hotelData.Data.AccommodationTypeName,
//         Amenities:   make([]string, 0),
//     }

//     // Extract bathroom count from first room (assuming it's the same for all rooms)
//     for _, room := range hotelData.Data.Rooms {
//         extractedInfo.Bathroom = room.PrivateBathroomCount
//         break
//     }

//     // Extract amenities
//     for _, facility := range hotelData.Data.FacilitiesBlock.Facilities {
//         if facility.Name != "" {
//             extractedInfo.Amenities = append(extractedInfo.Amenities, facility.Name)
//         }
//     }

//     // Convert to JSON and print
//     outputJSON, err := json.MarshalIndent(extractedInfo, "", "    ")
//     if err != nil {
//         log.Fatalf("Error converting to JSON: %v", err)
//     }

//     fmt.Println(string(outputJSON))
// }