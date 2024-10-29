// SMS/app.go
package main

import (
	"context"
	"encoding/base64"
	"sync"

	// "encoding/base64"
	"fmt"
	"myproject/pdf"
	"myproject/sms"
	// "net/http"
)
var once sync.Once
// App struct
type App struct {
    ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
    return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
    a.ctx = ctx
    sms.InitDb() // Initialize the database
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
    return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CreateClass(name string, year int64) (map[string]interface{}, error) {
    id, err := sms.CreateClass(name, year)
    if err != nil {
        if err.Error() == "class already exists" {
            return map[string]interface{}{"message": "Class already exists"}, nil
        }
        return nil, err
    }
    return map[string]interface{}{"id": id}, nil
}

func (a *App) GetClass(id int64) (*sms.ClassInfo, error) {
    return sms.GetClass(id)
}

func (a *App) UpdateClass(id int64, name string, year int) error {
    return sms.UpdateClass(id, name, year)
}

func (a *App) DeleteClass(id int64) error {
    return sms.DeleteClass(id)
}

func (a *App) ListClasses() ([]sms.ClassInfo, error) {
    return sms.ListClasses()
}

// func (a *App) ListStudents(tableName string) ([]sms.Student, error) {
//     return sms.ListStudents(tableName)
// }
func (a *App) ListStudents12(tableName string) ([]sms.Student12, error) {
    return sms.ListStudents12(tableName)
}

func (a *App) ListStudents345(tableName string) ([]sms.Student345, error) {
    return sms.ListStudents345(tableName)
}

func (a *App) AddStudent(tableName, name string, roll int64, student_id, mother_name, father_name, guardian_name string) (int64, error) {
    return sms.AddStudent(tableName, name, roll, student_id, mother_name, father_name, guardian_name)
}

func (a *App) UpdateStudent12(tableName string, id int64, name string, roll int64, student_id string, mother_name string, father_name string, guardian_name string, term1atcm int64, term1atco int64, term1aips int64, term1aimpc int64, term2atcm int64, term2atco int64, term2aips int64, term2aimpc int64, term3atcm int64, term3atco int64, term3aips int64, term3aimpc int64) error {
    return sms.UpdateStudent12(tableName, id, name, roll, student_id, mother_name, father_name, guardian_name, term1atcm, term1atco, term1aips, term1aimpc, term2atcm, term2atco, term2aips, term2aimpc, term3atcm, term3atco, term3aips, term3aimpc)
}

func (a *App) UpdateStudent345(tableName string, id int64, name string, roll int64, student_id string, mother_name string, father_name string, guardian_name string, term1fl int64, term1sl int64, term1m int64, term1e int64, term1hpeth int64, term1hpepr int64, term1aweth int64, term1awepr int64, term2fl int64, term2sl int64, term2m int64, term2e int64, term2hpeth int64, term2hpepr int64, term2aweth int64, term2awepr int64, term3fl int64, term3sl int64, term3m int64, term3e int64, term3hpeth int64, term3hpepr int64, term3aweth int64, term3awepr int64) error {
    return sms.UpdateStudent345(tableName, id, name, roll, student_id, mother_name, father_name, guardian_name, term1fl, term1sl, term1m, term1e, term1hpeth, term1hpepr, term1aweth, term1awepr, term2fl, term2sl, term2m, term2e, term2hpeth, term2hpepr, term2aweth, term2awepr, term3fl, term3sl, term3m, term3e, term3hpeth, term3hpepr, term3aweth, term3awepr)
}

func (a *App) DeleteStudent(tableName string, id int64) error {
    return sms.DeleteStudent(tableName, id)
}

// func (a *App) GetStudent(tableName string, id int64) (*sms.Student, error) {
//     return sms.GetStudent(tableName, id)
// }
// func (a *App) GenerateAndDownloadPDF() (map[string]interface{}, error) {
//     // Generate PDF bytes
//     pdfBytes, err := pdf.GenerateStudentReport()
//     if err != nil {
//         return nil, err
//     }

//     // Convert bytes to base64 string
//     base64String := base64.StdEncoding.EncodeToString(pdfBytes)

//     // Return both the base64 string and filename
//     return map[string]interface{}{
//         "filename": "student_report.pdf",
//         "data":     base64String,
//     }, nil
// }





func (a *App) GenerateAndDownloadPDF(student_name string,class_name string,class_id int64,student_id int64) (map[string]interface{}, error) {
    // Generate PDF bytes
    pdfBytes, err := pdf.GenerateStudentReport(class_id,student_id)
    if err != nil {
        return nil, err
    }

    // Convert bytes to base64 string
    base64String := base64.StdEncoding.EncodeToString(pdfBytes)

    // Return both the base64 string and filename
    return map[string]interface{}{
        "filename": fmt.Sprintf("%s_%s.pdf", student_name,class_name),
        "data":     base64String,
    }, nil
}













// func (a *App) GeneratePDFURL() (string, error) {
//     // Start an HTTP server to serve the PDF at localhost
//     go func() {
//         http.HandleFunc("/download/student_report.pdf", pdf.ServePDF)
//         http.ListenAndServe(":8080", nil)
//     }()

//     // Return the URL for the frontend to redirect to
//     url := "http://localhost:8080/download/student_report.pdf"
//     return url, nil
// }







// Initialize the server once on application start




/* 
func startPDFServer() {
    http.HandleFunc("/download/student_report.pdf", pdf.ServePDF)
    go http.ListenAndServe(":5500", nil)
}

// GeneratePDFURL simply returns the URL after ensuring the server has started
func (a *App) GeneratePDFURL() (string, error) {
    // Ensure the server is only started once
    once.Do(startPDFServer)

    // Return the URL for the frontend to open in the browser
    url := "http://localhost:5500/download/student_report.pdf"
    return url, nil
}
*/