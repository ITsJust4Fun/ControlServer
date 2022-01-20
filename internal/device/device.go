package device

import (
	"ControlServer/pkg/database"
	"encoding/json"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	OS      string             `json:"os" bson:"os"`
	Volumes []string           `json:"volumes" bson:"volumes"`
	IsVM    bool               `json:"is_vm" bson:"is_vm"`
	Token   string             `json:"token" bson:"token"`
}

type ProcessorInfo struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	DeviceId          primitive.ObjectID `json:"device_id" bson:"device_id"`
	Manufacturer      string             `json:"manufacturer" bson:"manufacturer"`
	Version           string             `json:"version" bson:"version"`
	CoreCount         string             `json:"core_count" bson:"core_count"`
	CoreEnabled       string             `json:"core_enabled" bson:"core_enabled"`
	ThreadCount       string             `json:"thread_count" bson:"thread_count"`
	SocketDesignation string             `json:"socket_designation" bson:"socket_designation"`
	ProcessorFamily   string             `json:"processor_family" bson:"processor_family"`
	ProcessorFamily2  string             `json:"processor_family_2" bson:"processor_family_2"`
	ProcessorID       string             `json:"processor_id" bson:"processor_id"`
}

type BiosInfo struct {
	ID                           primitive.ObjectID `json:"id" bson:"_id"`
	DeviceId                     primitive.ObjectID `json:"device_id" bson:"device_id"`
	Vendor                       string             `json:"vendor" bson:"vendor"`
	Version                      string             `json:"version" bson:"version"`
	StartingSegment              string             `json:"starting_segment" bson:"starting_segment"`
	ReleaseDate                  string             `json:"release_date" bson:"release_date"`
	ROMSize                      string             `json:"rom_size" bson:"rom_size"`
	SystemBIOSMajorRelease       string             `json:"system_bios_major_release" bson:"system_bios_major_release"`
	SystemBIOSMinorRelease       string             `json:"system_bios_minor_release" bson:"system_bios_minor_release"`
	EmbeddedFirmwareMajorRelease string             `json:"embedded_firmware_major_release" bson:"embedded_firmware_major_release"`
	EmbeddedFirmwareMinorRelease string             `json:"embedded_firmware_minor_release" bson:"embedded_firmware_minor_release"`
}

type SysInfo struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	DeviceId     primitive.ObjectID `json:"device_id" bson:"device_id"`
	Manufacturer string             `json:"manufacturer" bson:"manufacturer"`
	ProductName  string             `json:"product_name" bson:"product_name"`
	Version      string             `json:"version" bson:"version"`
	SerialNumber string             `json:"serial_number" bson:"serial_number"`
	UUID         string             `json:"uuid" bson:"uuid"`
	SKUNumber    string             `json:"sku_number" bson:"sku_number"`
	Family       string             `json:"family" bson:"family"`
}

type BaseBoardInfo struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	DeviceId          primitive.ObjectID `json:"device_id" bson:"device_id"`
	Manufacturer      string             `json:"manufacturer" bson:"manufacturer"`
	Product           string             `json:"product" bson:"product"`
	Version           string             `json:"version" bson:"version"`
	SerialNumber      string             `json:"serial_number" bson:"serial_number"`
	AssetTag          string             `json:"asset_tag" bson:"asset_tag"`
	LocationInChassis string             `json:"location_in_chassis" bson:"location_in_chassis"`
	ChassisHandle     string             `json:"chassis_handle" bson:"chassis_handle"`
	BoardType         string             `json:"board_type" bson:"board_type"`
}

type SysEnclosureInfo struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	DeviceId        primitive.ObjectID `json:"device_id" bson:"device_id"`
	Manufacturer    string             `json:"manufacturer" bson:"manufacturer"`
	Version         string             `json:"version" bson:"version"`
	SerialNumber    string             `json:"serial_number" bson:"serial_number"`
	AssetTag        string             `json:"asset_tag" bson:"asset_tag"`
	ContainedCount  string             `json:"contained_count" bson:"contained_count"`
	ContainedLength string             `json:"contained_length" bson:"contained_length"`
	SKUNumber       string             `json:"sku_number" bson:"sku_number"`
}

type SysSlotInfo struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	DeviceId           primitive.ObjectID `json:"device_id" bson:"device_id"`
	SlotDesignation    string             `json:"slot_designation" bson:"slot_designation"`
	SlotType           string             `json:"slot_type" bson:"slot_type"`
	SlotDataBusWidth   string             `json:"slot_data_bus_width" bson:"slot_data_bus_width"`
	SlotID             string             `json:"slot_id" bson:"slot_id"`
	SegmentGroupNumber string             `json:"segment_group_number" bson:"segment_group_number"`
	BusNumber          string             `json:"bus_number" bson:"bus_number"`
}

type PhysMemInfo struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	DeviceId           primitive.ObjectID `json:"device_id" bson:"device_id"`
	Use                string             `json:"use" bson:"use"`
	NumberDevices      string             `json:"number_devices" bson:"number_devices"`
	MaximumCapacity    string             `json:"maximum_capacity" bson:"maximum_capacity"`
	ExtMaximumCapacity string             `json:"ext_maximum_capacity" bson:"ext_maximum_capacity"`
}

type MemoryInfo struct {
	ID                   primitive.ObjectID `json:"id" bson:"_id"`
	DeviceId             primitive.ObjectID `json:"device_id" bson:"device_id"`
	DeviceLocator        string             `json:"device_locator" bson:"device_locator"`
	BankLocator          string             `json:"bank_locator" bson:"bank_locator"`
	Speed                string             `json:"speed" bson:"speed"`
	Manufacturer         string             `json:"manufacturer" bson:"manufacturer"`
	SerialNumber         string             `json:"serial_number" bson:"serial_number"`
	AssetTagNumber       string             `json:"asset_tag_number" bson:"asset_tag_number"`
	PartNumber           string             `json:"part_number" bson:"part_number"`
	Size                 string             `json:"size" bson:"size"`
	ExtendedSize         string             `json:"extended_size" bson:"extended_size"`
	ConfiguredClockSpeed string             `json:"configured_clock_speed" bson:"configured_clock_speed"`
}

type OemstringsInfo struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	DeviceId primitive.ObjectID `json:"device_id" bson:"device_id"`
	Count    string             `json:"count" bson:"count"`
	Values   string             `json:"values" bson:"values"`
}

func Auth(messageBytes []byte, messageType int, conn *websocket.Conn) error {
	var authRequest AuthRequest
	err := json.Unmarshal(messageBytes, &authRequest)

	if err != nil {
		log.Println(err)
		return err
	}

	deviceInfo, err := database.InsertOne(authRequest.DeviceInfo, "device")

	if err != nil {
		log.Println(err)
		return err
	}

	message := []byte(deviceInfo.(Device).ID.String())

	if err = conn.WriteMessage(messageType, message); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
