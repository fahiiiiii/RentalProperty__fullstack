// controllers/booking_controller.go
package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
	"math"
    "property-listing/conf"
	"property-listing/models" // You'll need to adjust this import path based on your module name
    "golang.org/x/time/rate"
	"context"
	"github.com/joho/godotenv"
	"os"
    // "gorm.io/gorm"
    // "regexp"
    // "strconv"
   
    
)

// BookingController handles all booking.com API related operations
// type BookingController struct {
// 	uniqueCountries map[string]bool
// 	uniqueCities    map[string]bool
// 	countryCities   map[string][]string
// 	cityProperties  map[string][]string
// 	mutex           sync.Mutex
// }
type BookingController struct {
    uniqueCountries map[string]bool
    uniqueCities    map[string]bool
    countryCities   map[string][]string
    cityProperties  map[string][]string
    mutex           sync.Mutex
    rateLimiter    *rate.Limiter
	rapidAPIKey     string //!added rapidApiKey
}
// NewBookingController creates a new instance of BookingController
func NewBookingController() *BookingController {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the RapidAPI key from the environment variable
	rapidAPIKey := os.Getenv("RAPIDAPI_KEY")
	if rapidAPIKey == "" {
		log.Fatalf("RAPIDAPI_KEY is not set in the environment")
	}
	
	//! Created a rate limiter with 5 requests per minute
    // The first parameter (5) is the rate limit
    // The second parameter (1) is the burst size
    return &BookingController{
        uniqueCountries: make(map[string]bool),
        uniqueCities:    make(map[string]bool),
        countryCities:   make(map[string][]string),
        cityProperties:  make(map[string][]string),
        rateLimiter:    rate.NewLimiter(rate.Every(12*time.Second), 1), // 5 requests per minute = 1 request per 12 seconds
		rapidAPIKey:     rapidAPIKey, //! Initialize the rapidAPIKey field
    }
}
//! helper method for rate-limited requests
func (c *BookingController) makeRateLimitedRequest(req *http.Request) (*http.Response, error) {
    // Wait for rate limiter
    err := c.rateLimiter.Wait(context.Background())
    if err != nil {
        return nil, fmt.Errorf("rate limiter error: %v", err)
    }

    client := &http.Client{
        Timeout: 10 * time.Second,
    }
    return client.Do(req)
}
// GetSummary returns the current state of all data
func (c *BookingController) GetSummary() interface{} {
	return struct {
		Countries      map[string]bool            `json:"countries"`
		Cities         map[string]bool            `json:"cities"`
		CountryCities  map[string][]string        `json:"country_cities"`
		CityProperties map[string][]string        `json:"city_properties"`
	}{
		Countries:      c.uniqueCountries,
		Cities:         c.uniqueCities,
		CountryCities:  c.countryCities,
		CityProperties: c.cityProperties,
	}
}



//! Modified ProcessAllCities to limit concurrent requests
func (c *BookingController) ProcessAllCities() error {
    queries := c.generateQueries()
    results := make(chan struct{}, len(queries))
    
    // Reduce concurrent requests to 1 to better control rate limiting
    semaphore := make(chan struct{}, 1)

    for _, query := range queries {
        semaphore <- struct{}{}
        go func(q string) {
            defer func() { <-semaphore }()
            c.processCities(q, results)
        }(query)
    }

    // Wait for all queries to complete
    for range queries {
        <-results
    }

    return nil
}

// ProcessAllProperties processes properties for all cities
func (c *BookingController) ProcessAllProperties() error {
	return c.processPropertiesForCities()
}

// Private methods

func (c *BookingController) generateQueries() []string {
	queries := []string{}
	
	// Alphabet queries
	for char := 'A'; char <= 'Z'; char++ {
		queries = append(queries, string(char))
	}

	// Common prefixes and patterns
	prefixes := []string{
		"a", "the", "new", "old", "big", "small", 
		"north", "south", "east", "west", "central",
	}

	for _, prefix := range prefixes {
		queries = append(queries, prefix)
	}

	return queries
}

func (c *BookingController) processCities(query string, results chan<- struct{}) {
	defer func() { results <- struct{}{} }()

	cities, err := c.fetchCities(query)
	if err != nil {
		log.Printf("Error fetching cities for query '%s': %v", query, err)
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, city := range cities {
		country := strings.TrimSpace(strings.ToUpper(city.Country))
		cityName := strings.TrimSpace(strings.ToUpper(city.CityName))

		if country != "" {
			c.uniqueCountries[country] = true
		}

		if cityName != "" {
			c.uniqueCities[cityName] = true
		}

		if country != "" && cityName != "" {
			if _, exists := c.countryCities[country]; !exists {
				c.countryCities[country] = []string{}
			}
			
			cityExists := false
			for _, existingCity := range c.countryCities[country] {
				if existingCity == cityName {
					cityExists = true
					break
				}
			}
			
			if !cityExists {
				c.countryCities[country] = append(c.countryCities[country], cityName)
			}
		}
	}
}

func (c *BookingController) fetchCities(query string) ([]models.City, error) {
	
	//!using rate limit
	apiURL := fmt.Sprintf("https://booking-com18.p.rapidapi.com/stays/auto-complete?query=%s", query)
    
    req, err := http.NewRequest("GET", apiURL, nil)
    if err != nil {
        return nil, fmt.Errorf("error creating request: %v", err)
    }

    req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", c.rapidAPIKey) // Use the stored RapidAPI key
    // req.Header.Add("x-rapidapi-key", "79d933f58amsh0baa13f673b03f0p16d4a2jsnb299a967d295")

    resp, err := c.makeRateLimitedRequest(req)
    if err != nil {
        return nil, fmt.Errorf("error sending request: %v", err)
    }
    defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", 
			resp.StatusCode, string(body))
	}

	var citiesResp struct {
		Data []models.City `json:"data"`
	}
	err = json.Unmarshal(body, &citiesResp)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return citiesResp.Data, nil
}

func (c *BookingController) processPropertiesForCities() error {
	//!using limit for req
	propertyResults := make(chan struct {
        City       models.CityKey
        Properties []models.Property
        Err        error
    }, len(c.uniqueCities))

    // Reduce concurrent requests to 1 to better control rate limiting
    semaphore := make(chan struct{}, 1)
    var wg sync.WaitGroup

    for country, cities := range c.countryCities {
        for _, cityName := range cities {
            wg.Add(1)
            
            go func(city, country string) {
                defer wg.Done()
                
                semaphore <- struct{}{} 
                defer func() { <-semaphore }()

                properties, err := c.fetchPropertiesWithRetry(city, country, 3)
                
                propertyResults <- struct {
                    City       models.CityKey
                    Properties []models.Property
                    Err        error
                }{
                    City:       models.CityKey{Name: city, Country: country},
                    Properties: properties,
                    Err:        err,
                }
            }(cityName, country)
        }
    }

	go func() {
		wg.Wait()
		close(propertyResults)
	}()

	for result := range propertyResults {
		if result.Err != nil {
			log.Printf("Error fetching properties for %s, %s: %v", 
				result.City.Name, result.City.Country, result.Err)
			continue
		}

		c.processPropertyResult(result)
	}

	return nil
}

func (c *BookingController) processPropertyResult(result struct {
	City       models.CityKey
	Properties []models.Property
	Err        error
}) {
	if len(result.Properties) == 0 {
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cityProperties[result.City.Name] = []string{}
	
	maxProperties := 20
	if len(result.Properties) < maxProperties {
		maxProperties = len(result.Properties)
	}

	// for _, prop := range result.Properties[:maxProperties] {
	// 	c.cityProperties[result.City.Name] = append(
	// 		c.cityProperties[result.City.Name], 
	// 		fmt.Sprintf("%s (Score: %s, Class: %.1f, City: %s, Country: %s)", 
	// 			prop.Name, 
	// 			prop.ReviewScoreWord, 
	// 			prop.PropertyClass,
	// 			prop.CityName,
	// 			prop.Country,
	// 		),
	// 	)
	// }
	for _, prop := range result.Properties[:maxProperties] {
		// Store only the property name without additional details
		c.cityProperties[result.City.Name] = append(
			c.cityProperties[result.City.Name], 
			prop.Name, // Only store the property name
		)
	}
}

func (c *BookingController) fetchPropertiesWithRetry(cityName, country string, maxRetries int) ([]models.Property, error) {
    for attempt := 0; attempt < maxRetries; attempt++ {
        properties, err := c.fetchPropertiesForCity(cityName, country)
        
        if err == nil {
            return properties, nil
        }

        if strings.Contains(err.Error(), "Too many requests") || 
           strings.Contains(err.Error(), "You are not subscribed") {
            waitTime := time.Duration(math.Pow(2, float64(attempt))) * time.Second
            time.Sleep(waitTime)
            continue
        }

        return nil, err
    }

    return nil, fmt.Errorf("failed to fetch properties after %d attempts", maxRetries)
}

func (c *BookingController) fetchPropertiesForCity(cityName, country string) ([]models.Property, error) {
    uniqueProperties := make(map[string]models.Property)
    searchQueries := []string{
        cityName,
        fmt.Sprintf("%s hotels", cityName),
        fmt.Sprintf("%s accommodation", cityName),
    }

    for _, query := range searchQueries {
        encodedQuery := url.QueryEscape(query)
        apiURL := fmt.Sprintf("https://booking-com18.p.rapidapi.com/stays/auto-complete?query=%s", encodedQuery)
        
        properties, err := c.fetchPropertyData(apiURL)
        if err != nil {
            continue
        }

        for _, prop := range properties {
            if prop.DestID != "" {
                uniqueProperties[prop.DestID] = prop
            }
        }
    }

    result := make([]models.Property, 0, len(uniqueProperties))
    for _, prop := range uniqueProperties {
        result = append(result, prop)
    }

    return result, nil
}

func (c *BookingController) fetchPropertyData(apiURL string) ([]models.Property, error) {
    
	//!using rate limit
	req, err := http.NewRequest("GET", apiURL, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
    // req.Header.Add("x-rapidapi-key", "79d933f58amsh0baa13f673b03f0p16d4a2jsnb299a967d295")
	req.Header.Add("x-rapidapi-key", c.rapidAPIKey) // Use the stored RapidAPI key

    resp, err := c.makeRateLimitedRequest(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode == 429 || strings.Contains(string(body), "Too many requests") {
        return nil, fmt.Errorf("rate limit exceeded")
    }

    var response struct {
        Data []models.Property `json:"data"`
    }
    
    err = json.Unmarshal(body, &response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}











func (c *BookingController) SaveToDatabase() error {
    // Create a slice to store all locations
    var locations []models.Location

    // Iterate through cityProperties to get properties
    for city, properties := range c.cityProperties {
        // Find the country for this city
        var country string
        for countryName, cities := range c.countryCities {
            for _, cityName := range cities {
                if cityName == city {
                    country = countryName
                    break
                }
            }
        }

        // Add each property to locations
        for _, property := range properties {
            locations = append(locations, models.Location{
                Property: property,
                City:    city,
                Country: country,
            })
        }
    }

    // Batch insert locations
    result := conf.DB.CreateInBatches(locations, 100)
    if result.Error != nil {
        return fmt.Errorf("failed to save locations: %v", result.Error)
    }

    log.Printf("Successfully saved %d locations to database", len(locations))
    return nil
}





















//!To only update the existing data in the database without running the full scraping process again, 
//!you can create a new endpoint or function that just updates the property names. 
func (c *BookingController) UpdatePropertyNames() error {
    var locations []models.Location
    
    // Get all locations from database
    result := conf.DB.Find(&locations)
    if result.Error != nil {
        return fmt.Errorf("failed to fetch locations: %v", result.Error)
    }

    // Update each location's property name
    for i := range locations {
        // Extract just the property name by taking everything before " (Score:"
        if idx := strings.Index(locations[i].Property, " (Score:"); idx != -1 {
            locations[i].Property = locations[i].Property[:idx]
            
            // Save the updated record
            if err := conf.DB.Save(&locations[i]).Error; err != nil {
                log.Printf("Error updating location %d: %v", locations[i].ID, err)
                continue
            }
        }
    }

    log.Printf("Successfully updated %d property names", len(locations))
    return nil
}








