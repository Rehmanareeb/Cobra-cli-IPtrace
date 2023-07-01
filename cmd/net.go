package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/ipinfo/go/v2/ipinfo"

	"github.com/spf13/cobra"
)

const token = "b2c1a0ba3dcc60"

var Netcmd = &cobra.Command{
	Use:   "Ping the provided ip address",
	Short: "not Ping the ip",
	Long:  `This net Command pings the provided IP(if it exists) and records the response and provides the response back`,

	Run: func(cmd *cobra.Command, args []string) {
		// params: httpClient, cache, token. `http.DefaultClient` and no cache will be used in case of `nil`.
		type IPInfo struct {
			IP      string `json:"ip"`
			Country string `json:"country"`
			Region  string `json:"region"`
			City    string `json:"city"`
		}

		client := ipinfo.NewClient(nil, nil, token)

		const ip_address = "175.107.212.168"
		info, err := client.GetIPInfo(net.ParseIP(ip_address))
		if err != nil {
			log.Fatal(err)
		} else {
			jsonData, err := json.Marshal(info)
			if err != nil {
				log.Fatal(err)
			} else {
				var customIPInfo IPInfo
				err = json.Unmarshal(jsonData, &customIPInfo)
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Printf("IP:%s\nCountry:%s\nRegion:%s\nCity:%s\n", customIPInfo.IP, customIPInfo.Country, customIPInfo.Region, customIPInfo.City)
				}

			}
		}
	},
}

func init() {
	rootCmd.AddCommand(Netcmd)
}
