package src

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type ReportData struct {
	ReportName string
	StartTime string
	EndTime string
	Duration string
	BrandLogo string
}

func CreateReport(reportName string,jFiles []string, r *ReportData)  {
	data := processExecutionData(createCombJunitRep(jFiles))
	if r !=nil {
		if r.ReportName != "" {data.ReportName=r.ReportName}else {data.ReportName="Test Execution Report"}
		if r.StartTime != "" {data.StartTime=r.StartTime}else {data.StartTime ="DD:MM:YYYY HH:MM:SS"}
		if r.EndTime != "" {data.EndTime=r.EndTime}else {data.EndTime="DD:MM:YYYY HH:MM:SS"}
		if r.Duration != "" {data.Duration=r.Duration}else {data.Duration="HH:MM:SS"}
		if r.BrandLogo != "" {data.BrandLogo=r.BrandLogo}else {data.BrandLogo="https://avatars.githubusercontent.com/u/43154620?v=4"}
	}else {
		data.ReportName="Test Execution Report"
		data.StartTime ="DD:MM:YYYY HH:MM:SS"
		data.EndTime="DD:MM:YYYY HH:MM:SS"
		data.Duration="HH:MM:SS"
		data.BrandLogo="https://avatars.githubusercontent.com/u/43154620?v=4"
	}

	t:=template.New("report")
	t,_=t.Parse(ReportTemplate)
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile(reportName,tpl.Bytes(),0644)
	log.Println("Report Generated: ", reportName)
}

func RenderReport(port string, jFiles []string, r *ReportData){
	data := processExecutionData(createCombJunitRep(jFiles))
	if r !=nil {
		if r.ReportName != "" {data.ReportName=r.ReportName}else {data.ReportName="Test Execution Report"}
		if r.StartTime != "" {data.StartTime=r.StartTime}else {data.StartTime ="DD:MM:YYYY HH:MM:SS"}
		if r.EndTime != "" {data.EndTime=r.EndTime}else {data.EndTime="DD:MM:YYYY HH:MM:SS"}
		if r.Duration != "" {data.Duration=r.Duration}else {data.Duration="HH:MM:SS"}
		if r.BrandLogo != "" {data.BrandLogo=r.BrandLogo}else {data.BrandLogo="https://avatars.githubusercontent.com/u/43154620?v=4"}
	}else {
		data.ReportName="Test Execution Report"
		data.StartTime ="DD:MM:YYYY HH:MM:SS"
		data.EndTime="DD:MM:YYYY HH:MM:SS"
		data.Duration="HH:MM:SS"
		data.BrandLogo="https://avatars.githubusercontent.com/u/43154620?v=4"
	}
	
	t:=template.New("report")
	t,_=t.Parse(ReportTemplate)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := t.Execute(w, data); err != nil {
			fmt.Println(err)
		}
	})

	http.ListenAndServe(":" + port, nil)
}
