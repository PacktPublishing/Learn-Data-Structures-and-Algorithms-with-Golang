//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt, log and math packages
import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

//Share class
type Share struct {
	Ticker   string
	StockExchange string
	CompanyName     string
}

//ShareTrend class
type ShareTrend struct {
	*Share
	Direction string
	Strength  int
	Volatility        float64
	VolatilityPercentage    float64
}

//ShareClose class
type ShareClose struct {
	*Share
	Close     float64
	AvgVolume float64
}
// GetShareTrends Method
func GetShareTrends(shareCloses []ShareClose, share Share, trendCategory string) (trendingShares []ShareTrend) {
	trendingShares = make([]ShareTrend, 0)

	var allCloses []float64
	var allVolumes []float64

	allCloses = make([]float64, 0)
	allVolumes = make([]float64, 0)
	var shareClose ShareClose
	for _, shareClose = range shareCloses {
		allCloses = append(allCloses, shareClose.Close)
		allVolumes = append(allVolumes, shareClose.AvgVolume)
	}

	if GetShareTrend(allCloses, allVolumes, "up", share.Ticker, trendCategory) {
		fmt.Printf("\t\t\tTrend is going UP for %s\n", share.Ticker)
		var volatility float64
		var volatilityPercentage float64
		volatility, volatilityPercentage = GetStandardDeviation(share.Ticker, 2, shareCloses)
		var trendingShare ShareTrend
		trendingShare = ShareTrend{&share, "up", 0, volatility, volatilityPercentage}
		trendingShares = append(trendingShares, trendingShare)
	} else if GetShareTrend(allCloses, allVolumes, "down", share.Ticker, trendCategory) {
		fmt.Printf("\t\t\tTrend is going DOWN for %s\n", share.Ticker)
		var volatility float64
		var volatilityPercentage float64
		volatility, volatilityPercentage = GetStandardDeviation(share.Ticker, 2, shareCloses)
		var trendingShare ShareTrend
		trendingShare = ShareTrend{&share, "down", 0, volatility, volatilityPercentage}
		trendingShares = append(trendingShares, trendingShare)
	}
	return
}

//GetShareTrend method
func GetShareTrend(closes []float64, volumes []float64, trendDirection string, ticker string, trendCategory string) (trending bool) {

	fmt.Printf("\t\t\t\tAnalysing %s trends in data: price: %f, %f, %f and volume: %f, %f, %f\n", ticker, closes[0], closes[1], closes[2], volumes[0], volumes[1], volumes[2])
	switch trendDirection {
	case "up":
		if trendCategory == "day" {
			if closes[0] > closes[1] && closes[1] > closes[2] && (volumes[0] > volumes[2] || volumes[0] > volumes[1]) {
				return true
			}
		} else if trendCategory == "hour" {
			if closes[0] > closes[1] && closes[1] > closes[2] && (volumes[0] > volumes[2] || volumes[0] > volumes[1]) {
				return true
			}
		}
		break
	case "down":
		if trendCategory == "day" {
			if closes[0] < closes[1] && closes[1] < closes[2] && (volumes[0] < volumes[2] || volumes[0] < volumes[1]) {
				return true
			}
		} else if trendCategory == "hour" {
			if closes[0] < closes[1] && closes[1] < closes[2] && (volumes[0] < volumes[2] || volumes[0] < volumes[1]) {
				return true
			}
		}
		break
	}

	return false
}
//GetStandardDeviation method
func GetStandardDeviation(ticker string, decimalPlaces int, shareCloses []ShareClose) (volatility float64, volatilityPercentage float64) {
	fmt.Println("Calculating standard deviation for" + ticker)
	var allCloses []float64
	allCloses = make([]float64, 0)
	var totalCloses float64
	var count float64
	count = 0
	var close ShareClose
	for _, close = range shareCloses {
		var shareClose float64
		shareClose = close.Close
		totalCloses += shareClose

		fmt.Printf("Close at count %f is %f\n", count, shareClose)
		allCloses = append(allCloses, shareClose)
		count++
	}

	fmt.Printf("Total closes %f\n", count)

	var mean float64
	mean = totalCloses / count
	fmt.Printf("Mean is %f\n", mean)

	var deviationsSquare float64
	deviationsSquare = 0.
	var cl float64
	for _, cl = range allCloses {
		var dev float64
		dev = cl - mean
		deviationsSquare += dev * dev
	}
	fmt.Printf("Deviations square is %f\n", deviationsSquare)

	var devSquareAvg float64
	devSquareAvg = deviationsSquare / count
	fmt.Printf("Deviations square average is %f\n", devSquareAvg)

	volatility = math.Sqrt(devSquareAvg)

	fmt.Printf("Volatility of share %s is %f\n", ticker, volatility)

	volatilityPercentage = (volatility / allCloses[int(count)-1]) * 100
	fmt.Printf("Volatility of sthare %s as percentage is %f\n", ticker, volatilityPercentage)

	if decimalPlaces != 0 {
		volatility = RoundDown(volatility, decimalPlaces)
		volatilityPercentage = RoundDown(volatilityPercentage, decimalPlaces)
	}

	return
}
//RoundDown method
func RoundDown(input float64, places int) (newVal float64) {
	var round float64
	var pow float64
	pow = math.Pow(10, float64(places))
	var digit float64
	digit = pow * input
	round = math.Floor(digit)
	newVal = round / pow
	return
}

// main method
func main() {

	var share Share
	share = Share{Ticker: "GM", StockExchange: "DOW", CompanyName: "General Motors"}

	var category string
	category = "day"

	var closes []ShareClose
	closes = make([]ShareClose, 10)
	var i int
	for i = range closes {
		var close ShareClose
		close = ShareClose{}
		close.Share = &share
		close.Close = 36.0 - float64(i) + rand.Float64()
		close.AvgVolume = 1000 - float64(1000*i) + rand.Float64()

		closes[i] = close

	}



	log.Print("trend analysis of ", share.Ticker)
	GetShareTrends(closes, share, category)

}
