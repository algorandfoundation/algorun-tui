package controls

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
)

// View renders the model's content if it is visible, aligning it horizontally and ensuring it fits within the specified width.
func (m Model) View() string {
	if !m.IsVisible {
		return ""
	}
	render := controlStyle.Render(m.Content)
	difference := m.Width - lipgloss.Width(render) - 2
	line := strings.Repeat("─", max(0, difference/2))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, render, line)
}
