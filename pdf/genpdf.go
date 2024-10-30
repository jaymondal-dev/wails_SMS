/*
package pdf

import (
	"bytes"
	"fmt"
	"myproject/sms"

	"github.com/jung-kurt/gofpdf"
)

func GenerateStudentReport(class_id int64, student_id int64) ([]byte, error) {
	var student1 sms.Student12
	var student2 sms.Student345

	class, err := sms.GetClass(class_id)
	if err != nil {
		return nil, err
	}
	if class.Name == "Class 1" || class.Name == "Class 2" {
		student, err := sms.GetStudent12(class.TableName, student_id)
		if err != nil {
			return nil, err
		}
		student1 = *student
	} else {
		student, err := sms.GetStudent345(class.TableName, student_id)
		if err != nil {
			return nil, err
		}
		student2 = *student
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set title
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(190, 10, "Student Report")
	pdf.Ln(20)

	pdf.Cell(190, 10, fmt.Sprintf("Name %s", class.Name))

	pdf.Ln(20)

	// Add content to PDF (customize this as needed)
	pdf.SetFont("Arial", "", 12)
	if class.Name == "Class 1" || class.Name == "Class 2" {
		pdf.MultiCell(190, 10, fmt.Sprintf("This is a sample report.%s", student1.Name), "", "", false)

	} else {
		pdf.MultiCell(190, 10, fmt.Sprintf("This is a sample report.%s", student2.Name), "", "", false)

	}

	// Output PDF to a byte buffer
	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
*/

/*
package pdf

import (
	"bytes"
	"fmt"
	"myproject/sms"

	"github.com/jung-kurt/gofpdf"
)

// Grade calculation constants
const (
	gradeA  = 90
	gradeB  = 80
	gradeCp = 70
	gradeC  = 60
	gradeD  = 50
	gradeF  = 0
)

func calculateGrade(percentage float64) string {
	switch {
	case percentage >= gradeA:
		return "A"
	case percentage >= gradeB:
		return "B"
	case percentage >= gradeCp:
		return "C+"
	case percentage >= gradeC:
		return "C"
	case percentage >= gradeD:
		return "D"
	default:
		return "F"
	}
}

func GenerateStudentReport(class_id int64, student_id int64) ([]byte, error) {
	var student1 sms.Student12
	var student2 sms.Student345

	class, err := sms.GetClass(class_id)
	if err != nil {
		return nil, err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Header
	pdf.SetFont("Arial", "B", 20)
	pdf.CellFormat(190, 10, "Student Report Card", "", 1, "C", false, 0, "")

	// School Info
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(190, 10, "School Name Here", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(190, 8, fmt.Sprintf("Class: %s", class.Name), "", 1, "L", false, 0, "")
	pdf.CellFormat(190, 8, fmt.Sprintf("Academic Year: %d", class.Year), "", 1, "L", false, 0, "")
	pdf.Ln(10)

	if class.Name == "Class 1" || class.Name == "Class 2" {
		student, err := sms.GetStudent12(class.TableName, student_id)
		if err != nil {
			return nil, err
		}
		student1 = *student
		generateClass12Report(pdf, student1)
	} else {
		student, err := sms.GetStudent345(class.TableName, student_id)
		if err != nil {
			return nil, err
		}
		student2 = *student
		generateClass345Report(pdf, student2)
	}

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func generateClass12Report(pdf *gofpdf.Fpdf, student sms.Student12) {
	// Student Details
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, "Student Details", "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(50, 8, "Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.Name, "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Roll Number:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, fmt.Sprintf("%d", student.Roll), "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Student ID:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.StudentId, "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Father's Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.FatherName, "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Mother's Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.MotherName, "", 1, "L", false, 0, "")
	pdf.Ln(5)

	// Academic Performance Header
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, "Academic Performance", "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// Table headers with proper formatting
	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(200, 200, 200)

	// Headers
	headers := []struct {
		width float64
		text  string
	}{
		{40, "Subject"},
		{20, "Term 1"},
		{20, "Term 2"},
		{20, "Term 3"},
		{25, "Total"},
		{25, "Max"},
		{20, "%"},
		{20, "Grade"},
	}

	// Draw header row
	for _, h := range headers {
		pdf.CellFormat(h.width, 8, h.text, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	// Function to add subject row
	addSubjectRow := func(subject string, term1, term2, term3, maxTotal float64) {
		total := term1 + term2 + term3
		percentage := (total / maxTotal) * 100
		grade := calculateGrade(percentage)

		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(40, 8, subject, "1", 0, "L", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", term1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", term2), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", term3), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, fmt.Sprintf("%.1f", total), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, fmt.Sprintf("%.0f", maxTotal), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", percentage), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, grade, "1", 1, "C", false, 0, "")
	}

	// Add subject rows
	maxATCM := float64(sms.Term1atcm + sms.Term2atcm + sms.Term3atcm)
	addSubjectRow("ATCM", float64(student.Term1atcm), float64(student.Term2atcm), float64(student.Term3atcm), maxATCM)

	maxATCO := float64(sms.Term1atco + sms.Term2atco + sms.Term3atco)
	addSubjectRow("ATCO", float64(student.Term1atco), float64(student.Term2atco), float64(student.Term3atco), maxATCO)

	maxAIPS := float64(sms.Term1aips + sms.Term2aips + sms.Term3aips)
	addSubjectRow("AIPS", float64(student.Term1aips), float64(student.Term2aips), float64(student.Term3aips), maxAIPS)

	maxAIMPC := float64(sms.Term1aimpc + sms.Term2aimpc + sms.Term3aimpc)
	addSubjectRow("AIMPC", float64(student.Term1aimpc), float64(student.Term2aimpc), float64(student.Term3aimpc), maxAIMPC)

	// Calculate overall performance
	totalObtained := float64(
		student.Term1atcm + student.Term2atcm + student.Term3atcm +
			student.Term1atco + student.Term2atco + student.Term3atco +
			student.Term1aips + student.Term2aips + student.Term3aips +
			student.Term1aimpc + student.Term2aimpc + student.Term3aimpc)

	totalMax := maxATCM + maxATCO + maxAIPS + maxAIMPC
	overallPercentage := (totalObtained / totalMax) * 100

	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, fmt.Sprintf("Overall Performance: %.1f%% (%s)",
		overallPercentage, calculateGrade(overallPercentage)), "", 1, "L", false, 0, "")

	// Add remarks
	pdf.Ln(5)
	pdf.CellFormat(190, 10, "Remarks:", "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(190, 8, generateRemarks(overallPercentage), "1", "L", false)
}

func generateClass345Report(pdf *gofpdf.Fpdf, student sms.Student345) {
	// Student Details
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, "Student Details", "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(50, 8, "Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.Name, "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Roll Number:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, fmt.Sprintf("%d", student.Roll), "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Student ID:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.StudentId, "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Father's Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.FatherName, "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Mother's Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.MotherName, "", 1, "L", false, 0, "")
	pdf.Ln(5)

	// Academic Performance Header
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, "Academic Performance", "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// Table headers
	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(200, 200, 200)

	headers := []struct {
		width float64
		text  string
	}{
		{40, "Subject"},
		{20, "Term 1"},
		{20, "Term 2"},
		{20, "Term 3"},
		{25, "Total"},
		{25, "Max"},
		{20, "%"},
		{20, "Grade"},
	}

	for _, h := range headers {
		pdf.CellFormat(h.width, 8, h.text, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	// Function to add subject row
	addSubjectRow := func(subject string, term1, term2, term3, maxTotal float64) {
		total := term1 + term2 + term3
		percentage := (total / maxTotal) * 100
		grade := calculateGrade(percentage)

		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(40, 8, subject, "1", 0, "L", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", term1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", term2), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", term3), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, fmt.Sprintf("%.1f", total), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, fmt.Sprintf("%.0f", maxTotal), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", percentage), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, grade, "1", 1, "C", false, 0, "")
	}

	// Add rows for each subject
	maxFL := float64(sms.Term1fl + sms.Term2fl + sms.Term3fl)
	addSubjectRow("First Lang", float64(student.Term1fl), float64(student.Term2fl), float64(student.Term3fl), maxFL)

	maxSL := float64(sms.Term1sl + sms.Term2sl + sms.Term3sl)
	addSubjectRow("Second Lang", float64(student.Term1sl), float64(student.Term2sl), float64(student.Term3sl), maxSL)

	maxMath := float64(sms.Term1M + sms.Term2M + sms.Term3M)
	addSubjectRow("Mathematics", float64(student.Term1M), float64(student.Term2M), float64(student.Term3M), maxMath)

	maxEng := float64(sms.Term1E + sms.Term2E + sms.Term3E)
	addSubjectRow("English", float64(student.Term1E), float64(student.Term2E), float64(student.Term3E), maxEng)

	// Calculate overall performance
	totalObtained := float64(
		student.Term1fl + student.Term2fl + student.Term3fl +
			student.Term1sl + student.Term2sl + student.Term3sl +
			student.Term1M + student.Term2M + student.Term3M +
			student.Term1E + student.Term2E + student.Term3E)

	totalMax := maxFL + maxSL + maxMath + maxEng
	overallPercentage := (totalObtained / totalMax) * 100

	// Add Overall Performance
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, fmt.Sprintf("Overall Performance: %.1f%% (%s)",
		overallPercentage, calculateGrade(overallPercentage)), "", 1, "L", false, 0, "")

	// Add remarks section
	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, "Remarks:", "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(190, 8, generateRemarks(overallPercentage), "1", "L", false)
}

// generateRemarks generates remarks based on the student's overall percentage.
func generateRemarks(percentage float64) string {
	switch {
	case percentage >= 90:
		return "Excellent performance. Keep up the great work!"
	case percentage >= 75:
		return "Good job! Continue to strive for excellence."
	case percentage >= 60:
		return "Fair performance, but there's room for improvement."
	case percentage >= 50:
		return "Needs improvement. Consider focusing more on your studies."
	default:
		return "Poor performance. Requires significant improvement."
	}
}
*/
package pdf

import (
	"bytes"
	"fmt"
	"myproject/sms"

	"github.com/jung-kurt/gofpdf"
)

// Grade calculation constants
const (
	gradeA  = 90
	gradeB  = 80
	gradeCp = 70
	gradeC  = 60
	gradeD  = 50
	gradeF  = 0
)

func calculateGrade(percentage float64) string {
	switch {
	case percentage >= gradeA:
		return "A"
	case percentage >= gradeB:
		return "B"
	case percentage >= gradeCp:
		return "C+"
	case percentage >= gradeC:
		return "C"
	case percentage >= gradeD:
		return "D"
	default:
		return "F"
	}
}

func GenerateStudentReport(class_id int64, student_id int64) ([]byte, error) {
	var student1 sms.Student12
	var student2 sms.Student345

	class, err := sms.GetClass(class_id)
	if err != nil {
		return nil, err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Header
	pdf.SetFont("Arial", "B", 20)
	pdf.CellFormat(190, 10, "Student Report Card", "", 1, "C", false, 0, "")

	// School Info
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(190, 10, fmt.Sprintf("%s",sms.SchoolName), "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(190, 8, fmt.Sprintf("Class: %s", class.Name), "", 1, "L", false, 0, "")
	pdf.CellFormat(190, 8, fmt.Sprintf("Academic Year: %d", class.Year), "", 1, "L", false, 0, "")
	pdf.Ln(10)

	if class.Name == "Class 1" || class.Name == "Class 2" {
		student, err := sms.GetStudent12(class.TableName, student_id)
		if err != nil {
			return nil, err
		}
		student1 = *student
		generateClass12Report(pdf, student1)
	} else {
		student, err := sms.GetStudent345(class.TableName, student_id)
		if err != nil {
			return nil, err
		}
		student2 = *student
		generateClass345Report(pdf, student2)
	}

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func generateClass12Report(pdf *gofpdf.Fpdf, student sms.Student12) {
	// Student Details
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, "Student Details", "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(50, 8, "Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.Name, "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Roll Number:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, fmt.Sprintf("%d", student.Roll), "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Student ID:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.StudentId, "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Father's Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.FatherName, "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Mother's Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.MotherName, "", 1, "L", false, 0, "")
	pdf.Ln(5)

	// Academic Performance Header
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, "Academic Performance", "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// Table headers with proper formatting
	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(200, 200, 200)

	// Headers
	headers := []struct {
		width float64
		text  string
	}{
		{40, "Subject"},
		{20, "Term 1"},
		{20, "Term 2"},
		{20, "Term 3"},
		{25, "Total"},
		{25, "Max"},
		{20, "%"},
		{20, "Grade"},
	}

	// Draw header row
	for _, h := range headers {
		pdf.CellFormat(h.width, 8, h.text, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	// Function to add subject row
	addSubjectRow := func(subject string, term1, term2, term3, maxTotal float64) {
		total := term1 + term2 + term3
		percentage := (total / maxTotal) * 100
		grade := calculateGrade(percentage)

		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(40, 8, subject, "1", 0, "L", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", term1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", term2), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", term3), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, fmt.Sprintf("%.1f", total), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, fmt.Sprintf("%.0f", maxTotal), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", percentage), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, grade, "1", 1, "C", false, 0, "")
	}

	// Add subject rows
	maxATCM := float64(sms.Term1atcm + sms.Term2atcm + sms.Term3atcm)
	addSubjectRow("ATCM", float64(student.Term1atcm), float64(student.Term2atcm), float64(student.Term3atcm), maxATCM)

	maxATCO := float64(sms.Term1atco + sms.Term2atco + sms.Term3atco)
	addSubjectRow("ATCO", float64(student.Term1atco), float64(student.Term2atco), float64(student.Term3atco), maxATCO)

	maxAIPS := float64(sms.Term1aips + sms.Term2aips + sms.Term3aips)
	addSubjectRow("AIPS", float64(student.Term1aips), float64(student.Term2aips), float64(student.Term3aips), maxAIPS)

	maxAIMPC := float64(sms.Term1aimpc + sms.Term2aimpc + sms.Term3aimpc)
	addSubjectRow("AIMPC", float64(student.Term1aimpc), float64(student.Term2aimpc), float64(student.Term3aimpc), maxAIMPC)

	// Calculate overall performance
	totalObtained := float64(
		student.Term1atcm + student.Term2atcm + student.Term3atcm +
			student.Term1atco + student.Term2atco + student.Term3atco +
			student.Term1aips + student.Term2aips + student.Term3aips +
			student.Term1aimpc + student.Term2aimpc + student.Term3aimpc)

	totalMax := maxATCM + maxATCO + maxAIPS + maxAIMPC
	overallPercentage := (totalObtained / totalMax) * 100

	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, fmt.Sprintf("Overall Performance: %.1f%% (%s)",
		overallPercentage, calculateGrade(overallPercentage)), "", 1, "L", false, 0, "")

	// Add remarks
	pdf.Ln(5)
	pdf.CellFormat(190, 10, "Remarks:", "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(190, 8, generateRemarks(overallPercentage), "1", "L", false)
}

func generateClass345Report(pdf *gofpdf.Fpdf, student sms.Student345) {
	// Student Details
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, "Student Details", "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(50, 8, "Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.Name, "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Roll Number:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, fmt.Sprintf("%d", student.Roll), "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Student ID:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.StudentId, "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Father's Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.FatherName, "", 1, "L", false, 0, "")

	pdf.CellFormat(50, 8, "Mother's Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(140, 8, student.MotherName, "", 1, "L", false, 0, "")
	pdf.Ln(5)

	// Academic Performance Header
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, "Academic Performance", "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// Table headers
	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(200, 200, 200)

	headers := []struct {
		width float64
		text  string
	}{
		{40, "Subject"},
		{20, "Term 1"},
		{20, "Term 2"},
		{20, "Term 3"},
		{25, "Total"},
		{25, "Max"},
		{20, "%"},
		{20, "Grade"},
	}

	for _, h := range headers {
		pdf.CellFormat(h.width, 8, h.text, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	// Function to add subject row
	addSubjectRow := func(subject string, term1, term2, term3, maxTotal float64) {
		total := term1 + term2 + term3
		percentage := (total / maxTotal) * 100
		grade := calculateGrade(percentage)

		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(40, 8, subject, "1", 0, "L", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", term1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", term2), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", term3), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, fmt.Sprintf("%.1f", total), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, fmt.Sprintf("%.0f", maxTotal), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", percentage), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, grade, "1", 1, "C", false, 0, "")
	}

	// Add rows for each subject
	maxFL := float64(sms.Term1fl + sms.Term2fl + sms.Term3fl)
	addSubjectRow("First Lang", float64(student.Term1fl), float64(student.Term2fl), float64(student.Term3fl), maxFL)

	maxSL := float64(sms.Term1sl + sms.Term2sl + sms.Term3sl)
	addSubjectRow("Second Lang", float64(student.Term1sl), float64(student.Term2sl), float64(student.Term3sl), maxSL)

	maxMath := float64(sms.Term1M + sms.Term2M + sms.Term3M)
	addSubjectRow("Mathematics", float64(student.Term1M), float64(student.Term2M), float64(student.Term3M), maxMath)

	maxEng := float64(sms.Term1E + sms.Term2E + sms.Term3E)
	addSubjectRow("English", float64(student.Term1E), float64(student.Term2E), float64(student.Term3E), maxEng)

	maxHPETH := float64(sms.Term1hpeth + sms.Term2hpeth + sms.Term3hpeth)
	addSubjectRow("HPET - Theory", float64(student.Term1hpeth), float64(student.Term2hpeth), float64(student.Term3hpeth), maxHPETH)

	maxHPEPR := float64(sms.Term1hpepr + sms.Term2hpepr + sms.Term3hpepr)
	addSubjectRow("HPET - Practical", float64(student.Term1hpepr), float64(student.Term2hpepr), float64(student.Term3hpepr), maxHPEPR)

	maxAWETH := float64(sms.Term1aweth + sms.Term2aweth + sms.Term3aweth)
	addSubjectRow("AWE - Theory", float64(student.Term1aweth), float64(student.Term2aweth), float64(student.Term3aweth), maxAWETH)

	maxAWEPR := float64(sms.Term1awepr + sms.Term2awepr + sms.Term3awepr)
	addSubjectRow("AWE - Practical", float64(student.Term1awepr), float64(student.Term2awepr), float64(student.Term3awepr), maxAWEPR)

	// Calculate overall performance
	totalObtained := float64(
		student.Term1fl + student.Term2fl + student.Term3fl +
			student.Term1sl + student.Term2sl + student.Term3sl +
			student.Term1M + student.Term2M + student.Term3M +
			student.Term1E + student.Term2E + student.Term3E +
			student.Term1hpeth + student.Term2hpeth + student.Term3hpeth +
			student.Term1hpepr + student.Term2hpepr + student.Term3hpepr +
			student.Term1aweth + student.Term2aweth + student.Term3aweth +
			student.Term1awepr + student.Term2awepr + student.Term3awepr)

	totalMax := maxFL + maxSL + maxMath + maxEng + maxHPETH + maxHPEPR + maxAWETH + maxAWEPR
	overallPercentage := (totalObtained / totalMax) * 100

	// Add Overall Performance
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, fmt.Sprintf("Overall Performance: %.1f%% (%s)",
		overallPercentage, calculateGrade(overallPercentage)), "", 1, "L", false, 0, "")

	// Add remarks section
	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, "Remarks:", "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(190, 8, generateRemarks(overallPercentage), "1", "L", false)
}

// generateRemarks generates remarks based on the student's overall percentage.
func generateRemarks(percentage float64) string {
	switch {
	case percentage >= 90:
		return "Excellent performance. Keep up the great work!"
	case percentage >= 75:
		return "Good job! Continue to strive for excellence."
	case percentage >= 60:
		return "Fair performance, but there's room for improvement."
	case percentage >= 50:
		return "Needs improvement. Consider focusing more on your studies."
	default:
		return "Poor performance. Requires significant improvement."
	}
}