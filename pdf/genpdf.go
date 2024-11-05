package pdf

import (
	"bytes"
	"fmt"
	"myproject/sms"

	"github.com/jung-kurt/gofpdf"
)

// Grade calculation constants

const (
	gradeAp = 90
	gradeA  = 80
	gradeBp = 70
	gradeB  = 60
	gradeCp = 45
	gradeC  = 25
)

// calculateGrade calculates the letter grade based on the percentage.
func calculateGrade(percentage float64) string {
	switch {
	case percentage >= gradeAp:
		return "A+"
	case percentage >= gradeA:
		return "A"
	case percentage >= gradeBp:
		return "B+"
	case percentage >= gradeB:
		return "B"
	case percentage >= gradeCp:
		return "C+"
	case percentage >= gradeC:
		return "C"
	default:
		return "D"
	}
}

// GenerateStudentReport generates the student report card in PDF format.
func GenerateStudentReport(class_id int64, student_id int64) ([]byte, error) {
	var student1 sms.Student12  // Student struct for classes 1 and 2
	var student2 sms.Student345 // Student struct for classes 3, 4, and 5

	// Fetch the class details
	class, err := sms.GetClass(class_id)
	if err != nil {
		return nil, err
	}

	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Add the report card header
	pdf.SetFont("Arial", "B", 20)
	pdf.CellFormat(190, 10, "Student Report Card", "", 1, "C", false, 0, "")

	// Add school information
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(190, 10, fmt.Sprintf("%s", sms.SchoolName), "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "B", 10)
	
	pdf.CellFormat(190, 10, fmt.Sprintf("Dise Code: %s", sms.SchoolDiseCode), "", 1, "C", false, 0, "")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(95, 8, fmt.Sprintf("Class: %s", class.Name), "", 0, "L", false, 0, "")

	pdf.CellFormat(95, 8, fmt.Sprintf("%s", sms.SchoolAddressl1), "", 1, "L", false, 0, "")
	pdf.CellFormat(95, 8, fmt.Sprintf("Academic Year: %d", class.Year), "", 0, "L", false, 0, "")

	pdf.CellFormat(95, 8, fmt.Sprintf("%s", sms.SchoolAddressl2), "", 1, "L", false, 0, "")
	pdf.CellFormat(95, 8, fmt.Sprintf(""), "", 0, "L", false, 0, "")

	pdf.CellFormat(95, 8, fmt.Sprintf("%s", sms.SchoolAddressl3), "", 1, "L", false, 0, "")

	pdf.Ln(3)

	// Determine the student structure based on the class
	if class.Name == "Class 1" || class.Name == "Class 2" ||class.Name=="PP" {
		// Fetch student data for classes 1 and 2
		student, err := sms.GetStudent12(class.TableName, student_id)
		if err != nil {
			return nil, err
		}
		student1 = *student                  // Assign the student data to the struct
		generateClass12Report(pdf, student1) // Generate the report for classes 1 and 2
	} else {
		// Fetch student data for classes 3, 4, and 5
		student, err := sms.GetStudent345(class.TableName, student_id)
		if err != nil {
			return nil, err
		}
		student2 = *student                   // Assign the student data to the struct
		generateClass345Report(pdf, student2) // Generate the report for classes 3, 4, and 5
	}

	// Output the PDF to a byte buffer
	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil // Return the generated PDF as bytes
}

// generateClass12Report generates the report for classes 1 and 2.
func generateClass12Report(pdf *gofpdf.Fpdf, student sms.Student12) {
	// Add student details section
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, "Student Details", "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(30, 8, "Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(65, 8, student.Name, "", 0, "L", false, 0, "")

	pdf.CellFormat(30, 8, "Roll Number:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(65, 8, fmt.Sprintf("%d", student.Roll), "", 1, "L", false, 0, "")

	pdf.CellFormat(30, 8, "Student ID:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(65, 8, student.StudentId, "", 0, "L", false, 0, "")

	pdf.CellFormat(30, 8, "Father's Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(65, 8, student.FatherName, "", 1, "L", false, 0, "")

	pdf.CellFormat(30, 8, "Guardian's Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(160, 8, student.GuardianName, "", 1, "L", false, 0, "")
	pdf.Ln(3)

	// Add Academic Performance Header
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, "Academic Performance", "", 1, "C", false, 0, "")
	pdf.CellFormat(190, 10, "Summative Evaluation", "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// Configure table headers with proper formatting
	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(200, 200, 200)

	// Define table headers
	headers := []struct {
		width float64
		text  string
	}{
		{90, "Subject"},
		{15, "Term 1"},
		{15, "Term 2"},
		{15, "Term 3"},
		{20, "Total"},
		
		{20, "Percentage"},
		{15, "Grade"},
	}

	// Draw the header row
	for _, h := range headers {
		pdf.CellFormat(h.width, 8, h.text, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	// Function to add a subject row to the table
	addSubjectRow := func(subject string, term1, term2, term3, maxTotal float64) {
		total := term1 + term2 + term3         // Calculate the total marks
		percentage := (total / maxTotal) * 100 // Calculate the percentage
		grade := calculateGrade(percentage)    // Get the letter grade

		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(90, 8, subject, "1", 0, "L", false, 0, "")
		pdf.CellFormat(15, 8, fmt.Sprintf("%.1f", term1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(15, 8, fmt.Sprintf("%.1f", term2), "1", 0, "C", false, 0, "")
		pdf.CellFormat(15, 8, fmt.Sprintf("%.1f", term3), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", total), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", percentage), "1", 0, "C", false, 0, "")
		pdf.CellFormat(15, 8, grade, "1", 1, "C", false, 0, "")
	}

	// Add subject rows to the table
	maxATCM := float64(sms.Term1atcm + sms.Term2atcm + sms.Term3atcm) // Calculate maximum marks for ATCM
	addSubjectRow(fmt.Sprintf("%s",sms.AtcmSubE), float64(student.Term1atcm), float64(student.Term2atcm), float64(student.Term3atcm), maxATCM)

	maxATCO := float64(sms.Term1atco + sms.Term2atco + sms.Term3atco) // Calculate maximum marks for ATCO
	addSubjectRow(fmt.Sprintf("%s",sms.AtcoSubE), float64(student.Term1atco), float64(student.Term2atco), float64(student.Term3atco), maxATCO)

	maxAIPS := float64(sms.Term1aips + sms.Term2aips + sms.Term3aips) // Calculate maximum marks for AIPS
	addSubjectRow(fmt.Sprintf("%s",sms.AipsSubE), float64(student.Term1aips), float64(student.Term2aips), float64(student.Term3aips), maxAIPS)

	maxAIMPC := float64(sms.Term1aimpc + sms.Term2aimpc + sms.Term3aimpc) // Calculate maximum marks for AIMPC
	addSubjectRow(fmt.Sprintf("%s",sms.AimpcSubE), float64(student.Term1aimpc), float64(student.Term2aimpc), float64(student.Term3aimpc), maxAIMPC)
    // Calculate the overall performance
	totalObtained := float64(
		student.Term1atcm + student.Term2atcm + student.Term3atcm +
			student.Term1atco + student.Term2atco + student.Term3atco +
			student.Term1aips + student.Term2aips + student.Term3aips +
			student.Term1aimpc + student.Term2aimpc + student.Term3aimpc) // Calculate total marks obtained

	totalMax := maxATCM + maxATCO + maxAIPS + maxAIMPC    // Calculate total maximum marks
	overallPercentage := (totalObtained / totalMax) * 100 // Calculate overall percentage
	
	
	pdf.SetFont("Arial", "B", 12)
	
	pdf.CellFormat(190, 10, "Formative Evaluation", "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// Configure table headers with proper formatting
	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(200, 200, 200)

	// Define table headers
	headers1 := []struct {
		width float64
		text  string
	}{
		{145, "Name of the Indicators"},
		{15, "Term 1"},
		{15, "Term 2"},
		{15, "Term 3"},
	}

	// Draw the header row
	for _, h := range headers1 {
		pdf.CellFormat(h.width, 8, h.text, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(145, 8,"Participaton", "1", 0, "L", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 1, "C", false, 0, "")
	pdf.CellFormat(145, 8,"Questioning and Experimentation", "1", 0, "L", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 1, "C", false, 0, "")
	pdf.CellFormat(145, 8,"Interpretation and Application", "1", 0, "L", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 1, "C", false, 0, "")
	pdf.CellFormat(145, 8,"Empathy and Cooperation", "1", 0, "L", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 1, "C", false, 0, "")
	pdf.CellFormat(145, 8,"Aesthetic and Creative Expression", "1", 0, "L", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 1, "C", false, 0, "")
	// Add Overall Performance section
	pdf.Ln(3)
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, fmt.Sprintf("Overall Performance: %.1f%% (%s)",
		overallPercentage, calculateGrade(overallPercentage)), "", 1, "L", false, 0, "")

	// Add Remarks section
	pdf.Ln(3)
	pdf.CellFormat(25, 10, "Remarks:", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(165, 10, generateRemarks(overallPercentage), "", "L", false)
	pdf.Ln(15)
	pdf.SetFont("Arial", "B", 10)
	
	pdf.CellFormat(190,10,"Signature of the Class Teacher with Date:", "", 1, "L", false, 0, "")
	pdf.CellFormat(190,10,"Signature of the Head of the Institution with Date:", "", 1, "L", false, 0, "")
	pdf.CellFormat(190,10,"Signature of the Guardian with Date:", "", 1, "L", false, 0, "")
	
	
}

// generateClass345Report generates the report for classes 3, 4, and 5.
func generateClass345Report(pdf *gofpdf.Fpdf, student sms.Student345) {
	// Add student details section
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, "Student Details", "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(30, 8, "Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(65, 8, student.Name, "", 0, "L", false, 0, "")

	pdf.CellFormat(30, 8, "Roll Number:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(65, 8, fmt.Sprintf("%d", student.Roll), "", 1, "L", false, 0, "")

	pdf.CellFormat(30, 8, "Student ID:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(65, 8, student.StudentId, "", 0, "L", false, 0, "")

	pdf.CellFormat(30, 8, "Father's Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(65, 8, student.FatherName, "", 1, "L", false, 0, "")

	pdf.CellFormat(30, 8, "Guardian's Name:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(160, 8, student.GuardianName, "", 1, "L", false, 0, "")
	pdf.Ln(3)

	// Add Academic Performance Header
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, "Academic Performance", "", 1, "C", false, 0, "")
	pdf.CellFormat(190, 10, "Summative Evaluation", "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// Configure table headers
	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(200, 200, 200)

	headers := []struct {
		width float64
		text  string
	}{
		{90, "Subject"},
		{15, "Term 1"},
		{15, "Term 2"},
		{15, "Term 3"},
		{20, "Total"},
		
		{20, "Percentage"},
		{15, "Grade"},
	}

	// Draw the header row
	for _, h := range headers {
		pdf.CellFormat(h.width, 8, h.text, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	// Function to add a subject row to the table
	addSubjectRow := func(subject string, term1, term2, term3, maxTotal float64) {
		total := term1 + term2 + term3         // Calculate the total marks
		percentage := (total / maxTotal) * 100 // Calculate the percentage
		grade := calculateGrade(percentage)    // Get the letter grade

		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(90, 8, subject, "1", 0, "L", false, 0, "")
		pdf.CellFormat(15, 8, fmt.Sprintf("%.1f", term1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(15, 8, fmt.Sprintf("%.1f", term2), "1", 0, "C", false, 0, "")
		pdf.CellFormat(15, 8, fmt.Sprintf("%.1f", term3), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", total), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%.1f", percentage), "1", 0, "C", false, 0, "")
		pdf.CellFormat(15, 8, grade, "1", 1, "C", false, 0, "")
	}

	// Add subject rows to the table
	maxFL := float64(sms.Term1fl + sms.Term2fl + sms.Term3fl) // Calculate maximum marks for First Lang
	addSubjectRow(fmt.Sprintf("%s",sms.Term1flSubE), float64(student.Term1fl), float64(student.Term2fl), float64(student.Term3fl), maxFL)

	maxSL := float64(sms.Term1sl + sms.Term2sl + sms.Term3sl) // Calculate maximum marks for Second Lang
	addSubjectRow(fmt.Sprintf("%s",sms.Term1slSubE), float64(student.Term1sl), float64(student.Term2sl), float64(student.Term3sl), maxSL)

	maxMath := float64(sms.Term1M + sms.Term2M + sms.Term3M) // Calculate maximum marks for Mathematics
	addSubjectRow(fmt.Sprintf("%s",sms.Term1MSubE), float64(student.Term1M), float64(student.Term2M), float64(student.Term3M), maxMath)

	maxEng := float64(sms.Term1E + sms.Term2E + sms.Term3E) // Calculate maximum marks for English
	addSubjectRow(fmt.Sprintf("%s",sms.Term1ESubE), float64(student.Term1E), float64(student.Term2E), float64(student.Term3E), maxEng)

	maxHPETH := float64(sms.Term1hpeth + sms.Term2hpeth + sms.Term3hpeth) // Calculate maximum marks for HPET - Theory
	addSubjectRow(fmt.Sprintf("%s",sms.Term1hpethSubE), float64(student.Term1hpeth), float64(student.Term2hpeth), float64(student.Term3hpeth), maxHPETH)

	maxHPEPR := float64(sms.Term1hpepr + sms.Term2hpepr + sms.Term3hpepr) // Calculate maximum marks for HPET - Practical
	addSubjectRow(fmt.Sprintf("%s",sms.Term1hpeprSubE), float64(student.Term1hpepr), float64(student.Term2hpepr), float64(student.Term3hpepr), maxHPEPR)

	maxAWETH := float64(sms.Term1aweth + sms.Term2aweth + sms.Term3aweth) // Calculate maximum marks for AWE - Theory
	addSubjectRow(fmt.Sprintf("%s",sms.Term1awethSubE), float64(student.Term1aweth), float64(student.Term2aweth), float64(student.Term3aweth), maxAWETH)

	maxAWEPR := float64(sms.Term1awepr + sms.Term2awepr + sms.Term3awepr) // Calculate maximum marks for AWE - Practical
	addSubjectRow(fmt.Sprintf("%s",sms.Term1aweprSubE), float64(student.Term1awepr), float64(student.Term2awepr), float64(student.Term3awepr), maxAWEPR)

	// Calculate the overall performance
	totalObtained := float64(
		student.Term1fl + student.Term2fl + student.Term3fl +
			student.Term1sl + student.Term2sl + student.Term3sl +
			student.Term1M + student.Term2M + student.Term3M +
			student.Term1E + student.Term2E + student.Term3E +
			student.Term1hpeth + student.Term2hpeth + student.Term3hpeth +
			student.Term1hpepr + student.Term2hpepr + student.Term3hpepr +
			student.Term1aweth + student.Term2aweth + student.Term3aweth +
			student.Term1awepr + student.Term2awepr + student.Term3awepr) // Calculate total marks obtained

	totalMax := maxFL + maxSL + maxMath + maxEng + maxHPETH + maxHPEPR + maxAWETH + maxAWEPR // Calculate total maximum marks
	overallPercentage := (totalObtained / totalMax) * 100                                    // Calculate overall percentage
	pdf.SetFont("Arial", "B", 12)
	
	pdf.CellFormat(190, 10, "Formative Evaluation", "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// Configure table headers with proper formatting
	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(200, 200, 200)

	// Define table headers
	headers1 := []struct {
		width float64
		text  string
	}{
		{145, "Name of the Indicators"},
		{15, "Term 1"},
		{15, "Term 2"},
		{15, "Term 3"},
	}

	// Draw the header row
	for _, h := range headers1 {
		pdf.CellFormat(h.width, 8, h.text, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(145, 8,"Participaton", "1", 0, "L", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 1, "C", false, 0, "")
	pdf.CellFormat(145, 8,"Questioning and Experimentation", "1", 0, "L", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 1, "C", false, 0, "")
	pdf.CellFormat(145, 8,"Interpretation and Application", "1", 0, "L", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 1, "C", false, 0, "")
	pdf.CellFormat(145, 8,"Empathy and Cooperation", "1", 0, "L", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 1, "C", false, 0, "")
	pdf.CellFormat(145, 8,"Aesthetic and Creative Expression", "1", 0, "L", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 8,"", "1", 1, "C", false, 0, "")
	// Add Overall Performance section
	pdf.Ln(3)
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 10, fmt.Sprintf("Overall Performance: %.1f%% (%s)",
		overallPercentage, calculateGrade(overallPercentage)), "", 1, "L", false, 0, "")

	// Add Remarks section
	pdf.Ln(3)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(25, 10, "Remarks:", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(165, 10, generateRemarks(overallPercentage), "", "L", false)
	pdf.Ln(15)
	pdf.SetFont("Arial", "B", 10)
	
	pdf.CellFormat(190,10,"Signature of the Class Teacher with Date:", "", 1, "L", false, 0, "")
	pdf.CellFormat(190,10,"Signature of the Head of the Institution with Date:", "", 1, "L", false, 0, "")
	pdf.CellFormat(190,10,"Signature of the Guardian with Date:", "", 1, "L", false, 0, "")
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
