// services/data_fetcher.go
package services

import (
    "log"
    "property-listing/models"
    "github.com/beego/beego/v2/client/orm"
)

type DataFetcher struct {
    bookingService *BookingService
    o              orm.Ormer
}

func NewDataFetcher(apiKey string) *DataFetcher {
    return &DataFetcher{
        bookingService: NewBookingService(apiKey),
        o:             orm.NewOrm(),
    }
}

func (df *DataFetcher) FetchAndStoreLocation(query string) error {
    locations, err := df.bookingService.GetLocations(query)
    if err != nil {
        return err
    }

    for _, loc := range locations {
        location := &models.Location{
            Name:        loc.Name,
            Description: "Located in " + loc.Name,
        }

        // Insert location
        _, err := df.o.Insert(location)
        if err != nil {
            log.Printf("Error inserting location %s: %v", loc.Name, err)
            continue
        }

        // Fetch and store properties for this location
        err = df.FetchAndStoreProperties(loc.DestID, location.ID)
        if err != nil {
            log.Printf("Error fetching properties for location %s: %v", loc.Name, err)
        }
    }

    return nil
}

func (df *DataFetcher) FetchAndStoreProperties(destID string, locationID int64) error {
    property, err := df.bookingService.GetPropertyDetails(destID)
    if err != nil {
        return err
    }

    // Create rental property
    rentalProperty := &models.RentalProperty{
        Name:      property.Name,
        Type:      property.Type,
        LocationID: locationID,
        Amenities: property.Amenities[0], // Store first amenity as example
    }

    if len(property.Rooms) > 0 {
        rentalProperty.Bedrooms = property.Rooms[0].Bedrooms
        rentalProperty.Bathrooms = property.Rooms[0].Bathrooms
    }

    // Insert property
    _, err = df.o.Insert(rentalProperty)
    if err != nil {
        return err
    }

    // Create property details
    details := &models.PropertyDetails{
        PropertyID:   rentalProperty.ID,
        Description:  property.Description,
        Images:      property.Images[0].URL, // Store first image as example
    }

    _, err = df.o.Insert(details)
    return err
}