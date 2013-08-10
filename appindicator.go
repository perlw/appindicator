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
type Category int

const (
	CategoryApplicationStatus Category = C.APP_INDICATOR_CATEGORY_APPLICATION_STATUS
	CategoryCommunications             = C.APP_INDICATOR_CATEGORY_COMMUNICATIONS
	CategorySystemServices             = C.APP_INDICATOR_CATEGORY_SYSTEM_SERVICES
	CategoryHardware                   = C.APP_INDICATOR_CATEGORY_HARDWARE
	CategoryOther                      = C.APP_INDICATOR_CATEGORY_OTHER
)

// AppIndicatorStatus
type Status int

const (
	StatusPassive   Status = C.APP_INDICATOR_STATUS_PASSIVE
	StatusActive           = C.APP_INDICATOR_STATUS_ACTIVE
	StatusAttention        = C.APP_INDICATOR_STATUS_ATTENTION
)

type AppIndicator struct {
	indicatorPtr *C.AppIndicator
}

// Creates a new AppIndicator.
func NewAppIndicator(id, iconName string, category Category) *AppIndicator {
	idString := (*C.gchar)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(idString))
	iconNameString := (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(iconNameString))

	indicator := C.app_indicator_new(idString, iconNameString, C.AppIndicatorCategory(category))

	return &AppIndicator{
		indicatorPtr: indicator,
	}
}

// Creates a new AppIndicator using a specific icon path.
func NewAppIndicatorWithPath(id, iconName, iconPath string, category int) *AppIndicator {
	idString := (*C.gchar)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(idString))
	iconNameString := (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(iconNameString))
	iconPathString := (*C.gchar)(unsafe.Pointer(C.CString(iconPath)))
	defer C.free(unsafe.Pointer(iconPathString))

	indicator := C.app_indicator_new_with_path(idString, iconNameString, C.AppIndicatorCategory(category), iconPathString)

	return &AppIndicator{
		indicatorPtr: indicator,
	}
}

// Sets the status of the indicator.
func (indicator *AppIndicator) SetStatus(status Status) {
	C.app_indicator_set_status(indicator.indicatorPtr, C.AppIndicatorStatus(status))
}

// Sets the attention icon of the indicator.
func (indicator *AppIndicator) SetAttentionIon(iconName, iconDescription string) {
	iconNameString := (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(iconNameString))
	iconDescriptionString := (*C.gchar)(unsafe.Pointer(C.CString(iconDescription)))
	defer C.free(unsafe.Pointer(iconDescriptionString))

	C.app_indicator_set_attention_icon_full(indicator.indicatorPtr, iconNameString, iconDescriptionString)
}

// Sets the menu that should be shown the indicator is clicked on in the panel. An application indicator will not be rendered unless it has a menu.
// This is the C version of the function and should only be used it not using GoGtk.
func (indicator *AppIndicator) C_SetMenu(menu unsafe.Pointer) {
	C.app_indicator_set_menu(indicator.indicatorPtr, (*C.GtkMenu)(menu))
}

func (indicator *AppIndicator) SetIcon(iconName, iconDescription string) {

}

func (indicator *AppIndicator) SetIconThemePath(iconPath string) {

}

func (indicator *AppIndicator) SetLabel(label, guide string) {

}

func (indicator *AppIndicator) SetOrderingIndex(index uint32) {

}

func (indicator *AppIndicator) C_SetSecondaryActivateTarget(menu unsafe.Pointer) {

}

func (indicator *AppIndicator) SetTitle(title string) {

}

func (indicator AppIndicator) GetId() string {
	return ""
}

func (indicator AppIndicator) GetCategory() Category {
	return -1
}

func (indicator AppIndicator) GetStatus() Status {
	return -1
}

// Gets the current icon and description that is associated with the indicator.
func (indicator AppIndicator) GetIcon() (string, string) {
	return "", ""
}

func (indicator AppIndicator) GetIconThemePath() string {
	return ""
}

// Gets the current attention icon and description that is associated with the indicator.
func (indicator AppIndicator) GetAttentionIcon() (string, string) {
	return "", ""
}

func (indicator AppIndicator) C_GetMenu() unsafe.Pointer {
	return nil
}

// Gets the current label and guide that is associated with the indicator.
func (indicator AppIndicator) GetLabel() (string, string) {
	return "", ""
}

func (indicator AppIndicator) GetOrderingIndex() uint32 {
	return 0
}

func (indicator AppIndicator) C_GetSecondaryActivateTarget() unsafe.Pointer {
	return nil
}

func (indicator AppIndicator) GetTitle() string {
	return ""
}

func (indicator *AppIndicator) BuildMenuFromDesktop(filePath, profile string) {

}
