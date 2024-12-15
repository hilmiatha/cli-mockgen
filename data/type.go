package data

const (
	TYPE_NAME = "name"
	TYPE_ADDRESS = "address"
	TYPE_PHONE = "phone"
	TYPE_DATE = "date"
)

var SUPPORTED_MOCK_DATA = map[string]bool{
	TYPE_NAME: true,
	TYPE_ADDRESS: true,
	TYPE_PHONE: true,
	TYPE_DATE: true,
}