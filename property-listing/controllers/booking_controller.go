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


    // "gorm.io/gorm"
    "regexp"
    "strconv"
   
    
)

// BookingController handles all booking.com API related operations
type BookingController struct {
	uniqueCountries map[string]bool
	uniqueCities    map[string]bool
	countryCities   map[string][]string
	cityProperties  map[string][]string
	mutex           sync.Mutex
}

// NewBookingController creates a new instance of BookingController
func NewBookingController() *BookingController {
	return &BookingController{
		uniqueCountries: make(map[string]bool),
		uniqueCities:    make(map[string]bool),
		countryCities:   make(map[string][]string),
		cityProperties:  make(map[string][]string),
	}
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

// ProcessAllCities processes all cities using the generated queries
func (c *BookingController) ProcessAllCities() error {
	queries := c.generateQueries()
	results := make(chan struct{}, len(queries))
	semaphore := make(chan struct{}, 10)

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
	apiURL := fmt.Sprintf("https://booking-com18.p.rapidapi.com/stays/auto-complete?query=%s", query)
	
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "79d933f58amsh0baa13f673b03f0p16d4a2jsnb299a967d295")

	client := &http.Client{}
	resp, err := client.Do(req)
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
	propertyResults := make(chan struct {
		City       models.CityKey
		Properties []models.Property
		Err        error
	}, len(c.uniqueCities))

	semaphore := make(chan struct{}, 5)
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

	for _, prop := range result.Properties[:maxProperties] {
		c.cityProperties[result.City.Name] = append(
			c.cityProperties[result.City.Name], 
			fmt.Sprintf("%s (Score: %s, Class: %.1f, City: %s, Country: %s)", 
				prop.Name, 
				prop.ReviewScoreWord, 
				prop.PropertyClass,
				prop.CityName,
				prop.Country,
			),
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
    req, err := http.NewRequest("GET", apiURL, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
    req.Header.Add("x-rapidapi-key", "79d933f58amsh0baa13f673b03f0p16d4a2jsnb299a967d295")

    client := &http.Client{
        Timeout: 10 * time.Second,
    }
    
    resp, err := client.Do(req)
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
    c.mutex.Lock()
    defer c.mutex.Unlock()

    // Begin transaction
    tx := conf.DB.Begin()

    for country, cities := range c.countryCities {
        for _, cityName := range cities {
            // Prepare location with properties
            location := models.Location{
                CityName:   cityName,
                Country:    country,
            }

            // Get properties for this city
            if properties, ok := c.cityProperties[cityName]; ok {
                // Parse and create property records
                for _, propStr := range properties {
                    // Parse the property string
                    // Assuming format: "Name (Score: ReviewScore, Class: PropertyClass, City: CityName, Country: Country)"
                    prop, err := parsePropertyString(propStr)
                    if err != nil {
                        log.Printf("Error parsing property string: %v", err)
                        continue
                    }

                    // Add property to location
                    location.Properties = append(location.Properties, prop)
                }
            }

            // Create location with associated properties in a single transaction
            if err := tx.Create(&location).Error; err != nil {
                tx.Rollback()
                return fmt.Errorf("failed to create location with properties: %v", err)
            }
        }
    }

    // Commit transaction
    if err := tx.Commit().Error; err != nil {
        return fmt.Errorf("failed to commit transaction: %v", err)
    }

    return nil
}

// Helper function to parse property string
func parsePropertyString(propStr string) (Property, error) {
    // Regular expression to extract details
    re := regexp.MustCompile(`^(.*) \(Score: (.*), Class: ([\d.]+), City: (.*), Country: (.*)\)$`)
    matches := re.FindStringSubmatch(propStr)
    
    if len(matches) < 6 {
        return Property{}, fmt.Errorf("invalid property string format: %s", propStr)
    }

    // Convert property class to float
    propertyClass, err := strconv.ParseFloat(matches[3], 64)
    if err != nil {
        return Property{}, fmt.Errorf("error parsing property class: %v", err)
    }

    return Property{
        Name:           matches[1],
        ReviewScore:    matches[2],
        PropertyClass:  propertyClass,
        // You might want to add more parsing logic for other fields
    }, nil
}

// Optional: Method to bulk insert with optimized performance
func (c *BookingController) BulkSaveToDatabase() error {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    // Prepare bulk locations and properties
    var locations []Location
    
    for country, cities := range c.countryCities {
        for _, cityName := range cities {
            location := Location{
                CityName: cityName,
                Country:  country,
            }

            if properties, ok := c.cityProperties[cityName]; ok {
                for _, propStr := range properties {
                    prop, err := parsePropertyString(propStr)
                    if err != nil {
                        log.Printf("Error parsing property string: %v", err)
                        continue
                    }
                    location.Properties = append(location.Properties, prop)
                }
            }

            locations = append(locations, location)
        }
    }

    // Bulk create locations with associated properties
    if err := conf.DB.Create(&locations).Error; err != nil {
        return fmt.Errorf("failed to bulk create locations: %v", err)
    }

    return nil
}