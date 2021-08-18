package src

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type node struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	Text     string     `xml:",chardata"`
	Children []node     `xml:",any"`
}

type testsuites struct {
	XMLName                   xml.Name     `xml:"testsuites"`
	Text                      string       `xml:",chardata"`
	NoNamespaceSchemaLocation string       `xml:"noNamespaceSchemaLocation,attr"`
	Xsi                       string       `xml:"xsi,attr"`
	Children                  []testsuites `xml:",any"`
	Testsuite                 []testsuite  `xml:"testsuite"`
}

type testsuite struct {
	Text       string `xml:",chardata"`
	Package    string `xml:"package,attr"`
	ID         string `xml:"id,attr"`
	Name       string `xml:"name,attr"`
	Timestamp  string `xml:"timestamp,attr"`
	Hostname   string `xml:"hostname,attr"`
	Tests      string `xml:"tests,attr"`
	Failures   string `xml:"failures,attr"`
	Errors     string `xml:"errors,attr"`
	Time       string `xml:"time,attr"`
	Properties struct {
		Text     string `xml:",chardata"`
		Property []struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Value string `xml:"value,attr"`
		} `xml:"property"`
	} `xml:"properties"`
	Testcase  []testcase `xml:"testcase"`
	SystemOut string     `xml:"system-out"`
	SystemErr string     `xml:"system-err"`
}

type testcase struct {
	Text      string `xml:",chardata"`
	Name      string `xml:"name,attr"`
	Classname string `xml:"classname,attr"`
	Time      string `xml:"time,attr"`
	Skipped   string `xml:"skipped"`
	Failure   struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"failure"`
}

func readXMLFile(filePath string) ([]byte, error) {

	xmlFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Successfully Opened: ", filePath)
	defer func(xmlFile *os.File) {
		err := xmlFile.Close()
		if err != nil {
			fmt.Println("not able to close the file")

		}
	}(xmlFile)

	// read xml file.
	byteValue, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		fmt.Println("error reading as byte array", err)
	}

	//populate the empty skipped tag with
	return []byte(strings.ReplaceAll(string(byteValue), "<skipped></skipped>", "<skipped>skipped</skipped>")), nil
}

func createCombJunitRep(files []string) string {
	var sb strings.Builder
	sb.WriteString("<testsuites>")
	for i := range files {
		xmlParser(files[i], &sb)
	}

	sb.WriteString("</testsuites>")
	xmlstr := xml.Header + sb.String()

	return xmlstr
}

func xmlParser(filePath string, sb *strings.Builder) {
	byteValue, _ := readXMLFile(filePath)
	var junitRep node

	xml.Unmarshal(byteValue, &junitRep)
	recursiveReplace(&junitRep)

	if junitRep.XMLName.Local == "testsuite" {
		buf, _ := xml.MarshalIndent(junitRep, "", "  ") // prefix, indent
		sb.Write(buf)
	} else if junitRep.XMLName.Local == "testsuites" {
		for i := range junitRep.Children {
			if junitRep.Children[i].XMLName.Local == "testsuite" {
				buf, _ := xml.MarshalIndent(junitRep.Children[i], "", "  ") // prefix, indent
				sb.Write(buf)
			}
		}
	}
}

var replacer = strings.NewReplacer("&#xA;", "", "&#x9;", "", "\n", "", "\t", "")

func recursiveReplace(n *node) {
	n.Text = replacer.Replace(n.Text)
	for i := range n.Children {
		recursiveReplace(&n.Children[i])
	}
}
