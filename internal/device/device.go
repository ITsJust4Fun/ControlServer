package device

import (
	"ControlServer/graph/model"
	"ControlServer/pkg/database"
	"ControlServer/pkg/jwt"
	"encoding/json"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"reflect"
	"strings"
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
	Oemstrings   []OemStringsInfo   `json:"oemstrings"`
}

type Device struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Os      string             `json:"os" bson:"os"`
	Volumes []string           `json:"volumes" bson:"volumes"`
	IsVM    bool               `json:"is_vm" bson:"is_vm"`
	Token   string             `json:"token" bson:"token"`
}

type ProcessorInfo struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	DeviceID          primitive.ObjectID `json:"device_id" bson:"device_id"`
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
	DeviceID                     primitive.ObjectID `json:"device_id" bson:"device_id"`
	Vendor                       string             `json:"vendor" bson:"vendor"`
	Version                      string             `json:"version" bson:"version"`
	StartingSegment              string             `json:"starting_segment" bson:"starting_segment"`
	ReleaseDate                  string             `json:"release_date" bson:"release_date"`
	ROMSize                      string             `json:"rom_size" bson:"rom_size"`
	SystemBIOSMajorRelease       string             `json:"system_bios_major_release" bson:"bios_major_release"`
	SystemBIOSMinorRelease       string             `json:"system_bios_minor_release" bson:"bios_minor_release"`
	EmbeddedFirmwareMajorRelease string             `json:"embedded_firmware_major_release" bson:"ef_major_release"`
	EmbeddedFirmwareMinorRelease string             `json:"embedded_firmware_minor_release" bson:"ef_minor_release"`
}

type SysInfo struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	DeviceID     primitive.ObjectID `json:"device_id" bson:"device_id"`
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
	DeviceID          primitive.ObjectID `json:"device_id" bson:"device_id"`
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
	DeviceID        primitive.ObjectID `json:"device_id" bson:"device_id"`
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
	DeviceID           primitive.ObjectID `json:"device_id" bson:"device_id"`
	SlotDesignation    string             `json:"slot_designation" bson:"slot_designation"`
	SlotType           string             `json:"slot_type" bson:"slot_type"`
	SlotDataBusWidth   string             `json:"slot_data_bus_width" bson:"slot_data_bus_width"`
	SlotID             string             `json:"slot_id" bson:"slot_id"`
	SegmentGroupNumber string             `json:"segment_group_number" bson:"segment_group_number"`
	BusNumber          string             `json:"bus_number" bson:"bus_number"`
}

type PhysMemInfo struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	DeviceID           primitive.ObjectID `json:"device_id" bson:"device_id"`
	Use                string             `json:"use" bson:"use"`
	NumberDevices      string             `json:"number_devices" bson:"number_devices"`
	MaximumCapacity    string             `json:"maximum_capacity" bson:"maximum_capacity"`
	ExtMaximumCapacity string             `json:"ext_maximum_capacity" bson:"ext_maximum_capacity"`
}

type MemoryInfo struct {
	ID                   primitive.ObjectID `json:"id" bson:"_id"`
	DeviceID             primitive.ObjectID `json:"device_id" bson:"device_id"`
	DeviceLocator        string             `json:"device_locator" bson:"device_locator"`
	BankLocator          string             `json:"bank_locator" bson:"bank_locator"`
	Speed                string             `json:"speed" bson:"speed"`
	Manufacturer         string             `json:"manufacturer" bson:"manufacturer"`
	SerialNumber         string             `json:"serial_number" bson:"serial_number"`
	AssetTagNumber       string             `json:"asset_tag_number" bson:"asset_tag_number"`
	PartNumber           string             `json:"part_number" bson:"part_number"`
	Size                 string             `json:"size" bson:"size"`
	ExtendedSize         string             `json:"extended_size" bson:"extended_size"`
	ConfiguredClockSpeed string             `json:"configured_clock_speed" bson:"conf_clock_speed"`
}

type OemStringsInfo struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	DeviceID primitive.ObjectID `json:"device_id" bson:"device_id"`
	Count    string             `json:"count" bson:"count"`
	Values   string             `json:"values" bson:"values"`
}

var onlineSockets = make(map[*websocket.Conn]*Device)

func Auth(messageBytes []byte, messageType int, conn *websocket.Conn) error {
	var authRequest AuthRequest
	err := json.Unmarshal(messageBytes, &authRequest)

	if err != nil {
		log.Println(err)
		return err
	}

	authRequest.DeviceInfo.GetDeviceId()

	isUpdateDevice := authRequest.DeviceInfo.ID.IsZero()

	if !isUpdateDevice && !authRequest.DeviceInfo.IsTokenValid() {
		RemoveSmbiosDocuments(authRequest.DeviceInfo.ID)
		err = database.RemoveOne(&authRequest.DeviceInfo, "device")

		if err != nil {
			return err
		}

		isUpdateDevice = true
	}

	if isUpdateDevice {
		authRequest.DeviceInfo.IsVM = authRequest.Smbios.IsVM()
		err = database.CreateNewDocument(&authRequest.DeviceInfo, "device")

		if err != nil {
			log.Println(err)
			return err
		}

		err = authRequest.DeviceInfo.SetToken()

		if err != nil {
			log.Println(err)
			return err
		}

		err = database.UpdateOne(&authRequest.DeviceInfo,
			bson.D{
				{"$set", bson.D{{"token", authRequest.DeviceInfo.Token}}},
			},
			"device")

		if err != nil {
			log.Println(err)
			return err
		}

		err = authRequest.Smbios.CreateSmbiosDocuments(authRequest.DeviceInfo.ID)

		if err != nil {
			log.Println(err)
			return err
		}
	}

	ConnectDevice(conn, &authRequest.DeviceInfo)
	message := []byte(authRequest.DeviceInfo.Token)

	if err = conn.WriteMessage(messageType, message); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func ConnectDevice(conn *websocket.Conn, device *Device) {
	onlineSockets[conn] = device
}

func DisconnectDevice(conn *websocket.Conn) {
	delete(onlineSockets, conn)
}

func (table *SmbiosTable) CreateSmbiosDocuments(deviceId primitive.ObjectID) error {
	value := reflect.ValueOf(table).Elem()
	valueType := reflect.ValueOf(table).Elem().Type()

	for i := 0; i < value.NumField(); i++ {
		array := value.Field(i)
		arrayType := valueType.Field(i)

		for j := 0; j < array.Len(); j++ {
			biosSec := array.Index(j)
			collectionName := GetCollectionNameByTag(string(arrayType.Tag))
			biosSecInterface := database.SetFieldToInterface(biosSec.Interface(), "DeviceID", deviceId)

			err := database.CreateNewDocument(biosSecInterface, collectionName)

			if err != nil {
				log.Println(err)
				return err
			}
		}
	}

	return nil
}

func RemoveSmbiosDocuments(deviceId primitive.ObjectID) {
	var value SmbiosTable
	reflectValue := reflect.ValueOf(value)
	valueType := reflect.TypeOf(value)

	for i := 0; i < reflectValue.NumField(); i++ {
		arrayType := valueType.Field(i)
		collectionName := GetCollectionNameByTag(string(arrayType.Tag))

		_, _ = database.RemoveByFilter(bson.M{"device_id": deviceId}, collectionName)
	}
}

func GetCollectionNameByTag(tag string) string {
	tag = strings.Replace(tag, `json:`, "", -1)
	tag = strings.Replace(tag, `"`, "", -1)

	return tag
}

func IsVMString(value string) bool {
	vmStrings := []string{"Virtual Machine", "VMware", "vbox", "VirtualBox"}

	for _, vmString := range vmStrings {
		if strings.Contains(value, vmString) {
			return true
		}
	}

	return false
}

func (table *SmbiosTable) IsVM() bool {
	value := reflect.ValueOf(table).Elem()

	for i := 0; i < value.NumField(); i++ {
		array := value.Field(i)

		for j := 0; j < array.Len(); j++ {
			biosSec := array.Index(j)

			for k := 0; k < biosSec.NumField(); k++ {
				biosField := biosSec.Field(k)

				if IsVMString(biosField.String()) {
					return true
				}
			}
		}
	}

	return false
}

func (device *Device) GetDeviceId() {
	for _, volume := range device.Volumes {
		deviceDecoded := &Device{}
		err := database.FindOne(deviceDecoded, bson.M{"volumes": bson.M{"$all": bson.A{volume}}}, "device")

		if err == nil {
			device.ID = deviceDecoded.ID
			device.Token = deviceDecoded.Token

			return
		}
	}

	return
}

func (device *Device) SetToken() error {
	token, err := jwt.GenerateTokenForDevice(device.ID.String())

	if err != nil {
		return err
	}

	device.Token = token

	return nil
}

func (device *Device) IsTokenValid() bool {
	idString, err := jwt.ParseTokenForDevice(device.Token)

	if err != nil {
		return false
	}

	if idString == device.ID.String() {
		return true
	}

	return false
}

func (device *Device) IsOnline() bool {
	for _, onlineSocket := range onlineSockets {
		if device.ID == onlineSocket.ID {
			return true
		}
	}

	return false
}

func GetDevices() ([]*model.Device, error) {
	var devices []Device
	err := database.GetAll(&devices, "device")

	if err != nil {
		return nil, err
	}

	var devicesModel []*model.Device

	for _, device := range devices {
		var deviceModel model.Device

		deviceModel.ID = device.ID.Hex()
		deviceModel.Os = device.Os
		deviceModel.IsVM = device.IsVM
		deviceModel.IsOnline = device.IsOnline()

		for _, volume := range device.Volumes {
			deviceModel.Volumes = append(deviceModel.Volumes, &volume)
		}

		devicesModel = append(devicesModel, &deviceModel)
	}

	return devicesModel, nil
}

func GetProcessors(id string) ([]*model.ProcessorInfo, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var processors []ProcessorInfo
	err = database.FindMany(&processors, bson.D{{"device_id", objectId}}, "processor")

	if err != nil {
		return nil, err
	}

	var processorsModel []*model.ProcessorInfo

	for _, processor := range processors {
		var processorModel model.ProcessorInfo

		processorModel.ID = processor.ID.Hex()
		processorModel.DeviceID = processor.DeviceID.Hex()
		processorModel.Manufacturer = processor.Manufacturer
		processorModel.Version = processor.Version
		processorModel.CoreCount = processor.CoreCount
		processorModel.CoreEnabled = processor.CoreEnabled
		processorModel.ThreadCount = processor.ThreadCount
		processorModel.SocketDesignation = processor.SocketDesignation
		processorModel.ProcessorFamily = processor.ProcessorFamily
		processorModel.ProcessorFamily2 = processor.ProcessorFamily2
		processorModel.ProcessorID = processor.ProcessorID

		processorsModel = append(processorsModel, &processorModel)
	}

	return processorsModel, nil
}

func GetBiosInfo(id string) ([]*model.BiosInfo, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var biosInfoArray []BiosInfo
	err = database.FindMany(&biosInfoArray, bson.D{{"device_id", objectId}}, "bios")

	if err != nil {
		return nil, err
	}

	var biosModelArray []*model.BiosInfo

	for _, biosInfo := range biosInfoArray {
		var biosModel model.BiosInfo

		biosModel.ID = biosInfo.ID.Hex()
		biosModel.DeviceID = biosInfo.DeviceID.Hex()
		biosModel.Vendor = biosInfo.Vendor
		biosModel.Version = biosInfo.Version
		biosModel.StartingSegment = biosInfo.StartingSegment
		biosModel.ReleaseDate = biosInfo.ReleaseDate
		biosModel.RomSize = biosInfo.ROMSize
		biosModel.SystemBIOSMajorRelease = biosInfo.SystemBIOSMajorRelease
		biosModel.SystemBIOSMinorRelease = biosInfo.SystemBIOSMinorRelease
		biosModel.EmbeddedFirmwareMajorRelease = biosInfo.EmbeddedFirmwareMajorRelease
		biosModel.EmbeddedFirmwareMinorRelease = biosInfo.EmbeddedFirmwareMinorRelease

		biosModelArray = append(biosModelArray, &biosModel)
	}

	return biosModelArray, nil
}

func GetSysInfo(id string) ([]*model.SysInfo, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var sysInfoArray []SysInfo
	err = database.FindMany(&sysInfoArray, bson.D{{"device_id", objectId}}, "sysinfo")

	if err != nil {
		return nil, err
	}

	var sysInfoModelArray []*model.SysInfo

	for _, sysInfo := range sysInfoArray {
		var sysInfoModel model.SysInfo

		sysInfoModel.ID = sysInfo.ID.Hex()
		sysInfoModel.DeviceID = sysInfo.DeviceID.Hex()
		sysInfoModel.Manufacturer = sysInfo.Manufacturer
		sysInfoModel.ProductName = sysInfo.ProductName
		sysInfoModel.Version = sysInfo.Version
		sysInfoModel.SerialNumber = sysInfo.SerialNumber
		sysInfoModel.UUID = sysInfo.UUID
		sysInfoModel.SkuNumber = sysInfo.SKUNumber
		sysInfoModel.Family = sysInfo.Family

		sysInfoModelArray = append(sysInfoModelArray, &sysInfoModel)
	}

	return sysInfoModelArray, nil
}

func GetBaseBoards(id string) ([]*model.BaseBoardInfo, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var baseBoardsArray []BaseBoardInfo
	err = database.FindMany(&baseBoardsArray, bson.D{{"device_id", objectId}}, "baseboard")

	if err != nil {
		return nil, err
	}

	var baseBoardModelArray []*model.BaseBoardInfo

	for _, baseBoard := range baseBoardsArray {
		var baseBoardModel model.BaseBoardInfo

		baseBoardModel.ID = baseBoard.ID.Hex()
		baseBoardModel.DeviceID = baseBoard.DeviceID.Hex()
		baseBoardModel.Manufacturer = baseBoard.Manufacturer
		baseBoardModel.Product = baseBoard.Product
		baseBoardModel.Version = baseBoard.Version
		baseBoardModel.SerialNumber = baseBoard.SerialNumber
		baseBoardModel.AssetTag = baseBoard.AssetTag
		baseBoardModel.LocationInChassis = baseBoard.LocationInChassis
		baseBoardModel.ChassisHandle = baseBoard.ChassisHandle
		baseBoardModel.BoardType = baseBoard.BoardType

		baseBoardModelArray = append(baseBoardModelArray, &baseBoardModel)
	}

	return baseBoardModelArray, nil
}

func GetSysEnclosure(id string) ([]*model.SysEnclosureInfo, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var sysEnclosureArray []SysEnclosureInfo
	err = database.FindMany(&sysEnclosureArray, bson.D{{"device_id", objectId}}, "sysenclosure")

	if err != nil {
		return nil, err
	}

	var sysEnclosureModelArray []*model.SysEnclosureInfo

	for _, sysEnclosure := range sysEnclosureArray {
		var sysEnclosureModel model.SysEnclosureInfo

		sysEnclosureModel.ID = sysEnclosure.ID.Hex()
		sysEnclosureModel.DeviceID = sysEnclosure.DeviceID.Hex()
		sysEnclosureModel.Manufacturer = sysEnclosure.Manufacturer
		sysEnclosureModel.Version = sysEnclosure.Version
		sysEnclosureModel.SerialNumber = sysEnclosure.SerialNumber
		sysEnclosureModel.AssetTag = sysEnclosure.AssetTag
		sysEnclosureModel.ContainedCount = sysEnclosure.ContainedCount
		sysEnclosureModel.ContainedLength = sysEnclosure.ContainedLength
		sysEnclosureModel.SkuNumber = sysEnclosure.SKUNumber

		sysEnclosureModelArray = append(sysEnclosureModelArray, &sysEnclosureModel)
	}

	return sysEnclosureModelArray, nil
}

func GetSysSlots(id string) ([]*model.SysSlotInfo, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var sysSlots []SysSlotInfo
	err = database.FindMany(&sysSlots, bson.D{{"device_id", objectId}}, "sysslot")

	if err != nil {
		return nil, err
	}

	var sysSlotsModel []*model.SysSlotInfo

	for _, sysSlot := range sysSlots {
		var sysSlotModel model.SysSlotInfo

		sysSlotModel.ID = sysSlot.ID.Hex()
		sysSlotModel.DeviceID = sysSlot.DeviceID.Hex()
		sysSlotModel.SlotDesignation = sysSlot.SlotDesignation
		sysSlotModel.SlotType = sysSlot.SlotType
		sysSlotModel.SlotDataBusWidth = sysSlot.SlotDataBusWidth
		sysSlotModel.SlotID = sysSlot.SlotID
		sysSlotModel.SegmentGroupNumber = sysSlot.SegmentGroupNumber
		sysSlotModel.BusNumber = sysSlot.BusNumber

		sysSlotsModel = append(sysSlotsModel, &sysSlotModel)
	}

	return sysSlotsModel, nil
}

func GetPhysMem(id string) ([]*model.PhysMemInfo, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var physMemArray []PhysMemInfo
	err = database.FindMany(&physMemArray, bson.D{{"device_id", objectId}}, "physmem")

	if err != nil {
		return nil, err
	}

	var physMemModelArray []*model.PhysMemInfo

	for _, physMem := range physMemArray {
		var physMemModel model.PhysMemInfo

		physMemModel.ID = physMem.ID.Hex()
		physMemModel.DeviceID = physMem.DeviceID.Hex()
		physMemModel.Use = physMem.Use
		physMemModel.NumberDevices = physMem.NumberDevices
		physMemModel.MaximumCapacity = physMem.MaximumCapacity
		physMemModel.ExtMaximumCapacity = physMem.ExtMaximumCapacity

		physMemModelArray = append(physMemModelArray, &physMemModel)
	}

	return physMemModelArray, nil
}

func GetMemory(id string) ([]*model.MemoryInfo, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var memoryArray []MemoryInfo
	err = database.FindMany(&memoryArray, bson.D{{"device_id", objectId}}, "memory")

	if err != nil {
		return nil, err
	}

	var memoryModelArray []*model.MemoryInfo

	for _, memory := range memoryArray {
		var memoryModel model.MemoryInfo

		memoryModel.ID = memory.ID.Hex()
		memoryModel.DeviceID = memory.DeviceID.Hex()
		memoryModel.DeviceLocator = memory.DeviceLocator
		memoryModel.BankLocator = memory.BankLocator
		memoryModel.Speed = memory.Speed
		memoryModel.Manufacturer = memory.Manufacturer
		memoryModel.SerialNumber = memory.SerialNumber
		memoryModel.AssetTagNumber = memory.AssetTagNumber
		memoryModel.PartNumber = memory.PartNumber
		memoryModel.Size = memory.Size
		memoryModel.ExtendedSize = memory.ExtendedSize
		memoryModel.ConfiguredClockSpeed = memory.ConfiguredClockSpeed

		memoryModelArray = append(memoryModelArray, &memoryModel)
	}

	return memoryModelArray, nil
}

func GetOemStrings(id string) ([]*model.OemStringsInfo, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var oemStringsArray []OemStringsInfo
	err = database.FindMany(&oemStringsArray, bson.D{{"device_id", objectId}}, "oemstrings")

	if err != nil {
		return nil, err
	}

	var oemStringsModelArray []*model.OemStringsInfo

	for _, oemStrings := range oemStringsArray {
		var oemStringsModel model.OemStringsInfo

		oemStringsModel.ID = oemStrings.ID.Hex()
		oemStringsModel.DeviceID = oemStrings.DeviceID.Hex()
		oemStringsModel.Count = oemStrings.Count
		oemStringsModel.Values = oemStrings.Values

		oemStringsModelArray = append(oemStringsModelArray, &oemStringsModel)
	}

	return oemStringsModelArray, nil
}
