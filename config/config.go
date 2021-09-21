package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"time"
)

type Configuration struct {
	Port  string        `toml:"port"`
	Dev   bool          `toml:"dev"`
	Aria2 aria2         `toml:"aria2"`
	Main  main          `toml:"main"`
	Bid   bid           `toml:"bid"`
}

type aria2 struct {
	Aria2DownloadDir             string      `toml:"aria2_download_dir"`
	Aria2Host                    string      `toml:"aria2_host"`
	Aria2Port                    int         `toml:"aria2_port"`
	Aria2Secret                  string      `toml:"aria2_secret"`
}

type main struct {
	SwanApiUrl               string        `toml:"api_url"`
	SwanApiKey               string        `toml:"api_key"`
	SwanAccessToken          string        `toml:"access_token"`
	SwanApiHeartbeatInterval time.Duration `toml:"api_heartbeat_interval"`
	MinerFid                 string        `toml:"miner_fid"`
	ExpectedSealingTime      int           `toml:"expected_sealing_time"`
	LotusImportInterval      time.Duration `toml:"import_interval"`
	LotusScanInterval        time.Duration `toml:"scan_interval"`
}

type bid struct {
	BidMode           int     `toml:"bid_mode"`
	StartEpoch        int     `toml:"start_epoch"`
	Price             string  `toml:"price"`
	VerifiedPrice     string  `toml:"verified_price"`
	MinPieceSize      string  `toml:"min_piece_size"`
	MaxPieceSize      string  `toml:"max_piece_size"`
	AutoBidTaskPerDay int     `toml:"auto_bid_task_per_day"`
}

var config *Configuration

func InitConfig() {
	//if strings.Trim(configFile, " ") == "" {
	configFile := "./config/config.toml"
	//}
	if metaData, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal("error:", err)
	} else {
		if !requiredFieldsAreGiven(metaData) {
			log.Fatal("required fields not given")
		}
	}
}

func GetConfig() Configuration {
	if config == nil {
		InitConfig()
	}
	return *config
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string {
		{"port"},
		{"dev"},

		{"aria2"},
		{"main"},
		{"bid"},

		{"aria2", "aria2_download_dir"},
		{"aria2", "aria2_host"},
		{"aria2", "aria2_port"},
		{"aria2", "aria2_secret"},

		{"main", "api_url"},
		{"main", "miner_fid"},
		{"main", "expected_sealing_time"},
		{"main", "import_interval"},
		{"main", "scan_interval"},
		{"main", "api_key"},
		{"main", "access_token"},
		{"main", "api_heartbeat_interval"},

		{"bid", "bid_mode"},
		{"bid", "start_epoch"},
		{"bid", "price"},
		{"bid", "verified_price"},
		{"bid", "min_piece_size"},
		{"bid", "max_piece_size"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("required conf fields ", v)
		}
	}

	return true
}
