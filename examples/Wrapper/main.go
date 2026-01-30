package main

import (
	"encoding/csv"
	"os"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// 1. Define Data Structure
type User struct {
	Name string
	City string
	Note string
}

// 2. Define Table Model
// Connects the data slice to the Walk TableView
type UserModel struct {
	walk.TableModelBase
	items []*User
}

func NewUserModel() *UserModel {
	return &UserModel{items: make([]*User, 0)}
}

// RowCount returns the number of rows in the list
func (m *UserModel) RowCount() int {
	return len(m.items)
}

// Value returns the value for a specific cell (row, col)
func (m *UserModel) Value(row, col int) interface{} {
	item := m.items[row]
	switch col {
	case 0:
		return item.Name
	case 1:
		return item.City
	case 2:
		return item.Note
	}
	return nil
}

// Global model instance
var myModel *UserModel

func main() {
	// Initialize model and load existing data
	myModel = NewUserModel()
	myModel.items = loadUsers()

	gui := NewGui("User Management System", 600, 600)
	
	// Set Icon (if file exists)
	gui.SetIcon("app.ico")

	var nameInp *walk.LineEdit
	var cityCombo *walk.ComboBox
	var descEdit *walk.TextEdit
	var tv *walk.TableView

	// --- Input Section ---
	gui.AddLabel("Full Name:")
	gui.AddInput(&nameInp)

	gui.AddLabel("City:")
	gui.AddCombo([]string{"New York", "London", "Berlin", "Tokyo", "Dubai"}, &cityCombo)

	gui.AddLabel("Notes:")
	gui.AddEdit(&descEdit)

	gui.AddSpacer()

	// --- Action Buttons ---
	gui.AddRow(func(row *App) {
		row.AddButton("Save to CSV", func() {
			u := &User{
				Name: nameInp.Text(),
				City: cityCombo.Text(),
				Note: descEdit.Text(),
			}

			if saveUserToCSV(u) {
				// Update model logic
				myModel.items = append(myModel.items, u)
				myModel.PublishRowsReset() // Refresh table
				
				Msg("Saved successfully!")
				
				// Clear inputs
				nameInp.SetText("")
				descEdit.SetText("")
			}
		})

		row.AddHSpacer() // Flexible space

		row.AddButton("Exit", func() {
			if Confirm("Are you sure you want to exit?") {
				walk.App().Exit(0)
			}
		})
	})

	gui.AddSpacer()

	// --- Table Section ---
	gui.AddLabel("Registered Users:")

	gui.AddTable(&tv, myModel, []TableViewColumn{
		{Title: "Name", Width: 150},
		{Title: "City", Width: 100},
		{Title: "Notes", Width: 200},
	})

	gui.Run()
}

// --- CSV Helpers ---

func saveUserToCSV(u *User) bool {
	f, err := os.OpenFile("users.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Msg("Error: " + err.Error())
		return false
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	return writer.Write([]string{u.Name, u.City, u.Note}) == nil
}

func loadUsers() []*User {
	var loaded []*User
	f, err := os.Open("users.csv")
	if err != nil {
		return loaded
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, _ := reader.ReadAll()

	for _, record := range records {
		if len(record) >= 3 {
			loaded = append(loaded, &User{
				Name: record[0],
				City: record[1],
				Note: record[2],
			})
		}
	}
	return loaded
}
// go build -ldflags="-H windowsgui" -o "MyProgram.exe"
