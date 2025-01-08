// // -----------------db cionnection---------------

// // package main

// // import (
// //     "database/sql"
// //     "encoding/json"
// //     "log"
// //     "net/http"
// //     "os"
// // 	"io"
// //     "github.com/lib/pq" // Import pq explicitly for array handling
// //     _ "github.com/lib/pq" // Import the PostgreSQL driver
// //     "property-listing/models" // Adjust the import path based on your project structure
// // )

// // const (
// //     apiKey = "ba3e9fc45bmshe789fe44546bf15p1effc1jsn5a5396d69765" // Replace with your RapidAPI key
// //     apiUrl = "https://booking-com.p.rapidapi.com/v1/property/list" // Example endpoint
// // )

// // func fetchData() ([]models.Location, []models.RentalProperty) {
// //     req, err := http.NewRequest("GET", apiUrl, nil)
// //     if err != nil {
// //         log.Fatal(err)
// //     }
// //     req.Header.Add("x-rapidapi-key", apiKey)
// //     req.Header.Add("x-rapidapi-host", "booking-com.p.rapidapi.com")

// //     resp, err := http.DefaultClient.Do(req)
// //     if err != nil {
// //         log.Fatal(err)
// //     }
// //     defer resp.Body.Close()


// // 	respBody, err := io.ReadAll(resp.Body)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	log.Println("Raw Response Body:", string(respBody)) // Log the raw response
// //     // var rentalProperties []models.RentalProperty
// //     // if err := json.NewDecoder(resp.Body).Decode(&rentalProperties); err != nil {
// //     //     log.Fatal(err)
// //     // }
// // 	var rentalProperties []models.RentalProperty
// // 	if err := json.Unmarshal(respBody, &rentalProperties); err != nil {
// // 		log.Fatal(err)
// // 	}

// //     locationMap := make(map[string]models.Location)
// //     for _, prop := range rentalProperties {
// //         locationMap[prop.Name] = models.Location{Name: prop.Name}
// //     }

// //     locations := make([]models.Location, 0, len(locationMap))
// //     for _, loc := range locationMap {
// //         locations = append(locations, loc)
// //     }

// //     return locations, rentalProperties
// // }

// // func main() {
// //     // Manually set environment variables
// //     os.Setenv("DB_HOST", "localhost")      // Change to your DB host if needed
// //     os.Setenv("DB_PORT", "5432")            // Default PostgreSQL port
// //     os.Setenv("DB_USER", "fahimah")         // Your DB username
// //     os.Setenv("DB_PASSWORD", "fahimah123")  // Your DB password
// //     os.Setenv("DB_NAME", "rental_db")       // Your DB name

// //     // Set up database connection
// //     connStr := "host=" + os.Getenv("DB_HOST") +
// //         " user=" + os.Getenv("DB_USER") +
// //         " password=" + os.Getenv("DB_PASSWORD") +
// //         " dbname=" + os.Getenv("DB_NAME") +
// //         " port=" + os.Getenv("DB_PORT") +
// //         " sslmode=disable"

// //     log.Println("Connection String:", connStr)

// //     db, err := sql.Open("postgres", connStr)
// //     if err != nil {
// //         log.Fatal(err)
// //     }
// //     defer db.Close()

// //     // Test the connection
// //     err = db.Ping()
// //     if err != nil {
// //         log.Fatal(err)
// //     }
// //     log.Println("Successfully connected to the database!")

// //     locations, rentalProperties := fetchData()

// //     // Insert locations into the database
// //     for _, location := range locations {
// //         _, err := db.Exec("INSERT INTO locations (name) VALUES ($1) ON CONFLICT (name) DO NOTHING", location.Name)
// //         if err != nil {
// //             log.Println("Error inserting location:", err)
// //         }
// //     }

// //     // Insert rental properties into the database
// //     for _, prop := range rentalProperties {
// //         var locationID int
// //         err := db.QueryRow("SELECT id FROM locations WHERE name = $1", prop.Name).Scan(&locationID)
// //         if err != nil {
// //             log.Println("Error fetching location ID:", err)
// //             continue
// //         }

// //         _, err = db.Exec("INSERT INTO rental_properties (name, type, bedrooms, bathrooms, amenities, location_id) VALUES ($1, $2, $3, $4, $5, $6)",
// //             prop.Name, prop.Type, prop.Bedrooms, prop.Bathrooms, pq.Array(prop.Amenities), locationID)
// //         if err != nil {
// //             log.Println("Error inserting rental property:", err)
// //         }
// //     }

// //     log.Println("Data inserted successfully!")
// // }









// // // ---------------country code extraction from lamguage--------------------
// // package main

// // import (
// // 	"encoding/json"
// // 	"fmt"
// // 	"io"
// // 	"log"
// // 	"net/http"
// // 	"strings"
// // 	"regexp"
// // )

// // // Language struct to match the API response
// // type Language struct {
// // 	Typename         string `json:"__typename"`
// // 	Code             string `json:"code"`
// // 	CodeAirportTaxis string `json:"codeAirportTaxis"`
// // 	CountryFlag      string `json:"countryFlag"`
// // 	Name             string `json:"name"`
// // }

// // // LanguagesResponse represents the structure of the API response
// // type LanguagesResponse struct {
// // 	Data []Language `json:"data"`
// // }

// // // Function to fetch languages from Booking.com API
// // func fetchLanguages() ([]Language, error) {
// // 	// API endpoint and headers
// // 	url := "https://booking-com18.p.rapidapi.com/languages"
	
// // 	// Create a new HTTP request
// // 	req, err := http.NewRequest("GET", url, nil)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error creating request: %v", err)
// // 	}

// // 	// Set headers
// // 	req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
// // 	req.Header.Add("x-rapidapi-key", "47ae2e1dd1mshc33c535e5f35902p1c98e3jsn28a41b10eef3")

// // 	// Create HTTP client and send request
// // 	client := &http.Client{}
// // 	resp, err := client.Do(req)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error sending request: %v", err)
// // 	}
// // 	defer resp.Body.Close()

// // 	// Read response body
// // 	body, err := io.ReadAll(resp.Body)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error reading response body: %v", err)
// // 	}

// // 	// Check response status
// // 	if resp.StatusCode != http.StatusOK {
// // 		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", 
// // 			resp.StatusCode, string(body))
// // 	}

// // 	// Parse JSON response
// // 	var languagesResp LanguagesResponse
// // 	err = json.Unmarshal(body, &languagesResp)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error parsing JSON: %v", err)
// // 	}

// // 	return languagesResp.Data, nil
// // }

// // // Function to extract country from language
// // func extractCountryFromLanguage(lang Language) string {
// // 	// Method 1: Extract from parentheses
// // 	re := regexp.MustCompile(`\(([^)]+)\)$`)
// // 	matches := re.FindStringSubmatch(lang.Name)
// // 	if len(matches) > 1 {
// // 		return strings.ToUpper(matches[1])
// // 	}

// // 	// Method 2: Use country flag
// // 	if lang.CountryFlag != "" {
// // 		return strings.ToUpper(lang.CountryFlag)
// // 	}

// // 	// Method 3: Extract from code
// // 	parts := strings.Split(lang.Code, "-")
// // 	if len(parts) > 1 {
// // 		return strings.ToUpper(parts[1])
// // 	}

// // 	return ""
// // }

// // func main() {
// // 	// Fetch languages
// // 	languages, err := fetchLanguages()
// // 	if err != nil {
// // 		log.Fatalf("Failed to fetch languages: %v", err)
// // 	}

// // 	// Process and print languages
// // 	fmt.Println("Total Languages:", len(languages))
	
// // 	for _, lang := range languages {
// // 		country := extractCountryFromLanguage(lang)
// // 		fmt.Printf("Language: %s, Code: %s, Country: %s\n", 
// // 			lang.Name, lang.Code, country)
// // 	}

// // 	// Optional: Export to JSON file
// // 	languagesJSON, err := json.MarshalIndent(languages, "", "  ")
// // 	if err != nil {
// // 		log.Fatalf("Failed to convert to JSON: %v", err)
// // 	}
	
// // 	fmt.Println("\nJSON Export:\n", string(languagesJSON))
// // }



// // // ----------------------------city etraction from auto-complete----------------------------package main
// // package main
// // import (
// // 	"encoding/json"
// // 	"fmt"
// // 	"io"
// // 	"log"
// // 	"net/http"
// // 	"strings"
// // 	"net/url"
// // )

// // // City struct to match the API response
// // type City struct {
// // 	CC1        string  `json:"cc1"`
// // 	ImageURL   string  `json:"image_url"`
// // 	Longitude  float64 `json:"longitude"`
// // 	CityName   string  `json:"city_name"`
// // 	DestID     string  `json:"dest_id"`
// // 	Timezone   string  `json:"timezone"`
// // 	Hotels     int     `json:"hotels"`
// // 	Label      string  `json:"label"`
// // 	Country    string  `json:"country"`
// // 	Region     string  `json:"region"`
// // 	DestType   string  `json:"dest_type"`
// // 	Name       string  `json:"name"`
// // 	Latitude   float64 `json:"latitude"`
// // 	Type       string  `json:"type"`
// // }

// // // CitiesResponse represents the structure of the API response
// // type CitiesResponse struct {
// // 	Data []City `json:"data"`
// // }

// // // Function to fetch cities from Booking.com API
// // func fetchCities(query string) ([]City, error) {
// // 	// Encode the query parameter
// // 	encodedQuery := url.QueryEscape(query)
	
// // 	// API endpoint and headers
// // 	apiURL := fmt.Sprintf("https://booking-com18.p.rapidapi.com/stays/auto-complete?query=%s", encodedQuery)
	
// // 	// Create a new HTTP request
// // 	req, err := http.NewRequest("GET", apiURL, nil)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error creating request: %v", err)
// // 	}

// // 	// Set headers
// // 	req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
// // 	req.Header.Add("x-rapidapi-key", "47ae2e1dd1mshc33c535e5f35902p1c98e3jsn28a41b10eef3")

// // 	// Create HTTP client and send request
// // 	client := &http.Client{}
// // 	resp, err := client.Do(req)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error sending request: %v", err)
// // 	}
// // 	defer resp.Body.Close()

// // 	// Read response body
// // 	body, err := io.ReadAll(resp.Body)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error reading response body: %v", err)
// // 	}

// // 	// Check response status
// // 	if resp.StatusCode != http.StatusOK {
// // 		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", 
// // 			resp.StatusCode, string(body))
// // 	}

// // 	// Parse JSON response
// // 	var citiesResp CitiesResponse
// // 	err = json.Unmarshal(body, &citiesResp)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error parsing JSON: %v", err)
// // 	}

// // 	return citiesResp.Data, nil
// // }

// // // Function to extract country code
// // func extractCountryCode(city City) string {
// // 	if city.CC1 != "" {
// // 		return city.CC1
// // 	}
// // 	return ""
// // }

// // func main() {
// // 	// Example queries
// // 	queries := []string{
// // 		"New York",
// // 		"Los Angeles",
// // 		"Chicago",
// // 		// Add more cities as needed
// // 	}

// // 	// Fetch and process cities for each query
// // 	for _, query := range queries {
// // 		cities, err := fetchCities(query)
// // 		if err != nil {
// // 			log.Printf("Failed to fetch cities for query '%s': %v", query, err)
// // 			continue
// // 		}

// // 		fmt.Printf("\nResults for query: %s\n", query)
// // 		fmt.Println("Total Cities:", len(cities))
		
// // 		for _, city := range cities {
// // 			countryCode := extractCountryCode(city)
// // 			fmt.Printf("City: %s, Country: %s, Country Code: %s, Hotels: %d\n", 
// // 				city.CityName, city.Country, countryCode, city.Hotels)
// // 		}

// // 		// Optional: Export to JSON
// // 		citiesJSON, err := json.MarshalIndent(cities, "", "  ")
// // 		if err != nil {
// // 			log.Printf("Failed to convert to JSON: %v", err)
// // 			continue
// // 		}
		
// // 		fmt.Println("\nJSON Export:\n", string(citiesJSON))
// // 	}
// // }

// // // Additional helper functions can be added as needed
// // func normalizeCountryCode(code string) string {
// // 	return strings.ToUpper(code)
// // }

// // // Struct for more detailed city information
// // type DetailedCity struct {
// // 	Name        string
// // 	Country     string
// // 	CountryCode string
// // 	Coordinates struct {
// // 		Latitude  float64
// // 		Longitude float64
// // 	}
// // 	HotelCount int
// // 	Timezone   string
// // }

// // // Convert City to DetailedCity
// // func (c City) ToDetailedCity() DetailedCity {
// // 	return DetailedCity{
// // 		Name:        c.CityName,
// // 		Country:     c.Country,
// // 		CountryCode: normalizeCountryCode(c.CC1),
// // 		Coordinates: struct {
// // 			Latitude  float64
// // 			Longitude float64
// // 		}{
// // 			Latitude:  c.Latitude,
// // 			Longitude: c.Longitude,
// // 		},
// // 		HotelCount: c.Hotels,
// // 		Timezone:   c.Timezone,
// // 	}
// // }



// // ----------------CITIES AND COUNTRIES ETR5ACTION AND THEIR NUMBER AND LIST----------------------------
// // package main

// // import (
// // 	"encoding/json"
// // 	"fmt"
// // 	"io"
// // 	"log"
// // 	"net/http"
// // 	"os"
// // 	"strings"
// // 	"sync"
// // )

// // // City struct to match the API response
// // type City struct {
// // 	CC1        string  `json:"cc1"`
// // 	ImageURL   string  `json:"image_url"`
// // 	Longitude  float64 `json:"longitude"`
// // 	CityName   string  `json:"city_name"`
// // 	DestID     string  `json:"dest_id"`
// // 	Timezone   string  `json:"timezone"`
// // 	Hotels     int     `json:"hotels"`
// // 	Label      string  `json:"label"`
// // 	Country    string  `json:"country"`
// // 	Region     string  `json:"region"`
// // 	DestType   string  `json:"dest_type"`
// // 	Name       string  `json:"name"`
// // 	Latitude   float64 `json:"latitude"`
// // 	Type       string  `json:"type"`
// // }

// // // Global maps to track unique countries and cities
// // var (
// // 	uniqueCountries = make(map[string]bool)
// // 	uniqueCities    = make(map[string]bool)
// // 	countryCities   = make(map[string][]string)
// // 	mutex           sync.Mutex
// // )

// // // Function to generate queries dynamically
// // func generateQueries() []string {
// // 	// Start with a comprehensive list of alphabet and common prefixes
// // 	queries := []string{}
	
// // 	// Alphabet queries
// // 	for char := 'A'; char <= 'Z'; char++ {
// // 		queries = append(queries, string(char))
// // 	}

// // 	// Common prefixes and patterns
// // 	prefixes := []string{
// // 		"a", "the", "new", "old", "big", "small", 
// // 		"north", "south", "east", "west", "central",
// // 	}

// // 	for _, prefix := range prefixes {
// // 		queries = append(queries, prefix)
// // 	}

// // 	return queries
// // }

// // // Function to fetch cities from Booking.com API
// // func fetchCities(query string) ([]City, error) {
// // 	// Encode the query parameter
// // 	apiURL := fmt.Sprintf("https://booking-com18.p.rapidapi.com/stays/auto-complete?query=%s", query)
	
// // 	// Create a new HTTP request
// // 	req, err := http.NewRequest("GET", apiURL, nil)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error creating request: %v", err)
// // 	}

// // 	// Set headers
// // 	req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
// // 	req.Header.Add("x-rapidapi-key", "47ae2e1dd1mshc33c535e5f35902p1c98e3jsn28a41b10eef3")

// // 	// Create HTTP client and send request
// // 	client := &http.Client{}
// // 	resp, err := client.Do(req)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error sending request: %v", err)
// // 	}
// // 	defer resp.Body.Close()

// // 	// Read response body
// // 	body, err := io.ReadAll(resp.Body)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error reading response body: %v", err)
// // 	}

// // 	// Check response status
// // 	if resp.StatusCode != http.StatusOK {
// // 		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", 
// // 			resp.StatusCode, string(body))
// // 	}

// // 	// Parse JSON response
// // 	var citiesResp struct {
// // 		Data []City `json:"data"`
// // 	}
// // 	err = json.Unmarshal(body, &citiesResp)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error parsing JSON: %v", err)
// // 	}

// // 	return citiesResp.Data, nil
// // }

// // // Function to process cities and update global maps
// // func processCities(query string, results chan<- struct{}) {
// // 	defer func() { results <- struct{}{} }()

// // 	cities, err := fetchCities(query)
// // 	if err != nil {
// // 		log.Printf("Error fetching cities for query '%s': %v", query, err)
// // 		return
// // 	}

// // 	mutex.Lock()
// // 	defer mutex.Unlock()

// // 	for _, city := range cities {
// // 		// Normalize country and city names
// // 		country := strings.TrimSpace(strings.ToUpper(city.Country))
// // 		cityName := strings.TrimSpace(strings.ToUpper(city.CityName))

// // 		// Track unique countries
// // 		if country != "" {
// // 			uniqueCountries[country] = true
// // 		}

// // 		// Track unique cities
// // 		if cityName != "" {
// // 			uniqueCities[cityName] = true
// // 		}

// // 		// Track cities per country
// // 		if country != "" && cityName != "" {
// // 			if _, exists := countryCities[country]; !exists {
// // 				countryCities[country] = []string{}
// // 			}
			
// // 			// Avoid duplicate cities
// // 			cityExists := false
// // 			for _, existingCity := range countryCities[country] {
// // 				if existingCity == cityName {
// // 					cityExists = true
// // 					break
// // 				}
// // 			}
			
// // 			if !cityExists {
// // 				countryCities[country] = append(countryCities[country], cityName)
// // 			}
// // 		}
// // 	}
// // }

// // // Function to save JSON summary to file
// // func saveJSONSummary(summaryJSON []byte) error {
// // 	return os.WriteFile("city_summary.json", summaryJSON, 0644)
// // }

// // // Utility function to get minimum of two integers
// // func min(a, b int) int {
// // 	if a < b {
// // 		return a
// // 	}
// // 	return b
// // }

// // func main() {
// // 	// Generate queries dynamically
// // 	queries := generateQueries()

// // 	// Channel for tracking goroutine completion
// // 	results := make(chan struct{}, len(queries))

// // 	// Limit concurrent goroutines to prevent overwhelming the API
// // 	semaphore := make(chan struct{}, 10)

// // 	// Process queries
// // 	for _, query := range queries {
// // 		semaphore <- struct{}{}
// // 		go func(q string) {
// // 			defer func() { <-semaphore }()
// // 			processCities(q, results)
// // 		}(query)
// // 	}

// // 	// Wait for all queries to complete
// // 	for range queries {
// // 		<-results
// // 	}

// // 	// Print detailed results
// // 	fmt.Println("Extraction Summary:")
// // 	fmt.Println("-------------------")
// // 	fmt.Printf("Total Unique Countries: %d\n", len(uniqueCountries))
// // 	fmt.Printf("Total Unique Cities: %d\n", len(uniqueCities))
	
// // 	// Print countries and their cities
// // 	fmt.Println("\nCountries and Cities:")
// // 	for country, cities := range countryCities {
// // 		fmt.Printf("%s (%d cities):\n", country, len(cities))
// // 		for _, city := range cities {
// // 			fmt.Printf("  - %s\n", city)
// // 		}
// // 		fmt.Println()
// // 	}

// // 	// Prepare summary data
// // 	summaryData := struct {
// // 		Countries     map[string]bool            `json:"countries"`
// // 		Cities        map[string]bool            `json:"cities"`
// // 		CountryCities map[string][]string        `json:"country_cities"`
// // 	}{
// // 		Countries:     uniqueCountries,
// // 		Cities:        uniqueCities,
// // 		CountryCities: countryCities,
// // 	}

// // 	// Convert to JSON
// // 	summaryJSON, err := json.MarshalIndent(summaryData, "", "  ")
// // 	if err != nil {
// // 		log.Fatalf("Failed to create JSON summary: %v", err)
// // 	}

// // 	// Save JSON to file
// // 	err = saveJSONSummary(summaryJSON)
// // 	if err != nil {
// // 		log.Printf("Failed to save JSON summary: %v", err)
// // 	} else {
// // 		fmt.Println("JSON Summary saved to city_summary.json")
// // 	}

// // 	// Optional: Print a snippet of the JSON
// // 	fmt.Println("\nJSON Summary Snippet:")
// // 	fmt.Println(string(summaryJSON[:min(len(summaryJSON), 500)]))
// // }



// // // ---------------properties under cities-------------------
// // package main

// // import (
// // 	"encoding/json"
// // 	"fmt"
// // 	"io"
// // 	"log"
// // 	"net/http"
// // 	"os"
// // 	// "strings"
// // 	"sync"
// // 	"net/url"
// //     "strings"
// // 	"math"
// //     "time"
// //     "net/url"
// //     "strings"
// // )

// // //! Existing City struct remains the same
// // // City struct to match the API response
// // type City struct {
// // 	CC1        string  `json:"cc1"`
// // 	ImageURL   string  `json:"image_url"`
// // 	Longitude  float64 `json:"longitude"`
// // 	CityName   string  `json:"city_name"`
// // 	DestID     string  `json:"dest_id"`
// // 	Timezone   string  `json:"timezone"`
// // 	Hotels     int     `json:"hotels"`
// // 	Label      string  `json:"label"`
// // 	Country    string  `json:"country"`
// // 	Region     string  `json:"region"`
// // 	DestType   string  `json:"dest_type"`
// // 	Name       string  `json:"name"`
// // 	Latitude   float64 `json:"latitude"`
// // 	Type       string  `json:"type"`
// // }


// // // New Property struct
// // type Property struct {
// // 	UFI               int64   `json:"ufi"`
// // 	CheckoutDate      string  `json:"checkoutDate"`
// // 	ReviewScoreWord   string  `json:"reviewScoreWord"`
// // 	Longitude         float64 `json:"longitude"`
// // 	IsPreferred       bool    `json:"isPreferred"`
// // 	CountryCode       string  `json:"countryCode"`
// // 	Latitude          float64 `json:"latitude"`
// // 	WishlistName      string  `json:"wishlistName"`
// // 	Name              string  `json:"name"`
// // 	PropertyClass     float64 `json:"accuratePropertyClass"`
// // }

// // // Global maps to track properties
// // var (
// // 	uniqueCountries = make(map[string]bool)
// // 	uniqueCities    = make(map[string]bool)
// // 	countryCities   = make(map[string][]string)
// // 	cityProperties  = make(map[string][]string)
// // 	mutex           sync.Mutex
// // )

// // //! Existing functions remain the same...

// // // Function to generate queries dynamically
// // func generateQueries() []string {
// // 	// Start with a comprehensive list of alphabet and common prefixes
// // 	queries := []string{}
	
// // 	// Alphabet queries
// // 	for char := 'A'; char <= 'Z'; char++ {
// // 		queries = append(queries, string(char))
// // 	}

// // 	// Common prefixes and patterns
// // 	prefixes := []string{
// // 		"a", "the", "new", "old", "big", "small", 
// // 		"north", "south", "east", "west", "central",
// // 	}

// // 	for _, prefix := range prefixes {
// // 		queries = append(queries, prefix)
// // 	}

// // 	return queries
// // }

// // // Function to fetch cities from Booking.com API
// // func fetchCities(query string) ([]City, error) {
// // 	// Encode the query parameter
// // 	apiURL := fmt.Sprintf("https://booking-com18.p.rapidapi.com/stays/auto-complete?query=%s", query)
	
// // 	// Create a new HTTP request
// // 	req, err := http.NewRequest("GET", apiURL, nil)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error creating request: %v", err)
// // 	}

// // 	// Set headers
// // 	req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
// // 	req.Header.Add("x-rapidapi-key", "11253ee205msh7def7446d6fd7a0p1fac3ejsnb76491e66cf7")

// // 	// Create HTTP client and send request
// // 	client := &http.Client{}
// // 	resp, err := client.Do(req)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error sending request: %v", err)
// // 	}
// // 	defer resp.Body.Close()

// // 	// Read response body
// // 	body, err := io.ReadAll(resp.Body)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error reading response body: %v", err)
// // 	}

// // 	// Check response status
// // 	if resp.StatusCode != http.StatusOK {
// // 		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", 
// // 			resp.StatusCode, string(body))
// // 	}

// // 	// Parse JSON response
// // 	var citiesResp struct {
// // 		Data []City `json:"data"`
// // 	}
// // 	err = json.Unmarshal(body, &citiesResp)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error parsing JSON: %v", err)
// // 	}

// // 	return citiesResp.Data, nil
// // }

// // // Function to process cities and update global maps
// // func processCities(query string, results chan<- struct{}) {
// // 	defer func() { results <- struct{}{} }()

// // 	cities, err := fetchCities(query)
// // 	if err != nil {
// // 		log.Printf("Error fetching cities for query '%s': %v", query, err)
// // 		return
// // 	}

// // 	mutex.Lock()
// // 	defer mutex.Unlock()

// // 	for _, city := range cities {
// // 		// Normalize country and city names
// // 		country := strings.TrimSpace(strings.ToUpper(city.Country))
// // 		cityName := strings.TrimSpace(strings.ToUpper(city.CityName))

// // 		// Track unique countries
// // 		if country != "" {
// // 			uniqueCountries[country] = true
// // 		}

// // 		// Track unique cities
// // 		if cityName != "" {
// // 			uniqueCities[cityName] = true
// // 		}

// // 		// Track cities per country
// // 		if country != "" && cityName != "" {
// // 			if _, exists := countryCities[country]; !exists {
// // 				countryCities[country] = []string{}
// // 			}
			
// // 			// Avoid duplicate cities
// // 			cityExists := false
// // 			for _, existingCity := range countryCities[country] {
// // 				if existingCity == cityName {
// // 					cityExists = true
// // 					break
// // 				}
// // 			}
			
// // 			if !cityExists {
// // 				countryCities[country] = append(countryCities[country], cityName)
// // 			}
// // 		}
// // 	}
// // }

// // // Function to save JSON summary to file
// // func saveJSONSummary(summaryJSON []byte) error {
// // 	return os.WriteFile("city_summary.json", summaryJSON, 0644)
// // }

// // // Utility function to get minimum of two integers
// // func min(a, b int) int {
// // 	if a < b {
// // 		return a
// // 	}
// // 	return b
// // }












// // func fetchPropertiesWithRetry(cityName, country string, maxRetries int) ([]Property, error) {
// //     for attempt := 0; attempt < maxRetries; attempt++ {
// //         properties, err := fetchPropertiesForCity(cityName, country)
        
// //         if err == nil {
// //             return properties, nil
// //         }

// //         // Check for specific error conditions
// //         if strings.Contains(err.Error(), "Too many requests") || 
// //            strings.Contains(err.Error(), "You are not subscribed") {
            
// //             // Exponential backoff
// //             waitTime := time.Duration(math.Pow(2, float64(attempt))) * time.Second
// //             log.Printf("Rate limit hit. Waiting %v before retry", waitTime)
// //             time.Sleep(waitTime)
            
// //             continue
// //         }

// //         // For other errors, return immediately
// //         return nil, err
// //     }

// //     return nil, fmt.Errorf("failed to fetch properties after %d attempts", maxRetries)
// // }
// // // New function to fetch properties for a specific location
// // // Function to fetch properties for a specific location
// // func fetchPropertiesForCity(cityName, country string) ([]Property, error) {
// //     // URL encode the city name
// //     encodedCityName := url.QueryEscape(cityName)
    
// //     // Construct the API URL dynamically
// //     apiURL := fmt.Sprintf("https://booking-com18.p.rapidapi.com/stays/auto-complete?query=%s", encodedCityName)
    
// //     // Create a new HTTP request
// //     req, err := http.NewRequest("GET", apiURL, nil)
// //     if err != nil {
// //         return nil, fmt.Errorf("error creating request: %v", err)
// //     }

// //     // Set headers
// //     req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
// //     req.Header.Add("x-rapidapi-key", "11253ee205msh7def7446d6fd7a0p1fac3ejsnb76491e66cf7")

// //     // Create HTTP client with timeout
// //     client := &http.Client{
// //         Timeout: 10 * time.Second,
// //     }
// //     resp, err := client.Do(req)
// //     if err != nil {
// //         return nil, fmt.Errorf("error sending request: %v", err)
// //     }
// //     defer resp.Body.Close()

// //     // Read response body
// //     body, err := io.ReadAll(resp.Body)
// //     if err != nil {
// //         return nil, fmt.Errorf("error reading response body: %v", err)
// //     }

// //     // Check for rate limiting or subscription errors
// //     if resp.StatusCode == 429 || strings.Contains(string(body), "Too many requests") {
// //         return nil, fmt.Errorf("rate limit exceeded: %s", string(body))
// //     }

// //     if resp.StatusCode == 403 || strings.Contains(string(body), "not subscribed") {
// //         return nil, fmt.Errorf("API subscription error: %s", string(body))
// //     }

// //     // Parse JSON response
// //     var response struct {
// //         Data []Property `json:"data"`
// //     }
    
// //     err = json.Unmarshal(body, &response)
// //     if err != nil {
// //         return nil, fmt.Errorf("error parsing JSON: %v. Raw response: %s", err, string(body))
// //     }

// //     return response.Data, nil
// // }
// // // func fetchPropertiesForCity(locationID string) ([]Property, error) {
// // // 	// Construct API URL
// // // 	apiURL := fmt.Sprintf(
// // // 		"https://booking-com18.p.rapidapi.com/stays/search?locationId=%s&checkinDate=2025-01-08&checkoutDate=2025-01-15&units=metric&temperature=c", 
// // // 		locationID,
// // // 	)
	
// // // 	// Create a new HTTP request
// // // 	req, err := http.NewRequest("GET", apiURL, nil)
// // // 	if err != nil {
// // // 		return nil, fmt.Errorf("error creating request: %v", err)
// // // 	}

// // // 	// Set headers
// // // 	req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
// // // 	req.Header.Add("x-rapidapi-key", "47ae2e1dd1mshc33c535e5f35902p1c98e3jsn28a41b10eef3")

// // // 	// Create HTTP client and send request
// // // 	client := &http.Client{}
// // // 	resp, err := client.Do(req)
// // // 	if err != nil {
// // // 		return nil, fmt.Errorf("error sending request: %v", err)
// // // 	}
// // // 	defer resp.Body.Close()

// // // 	// Read response body
// // // 	body, err := io.ReadAll(resp.Body)
// // // 	if err != nil {
// // // 		return nil, fmt.Errorf("error reading response body: %v", err)
// // // 	}

// // // 	// Check response status
// // // 	if resp.StatusCode != http.StatusOK {
// // // 		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", 
// // // 			resp.StatusCode, string(body))
// // // 	}

// // // 	// Parse JSON response
// // // 	var propertiesResp struct {
// // // 		Data []Property `json:"data"`
// // // 	}
// // // 	err = json.Unmarshal(body, &propertiesResp)
// // // 	if err != nil {
// // // 		return nil, fmt.Errorf("error parsing JSON: %v", err)
// // // 	}

// // // 	return propertiesResp.Data, nil
// // // }

// // // Function to extract and store properties for cities
// // // Function to process properties for cities
// // func processPropertiesForCities() {
// //     // Create a channel for concurrent processing with controlled concurrency
// //     propertyResults := make(chan struct {
// //         City       CityKey
// //         Properties []Property
// //         Err        error
// //     }, len(uniqueCities))

// //     // Use a semaphore to limit concurrent API calls
// //     semaphore := make(chan struct{}, 5) // Limit to 5 concurrent requests

// //     // Track processed cities
// //     processedCities := 0

// //     // Process cities concurrently with rate limiting
// //     for country, cities := range countryCities {
// //         for _, cityName := range cities {
// //             processedCities++
            
// //             go func(city, country string) {
// //                 semaphore <- struct{}{} // Acquire semaphore
// //                 defer func() { <-semaphore }() // Release semaphore

// //                 properties, err := fetchPropertiesWithRetry(city, country, 3)
// //                 propertyResults <- struct {
// //                     City       CityKey
// //                     Properties []Property
// //                     Err        error
// //                 }{
// //                     City:       CityKey{Name: city, Country: country},
// //                     Properties: properties,
// //                     Err:        err,
// //                 }
// //             }(cityName, country)
// //         }
// //     }

// //     // Collect and process results
// //     for i := 0; i < processedCities; i++ {
// //         result := <-propertyResults
        
// //         if result.Err != nil {
// //             log.Printf("Error fetching properties for %s, %s: %v", 
// //                 result.City.Name, result.City.Country, result.Err)
// //             continue
// //         }

// //         if len(result.Properties) == 0 {
// //             log.Printf("No properties found for %s, %s", 
// //                 result.City.Name, result.City.Country)
// //             continue
// //         }

// //         // Store properties
// //         mutex.Lock()
// //         cityProperties[result.City.Name] = []string{}
// //         for _, prop := range result.Properties {
// //             cityProperties[result.City.Name] = append(
// //                 cityProperties[result.City.Name], 
// //                 fmt.Sprintf("%s (Score: %s)", prop.Name, prop.ReviewScoreWord)
// //             )
// //         }
// //         mutex.Unlock()
// //     }
// // }
// // // func processPropertiesForCities() {
// // // 	// Iterate through unique cities
// // // 	for country, cities := range countryCities {
// // // 		for _, cityName := range cities {
// // // 			// Find the location ID (you might need to modify this based on your city data)
// // // 			locationID := generateLocationID(cityName, country)
			
// // // 			// Fetch properties
// // // 			properties, err := fetchPropertiesForCity(locationID)
// // // 			if err != nil {
// // // 				log.Printf("Error fetching properties for %s, %s: %v", cityName, country, err)
// // // 				continue
// // // 			}

// // // 			// Store property names
// // // 			mutex.Lock()
// // // 			cityProperties[cityName] = []string{}
// // // 			for _, prop := range properties {
// // // 				cityProperties[cityName] = append(cityProperties[cityName], prop.Name)
// // // 			}
// // // 			mutex.Unlock()
// // // 		}
// // // 	}
// // // }

// // // Utility function to generate location ID (you might need to adjust this)
// // func generateLocationID(cityName, country string) string {
// // 	// This is a placeholder - you'll need to implement proper location ID generation
// // 	return fmt.Sprintf("eyJjaXR5X25hbWUiOiIlc1wiLCJjb3VudHJ5XCI6XCIlc1wiLCJkZXN0X2lkXCI6XCIyMDA4ODMyNVwiLCJkZXN0X3R5cGVcIjpcImNpdHlcIn0=", 
// // 		cityName, country)
// // }
// // func main() {
// // 	// Existing main function code for city extraction...
// // 	queries := generateQueries()

// // 	// Channel for tracking goroutine completion
// // 	results := make(chan struct{}, len(queries))

// // 	// Limit concurrent goroutines to prevent overwhelming the API
// // 	semaphore := make(chan struct{}, 10)

// // 	// Process queries
// // 	for _, query := range queries {
// // 		semaphore <- struct{}{}
// // 		go func(q string) {
// // 			defer func() { <-semaphore }()
// // 			processCities(q, results)
// // 		}(query)
// // 	}

// // 	// Wait for all queries to complete
// // 	for range queries {
// // 		<-results
// // 	}

// // 	// Print detailed results
// // 	fmt.Println("Extraction Summary:")
// // 	fmt.Println("-------------------")
// // 	fmt.Printf("Total Unique Countries: %d\n", len(uniqueCountries))
// // 	fmt.Printf("Total Unique Cities: %d\n", len(uniqueCities))
	
// // 	// Print countries and their cities
// // 	fmt.Println("\nCountries and Cities:")
// // 	for country, cities := range countryCities {
// // 		fmt.Printf("%s (%d cities):\n", country, len(cities))
// // 		for _, city := range cities {
// // 			fmt.Printf("  - %s\n", city)
// // 		}
// // 		fmt.Println()
// // 	}

// // 	// After city extraction, process properties
// // 	processPropertiesForCities()
// // 	// Print properties for each city
// // 	fmt.Println("\nProperties by City:")
// // 	for city, properties := range cityProperties {
// // 		fmt.Printf("%s (%d properties):\n", city, len(properties))
// // 		for _, propName := range properties {
// // 			fmt.Printf("  - %s\n", propName)
// // 		}
// // 		fmt.Println()
// // 	}
// // 	// // Print properties for each city
// // 	// fmt.Println("\nProperties by City:")
// // 	// for city, properties := range cityProperties {
// // 	// 	fmt.Printf("%s (%d properties):\n", city, len(properties))
// // 	// 	for _, propName := range properties {
// // 	// 		fmt.Printf("  - %s\n", propName)
// // 	// 	}
// // 	// 	fmt.Println()
// // 	// }

// // 	// Prepare final summary data
// // 	var summaryData = struct {
// // 		Countries      map[string]bool            `json:"countries"`
// // 		Cities         map[string]bool            `json:"cities"`
// // 		CountryCities  map[string][]string        `json:"country_cities"`
// // 		CityProperties map[string][]string        `json:"city_properties"`
// // 	}{
// // 		Countries:      uniqueCountries,
// // 		Cities:         uniqueCities,
// // 		CountryCities:  countryCities,
// // 		CityProperties: cityProperties,
// // 	}

// // 	// Convert to JSON
// // 	var summaryJSON []byte
// // 	summaryJSON, err := json.MarshalIndent(summaryData, "", "  ")
// // 	if err != nil {
// // 		log.Fatalf("Failed to create JSON summary: %v", err)
// // 	}

// // 	// Save JSON to file
// // 	err = saveJSONSummary(summaryJSON)
// // 	if err != nil {
// // 		log.Printf("Failed to save JSON summary: %v", err)
// // 	} else {
// // 		fmt.Println("JSON Summary saved to city_summary.json")
// // 	}

// // 	// Optional: Print a snippet of the JSON
// // 	fmt.Println("\nJSON Summary Snippet:")
// // 	fmt.Println(string(summaryJSON[:min(len(summaryJSON), 500)]))
// // }
// // func main() {
// // 	// Existing main function code for city extraction...
// // 		queries := generateQueries()

// // 	// Channel for tracking goroutine completion
// // 	results := make(chan struct{}, len(queries))

// // 	// Limit concurrent goroutines to prevent overwhelming the API
// // 	semaphore := make(chan struct{}, 10)

// // 	// Process queries
// // 	for _, query := range queries {
// // 		semaphore <- struct{}{}
// // 		go func(q string) {
// // 			defer func() { <-semaphore }()
// // 			processCities(q, results)
// // 		}(query)
// // 	}

// // 	// Wait for all queries to complete
// // 	for range queries {
// // 		<-results
// // 	}

// // 	// Print detailed results
// // 	fmt.Println("Extraction Summary:")
// // 	fmt.Println("-------------------")
// // 	fmt.Printf("Total Unique Countries: %d\n", len(uniqueCountries))
// // 	fmt.Printf("Total Unique Cities: %d\n", len(uniqueCities))
	
// // 	// Print countries and their cities
// // 	fmt.Println("\nCountries and Cities:")
// // 	for country, cities := range countryCities {
// // 		fmt.Printf("%s (%d cities):\n", country, len(cities))
// // 		for _, city := range cities {
// // 			fmt.Printf("  - %s\n", city)
// // 		}
// // 		fmt.Println()
// // 	}

// // 	// Prepare summary data
// // 	summaryData := struct {
// // 		Countries     map[string]bool            `json:"countries"`
// // 		Cities        map[string]bool            `json:"cities"`
// // 		CountryCities map[string][]string        `json:"country_cities"`
// // 	}{
// // 		Countries:     uniqueCountries,
// // 		Cities:        uniqueCities,
// // 		CountryCities: countryCities,
// // 	}

// // 	// Convert to JSON
// // 	summaryJSON, err := json.MarshalIndent(summaryData, "", "  ")
// // 	if err != nil {
// // 		log.Fatalf("Failed to create JSON summary: %v", err)
// // 	}

// // 	// Save JSON to file
// // 	err = saveJSONSummary(summaryJSON)
// // 	if err != nil {
// // 		log.Printf("Failed to save JSON summary: %v", err)
// // 	} else {
// // 		fmt.Println("JSON Summary saved to city_summary.json")
// // 	}

// // 	// Optional: Print a snippet of the JSON
// // 	fmt.Println("\nJSON Summary Snippet:")
// // 	fmt.Println(string(summaryJSON[:min(len(summaryJSON), 500)]))











// // 	// After city extraction, process properties
// // 	processPropertiesForCities()

// // 	// Print properties for each city
// // 	fmt.Println("\nProperties by City:")
// // 	for city, properties := range cityProperties {
// // 		fmt.Printf("%s (%d properties):\n", city, len(properties))
// // 		for _, propName := range properties {
// // 			fmt.Printf("  - %s\n", propName)
// // 		}
// // 		fmt.Println()
// // 	}

// // 	// Update summary data to include properties
// // 	summaryData := struct {
// // 		Countries     map[string]bool            `json:"countries"`
// // 		Cities        map[string]bool            `json:"cities"`
// // 		CountryCities map[string][]string        `json:"country_cities"`
// // 		CityProperties map[string][]string       `json:"city_properties"`
// // 	}{
// // 		Countries:     uniqueCountries,
// // 		Cities:        uniqueCities,
// // 		CountryCities: countryCities,
// // 		CityProperties: cityProperties,
// // 	}

// // 	// Convert to JSON
// // 	summaryJSON, err := json.MarshalIndent(summaryData, "", "  ")
// // 	if err != nil {
// // 		log.Fatalf("Failed to create JSON summary: %v", err)
// // 	}

// // 	// Save JSON to file
// // 	err = saveJSONSummary(summaryJSON)
// // 	if err != nil {
// // 		log.Printf("Failed to save JSON summary: %v", err)
// // 	} else {
// // 		fmt.Println("JSON Summary saved to city_summary.json")
// // 	}
// // }


































































































































// // -=----------------------------
// package main


// import (
// 	"encoding/json"
//     "fmt"
//     "io"
//     "log"
//     "net/http"
//     "os"
//     "sync"
//     "net/url"
//     "strings"
//     "math"
//     "time"
// 	"sort"
// )

// //! Existing City struct remains the same
// // City struct to match the API response
// type City struct {
// 	CC1        string  `json:"cc1"`
// 	ImageURL   string  `json:"image_url"`
// 	Longitude  float64 `json:"longitude"`
// 	CityName   string  `json:"city_name"`
// 	DestID     string  `json:"dest_id"`
// 	Timezone   string  `json:"timezone"`
// 	Hotels     int     `json:"hotels"`
// 	Label      string  `json:"label"`
// 	Country    string  `json:"country"`
// 	Region     string  `json:"region"`
// 	DestType   string  `json:"dest_type"`
// 	Name       string  `json:"name"`
// 	Latitude   float64 `json:"latitude"`
// 	Type       string  `json:"type"`
// }


// type Property struct {
//     UFI               int64   `json:"ufi"`
//     CheckoutDate      string  `json:"checkoutDate"`
//     ReviewScoreWord   string  `json:"reviewScoreWord"`
//     Longitude         float64 `json:"longitude"`
//     IsPreferred       bool    `json:"isPreferred"`
//     CountryCode       string  `json:"countryCode"`
//     Latitude          float64 `json:"latitude"`
//     WishlistName      string  `json:"wishlistName"`
//     Name              string  `json:"name"`
//     PropertyClass     float64 `json:"accuratePropertyClass"`
//     DestID            string  `json:"dest_id"`     // Add destination ID
//     CityName          string  `json:"city_name"`   // Add city name
//     Country           string  `json:"country"`     // Add country
// }

// // Global maps to track properties
// var (
// 	uniqueCountries = make(map[string]bool)
// 	uniqueCities    = make(map[string]bool)
// 	countryCities   = make(map[string][]string)
// 	cityProperties  = make(map[string][]string)
// 	mutex           sync.Mutex
// )

// //! Existing functions remain the same...

// // Function to generate queries dynamically
// func generateQueries() []string {
// 	// Start with a comprehensive list of alphabet and common prefixes
// 	queries := []string{}
	
// 	// Alphabet queries
// 	for char := 'A'; char <= 'Z'; char++ {
// 		queries = append(queries, string(char))
// 	}

// 	// Common prefixes and patterns
// 	prefixes := []string{
// 		"a", "the", "new", "old", "big", "small", 
// 		"north", "south", "east", "west", "central",
// 	}

// 	for _, prefix := range prefixes {
// 		queries = append(queries, prefix)
// 	}

// 	return queries
// }

// // Function to fetch cities from Booking.com API
// func fetchCities(query string) ([]City, error) {
// 	// Encode the query parameter
// 	apiURL := fmt.Sprintf("https://booking-com18.p.rapidapi.com/stays/auto-complete?query=%s", query)
	
// 	// Create a new HTTP request
// 	req, err := http.NewRequest("GET", apiURL, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating request: %v", err)
// 	}

// 	// Set headers
// 	req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
// 	req.Header.Add("x-rapidapi-key", "79d933f58amsh0baa13f673b03f0p16d4a2jsnb299a967d295")

// 	// Create HTTP client and send request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("error sending request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	// Read response body
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("error reading response body: %v", err)
// 	}

// 	// Check response status
// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", 
// 			resp.StatusCode, string(body))
// 	}

// 	// Parse JSON response
// 	var citiesResp struct {
// 		Data []City `json:"data"`
// 	}
// 	err = json.Unmarshal(body, &citiesResp)
// 	if err != nil {
// 		return nil, fmt.Errorf("error parsing JSON: %v", err)
// 	}

// 	return citiesResp.Data, nil
// }

// // Function to process cities and update global maps
// func processCities(query string, results chan<- struct{}) {
// 	defer func() { results <- struct{}{} }()

// 	cities, err := fetchCities(query)
// 	if err != nil {
// 		log.Printf("Error fetching cities for query '%s': %v", query, err)
// 		return
// 	}

// 	mutex.Lock()
// 	defer mutex.Unlock()

// 	for _, city := range cities {
// 		// Normalize country and city names
// 		country := strings.TrimSpace(strings.ToUpper(city.Country))
// 		cityName := strings.TrimSpace(strings.ToUpper(city.CityName))

// 		// Track unique countries
// 		if country != "" {
// 			uniqueCountries[country] = true
// 		}

// 		// Track unique cities
// 		if cityName != "" {
// 			uniqueCities[cityName] = true
// 		}

// 		// Track cities per country
// 		if country != "" && cityName != "" {
// 			if _, exists := countryCities[country]; !exists {
// 				countryCities[country] = []string{}
// 			}
			
// 			// Avoid duplicate cities
// 			cityExists := false
// 			for _, existingCity := range countryCities[country] {
// 				if existingCity == cityName {
// 					cityExists = true
// 					break
// 				}
// 			}
			
// 			if !cityExists {
// 				countryCities[country] = append(countryCities[country], cityName)
// 			}
// 		}
// 	}
// }

// // Function to save JSON summary to file
// func saveJSONSummary(summaryJSON []byte) error {
// 	return os.WriteFile("city_summary.json", summaryJSON, 0644)
// }

// // Utility function to get minimum of two integers
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }


// func fetchPropertiesWithRetry(cityName, country string, maxRetries int) ([]Property, error) {
//     for attempt := 0; attempt < maxRetries; attempt++ {
//         properties, err := fetchPropertiesForCity(cityName, country)
        
//         if err == nil {
//             return properties, nil
//         }

//         // Check for specific error conditions
//         if strings.Contains(err.Error(), "Too many requests") || 
//            strings.Contains(err.Error(), "You are not subscribed") {
            
//             // Exponential backoff
//             waitTime := time.Duration(math.Pow(2, float64(attempt))) * time.Second
//             log.Printf("Rate limit hit. Waiting %v before retry", waitTime)
//             time.Sleep(waitTime)
            
//             continue
//         }

//         // For other errors, return immediately
//         return nil, err
//     }

//     return nil, fmt.Errorf("failed to fetch properties after %d attempts", maxRetries)
// }
// // New function to fetch properties for a specific location
// // Function to fetch properties for a specific location
// func fetchPropertiesForCity(cityName, country string) ([]Property, error) {
//     // Create a map to track properties by their destination ID
//     uniqueProperties := make(map[string]Property)
    
//     // Multiple search strategies to capture more properties
//     searchQueries := []string{
//         cityName,
//         fmt.Sprintf("%s hotels", cityName),
//         fmt.Sprintf("%s accommodation", cityName),
//     }

//     for _, query := range searchQueries {
//         // URL encode the query
//         encodedQuery := url.QueryEscape(query)
        
//         // Construct the API URL dynamically
//         apiURL := fmt.Sprintf("https://booking-com18.p.rapidapi.com/stays/auto-complete?query=%s", encodedQuery)
        
//         // Create a new HTTP request
//         req, err := http.NewRequest("GET", apiURL, nil)
//         if err != nil {
//             log.Printf("Error creating request for %s: %v", query, err)
//             continue
//         }

//         // Set headers
//         req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
//         req.Header.Add("x-rapidapi-key", "79d933f58amsh0baa13f673b03f0p16d4a2jsnb299a967d295")

//         // Create HTTP client with timeout
//         client := &http.Client{
//             Timeout: 10 * time.Second,
//         }
//         resp, err := client.Do(req)
//         if err != nil {
//             log.Printf("Error sending request for %s: %v", query, err)
//             continue
//         }
//         defer resp.Body.Close()

//         // Read response body
//         body, err := io.ReadAll(resp.Body)
//         if err != nil {
//             log.Printf("Error reading response body for %s: %v", query, err)
//             continue
//         }

//         // Check for rate limiting or subscription errors
//         if resp.StatusCode == 429 || strings.Contains(string(body), "Too many requests") {
//             return nil, fmt.Errorf("rate limit exceeded for query %s", query)
//         }

//         // Parse JSON response
//         var response struct {
//             Data []Property `json:"data"`
//         }
        
//         err = json.Unmarshal(body, &response)
//         if err != nil {
//             log.Printf("Error parsing JSON for %s: %v. Raw response: %s", query, err, string(body))
//             continue
//         }

//         // Add unique properties to the map using DestID
//         for _, prop := range response.Data {
//             // Use DestID as the unique key
//             if prop.DestID != "" {
//                 uniqueProperties[prop.DestID] = prop
//             }
//         }
//     }

//     // Convert map to slice
//     properties := make([]Property, 0, len(uniqueProperties))
//     for _, prop := range uniqueProperties {
//         properties = append(properties, prop)
//     }

//     return properties, nil
// }

// func processPropertiesForCities() {
//     // Create a channel for concurrent processing with controlled concurrency
//     propertyResults := make(chan struct {
//         City       CityKey
//         Properties []Property
//         Err        error
//     }, len(uniqueCities))

//     // Use a semaphore to limit concurrent API calls
//     semaphore := make(chan struct{}, 5) // Limit to 5 concurrent requests

//     // Track processed cities
//     processedCities := 0

//     // Create a wait group to ensure all goroutines complete
//     var wg sync.WaitGroup

//     // Process cities concurrently with rate limiting
//     for country, cities := range countryCities {
//         for _, cityName := range cities {
//             processedCities++
            
//             // Increment wait group counter
//             wg.Add(1)
            
//             go func(city, country string) {
//                 // Ensure wait group is decremented
//                 defer wg.Done()
                
//                 // Acquire and release semaphore
//                 semaphore <- struct{}{} 
//                 defer func() { <-semaphore }()

//                 // Fetch properties with retry
//                 properties, err := fetchPropertiesWithRetry(city, country, 3)
                
//                 // Send results to channel
//                 propertyResults <- struct {
//                     City       CityKey
//                     Properties []Property
//                     Err        error
//                 }{
//                     City:       CityKey{Name: city, Country: country},
//                     Properties: properties,
//                     Err:        err,
//                 }
//             }(cityName, country)
//         }
//     }

//     // Close results channel when all goroutines are done
//     go func() {
//         wg.Wait()
//         close(propertyResults)
//     }()

//     // Collect and process results
//     for result := range propertyResults {
//         if result.Err != nil {
//             log.Printf("Error fetching properties for %s, %s: %v", 
//                 result.City.Name, result.City.Country, result.Err)
//             continue
//         }

//         if len(result.Properties) == 0 {
//             log.Printf("No properties found for %s, %s", 
//                 result.City.Name, result.City.Country)
//             continue
//         }

//         // Store properties
//         mutex.Lock()
//         cityProperties[result.City.Name] = []string{}
        
//         // Sort properties by review score or preference
//         sort.Slice(result.Properties, func(i, j int) bool {
//             // Prioritize preferred properties
//             if result.Properties[i].IsPreferred != result.Properties[j].IsPreferred {
//                 return result.Properties[i].IsPreferred
//             }
            
//             // If review scores differ, sort by score
//             return result.Properties[i].PropertyClass > result.Properties[j].PropertyClass
//         })

//         // Limit to top 20 properties
//         maxProperties := 20
//         if len(result.Properties) < maxProperties {
//             maxProperties = len(result.Properties)
//         }

//         for _, prop := range result.Properties[:maxProperties] {
//             cityProperties[result.City.Name] = append(
//                 cityProperties[result.City.Name], 
//                 fmt.Sprintf("%s (Score: %s, Class: %.1f, City: %s, Country: %s)", 
//                     prop.Name, 
//                     prop.ReviewScoreWord, 
//                     prop.PropertyClass,
//                     prop.CityName,
//                     prop.Country,
//                 ),
//             )
//         }
//         mutex.Unlock()
//     }
// }

// // Add this type definition before the function
// type CityKey struct {
//     Name    string
//     Country string
// }


// // Utility function to generate location ID (you might need to adjust this)
// func generateLocationID(cityName, country string) string {
// 	// This is a placeholder - you'll need to implement proper location ID generation
// 	return fmt.Sprintf("eyJjaXR5X25hbWUiOiIlc1wiLCJjb3VudHJ5XCI6XCIlc1wiLCJkZXN0X2lkXCI6XCIyMDA4ODMyNVwiLCJkZXN0X3R5cGVcIjpcImNpdHlcIn0=", 
// 		cityName, country)
// }
// func main() {
// 	// Existing main function code for city extraction...
// 	queries := generateQueries()

// 	// Channel for tracking goroutine completion
// 	results := make(chan struct{}, len(queries))

// 	// Limit concurrent goroutines to prevent overwhelming the API
// 	semaphore := make(chan struct{}, 10)

// 	// Process queries
// 	for _, query := range queries {
// 		semaphore <- struct{}{}
// 		go func(q string) {
// 			defer func() { <-semaphore }()
// 			processCities(q, results)
// 		}(query)
// 	}

// 	// Wait for all queries to complete
// 	for range queries {
// 		<-results
// 	}

// 	// Print detailed results
// 	fmt.Println("Extraction Summary:")
// 	fmt.Println("-------------------")
// 	fmt.Printf("Total Unique Countries: %d\n", len(uniqueCountries))
// 	fmt.Printf("Total Unique Cities: %d\n", len(uniqueCities))
	
// 	// Print countries and their cities
// 	fmt.Println("\nCountries and Cities:")
// 	for country, cities := range countryCities {
// 		fmt.Printf("%s (%d cities):\n", country, len(cities))
// 		for _, city := range cities {
// 			fmt.Printf("  - %s\n", city)
// 		}
// 		fmt.Println()
// 	}

// 	// After city extraction, process properties
// 	processPropertiesForCities()
// 	// Print properties for each city
// 	fmt.Println("\nProperties by City:")
// 	for city, properties := range cityProperties {
// 		fmt.Printf("%s (%d properties):\n", city, len(properties))
// 		for _, propName := range properties {
// 			fmt.Printf("  - %s\n", propName)
// 		}
// 		fmt.Println()
// 	}
	
// 	// Prepare final summary data
// 	var summaryData = struct {
// 		Countries      map[string]bool            `json:"countries"`
// 		Cities         map[string]bool            `json:"cities"`
// 		CountryCities  map[string][]string        `json:"country_cities"`
// 		CityProperties map[string][]string        `json:"city_properties"`
// 	}{
// 		Countries:      uniqueCountries,
// 		Cities:         uniqueCities,
// 		CountryCities:  countryCities,
// 		CityProperties: cityProperties,
// 	}

// 	// Convert to JSON
// 	var summaryJSON []byte
// 	summaryJSON, err := json.MarshalIndent(summaryData, "", "  ")
// 	if err != nil {
// 		log.Fatalf("Failed to create JSON summary: %v", err)
// 	}

// 	// Save JSON to file
// 	err = saveJSONSummary(summaryJSON)
// 	if err != nil {
// 		log.Printf("Failed to save JSON summary: %v", err)
// 	} else {
// 		fmt.Println("JSON Summary saved to city_summary.json")
// 	}

// 	// Optional: Print a snippet of the JSON
// 	fmt.Println("\nJSON Summary Snippet:")
// 	fmt.Println(string(summaryJSON[:min(len(summaryJSON), 500)]))
// }


//!main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	
	"property-listing/conf"
	"property-listing/controllers"
)

func main() {
	// Initialize database
	conf.InitDB()
	// config.InitDB()

	// Create a new controller instance
	controller := controllers.NewBookingController()

	// Process all cities
	fmt.Println("Processing cities...")
	err := controller.ProcessAllCities()
	if err != nil {
		log.Fatalf("Failed to process cities: %v", err)
	}

	// Process properties for all cities
	fmt.Println("Processing properties...")
	err = controller.ProcessAllProperties()
	if err != nil {
		log.Fatalf("Failed to process properties: %v", err)
	}

	// Save data to database
	fmt.Println("Saving to database...")
	err = controller.SaveToDatabase()
	if err != nil {
		log.Fatalf("Failed to save to database: %v", err)
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
}