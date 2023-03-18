package gooxml

import (
	"gooxml/color"
	"gooxml/document"
	"testing"
)

/*
* @Author: zouyx
* @Email: zouyx@knowsec.com
* @Date:   2023/3/16 16:25
* @Package:
 */

func Test_Copy(t *testing.T) {
	filePath := "/Users/zouyuxi/workspace/third/gooxml@v1.0.1/document/testdata/template.docx"
	doc, err := document.Open(filePath)
	if nil != err {
		panic(err)
	}
	row := doc.Tables()[0].AddRow()
	row.AddCell().AddParagraph().AddRun().AddText("www.baidu.com")
	row.AddCell().AddParagraph().AddRun().AddText("检测链接数")
	row.AddCell().AddParagraph().AddRun().AddText("违规链接数")
	result := row.AddCell().AddParagraph().AddRun()
	result.Properties().SetColor(color.RGB(255, 0, 0))
	result.AddText("色情")

	doc.AddParagraph().AddRun().AddText("new table:")
	t1 := doc.AddTable()
	t1.Properties().SetWidth(1)
	cell := t1.AddRow().AddCell()
	cell.AddParagraph().AddRun().AddText("www.baidu.com")
	cell.AddParagraph().AddRun().AddText("www.baidu.com")

	//doc.AddTableFromTemplate()
	//t := doc.AddTable()
	//tb := t.Properties().Borders()
	//
	//row2 := t.AddRow()
	//row2.AddCell().AddParagraph().AddRun().AddText("www.baidu.com")
	//row2.AddCell().AddParagraph().AddRun().AddText("检测链接数")
	//row2.AddCell().AddParagraph().AddRun().AddText("违规链接数")
	//result2 := row2.AddCell().AddParagraph().AddRun()
	//result2.Properties().SetColor(color.RGB(255, 0, 0))
	//result2.AddText("色情")
	//
	doc.SaveToFile("./test.docx")
}
