package appindicator

//#cgo CFLAGS: -I/usr/include/libappindicator3-0.1
//#cgo LDFLAGS: -lappindicator3
//#cgo pkg-config: gtk+-3.0
//#include <stdlib.h>
//#include <gtk/gtk.h>
//#include <libappindicator/app-indicator.h>
import "C"
import "unsafe"

// AppIndicatorCatogory
const (
	CategoryApplicationStatus = C.APP_INDICATOR_CATEGORY_APPLICATION_STATUS
	CategoryCommunications    = C.APP_INDICATOR_CATEGORY_COMMUNICATIONS
	CategorySystemServices    = C.APP_INDICATOR_CATEGORY_SYSTEM_SERVICES
	CategoryHardware          = C.APP_INDICATOR_CATEGORY_HARDWARE
	CategoryOther             = C.APP_INDICATOR_CATEGORY_OTHER
)

// AppIndicatorStatus
const (
	StatusPassive   = C.APP_INDICATOR_STATUS_PASSIVE
	StatusActive    = C.APP_INDICATOR_STATUS_ACTIVE
	StatusAttention = C.APP_INDICATOR_STATUS_ATTENTION
)

type AppIndicator struct {
	indicatorPtr *C.AppIndicator
}

func NewAppIndicator(title, icon string, category int) *AppIndicator {
	titleString := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(titleString))
	iconString := (*C.gchar)(unsafe.Pointer(C.CString(icon)))
	defer C.free(unsafe.Pointer(iconString))

	indicator := C.app_indicator_new(titleString, iconString, C.AppIndicatorCategory(category))

	return &AppIndicator{
		indicatorPtr: indicator,
	}
}

func (indicator *AppIndicator) SetStatus(status int) {
	C.app_indicator_set_status(indicator.indicatorPtr, C.AppIndicatorStatus(status))
}

func (indicator *AppIndicator) SetMenu(menu unsafe.Pointer) {
	C.app_indicator_set_menu(indicator.indicatorPtr, (*C.GtkMenu)(menu))
}
