package databases

//go:generate sh -c "cd ../../ && sqlboiler sqlite3 -o pkg/sql/generated/mac2vendor -c pkg/sql/mac2vendor.toml"

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	mac2vendorModels "github.com/pojntfx/liwasc/pkg/sql/generated/mac2vendor"
)

type MAC2VendorDatabase struct {
	*SQLiteDatabase
}

func NewMAC2VendorDatabase(dbPath string) *MAC2VendorDatabase {
	return &MAC2VendorDatabase{&SQLiteDatabase{dbPath, nil}}
}

func (d *MAC2VendorDatabase) GetVendor(mac string) (*mac2vendorModels.Vendordb, error) {
	oui, err := GetOUI(mac)
	if err != nil {
		return nil, err
	}

	vendor, err := mac2vendorModels.FindVendordb(context.Background(), d.db, int64(oui))
	if err != nil {
		return nil, err
	}

	return vendor, nil
}

func GetOUI(mac string) (uint64, error) {
	parsedMAC := strings.Split(mac, ":")
	if len(parsedMAC) < 4 {
		return 0, fmt.Errorf("invalid MAC Address: %v", mac)
	}

	res, err := strconv.ParseUint(strings.Join(parsedMAC[0:3], ""), 16, 64)
	if err != nil {
		return 0, err
	}

	return uint64(res), err
}