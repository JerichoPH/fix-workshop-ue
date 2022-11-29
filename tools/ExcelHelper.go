package tools

import (
	"fix-workshop-ue/wrongs"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/xuri/excelize/v2"
	"strconv"
)

// ExcelReader Excel读取器
type ExcelReader struct {
	data        map[string][]string
	excel       *excelize.File
	sheetName   string
	originalRow int
	finishedRow int
	titleRow    int
	titles      []string
	content     [][]string
}

// ToList 获取数据（数组类型）
func (cls *ExcelReader) ToList() map[string][]string {
	return cls.data
}

// ToMap 获取数据（map类型）
func (cls *ExcelReader) ToMap() map[string]map[string]string {
	if len(cls.GetTitle()) == 0 {
		wrongs.PanicEmpty("未设置表头")
	}

	_data := make(map[string]map[string]string)

	for rowNumber, row := range cls.ToList() {
		if len(cls.GetTitle()) != len(row) {
			wrongs.PanicForbidden(fmt.Sprintf("表头数量与实际数据列不匹配（第%s行）", rowNumber))
		}

		_row := make(map[string]string)
		for k, v := range row {
			_row[cls.GetTitle()[k]] = v
		}
		_data[rowNumber] = make(map[string]string)
		_data[rowNumber] = _row
	}

	return _data
}

// SetDataByRow 设置单行数据
func (cls *ExcelReader) SetDataByRow(rowNumber int, data []string) *ExcelReader {
	cls.data[strconv.Itoa(rowNumber+1)] = data
	return cls
}

// GetSheetName 获取工作表名称
func (cls *ExcelReader) GetSheetName() string {
	return cls.sheetName
}

// SetSheetName 设置工作表名称
func (cls *ExcelReader) SetSheetName(sheetName string) *ExcelReader {
	cls.sheetName = sheetName
	return cls
}

// GetOriginalRow 获取读取起始行
func (cls *ExcelReader) GetOriginalRow() int {
	return cls.originalRow
}

// SetOriginalRow 设置读取起始行
func (cls *ExcelReader) SetOriginalRow(originalRow int) *ExcelReader {
	cls.originalRow = originalRow - 1
	return cls
}

// GetFinishedRow 获取读取终止行
func (cls *ExcelReader) GetFinishedRow() int {
	return cls.finishedRow
}

// SetFinishedRow 设置读取终止行
func (cls *ExcelReader) SetFinishedRow(finishedRow int) *ExcelReader {
	cls.finishedRow = finishedRow - 1
	return cls
}

// GetTitleRow 获取表头行
func (cls *ExcelReader) GetTitleRow() int {
	return cls.titleRow
}

// SetTitleRow 设置表头行
func (cls *ExcelReader) SetTitleRow(titleRow int) *ExcelReader {
	cls.titleRow = titleRow - 1
	return cls
}

// GetTitle 获取表头
func (cls *ExcelReader) GetTitle() []string {
	return cls.titles
}

// SetTitle 设置表头
func (cls *ExcelReader) SetTitle(titles []string) *ExcelReader {
	if len(titles) == 0 {
		wrongs.PanicForbidden("表头不能为空")
	}
	cls.titles = titles
	return cls
}

// OpenFile 打开文件
func (cls *ExcelReader) OpenFile(filename string) *ExcelReader {
	if filename == "" {
		wrongs.PanicEmpty("文件名不能为空")
	}
	f, err := excelize.OpenFile(filename)
	if err != nil {
		wrongs.PanicForbidden(fmt.Sprintf("打开文件错误：%s", err.Error()))
	}
	cls.excel = f

	defer func() {
		if err := cls.excel.Close(); err != nil {
			wrongs.PanicForbidden("文件关闭错误")
		}
	}()

	cls.SetTitleRow(1)
	cls.SetOriginalRow(2)
	cls.data = make(map[string][]string)

	return cls
}

// ReadTitle 读取表头
func (cls *ExcelReader) ReadTitle() *ExcelReader {
	if cls.GetSheetName() == "" {
		wrongs.PanicEmpty("未设置工作表名称")
	}

	if rows, err := cls.excel.GetRows(cls.GetSheetName()); err != nil {
		wrongs.PanicForbidden("读取数据错误" + err.Error())
	} else {
		cls.SetTitle(rows[cls.GetTitleRow()])
	}

	return cls
}

// Read 读取Excel
func (cls *ExcelReader) Read() *ExcelReader {
	if cls.GetSheetName() == "" {
		wrongs.PanicEmpty("未设置工作表名称")
	}

	if rows, err := cls.excel.GetRows(cls.GetSheetName()); err != nil {
		wrongs.PanicForbidden("读取数据错误")
	} else {
		if cls.finishedRow == 0 {
			cls.content = rows[cls.GetOriginalRow():]
		} else {
			cls.content = rows[cls.GetOriginalRow():cls.GetFinishedRow()]
		}

		for rowNumber, row := range cls.content {
			cls.SetDataByRow(rowNumber, row)
		}
	}

	return cls
}

// GetByDataFrameUseDefaultType 获取DataFrame类型数据 通过Excel表头自定义数据类型
func (cls *ExcelReader) GetByDataFrameUseDefaultType() dataframe.DataFrame {
	titleWithType := make(map[string]series.Type)
	for _, title := range cls.GetTitle() {
		titleWithType[title] = series.String
	}

	return cls.GetByDataFrame(titleWithType)
}

// GetByDataFrame 获取DataFrame类型数据
func (cls *ExcelReader) GetByDataFrame(titleWithType map[string]series.Type) dataframe.DataFrame {
	if cls.GetSheetName() == "" {
		wrongs.PanicEmpty("未设置工作表名称")
	}

	var _content [][]string

	if rows, err := cls.excel.GetRows(cls.GetSheetName()); err != nil {
		wrongs.PanicForbidden("读取数据错误")
	} else {
		if cls.finishedRow == 0 {
			_content = rows[cls.GetTitleRow():]
		} else {
			_content = rows[cls.GetTitleRow():cls.GetFinishedRow()]
		}
	}

	return dataframe.LoadRecords(
		_content,
		dataframe.DetectTypes(false),
		dataframe.DefaultType(series.String),
		dataframe.WithTypes(titleWithType),
	)
}

// ExcelWriter Excel写入器
type ExcelWriter struct {
	filename  string
	excel     *excelize.File
	sheetName string
}

// GetFilename 获取文件名
func (cls *ExcelWriter) GetFilename() string {
	return cls.filename
}

// SetFilename 设置文件名
func (cls *ExcelWriter) SetFilename(filename string) *ExcelWriter {
	cls.filename = filename
	return cls
}

// Init 初始化
func (cls *ExcelWriter) Init(filename string) *ExcelWriter {
	if filename == "" {
		wrongs.PanicEmpty("文件名不能为空")
	}
	cls.filename = filename
	cls.excel = excelize.NewFile()

	return cls
}

// CreateSheet 创建工作表
func (cls *ExcelWriter) CreateSheet(sheetName string) *ExcelWriter {
	if sheetName == "" {
		wrongs.PanicEmpty("工作表名称不能为空")
	}
	sheetIndex := cls.excel.NewSheet(sheetName)
	cls.excel.SetActiveSheet(sheetIndex)
	cls.sheetName = cls.excel.GetSheetName(sheetIndex)

	return cls
}

//ActiveSheetByName 选择工作表（根据名称）
func (cls *ExcelWriter) ActiveSheetByName(sheetName string) *ExcelWriter {
	if sheetName == "" {
		wrongs.PanicEmpty("工作表名称不能为空")
	}
	sheetIndex := cls.excel.GetSheetIndex(sheetName)
	cls.excel.SetActiveSheet(sheetIndex)
	cls.sheetName = sheetName

	return cls
}

// ActiveSheetByIndex 选择工作表（根据编号）
func (cls *ExcelWriter) ActiveSheetByIndex(sheetIndex int) *ExcelWriter {
	if sheetIndex < 0 {
		wrongs.PanicEmpty("工作表索引不能小于0")
	}
	cls.excel.SetActiveSheet(sheetIndex)
	cls.sheetName = cls.excel.GetSheetName(sheetIndex)
	return cls
}

// setStyleFont 设置字体
func (cls *ExcelWriter) setStyleFont(cell *ExcelCell) {
	if style, err := cls.excel.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:   cell.GetFontBold(),
			Italic: cell.GetFontItalic(),
			Family: cell.GetFontFamily(),
			Size:   cell.GetFontSize(),
			Color:  cell.GetFontColor(),
		},
	}); err != nil {
		wrongs.PanicForbidden(fmt.Sprintf("设成字体错误：%s", cell.GetCoordinate()))
	} else {
		if err = cls.excel.SetCellStyle(cls.sheetName, cell.GetCoordinate(), cell.GetCoordinate(), style); err != nil {
			wrongs.PanicForbidden(fmt.Sprintf("设置字体错误：%s", cell.GetCoordinate()))
		}
	}
}

// SetRows 设置行数据
func (cls *ExcelWriter) SetRows(excelRows []*ExcelRow) *ExcelWriter {
	for _, row := range excelRows {
		for _, cell := range row.GetCells() {
			if err := cls.excel.SetCellValue(cls.sheetName, cell.GetCoordinate(), cell.GetContent()); err != nil {
				wrongs.PanicForbidden(fmt.Sprintf("写入数据错误：%s", cell.GetCoordinate()))
			}

			cls.setStyleFont(cell)
		}
	}

	return cls
}

// Save 保存文件
func (cls *ExcelWriter) Save() error {
	if cls.filename == "" {
		wrongs.PanicEmpty("未设置文件名")
	}
	return cls.excel.SaveAs(cls.filename)
}

// ExcelRow Excel行
type ExcelRow struct {
	cells        []*ExcelCell
	rowNumber    int
	excelColumns []string
}

// GetCells 获取单元格组
func (cls *ExcelRow) GetCells() []*ExcelCell {
	return cls.cells
}

// SetCells 设置单元格组
func (cls *ExcelRow) SetCells(cells []*ExcelCell) *ExcelRow {
	if cls.GetRowNumber() == 0 {
		wrongs.PanicForbidden("行标必须大于0")
	}

	for colNumber, cell := range cells {
		if colText, err := excelize.ColumnNumberToName(colNumber + 1); err != nil {
			wrongs.PanicForbidden(fmt.Sprintf("列索引转列文字失败：%d，%d", cls.GetRowNumber(), colNumber+1))
		} else {
			cell.SetCoordinate(fmt.Sprintf("%s%d", colText, cls.GetRowNumber()))
		}
	}
	cls.cells = cells

	return cls
}

// GetRowNumber 获取行标
func (cls *ExcelRow) GetRowNumber() int {
	return cls.rowNumber
}

// SetRowNumber 设置行标
func (cls *ExcelRow) SetRowNumber(rowNumber int) *ExcelRow {
	cls.rowNumber = rowNumber
	return cls
}

// ExcelCell Excel单元格
type ExcelCell struct {
	content    interface{}
	coordinate string
	fontColor  string
	fontBold   bool
	fontItalic bool
	fontFamily string
	fontSize   float64
}

// GetFontColor 获取字体颜色
func (cls *ExcelCell) GetFontColor() string {
	return cls.fontColor
}

// SetFontColor 设置字体颜色
func (cls *ExcelCell) SetFontColor(fontColor string, condition bool) *ExcelCell {
	if condition {
		cls.fontColor = fontColor
	}
	return cls
}

// GetFontBold 获取字体粗体
func (cls *ExcelCell) GetFontBold() bool {
	return cls.fontBold
}

// SetFontBold 设置字体粗体
func (cls *ExcelCell) SetFontBold(fontBold bool, condition bool) *ExcelCell {
	if condition {
		cls.fontBold = fontBold
	}
	return cls
}

// GetFontItalic 获取字体斜体
func (cls *ExcelCell) GetFontItalic() bool {
	return cls.fontItalic
}

// SetFontItalic 设置字体斜体
func (cls *ExcelCell) SetFontItalic(fontItalic bool, condition bool) *ExcelCell {
	if condition {
		cls.fontItalic = fontItalic
	}
	return cls
}

// GetFontFamily 获取字体
func (cls *ExcelCell) GetFontFamily() string {
	return cls.fontFamily
}

// SetFontFamily 设置字体
func (cls *ExcelCell) SetFontFamily(fontFamily string, condition bool) *ExcelCell {
	if condition {
		cls.fontFamily = fontFamily
	}
	return cls
}

// GetFontSize 获取字体字号
func (cls *ExcelCell) GetFontSize() float64 {
	return cls.fontSize
}

// SetFontSize 设置字体字号
func (cls *ExcelCell) SetFontSize(fontSize float64) *ExcelCell {
	cls.fontSize = fontSize
	return cls
}

// Init 初始化
func (cls *ExcelCell) Init(content interface{}) *ExcelCell {
	return cls
}

// GetContent 获取内容
func (cls *ExcelCell) GetContent() interface{} {
	return cls.content
}

// SetContent 设置内容
func (cls *ExcelCell) SetContent(content interface{}) *ExcelCell {
	cls.content = content
	return cls
}

// GetCoordinate 获取单元格坐标
func (cls *ExcelCell) GetCoordinate() string {
	return cls.coordinate
}

// SetCoordinate 设置单元格坐标
func (cls *ExcelCell) SetCoordinate(coordinate string) *ExcelCell {
	cls.coordinate = coordinate
	return cls
}
