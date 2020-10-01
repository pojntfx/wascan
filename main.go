package main

import (
	"flag"
	"log"

	"github.com/pojntfx/liwasc/pkg/databases"
	"github.com/pojntfx/liwasc/pkg/servers"
	"github.com/pojntfx/liwasc/pkg/services"
	"golang.org/x/sync/semaphore"
)

func main() {
	// Parse flags
	deviceName := flag.String("deviceName", "eth0", "Network device name")
	mac2vendorDatabasePath := flag.String("mac2vendorDatabasePath", "/etc/liwasc/oui-database.sqlite", "Path to the mac2vendor database. Download from https://mac2vendor.com/articles/download")
	networkAndNodeScanDatabasePath := flag.String("networkAndNodeScanDatabasePath", "/var/liwasc/network_and_node_scan.sqlite", "Path to the persistence database for the network and node scan service.")
	serviceNamesPortNumbersDatabasePath := flag.String("serviceNamesPortNumbersDatabasePath", "/etc/liwasc/service-names-port-numbers.csv", "Path to the CSV input file containing the registered services. Download from https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.xhtml")
	ports2PacketsDatabasePath := flag.String("ports2PacketsDatabasePath", "/etc/liwasc/ports2packets.csv", "Path to the ports2packets database. Download from https://github.com/pojntfx/ports2packets/releases")
	listenAddress := flag.String("listenAddress", "0.0.0.0:15123", "Listen address.")
	maxConcurrentPortScans := flag.Uint("maxConcurrentPortScans", 1000, "Maximum concurrent port scans. Be sure to set this value to something lower than the systems ulimit or increase the latter.")

	flag.Parse()

	// Create instances
	mac2VendorDatabase := databases.NewMAC2VendorDatabase(*mac2vendorDatabasePath)
	networkAndNodeScanDatabase := databases.NewNetworkAndNodeScanDatabase(*networkAndNodeScanDatabasePath)
	serviceNamesPortNumbersDatabase := databases.NewServiceNamesPortNumbersDatabase(*serviceNamesPortNumbersDatabasePath)
	ports2PacketsDatabase := databases.NewPorts2PacketDatabase(*ports2PacketsDatabasePath)
	networkAndNodeScanService := services.NewNetworkAndNodeScanService(
		*deviceName,
		mac2VendorDatabase,
		serviceNamesPortNumbersDatabase,
		ports2PacketsDatabase,
		networkAndNodeScanDatabase,
		semaphore.NewWeighted(int64(*maxConcurrentPortScans)),
	)
	liwascServer := servers.NewLiwascServer(*listenAddress, networkAndNodeScanService)

	// Open instances
	if err := mac2VendorDatabase.Open(); err != nil {
		log.Fatal("could not open mac2VendorDatabase", err)
	}

	if err := serviceNamesPortNumbersDatabase.Open(); err != nil {
		log.Fatal("could not open serviceNamesPortNumbersDatabase", err)
	}

	if err := ports2PacketsDatabase.Open(); err != nil {
		log.Fatal("could not open ports2PacketsDatabase", err)
	}

	if err := networkAndNodeScanDatabase.Open(); err != nil {
		log.Fatal("could not open networkAndNodeScanDatabase", err)
	}

	log.Printf("Listening on %v", *listenAddress)

	if err := liwascServer.Open(); err != nil {
		log.Fatal("could not open liwasc server", err)
	}
}
