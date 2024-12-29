package table

import (
	"fmt"

	"github.com/chentiangang/marketdata/model"
	"github.com/chentiangang/marketdata/util"

	"github.com/charmbracelet/lipgloss"
)

func Style(rows []Row) Model {
	columns := []Column{
		{Title: "代码", Width: 10, Align: lipgloss.Left},
		{Title: "名称", Width: 18, Align: lipgloss.Center},
		{Title: "价格", Width: 10, Align: lipgloss.Right},
		{Title: "涨跌额", Width: 10, Align: lipgloss.Right},
		{Title: "流通市值", Width: 13, Align: lipgloss.Right},
		{Title: "总市值", Width: 13, Align: lipgloss.Right},
		{Title: "换手率", Width: 10, Align: lipgloss.Right},
		//{Title: "rsi(45/14)", Width: 10, Align: lipgloss.Center},
		{Title: "涨跌幅", Width: 15, Align: lipgloss.Right},
	}
	t := New(
		WithColumns(columns),
		WithRows(rows),
		WithFocused(true),
		//WithHeight(config.GetTableHeight()),
		WithHeight(6),
	)

	s := DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("253")).
		//Background(lipgloss.Color("#871F78")).
		Background(lipgloss.Color("#545454")).
		Bold(false)
	t.EnableColumnMovement = true
	t.SetStyles(s)
	return t
}

func RenderTable(list []model.Quote) []Row {
	var rows []Row
	// 根据涨跌幅决定颜色
	for _, s := range list {
		var row Row
		defaultStyle := lipgloss.NewStyle()
		//diffStyle := lipgloss.NewStyle()
		changePercentStyle := lipgloss.NewStyle()
		//涨跌额和涨跌幅大于0为红色，小于0为绿色，等于0保持默认颜色
		if s.PriceLimit > 0 {
			changePercentStyle = changePercentStyle.Foreground(lipgloss.Color("9")) // 红色
		} else if s.PriceLimit < 0 {
			changePercentStyle = changePercentStyle.Foreground(lipgloss.Color("10")) // 绿色
		}
		row = append(row, s.Symbol)
		row = append(row, s.Name)

		if s.Exchange >= 100 {
			row = append(row, defaultStyle.Render(fmt.Sprintf("%.3f", s.Price/10)))
			row = append(row, defaultStyle.Render(fmt.Sprintf("%.3f", s.DifferenceValue/10)))
			row = append(row, defaultStyle.Render(util.ConvertToLargeUnit(s.CirculatingValue)+"$"))
			row = append(row, defaultStyle.Render(util.ConvertToFormattedUnit(s.TotalValue)+"$"))

		} else {
			row = append(row, defaultStyle.Render(fmt.Sprintf("%.2f", s.Price)))
			row = append(row, defaultStyle.Render(fmt.Sprintf("%.2f", s.DifferenceValue)))
			row = append(row, defaultStyle.Render(util.ConvertToLargeUnit(s.CirculatingValue)))
			row = append(row, defaultStyle.Render(util.ConvertToFormattedUnit(s.TotalValue)))
		}
		row = append(row, defaultStyle.Render(fmt.Sprintf("%.2f%%", s.TurnoverRate)))
		row = append(row, changePercentStyle.Render(fmt.Sprintf("%.2f%%", s.PriceLimit)))
		rows = append(rows, row)
	}
	return rows
}
