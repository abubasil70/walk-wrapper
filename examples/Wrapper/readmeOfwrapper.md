# Walk Wrapper for Go

A simplified, declarative wrapper for the [lxn/walk](https://github.com/lxn/walk) Windows GUI library. 

This project aims to streamline the process of building Windows applications in Go by reducing boilerplate code and providing an intuitive helper struct for layout management and widget creation.

## Features

- **Simplified Syntax:** Create windows, inputs, and buttons with single-line commands.
- **Layout Management:** Easy vertical (`VBox`) and horizontal (`AddRow / HBox`) layout handling.
- **Table Support:** Integrated helper for `TableView` creation.
- **Safety:** Built-in checks for resources (like Icons) to prevent runtime crashes.
- **Dialog Helpers:** Quick access to standard `MsgBox` and `Confirmation` dialogs.

## Installation

Ensure you have Go installed, then get the original walk library:

```bash
go get [github.com/lxn/walk](https://github.com/lxn/walk)
