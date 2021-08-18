package main

import (
	"Junit2htmlGo/src"
)

func main() {
	var jFiles []string
	jFiles = append(jFiles, "testdata/junit1.xml", "testdata/junit2.xml","testdata/junit3.xml","testdata/junit4.xml")
	data:=src.ReportData{
		ReportName: "Sample Report Name",
		StartTime:  "Start time of execution",
		EndTime:    "end time of execution",
		Duration:   "test duration",
	}

	src.CreateReport("staticReport.html",jFiles,&data)
	src.RenderReport("8080",jFiles,nil)
}



