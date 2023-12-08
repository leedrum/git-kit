package components

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func RenderTableUI(table *widgets.Table) {
	table.SetRect(0, 0, 80, 24)
	ui.Render(table)
}

func RenderListUI(list *widgets.List) {
	list.SetRect(0, 0, 80, 24)
	ui.Render(list)
}
