package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Name           string
	Genre          string
	LeadStudio     string
	AudienceScore  string
	Profitability  string
	RottenTomatoes string
	WorldwideGross string
	Year           string
}

var Data []Movie

func populateData() {
	records := readCsvFile("./movies.csv")
	for i := 1; i < len(records); i++ {
		movie := Movie{
			Name:           records[i][0],
			Genre:          records[i][1],
			LeadStudio:     records[i][2],
			AudienceScore:  records[i][3],
			Profitability:  records[i][4],
			RottenTomatoes: records[i][5],
			WorldwideGross: records[i][6],
			Year:           records[i][7],
		}
		Data = append(Data, movie)
	}
}
func main() {
	populateData()
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRoutes(r *gin.Engine) {
	r.GET("/year/:year", Dummy)
	r.GET("/rating/:rating", Dummy1)
	r.GET("/genre/:genre", Dummy2)
}

//Dummy function
func Dummy(c *gin.Context) {

	//records := readCsvFile("./movies.csv")
	year, ok := c.Params.Get("year")
	genre, ok := c.GetQuery("genre")
	movieNames := getYear(year, genre)
	if ok == false {
		res := gin.H{
			"error": "name is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	res := gin.H{
		"year":   year,
		"movies": movieNames,
		"count":  len(movieNames),
	}
	c.JSON(http.StatusOK, res)
}

//Dummy function
func Dummy1(c *gin.Context) {

	records := readCsvFile("./movies.csv")
	rating, ok := c.Params.Get("rating")
	moviesName := getRating(records, rating)
	if ok == false {
		res := gin.H{
			"error": "name is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	res := gin.H{
		"Name":   moviesName,
		"Rating": rating,
	}
	c.JSON(http.StatusOK, res)
}

//Dummy function
func Dummy2(c *gin.Context) {

	records := readCsvFile("./movies.csv")
	genre, ok := c.Params.Get("genre")
	moviesName := getGenre(records, genre)
	if ok == false {
		res := gin.H{
			"error": "name is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	res := gin.H{
		"Name":  moviesName,
		"Genre": genre,
	}
	c.JSON(http.StatusOK, res)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func getYear(year, genre string) []Movie {
	movies := []Movie{}
	for i := 0; i < len(Data); i++ {
		if Data[i].Year == year {
			if genre != "" {
				if Data[i].Genre == genre {
					movies = append(movies, Data[i])
				}
			} else {
				movies = append(movies, Data[i])
			}
		}
	}
	return movies
}
func getRating(records [][]string, rating string) []string {
	movieName := []string{}
	for i := 0; i < len(records); i++ {
		if records[i][5] >= rating {
			movieName = append(movieName, records[i][0])

		}

	}
	return movieName
}
func getGenre(records [][]string, genre string) []string {
	movieName := []string{}
	for i := 0; i < len(records); i++ {
		if records[i][1] == genre {
			movieName = append(movieName, records[i][0])
		}
	}
	return movieName

}
