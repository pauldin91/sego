# SeGo 
A lightweight and intuitive image segmentation annotation tool built with [Fyne](https://fyne.io/) in Go.

SeGo lets you annotate images, export label masks, and manage datasets for machine learning workflows — all in a clean, cross-platform GUI.

## Features

- Load and save images from local folders
- Assign and manage class labels (TODO)
- Export annotations PNG masks
- Native cross-platform GUI (Windows, Linux, macOS)
- Built with Go and [Fyne](https://fyne.io/) for performance and simplicity

## Dependencies

SeGo is written in pure Go and uses the [Fyne](https://github.com/fyne-io/fyne) UI toolkit.

### Prerequisites

- Go 1.20 or later installed: https://go.dev/dl/
- C compiler installed (e.g., `gcc`, `clang` – required by Fyne)

## Installation

Clone the repository and build the application:

```
git clone https://github.com/pauldin91/sego.git
cd sego
go mod tidy
go run main.go
