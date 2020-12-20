# ibb- A Go wrapper for the IBB (Istanbul Metropolitan Municipality) Open Data Portal




## ⚙️ Installation

```go
go get -u github.com/ycd/ibb
```

Or install with go modules.

```go
import (
    "github.com/ycd/ibb"
)
```

## Available methods

```go
IETTTicketPrices(ctx context.Context, year int) (*ibb.TicketPrices, error)
```

IETTTicketPrices Contains the price information of the public transportation tickets used in Istanbul.

* Type   
* Name 
* Pricing 


```go
RailSystemsTimeline(ctx context.Context) (*ibb.RailSystemsTimeline, error) 
```
func(ctx context.Context) (*ibb.RailSystemsTimeline, error)

RailSystemsTimeline contains data on the number of voyages of the rail system lines in Istanbul. The data are in daily, monthly and yearly formats. 

* Year
* Period
* M1Hatti 
* M2Hatti 
* M3Hatti 
* M4Hatti 
* M5Hatti 
* M6Hatti 
* T1Hatti 
* T3Hatti 
* T4Hatti 
* F1Hatti 
* TF1Hatti
* TF2Hatti

```go
DailyMaximumJourneyByRailSystem(ctx context.Context) (*ibb.RailSystemsDailyMaximumJourneys, error)
```
DailyMaximumJourneyByRailSystem contains the maximum number of daily trips using rail system lines in Istanbul. Metro Istanbul Inc. and MARMARAY maximum line-based daily trips are included.

* HatAdi 
* Year from 2004-2021
* IsletmeTuru 


```go
IsparkParkingLots(ctx context.Context) (*ibb.ParkingLots, error)
```

IsparkParkingLots contains information about:

* Park ID 
* Park Name
* Location2 ID
* Location Code
* Location Name
* Park Type ID
* Park Type
* Park Capacity
* Hours of Operation
* Region ID
* Region
* Sub Area ID
* Sub Area
* District ID
* District
* Address
* Latitude / Longitude
* Polygon Data
* Latitude
* Longitude
* Monthly Subscription Fee
* Free Parking Time
* Timetable
* Park Continue Point



```go
MunicipalityPopulation2019(ctx context.Context) (*ibb.MunicipalityPopulation2019, error)
```

MunicipalityPopulation2019 contains the data about the population based on municipality's on Istanbul in 2019.

* Municipality
* Population


```go
WasteAmounts(ctx context.Context) (*ibb.WasteAmount, error)
```

WasteAmounts contains information on the amount of waste by district, year and type of waste (solid waste collection activities, Medical waste collection activities, mechanical sweeping activities to prevent pollution in the main arteries and squares) 

```go
IGDASSubscribers(ctx context.Context) (*ibb.IGDASSubscribers, error)
```

IGDASSubscribers contains IGDAS (gas company) subscription numbers in the 39 districts of Istanbul. 

* Ilce       
* AboneSayisi


## Example

```go
package main

import (
    "context"
    "github.com/ycd/ibb"
)

func main() {

    // Initialize the client
    ibb := ibb.NewClient()

    ctx := context.Background()

    // Get information about ticket prices of 2020
	ticketP2020, _ := ibb.IETTTicketPrices(ctx, 2020)
	for _, v := range ticketP2020.Records {
		fmt.Println(v)
    }
    
    
    // Get information about ISPARK's parking lots
    ispark, _ := ibb.IsparkParkingLots(ctx)
    for _,v := range ispark.Records {
        fmt.Println(v)
    }
    
}
```


## Licence
Package is distributed under MIT Licence.