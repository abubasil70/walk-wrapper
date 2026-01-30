package main

import (
	"os"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// App represents the main application structure holding configuration and widgets.
type App struct {
	Title    string
	Width    int
	Height   int
	IconPath string
	Children []Widget      // List of widgets to be rendered
	Window   *walk.MainWindow
}

// NewGui initializes a new application instance.
func NewGui(title string, w, h int) *App {
	return &App{
		Title:  title,
		Width:  w,
		Height: h,
	}
}

// SetIcon sets the application icon safely.
// It checks if the file exists to prevent runtime errors.
func (a *App) SetIcon(path string) {
	if _, err := os.Stat(path); err == nil {
		a.IconPath = path
	}
}

// Run constructs the window and starts the main event loop.
func (a *App) Run() {
	mw := MainWindow{
		AssignTo: &a.Window,
		Title:    a.Title,
		MinSize:  Size{Width: a.Width, Height: a.Height},
		Layout:   VBox{}, // Default main layout is Vertical
		Children: a.Children,
	}

	// Apply icon only if a valid path was set
	if a.IconPath != "" {
		mw.Icon = a.IconPath
	}

	mw.Run()
}

// ---------------------------------------------------------
// Layout Tools
// ---------------------------------------------------------

// AddRow creates a horizontal container (HBox) for grouping widgets side-by-side.
// It accepts a closure function to build the internal widgets.
func (a *App) AddRow(buildFunc func(row *App)) {
	// Create a temporary "mini-app" to collect child widgets
	tempRow := &App{}
	buildFunc(tempRow)

	// Append the children as a Composite with Horizontal Layout
	a.Children = append(a.Children, Composite{
		Layout:   HBox{MarginsZero: true},
		Children: tempRow.Children,
	})
}

// ---------------------------------------------------------
// Widgets
// ---------------------------------------------------------

// AddLabel adds a static text label.
func (a *App) AddLabel(txt string) {
	a.Children = append(a.Children, Label{Text: txt})
}

// AddInput adds a single-line text input field.
func (a *App) AddInput(ptr **walk.LineEdit) {
	a.Children = append(a.Children, LineEdit{AssignTo: ptr})
}

// AddEdit adds a multi-line text editing area.
func (a *App) AddEdit(ptr **walk.TextEdit) {
	a.Children = append(a.Children, TextEdit{AssignTo: ptr, VScroll: true})
}

// AddButton adds a push button with a click handler.
func (a *App) AddButton(txt string, onClick func()) {
	a.Children = append(a.Children, PushButton{
		Text:      txt,
		OnClicked: onClick,
	})
}

// AddCheck adds a checkbox widget.
func (a *App) AddCheck(txt string, ptr **walk.CheckBox) {
	a.Children = append(a.Children, CheckBox{Text: txt, AssignTo: ptr})
}

// AddCombo adds a dropdown list (ComboBox).
func (a *App) AddCombo(items []string, ptr **walk.ComboBox) {
	a.Children = append(a.Children, ComboBox{
		AssignTo:     ptr,
		Model:        items,
		CurrentIndex: 0,
	})
}

// AddSpacer adds a vertical spacer to push subsequent widgets down.
func (a *App) AddSpacer() {
	a.Children = append(a.Children, VSpacer{})
}

// AddHSpacer adds a horizontal spacer (to be used inside AddRow).
func (a *App) AddHSpacer() {
	a.Children = append(a.Children, HSpacer{})
}

// ---------------------------------------------------------
// Table Support
// ---------------------------------------------------------

// AddTable adds a data grid (TableView).
// ptr: Pointer to the TableView for future reference.
// model: The data model struct (must implement walk.TableModel or use reflection).
// cols: Definition of columns.
func (a *App) AddTable(ptr **walk.TableView, model interface{}, cols []TableViewColumn) {
	a.Children = append(a.Children, TableView{
		AssignTo:         ptr,
		AlternatingRowBG: true,
		Columns:          cols,
		Model:            model,
	})
}

// ---------------------------------------------------------
// Helper Functions
// ---------------------------------------------------------

// Msg shows a simple information message box.
func Msg(txt string) {
	walk.MsgBox(nil, "Information", txt, walk.MsgBoxIconInformation)
}

// Confirm shows a confirmation dialog (OK/Cancel).
// Returns true if the user clicked OK.
func Confirm(txt string) bool {
	return walk.MsgBox(nil, "Confirmation", txt, walk.MsgBoxOKCancel|walk.MsgBoxIconQuestion) == walk.DlgCmdOK
}
