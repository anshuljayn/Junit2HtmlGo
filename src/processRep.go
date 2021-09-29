package src

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
)

type pageData struct {
	SuiteCollections       template.HTML
	SuiteName              string
	SuiteContent           template.HTML
	SummaryViewSuiteTables template.HTML
	ErrorViewSuiteTables   template.HTML
	ReportName             string
	NoOfTestSuites         int
	NoOfTestCases          int
	NoOfTestCasesPass      int
	NoOfTestCasesFail      int
	NoOfTestCasesSkip      int
	NoOfTSPass             int
	NoOfTSFail             int
	NoOfTSSkip             int
	StartTime              string
	EndTime                string
	Duration               string
	BrandLogo              string
}

type div struct {
	XMLName xml.Name `xml:"div,omitempty"`
	Class   string   `xml:"class,attr,omitempty"`
	Table   []table  `xml:"table,omitempty"`
	Text    string   `xml:",chardata"`
	Div     []*div   `xml:",omitempty"`
	Span    []span   `xml:"span,omitempty"`
}

type span struct {
	Class string `xml:"class,attr"`
	Text  string `xml:",chardata"`
}
type table struct {
	Tr    []tr   `xml:"tr"`
	Class string `xml:"class,attr,omitempty"`
}

type tr struct {
	Td    []td   `xml:"td"`
	Th    []th   `xml:"th,omitempty"`
	Class string `xml:"class,attr,omitempty"`
}
type th struct {
	Class string `xml:"class,attr,omitempty"`
	Text  string `xml:",chardata"`
}

type td struct {
	Class string `xml:"class,attr,omitempty"`
	Text  string `xml:",chardata"`
}

func processExecutionData(xmlJunitRep string) pageData {

	pageData := pageData{}
	var fr testsuites
	xml.Unmarshal([]byte(xmlJunitRep), &fr)

	tdH1 := td{Class: "fitwidth headColumn", Text: "Test Case:"}
	tdH2 := td{Class: "headColumn", Text: "Result:"}
	tdH3 := td{Class: "headColumn", Text: "Duration:"}
	tdH4 := td{Class: "fitwidth headColumn", Text: "Test Data:"}
	
	var divSCs []div
	var tableSVs, tableEVs []table
	var NoOfTestCasesPass, NoOfTestCasesFail, NoOfTestCasesSkip, NoOfTSPass, NoOfTSFail, NoOfTSSkip int

	for i := range fr.Testsuite {
		//Test Content node +++
		var TCTables []table
		var trSVs, trEVs []tr

		//SummaryView +++
		trSVs = append(trSVs, tr{Class: "suite", Td: []td{{Text: fr.Testsuite[i].Name}}})
		//SummaryView ---

		var tcPass, tcFail, tcSkip int

		//loop the test case in the test suite
		for j := range fr.Testsuite[i].Testcase {
			var status string

			if len(fr.Testsuite[i].Testcase[j].Skipped) > 0 {
				status = "skipped"
				NoOfTestCasesSkip++
				tcSkip++
			} else if len(fr.Testsuite[i].Testcase[j].Failure.Text) > 0 {
				status = "failed"
				NoOfTestCasesFail++
				tcFail++
			} else {
				status = "passed"
				NoOfTestCasesPass++
				tcPass++
			}

			tdC1 := td{Text: fr.Testsuite[i].Testcase[j].Name}
			tdC2 := td{Class: status, Text: status}
			tdC3 := td{Text: fr.Testsuite[i].Testcase[j].Time}

			tr1 := tr{Td: []td{tdH1, tdC1}}
			tr2 := tr{Td: []td{tdH2, tdC2}}
			tr3 := tr{Td: []td{tdH3, tdC3}}

			var trs []tr

			if len(fr.Testsuite[i].Testcase[j].Tdata) >0{
				tdC4 := td{Text: fr.Testsuite[i].Testcase[j].Tdata}
				tr4 := tr{Td: []td{tdH4, tdC4}}
				trs = []tr{tr1, tr2, tr3,tr4}
			}else{
				trs = []tr{tr1, tr2, tr3}
			}

			TCtable := table{Tr: trs, Class: "testcase"}
			TCTables = append(TCTables, TCtable)

			//SummaryView
			trSVs = append(trSVs, tr{Td: []td{{Class: status, Text: fr.Testsuite[i].Testcase[j].Name}}})

			//Failure View
			if status == "failed" {
				th1 := td{Class: "fitwidth", Text: "Test case:"}
				td1 := td{Text: fr.Testsuite[i].Testcase[j].Name}
				tr1 := tr{Td: []td{th1, td1}}

				th2 := td{Text: "Failure:"}
				td2 := td{Class: "error", Text: fr.Testsuite[i].Testcase[j].Failure.Text}

				tr2 := tr{Td: []td{th2, td2}}

				trEVs = append(trEVs, tr1, tr2)
			}
		}

		var testSuiteStatus string
		//DashBoad View data +++
		if tcFail == 0 {
			if tcPass == 0 {
				NoOfTSSkip++
				testSuiteStatus = "skipped"
			} else {
				NoOfTSPass++
				testSuiteStatus = "passed"
			}
		} else {
			NoOfTSFail++
			testSuiteStatus = "failed"
		}
		//DashBoad View data ---

		//Test Heading node +++++
		divH := div{Class: "heading", Text: fr.Testsuite[i].Name}
		divSH := div{
			Class: "sub-heading",
			Span: []span{{
				Class: "duration",
				Text:  fr.Testsuite[i].Time},
				{
					Class: "right " + testSuiteStatus,
					Text:  testSuiteStatus},
			},
		}
		divTH := div{}
		divTH.Div = []*div{&divH, &divSH}
		if i == 0 {
			divTH.Class = "test-heading active"
			pageData.SuiteName = fr.Testsuite[i].Name
		} else {
			divTH.Class = "test-heading"
		}
		//Test Heading node -----

		divTC := div{
			Class: "test-content hide",
			Table: TCTables,
		}

		if i == 0 {
			divTCr := divTC
			divTCr.Class = "test-content"
			pageData.SuiteContent = getHTML(divTCr)
		}

		//Test Content node -----
		//Suite Collection node +++++
		divSC := div{
			Class: "suite_collection",
			Div: []*div{
				&divTH, &divTC,
			},
		}

		divSCs = append(divSCs, divSC)
		//Suite Collection node -----
		//detail view ---

		//SummaryView +++
		tableSVs = append(tableSVs, table{Class: "suite-table", Tr: trSVs})
		//SummaryView ---

		//Failure view +++
		if len(trEVs) > 0 {
			headRow := tr{Class: "suite", Td: []td{
				{Text: "Suite:"},
				{Text: fr.Testsuite[i].Name},
			}}

			var tableRowEV []tr
			tableRowEV = append(tableRowEV, headRow)
			tableRowEV = append(tableRowEV, trEVs...)

			tableEVs = append(tableEVs, table{Class: "suite-table", Tr: tableRowEV})
			//Failure view ---ß
		}
	}

	//set the template page data
	pageData.SuiteCollections = getHTML(divSCs)
	pageData.SummaryViewSuiteTables = getHTML(tableSVs)
	pageData.ErrorViewSuiteTables = getHTML(tableEVs)
	pageData.NoOfTestSuites = len(fr.Testsuite)
	pageData.NoOfTestCases = NoOfTestCasesPass + NoOfTestCasesFail + NoOfTestCasesSkip
	pageData.NoOfTestCasesPass = NoOfTestCasesPass
	pageData.NoOfTestCasesFail = NoOfTestCasesFail
	pageData.NoOfTestCasesSkip = NoOfTestCasesSkip
	pageData.NoOfTSPass = NoOfTSPass
	pageData.NoOfTSFail = NoOfTSFail
	pageData.NoOfTSSkip = NoOfTSSkip

	return pageData
}

func getHTML(div interface{}) template.HTML {
	b := new(bytes.Buffer)

	enc := xml.NewEncoder(b)
	enc.Indent("  ", "    ")
	err := enc.Encode(div)
	if err != nil {
		fmt.Println(err)
	}
	return template.HTML(b.String())
}
