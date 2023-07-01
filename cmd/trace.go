package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Ip struct represents the IP data
type Ip struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
}

// TraceCmd represents the trace command
var TraceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  `Trace the IP.`,
	Run:   runTrace,
}

func runTrace(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		for _, ip := range args {
			fmt.Println("We are showing data for the following ip:", ip)
			showData(ip)
		}
	} else {
		fmt.Println("Please provide IP to trace.")
	}
}

func showData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := GetData(url)

	if responseByte == nil {
		log.Fatal(responseByte)
		return
	}

	fetchedData := Ip{}
	err := json.Unmarshal(responseByte, &fetchedData)
	if err != nil {
		err := errors.New("Unable to Unmarshal the response")
		log.Fatalln(err)
		return
	}

	fmt.Println("Displaying the Data:")
	time.Sleep(2 * time.Second)
	c := color.New(color.FgRed).Add(color.Underline).Add(color.Bold)

	c.Printf(
		"IP Address:%s\n city:%s\n region:%s\n country:%s\n location:%s\n timezone:%s\n postal:%s\n",
		fetchedData.IP,
		fetchedData.City,
		fetchedData.Region,
		fetchedData.Country,
		fetchedData.Loc,
		fetchedData.Timezone,
		fetchedData.Postal,
	)
}

func GetData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get the response")
		return nil
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable to read the response")
		return nil
	}

	return responseByte
}

func init() {
	rootCmd.AddCommand(TraceCmd)
}
