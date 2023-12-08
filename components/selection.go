package components

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	utils "github.com/leedrum/git-kit/utils"
)

func HandleBranchSelection(branchNames []string) (selectedBranches []string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	l := widgets.NewList()
	l.Title = "List branches"
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorGreen)
	l.WrapText = false
	l.SelectedRow = 0

	selectedIndex := 0
	l.Rows = append(l.Rows, branchNames...)

	RenderListUI(l)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			utils.PrintSlice(selectedBranches)
			ui.Clear()
			return selectedBranches
		case "<Down>":
			selectedIndex = l.SelectedRow + 1
			l.ScrollDown()

		case "<Up>":
			selectedIndex = l.SelectedRow - 1
			l.ScrollUp()

		case "<Space>":
			branchName := branchNames[selectedIndex]
			if utils.IsSliceContain(selectedBranches, branchNames[selectedIndex]) {
				selectedBranches = utils.RemoveSliceElm(selectedBranches, branchName)
			} else {
				selectedBranches = append(selectedBranches, branchName)
			}
		}

		RenderListUI(l)
	}
}
