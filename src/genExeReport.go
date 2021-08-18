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
}

func CreateReport(reportName string,jFiles []string, r *ReportData)  {
	data := processExecutionData(createCombJunitRep(jFiles))
	if r !=nil {
		data.ReportName=r.ReportName
		data.StartTime =r.StartTime
		data.EndTime=r.EndTime
		data.Duration=r.Duration
	}else {
		data.ReportName="Test Execution Report"
		data.StartTime ="DD:MM:YYYY HH:MM:SS"
		data.EndTime="DD:MM:YYYY HH:MM:SS"
		data.Duration="HH:MM:SS"
	}

	tmpl := template.Must(template.ParseFiles( "template.html"))
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile(reportName,tpl.Bytes(),0644)
	log.Println("Report Generated: ", reportName)
}

func RenderReport(port string, jFiles []string, r *ReportData){
	data := processExecutionData(createCombJunitRep(jFiles))
	if r !=nil {
		data.ReportName=r.ReportName
		data.StartTime =r.StartTime
		data.EndTime=r.EndTime
		data.Duration=r.Duration
	}else {
		data.ReportName="Test Execution Report"
		data.StartTime ="DD:MM:YYYY HH:MM:SS"
		data.EndTime="DD:MM:YYYY HH:MM:SS"
		data.Duration="HH:MM:SS"
	}

	tmpl := template.Must(template.ParseFiles( "template.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.Execute(w, data); err != nil {
			fmt.Println(err)
		}
	})

	http.ListenAndServe(":" + port, nil)
}