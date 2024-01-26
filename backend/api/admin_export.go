package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"exporttavanir"

	"tavanir"

	ptime "github.com/yaa110/go-persian-calendar"

	"github.com/spf13/cast"

	"github.com/360EntSecGroup-Skylar/excelize"

	"models"

	"github.com/gin-gonic/gin"
	"export"
	"tools"
)

func adminExportRequest(c *gin.Context) {
	exFilename := fmt.Sprintf("export_%d.xlsx", time.Now().Unix())
	d := db.Model(&models.Request{}).Where("reference_code!=?", "")
	filterKeys := []string{
		"status",
		"province",
		"city",
		"sum_damage_amount",
		"casuality_date",
		"id",
		"damage_type",
		"national_code",
		"reference_code",
		"bill_identifier",
		"expert_status",
	}

	for _, fk := range filterKeys {
		k := fmt.Sprintf("%s[]", fk)

		if values, ok := c.Request.URL.Query()[k]; ok {
			q := fmt.Sprintf("%s in (?)", fk)
			d = d.Where(q, values)
		}
	}

	if q, ok := c.Request.URL.Query()["query"]; ok {
		d = d.Where(q[0])
	}
	rows, err := d.Rows()
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	defer rows.Close()

	f := excelize.NewFile()

	sheetID := "Sheet1"
	f.SetSheetViewOptions(sheetID, -1, excelize.RightToLeft(true))

	titleFormat, err := f.NewStyle(`{"font":{"bold":true,"family":"B Nazanin"}}`)
	if err != nil {
		log.Println(err.Error())
	}

	titles := []string{"ردیف",
		"نام",
		"نام خانوادگی",
		"وضعیت",
		"خسارت مورد ادعا",
		"خسارت تایید شده",
		"تاریخ حادثه",
		"شناسه قبض",
		"کد ملی",
		"نوع خسارت",
		"کد رهگیری",
		"استان",
		"شهر",
		"",
		"",
		""}

	for j, t := range titles {
		f.SetCellValue(sheetID, toCharStr(j+1)+"1", t)
		f.SetCellStyle(sheetID, toCharStr(j+1)+"1", toCharStr(j+1)+"1", titleFormat)
	}

	i := 2

	numberFormat, err := f.NewStyle(`{"number_format": 3}`)
	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		r := &models.Request{}
		db.ScanRows(rows, r)

		f.SetCellValue(sheetID, fmt.Sprintf("A%d", i), i-1)
		f.SetCellValue(sheetID, fmt.Sprintf("B%d", i), r.Firstname)
		f.SetCellValue(sheetID, fmt.Sprintf("C%d", i), r.Surname)
		f.SetCellValue(sheetID, fmt.Sprintf("D%d", i), r.GetStatusMessage())
		f.SetCellValue(sheetID, fmt.Sprintf("E%d", i), r.SumDamageAmount)
		f.SetCellStyle(sheetID, fmt.Sprintf("E%d", i), fmt.Sprintf("E%d", i), numberFormat)
		f.SetCellValue(sheetID, fmt.Sprintf("F%d", i), r.AcceptedAmount)
		f.SetCellStyle(sheetID, fmt.Sprintf("F%d", i), fmt.Sprintf("F%d", i), numberFormat)
		if r.CasualityDate.Year() > 2000 {
			f.SetCellValue(sheetID, fmt.Sprintf("G%d", i), ptime.New(r.CasualityDate).Format("yyyy/MM/dd"))
		}
		f.SetCellValue(sheetID, fmt.Sprintf("H%d", i), r.BillIdentifier)
		f.SetCellValue(sheetID, fmt.Sprintf("I%d", i), r.NationalCode)
		f.SetCellValue(sheetID, fmt.Sprintf("J%d", i), r.GetDamageTypeMessage())
		f.SetCellValue(sheetID, fmt.Sprintf("K%d", i), r.ReferenceCode)
		f.SetCellValue(sheetID, fmt.Sprintf("L%d", i), r.Province)
		f.SetCellValue(sheetID, fmt.Sprintf("M%d", i), r.City)
		i++
	}

	istr := cast.ToString(i)
	for j, t := range titles {
		f.SetCellValue(sheetID, toCharStr(j+1)+istr, t)
		f.SetCellStyle(sheetID, toCharStr(j+1)+istr, toCharStr(j+1)+istr, titleFormat)
	}

	i++

	f.SetCellFormula(sheetID, fmt.Sprintf("E%d", i), fmt.Sprintf("=SUM(E2:E%d)", i-2))
	f.SetCellStyle(sheetID, fmt.Sprintf("E%d", i), fmt.Sprintf("E%d", i), numberFormat)
	f.SetCellFormula(sheetID, fmt.Sprintf("F%d", i), fmt.Sprintf("=SUM(F2:F%d)", i-2))
	f.SetCellStyle(sheetID, fmt.Sprintf("F%d", i), fmt.Sprintf("F%d", i), numberFormat)

	// Save xlsx file by the given path.
	if err := f.SaveAs(exFilename); err != nil {
		fmt.Println(err)
	}

	writeExcel(c, exFilename)

	os.Remove(exFilename)

}

func adminExportTavanirRequest(c *gin.Context) {
	exFilename := fmt.Sprintf("export_tavanir_%d.xlsx", time.Now().Unix())
	d := db.Model(&tavanir.Case{})
	filterKeys := []string{
		"status",
		"province",
		"city",
		"sum_damage_amount",
		"casuality_date",
		"id",
		"damage_type",
		"national_code",
		"reference_code",
		"bill_identifier",
		"expert_status",
	}

	for _, fk := range filterKeys {
		k := fmt.Sprintf("%s[]", fk)

		if values, ok := c.Request.URL.Query()[k]; ok {
			q := fmt.Sprintf("%s in (?)", fk)
			d = d.Where(q, values)
		}
	}

	if q, ok := c.Request.URL.Query()["query"]; ok {
		d = d.Where(q[0])
	}
	rows, err := d.Rows()
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	defer rows.Close()

	f := excelize.NewFile()

	sheetID := "Sheet1"
	f.SetSheetViewOptions(sheetID, -1, excelize.RightToLeft(true))

	titleFormat, err := f.NewStyle(`{"font":{"bold":true,"family":"B Nazanin"}}`)
	if err != nil {
		log.Println(err.Error())
	}

	titles := []string{"ردیف",
		"نام",
		"نام خانوادگی",
		"وضعیت",
		"خسارت مورد ادعا",
		"خسارت تایید شده",
		"تاریخ حادثه",
		"شناسه قبض",
		"کد ملی",
		"نوع خسارت",
		"کد رهگیری",
		"استان",
		"شهر",
		//"",
		//"",
		//"",
	}

	for j, t := range titles {
		f.SetCellValue(sheetID, toCharStr(j+1)+"1", t)
		f.SetCellStyle(sheetID, toCharStr(j+1)+"1", toCharStr(j+1)+"1", titleFormat)
	}

	i := 2

	numberFormat, err := f.NewStyle(`{"number_format": 3}`)
	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		r := &tavanir.Case{}
		db.ScanRows(rows, r)

		f.SetCellValue(sheetID, fmt.Sprintf("A%d", i), i-1)
		f.SetCellValue(sheetID, fmt.Sprintf("B%d", i), r.UserName)
		f.SetCellValue(sheetID, fmt.Sprintf("C%d", i), "")
		f.SetCellValue(sheetID, fmt.Sprintf("D%d", i), tavanir.TavanirStatus[r.Status])
		f.SetCellValue(sheetID, fmt.Sprintf("E%d", i), cast.ToInt64(r.Amount))
		f.SetCellStyle(sheetID, fmt.Sprintf("E%d", i), fmt.Sprintf("E%d", i), numberFormat)
		f.SetCellValue(sheetID, fmt.Sprintf("F%d", i), cast.ToInt64(r.AcceptedAmount))
		f.SetCellStyle(sheetID, fmt.Sprintf("F%d", i), fmt.Sprintf("F%d", i), numberFormat)
		f.SetCellValue(sheetID, fmt.Sprintf("G%d", i), r.EventDate)

		f.SetCellValue(sheetID, fmt.Sprintf("H%d", i), r.BillID)
		f.SetCellValue(sheetID, fmt.Sprintf("I%d", i), r.NationalID)
		f.SetCellValue(sheetID, fmt.Sprintf("J%d", i), tavanir.TavanirDamageTypes[r.CompensationTypeID])
		f.SetCellValue(sheetID, fmt.Sprintf("K%d", i), r.TrackingID)
		f.SetCellValue(sheetID, fmt.Sprintf("L%d", i), r.StateName)
		f.SetCellValue(sheetID, fmt.Sprintf("M%d", i), r.CityName)
		//f.SetCellValue(sheetID, fmt.Sprintf("N%d", i), r.CompanyID)
		//f.SetCellValue(sheetID, fmt.Sprintf("O%d", i), r.State)
		i++
	}

	istr := cast.ToString(i)
	for j, t := range titles {
		f.SetCellValue(sheetID, toCharStr(j+1)+istr, t)
		f.SetCellStyle(sheetID, toCharStr(j+1)+istr, toCharStr(j+1)+istr, titleFormat)
	}

	i++

	f.SetCellFormula(sheetID, fmt.Sprintf("E%d", i), fmt.Sprintf("=SUM(E2:E%d)", i-2))
	f.SetCellStyle(sheetID, fmt.Sprintf("E%d", i), fmt.Sprintf("E%d", i), numberFormat)
	f.SetCellFormula(sheetID, fmt.Sprintf("F%d", i), fmt.Sprintf("=SUM(F2:F%d)", i-2))
	f.SetCellStyle(sheetID, fmt.Sprintf("F%d", i), fmt.Sprintf("F%d", i), numberFormat)

	// Save xlsx file by the given path.
	if err := f.SaveAs(exFilename); err != nil {
		fmt.Println(err)
	}

	writeExcel(c, exFilename)

	os.Remove(exFilename)

}

func toCharStr(i int) string {
	return string('A' - 1 + i)
}

func adminExport(c *gin.Context) {
	filename := fmt.Sprintf("result_98_%d.xlsx", time.Now().Unix())
	export.Runner(filename, getUserFromSession(c).State, c.Query("from"), c.Query("to"))
	mediaPrefix := tools.GetEnv("MEDIA_ROOT", "./media")

	Openfile, err := os.Open(mediaPrefix + "/storage/" + filename)
	defer Openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		http.Error(c.Writer, "File not found.", 404)
		return
	}

	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+filename)
	c.Writer.Header().Set("Content-Type", FileContentType)
	c.Writer.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(c.Writer, Openfile) //'Copy' the file to the client
	return
}

func adminExportTavanir(c *gin.Context) {
	filename := fmt.Sprintf("result_99_%d.xlsx", time.Now().Unix())
	exporttavanir.Runner(filename, getUserFromSession(c).State, c.Query("from"), c.Query("to"))
	mediaPrefix := tools.GetEnv("MEDIA_ROOT", "./media")

	Openfile, err := os.Open(mediaPrefix + "/storage/" + filename)
	defer Openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		http.Error(c.Writer, "File not found.", 404)
		return
	}

	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+filename)
	c.Writer.Header().Set("Content-Type", FileContentType)
	c.Writer.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(c.Writer, Openfile) //'Copy' the file to the client
	return
}

func writeExcel(c *gin.Context, file string) {
	Openfile, err := os.Open(file)
	defer Openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		http.Error(c.Writer, "File not found.", 404)
		return
	}

	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+file)
	c.Writer.Header().Set("Content-Type", FileContentType)
	c.Writer.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(c.Writer, Openfile)
}
