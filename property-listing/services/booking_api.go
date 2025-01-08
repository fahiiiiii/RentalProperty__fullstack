// services/booking_api.go
package services

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type BookingService struct {
    ApiKey  string
    BaseURL string
}

// Response structures for the API
type LocationResponse struct {
    DestID       string `json:"dest_id"`
    Name         string `json:"name"`
    Type         string `json:"type"`
    HotelCount   int    `json:"hotels"`
    CountryCode  string `json:"country"`
}

type PropertyResponse struct {
    HotelID      string   `json:"hotel_id"`
    Name         string   `json:"name"`
    Type         string   `json:"type"`
    Address      string   `json:"address"`
    Description  string   `json:"description"`
    Amenities    []string `json:"amenities"`
    Images       []struct {
        URL string `json:"url"`
    } `json:"images"`
    Rooms []struct {
        Type      string `json:"type"`
        Bedrooms  int    `json:"bedrooms"`
        Bathrooms int    `json:"bathrooms"`
    } `json:"rooms"`
}

func NewBookingService(apiKey string) *BookingService {
    return &BookingService{
        ApiKey:  apiKey,
        BaseURL: "https://booking-com.p.rapidapi.com",
    }
}

func (s *BookingService) GetLocations(query string) ([]LocationResponse, error) {
    url := fmt.Sprintf("%s/v1/hotels/locations?name=%s&locale=en-gb", s.BaseURL, query)
    
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    
    req.Header.Add("X-RapidAPI-Key", s.ApiKey)
    req.Header.Add("X-RapidAPI-Host", "booking-com.p.rapidapi.com")
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    
    var locations []LocationResponse
    err = json.Unmarshal(body, &locations)
    return locations, err
}

func (s *BookingService) GetPropertyDetails(hotelID string) (*PropertyResponse, error) {
    url := fmt.Sprintf("%s/v1/hotels/details?hotel_id=%s", s.BaseURL, hotelID)
    
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    
    req.Header.Add("X-RapidAPI-Key", s.ApiKey)
    req.Header.Add("X-RapidAPI-Host", "booking-com.p.rapidapi.com")
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    
    var property PropertyResponse
    err = json.Unmarshal(body, &property)
    return &property, err
}