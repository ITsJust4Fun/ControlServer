package device

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type AuthRequest struct {
	Method     string      `json:"method"`
	DeviceInfo Device      `json:"device"`
	Smbios     SmbiosTable `json:"smbios"`
}

type SmbiosTable struct {
	Bios         []BiosInfo         `json:"bios"`
	Sysinfo      []SysInfo          `json:"sysinfo"`
	Baseboard    []BaseBoardInfo    `json:"baseboard"`
	Sysenclosure []SysEnclosureInfo `json:"sysenclosure"`
	Processor    []ProcessorInfo    `json:"processor"`
	Sysslot      []SysSlotInfo      `json:"sysslot"`
	Physmem      []PhysMemInfo      `json:"physmem"`
	Memory       []MemoryInfo       `json:"memory"`
	Oemstrings   []OemstringsInfo   `json:"oemstrings"`
}

type Device struct {
	ID      string   `json:"_id" bson:"_id"`
	OS      string   `json:"os"`
	Volumes []string `json:"volumes"`
	IsVM    bool     `json:"is_vm"`
	Token   string   `json:"token"`
}

type ProcessorInfo struct {
	ID                string `json:"_id" bson:"_id"`
	DeviceId          string `json:"device_id"`
	Manufacturer      string `json:"manufacturer"`
	Version           string `json:"version"`
	CoreCount         string `json:"core_count"`
	CoreEnabled       string `json:"core_enabled"`
	ThreadCount       string `json:"thread_count"`
	SocketDesignation string `json:"socket_designation"`
	ProcessorFamily   string `json:"processor_family"`
	ProcessorFamily2  string `json:"processor_family_2"`
	ProcessorID       string `json:"processor_id"`
}

type BiosInfo struct {
	ID                           string `json:"_id" bson:"_id"`
	DeviceId                     string `json:"device_id"`
	Vendor                       string `json:"vendor"`
	Version                      string `json:"version"`
	StartingSegment              string `json:"starting_segment"`
	ReleaseDate                  string `json:"release_date"`
	ROMSize                      string `json:"rom_size"`
	SystemBIOSMajorRelease       string `json:"system_bios_major_release"`
	SystemBIOSMinorRelease       string `json:"system_bios_minor_release"`
	EmbeddedFirmwareMajorRelease string `json:"embedded_firmware_major_release"`
	EmbeddedFirmwareMinorRelease string `json:"embedded_firmware_minor_release"`
}

type SysInfo struct {
	ID           string `json:"_id" bson:"_id"`
	DeviceId     string `json:"device_id"`
	Manufacturer string `json:"manufacturer"`
	ProductName  string `json:"product_name"`
	Version      string `json:"version"`
	SerialNumber string `json:"serial_number"`
	UUID         string `json:"uuid"`
	SKUNumber    string `json:"sku_number"`
	Family       string `json:"family"`
}

type BaseBoardInfo struct {
	ID                string `json:"_id" bson:"_id"`
	DeviceId          string `json:"device_id"`
	Manufacturer      string `json:"manufacturer"`
	Product           string `json:"product"`
	Version           string `json:"version"`
	SerialNumber      string `json:"serial_number"`
	AssetTag          string `json:"asset_tag"`
	LocationInChassis string `json:"location_in_chassis"`
	ChassisHandle     string `json:"chassis_handle"`
	BoardType         string `json:"board_type"`
}

type SysEnclosureInfo struct {
	ID              string `json:"_id" bson:"_id"`
	DeviceId        string `json:"device_id"`
	Manufacturer    string `json:"manufacturer"`
	Version         string `json:"version"`
	SerialNumber    string `json:"serial_number"`
	AssetTag        string `json:"asset_tag"`
	ContainedCount  string `json:"contained_count"`
	ContainedLength string `json:"contained_length"`
	SKUNumber       string `json:"sku_number"`
}

type SysSlotInfo struct {
	ID                 string `json:"_id" bson:"_id"`
	DeviceId           string `json:"device_id"`
	SlotDesignation    string `json:"slot_designation"`
	SlotType           string `json:"slot_type"`
	SlotDataBusWidth   string `json:"slot_data_bus_width"`
	SlotID             string `json:"slot_id"`
	SegmentGroupNumber string `json:"segment_group_number"`
	BusNumber          string `json:"bus_number"`
}

type PhysMemInfo struct {
	ID                 string `json:"_id" bson:"_id"`
	DeviceId           string `json:"device_id"`
	Use                string `json:"use"`
	NumberDevices      string `json:"number_devices"`
	MaximumCapacity    string `json:"maximum_capacity"`
	ExtMaximumCapacity string `json:"ext_maximum_capacity"`
}

type MemoryInfo struct {
	ID                   string `json:"_id" bson:"_id"`
	DeviceId             string `json:"device_id"`
	DeviceLocator        string `json:"device_locator"`
	BankLocator          string `json:"bank_locator"`
	Speed                string `json:"speed"`
	Manufacturer         string `json:"manufacturer"`
	SerialNumber         string `json:"serial_number"`
	AssetTagNumber       string `json:"asset_tag_number"`
	PartNumber           string `json:"part_number"`
	Size                 string `json:"size"`
	ExtendedSize         string `json:"extended_size"`
	ConfiguredClockSpeed string `json:"configured_clock_speed"`
}

type OemstringsInfo struct {
	ID       string `json:"_id" bson:"_id"`
	DeviceId string `json:"device_id"`
	Count    string `json:"count"`
	Values   string `json:"values"`
}

func Auth(messageBytes []byte, messageType int, conn *websocket.Conn) error {
	var authRequest AuthRequest
	err := json.Unmarshal(messageBytes, &authRequest)

	if err != nil {
		log.Println(err)
		return err
	}

	test, _ := json.Marshal(authRequest)
	log.Println(string(test))

	message := []byte("ok")

	if err = conn.WriteMessage(messageType, message); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
