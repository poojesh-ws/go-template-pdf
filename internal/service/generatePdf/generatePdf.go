package generatePdf

import (
	"bytes"
	"html/template"
	"os"
	"strconv"
	"time"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func ParseTemplate(templatePath string) (string, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, nil); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func GeneratePDF(fileName, htmlBody string) error {
	nowTime := time.Now().Unix()
	// write whole the body

	if _, err := os.Stat("cloneTemplate/"); os.IsNotExist(err) {
		errDir := os.Mkdir("cloneTemplate/", 0777)
		if errDir != nil {
			return errDir
		}
	}
	err1 := os.WriteFile("cloneTemplate/"+strconv.FormatInt(int64(nowTime), 10)+".html", []byte(htmlBody), 0644)
	if err1 != nil {
		return err1
	}

	f, err := os.Open("cloneTemplate/" + strconv.FormatInt(int64(nowTime), 10) + ".html")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		return err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		return err
	}

	err = pdfg.WriteFile(fileName)
	if err != nil {
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	defer os.RemoveAll(dir + "/cloneTemplate")

	return nil
}
