package gotk3

//#cgo CFLAGS: -I/usr/include/libappindicator3-0.1
//#cgo LDFLAGS: -lappindicator3
//#cgo pkg-config: gtk+-3.0
//#include <libappindicator/app-indicator.h>
import "C"
import "unsafe"
import "github.com/gotk3/gotk3/gtk"
import "github.com/doxxan/appindicator"

type AppIndicatorGotk3 struct {
	appindicator.AppIndicator
}

func NewAppIndicator(id, iconName string, category appindicator.Category) *AppIndicatorGotk3 {
	return &AppIndicatorGotk3{(*appindicator.NewAppIndicator(id, iconName, category))}
}

func NewAppIndicatorWithPath(id, iconName, iconPath string, category int) *AppIndicatorGotk3 {
	return &AppIndicatorGotk3{(*appindicator.NewAppIndicatorWithPath(id, iconName, iconPath, category))}
}

func (indicator *AppIndicatorGotk3) SetMenu(menu *gtk.Menu) {
	C.app_indicator_set_menu((*C.AppIndicator)(indicator.IndicatorPtr), (*C.GtkMenu)(unsafe.Pointer(menu.Native())))
}

func (indicator *AppIndicatorGotk3) SetSecondaryActivateTarget(menuItem *gtk.Widget) {
	C.app_indicator_set_secondary_activate_target((*C.AppIndicator)(indicator.IndicatorPtr), (*C.GtkWidget)(unsafe.Pointer(menuItem.Native())))
}
