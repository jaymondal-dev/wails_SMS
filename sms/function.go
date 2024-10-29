// SMS/sms/function.go
package sms

import (
	"database/sql"
	"fmt"
)

func CreateClass(name string, year int64) (int64, error) {
	if name == "" {
		return 0, fmt.Errorf("class name cannot be empty")
	}

	if year <= 0 {
		return 0, fmt.Errorf("invalid year")
	}

	var exists bool
	err := DB.Get(&exists, "SELECT EXISTS(SELECT 1 FROM classes WHERE name = ? AND year = ?)", name, year)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, fmt.Errorf("class already exists")
	}

	tableName := Generate16CharAlphaString()
	var query string
	if name == "Class 3" || name == "Class 4" || name == "Class 5" {
		query = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name VARCHAR(255) NOT NULL,
				roll INTEGER NOT NULL,
				student_id VARCHAR(50) NOT NULL,
				mother_name VARCHAR(255),
				father_name VARCHAR(255),
				guardian_name VARCHAR(255),
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				term1fl INTEGER DEFAULT 0,
				term1sl INTEGER DEFAULT 0,
				term1m INTEGER DEFAULT 0,
				term1e INTEGER DEFAULT 0,
				term1hpeth INTEGER DEFAULT 0,
				term1hpepr INTEGER DEFAULT 0,
				term1aweth INTEGER DEFAULT 0,
				term1awepr INTEGER DEFAULT 0,
				term2fl INTEGER DEFAULT 0,
				term2sl INTEGER DEFAULT 0,
				term2m INTEGER DEFAULT 0,
				term2e INTEGER DEFAULT 0,
				term2hpeth INTEGER DEFAULT 0,
				term2hpepr INTEGER DEFAULT 0,
				term2aweth INTEGER DEFAULT 0,
				term2awepr INTEGER DEFAULT 0,
				term3fl INTEGER DEFAULT 0,
				term3sl INTEGER DEFAULT 0,
				term3m INTEGER DEFAULT 0,
				term3e INTEGER DEFAULT 0,
				term3hpeth INTEGER DEFAULT 0,
				term3hpepr INTEGER DEFAULT 0,
				term3aweth INTEGER DEFAULT 0,
				term3awepr INTEGER DEFAULT 0
								)`, tableName)
	} else if name == "Class 1" || name == "Class 2" {
		query = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name VARCHAR(255) NOT NULL,
				roll INTEGER NOT NULL,
				student_id VARCHAR(50) NOT NULL,
				mother_name VARCHAR(255),
				father_name VARCHAR(255),
				guardian_name VARCHAR(255),
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				term1atcm INTEGER DEFAULT 0,
				term1atco INTEGER DEFAULT 0,
				term1aips INTEGER DEFAULT 0,
				term1aimpc INTEGER DEFAULT 0,
				term2atcm INTEGER DEFAULT 0,
				term2atco INTEGER DEFAULT 0,
				term2aips INTEGER DEFAULT 0,
				term2aimpc INTEGER DEFAULT 0,
				term3atcm INTEGER DEFAULT 0,
				term3atco INTEGER DEFAULT 0,
				term3aips INTEGER DEFAULT 0,
				term3aimpc INTEGER DEFAULT 0
							)`, tableName)
	} else {
		return 0, fmt.Errorf("invalid class name: must be Class 1-5")
	}

	tx, err := DB.Beginx()
	if err != nil {
		return 0, err
	}

	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	result, err := tx.Exec("INSERT INTO classes (name, year, tablename) VALUES (?, ?, ?)", name, year, tableName)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func ListClasses() ([]ClassInfo, error) {
	var classes []ClassInfo
	err := DB.Select(&classes, "SELECT id, name, year, tablename, created_at FROM classes")
	return classes, err
}

func GetClass(id int64) (*ClassInfo, error) {
	var class ClassInfo
	err := DB.Get(&class, "SELECT id, name, year, tablename, created_at FROM classes WHERE id = ?", id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &class, err
}

func UpdateClass(id int64, name string, year int) error {
	_, err := DB.Exec("UPDATE classes SET name = ?, year = ? WHERE id = ?", name, year, id)
	return err
}

func DeleteClass(id int64) error {
	tx, err := DB.Beginx()
	if err != nil {
		return err
	}

	var tableName string
	err = tx.Get(&tableName, "SELECT tablename FROM classes WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM classes WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName))
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func ListStudents12(tableName string) ([]Student12, error) {
	var students []Student12
	query := fmt.Sprintf("SELECT id, name,roll,student_id,mother_name,father_name,guardian_name, created_at ,term1atcm,term1atco,term1aips,term1aimpc,term2atcm,term2atco,term2aips,term2aimpc,term3atcm,term3atco,term3aips,term3aimpc FROM %s", tableName)
	err := DB.Select(&students, query)
	return students, err
}
func ListStudents345(tableName string) ([]Student345, error) {
	var students []Student345
	query := fmt.Sprintf("SELECT id, name,roll,student_id,mother_name,father_name,guardian_name, created_at ,term1fl,term1sl,term1m,term1e,term1hpeth,term1hpepr,term1aweth,term1awepr,term2fl,term2sl,term2m,term2e,term2hpeth,term2hpepr,term2aweth,term2awepr,term3fl,term3sl,term3m,term3e,term3hpeth,term3hpepr,term3aweth,term3awepr FROM %s", tableName)
	err := DB.Select(&students, query)
	return students, err
}

func AddStudent(tableName, name string, roll int64, student_id string, mother_name string, father_name string, guardian_name string) (int64, error) {

	if name == "" {
		return 0, fmt.Errorf("student name cannot be empty")
	}
	query := fmt.Sprintf("INSERT INTO %s (name,roll,student_id,mother_name,father_name,guardian_name) VALUES (?,?,?,?,?,?)", tableName)
	result, err := DB.Exec(query, name, roll, student_id, mother_name, father_name, guardian_name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func UpdateStudent12(tableName string, id int64, name string, roll int64, student_id string, mother_name string, father_name string, guardian_name string, term1atcm int64, term1atco int64, term1aips int64, term1aimpc int64, term2atcm int64, term2atco int64, term2aips int64, term2aimpc int64, term3atcm int64, term3atco int64, term3aips int64, term3aimpc int64) error {
	// Validate marks against constants
	if term1atcm < 0 || term1atcm > Term1atcm {
		return fmt.Errorf("term1atcm must be between 0 and %d", Term1atcm)
	}
	if term1atco < 0 || term1atco > Term1atco {
		return fmt.Errorf("term1atco must be between 0 and %d", Term1atco)
	}
	if term1aips < 0 || term1aips > Term1aips {
		return fmt.Errorf("term1aips must be between 0 and %d", Term1aips)
	}
	if term1aimpc < 0 || term1aimpc > Term1aimpc {
		return fmt.Errorf("term1aimpc must be between 0 and %d", Term1aimpc)
	}

	if term2atcm < 0 || term2atcm > Term2atcm {
		return fmt.Errorf("term2atcm must be between 0 and %d", Term2atcm)
	}
	if term2atco < 0 || term2atco > Term2atco {
		return fmt.Errorf("term2atco must be between 0 and %d", Term2atco)
	}
	if term2aips < 0 || term2aips > Term2aips {
		return fmt.Errorf("term2aips must be between 0 and %d", Term2aips)
	}
	if term2aimpc < 0 || term2aimpc > Term2aimpc {
		return fmt.Errorf("term2aimpc must be between 0 and %d", Term2aimpc)
	}

	if term3atcm < 0 || term3atcm > Term3atcm {
		return fmt.Errorf("term3atcm must be between 0 and %d", Term3atcm)
	}
	if term3atco < 0 || term3atco > Term3atco {
		return fmt.Errorf("term3atco must be between 0 and %d", Term3atco)
	}
	if term3aips < 0 || term3aips > Term3aips {
		return fmt.Errorf("term3aips must be between 0 and %d", Term3aips)
	}
	if term3aimpc < 0 || term3aimpc > Term3aimpc {
		return fmt.Errorf("term3aimpc must be between 0 and %d", Term3aimpc)
	}

	query := fmt.Sprintf("UPDATE %s SET name = ?, roll = ?, student_id = ?, mother_name = ?, father_name = ?, guardian_name = ?, term1atcm = ?, term1atco = ?, term1aips = ?, term1aimpc = ?, term2atcm = ?, term2atco = ?, term2aips = ?, term2aimpc = ?, term3atcm = ?, term3atco = ?, term3aips = ?, term3aimpc = ? WHERE id = ?", tableName)
	_, err := DB.Exec(query, name, roll, student_id, mother_name, father_name, guardian_name, term1atcm, term1atco, term1aips, term1aimpc, term2atcm, term2atco, term2aips, term2aimpc, term3atcm, term3atco, term3aips, term3aimpc, id)
	return err
}

func UpdateStudent345(tableName string, id int64, name string, roll int64, student_id string, mother_name string, father_name string, guardian_name string, term1fl int64, term1sl int64, term1m int64, term1e int64, term1hpeth int64, term1hpepr int64, term1aweth int64, term1awepr int64, term2fl int64, term2sl int64, term2m int64, term2e int64, term2hpeth int64, term2hpepr int64, term2aweth int64, term2awepr int64, term3fl int64, term3sl int64, term3m int64, term3e int64, term3hpeth int64, term3hpepr int64, term3aweth int64, term3awepr int64) error {
	// Validate Term 1 marks
	if term1fl < 0 || term1fl > Term1fl {
		return fmt.Errorf("term1fl must be between 0 and %d", Term1fl)
	}
	if term1sl < 0 || term1sl > Term1sl {
		return fmt.Errorf("term1sl must be between 0 and %d", Term1sl)
	}
	if term1m < 0 || term1m > Term1M {
		return fmt.Errorf("term1m must be between 0 and %d", Term1M)
	}
	if term1e < 0 || term1e > Term1E {
		return fmt.Errorf("term1e must be between 0 and %d", Term1E)
	}
	if term1hpeth < 0 || term1hpeth > Term1hpeth {
		return fmt.Errorf("term1hpeth must be between 0 and %d", Term1hpeth)
	}
	if term1hpepr < 0 || term1hpepr > Term1hpepr {
		return fmt.Errorf("term1hpepr must be between 0 and %d", Term1hpepr)
	}
	if term1aweth < 0 || term1aweth > Term1aweth {
		return fmt.Errorf("term1aweth must be between 0 and %d", Term1aweth)
	}
	if term1awepr < 0 || term1awepr > Term1awepr {
		return fmt.Errorf("term1awepr must be between 0 and %d", Term1awepr)
	}

	// Validate Term 2 marks
	if term2fl < 0 || term2fl > Term2fl {
		return fmt.Errorf("term2fl must be between 0 and %d", Term2fl)
	}
	if term2sl < 0 || term2sl > Term2sl {
		return fmt.Errorf("term2sl must be between 0 and %d", Term2sl)
	}
	if term2m < 0 || term2m > Term2M {
		return fmt.Errorf("term2m must be between 0 and %d", Term2M)
	}
	if term2e < 0 || term2e > Term2E {
		return fmt.Errorf("term2e must be between 0 and %d", Term2E)
	}
	if term2hpeth < 0 || term2hpeth > Term2hpeth {
		return fmt.Errorf("term2hpeth must be between 0 and %d", Term2hpeth)
	}
	if term2hpepr < 0 || term2hpepr > Term2hpepr {
		return fmt.Errorf("term2hpepr must be between 0 and %d", Term2hpepr)
	}
	if term2aweth < 0 || term2aweth > Term2aweth {
		return fmt.Errorf("term2aweth must be between 0 and %d", Term2aweth)
	}
	if term2awepr < 0 || term2awepr > Term2awepr {
		return fmt.Errorf("term2awepr must be between 0 and %d", Term2awepr)
	}

	// Validate Term 3 marks
	if term3fl < 0 || term3fl > Term3fl {
		return fmt.Errorf("term3fl must be between 0 and %d", Term3fl)
	}
	if term3sl < 0 || term3sl > Term3sl {
		return fmt.Errorf("term3sl must be between 0 and %d", Term3sl)
	}
	if term3m < 0 || term3m > Term3M {
		return fmt.Errorf("term3m must be between 0 and %d", Term3M)
	}
	if term3e < 0 || term3e > Term3E {
		return fmt.Errorf("term3e must be between 0 and %d", Term3E)
	}
	if term3hpeth < 0 || term3hpeth > Term3hpeth {
		return fmt.Errorf("term3hpeth must be between 0 and %d", Term3hpeth)
	}
	if term3hpepr < 0 || term3hpepr > Term3hpepr {
		return fmt.Errorf("term3hpepr must be between 0 and %d", Term3hpepr)
	}
	if term3aweth < 0 || term3aweth > Term3aweth {
		return fmt.Errorf("term3aweth must be between 0 and %d", Term3aweth)
	}
	if term3awepr < 0 || term3awepr > Term3awepr {
		return fmt.Errorf("term3awepr must be between 0 and %d", Term3awepr)
	}

	query := fmt.Sprintf("UPDATE %s SET name = ?, roll = ?, student_id = ?, mother_name = ?, father_name = ?, guardian_name = ?, term1fl = ?, term1sl = ?, term1m = ?, term1e = ?, term1hpeth = ?, term1hpepr = ?, term1aweth = ?, term1awepr = ?, term2fl = ?, term2sl = ?, term2m = ?, term2e = ?, term2hpeth = ?, term2hpepr = ?, term2aweth = ?, term2awepr = ?, term3fl = ?, term3sl = ?, term3m = ?, term3e = ?, term3hpeth = ?, term3hpepr = ?, term3aweth = ?, term3awepr = ? WHERE id = ?", tableName)
	_, err := DB.Exec(query, name, roll, student_id, mother_name, father_name, guardian_name, term1fl, term1sl, term1m, term1e, term1hpeth, term1hpepr, term1aweth, term1awepr, term2fl, term2sl, term2m, term2e, term2hpeth, term2hpepr, term2aweth, term2awepr, term3fl, term3sl, term3m, term3e, term3hpeth, term3hpepr, term3aweth, term3awepr, id)
	return err
}

func DeleteStudent(tableName string, id int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", tableName)
	_, err := DB.Exec(query, id)
	return err
}

func GetStudent12(tableName string, id int64) (*Student12, error) {
	var student Student12
	query := fmt.Sprintf("SELECT id, name,roll,student_id,mother_name,father_name,guardian_name, created_at ,term1atcm,term1atco,term1aips,term1aimpc,term2atcm,term2atco,term2aips,term2aimpc,term3atcm,term3atco,term3aips,term3aimpc FROM %s WHERE id = ?", tableName)
	err := DB.Get(&student, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &student, err
}
func GetStudent345(tableName string, id int64) (*Student345, error) {
	var student Student345
	query := fmt.Sprintf("SELECT id, name,roll,student_id,mother_name,father_name,guardian_name, created_at ,term1fl,term1sl,term1m,term1e,term1hpeth,term1hpepr,term1aweth,term1awepr,term2fl,term2sl,term2m,term2e,term2hpeth,term2hpepr,term2aweth,term2awepr,term3fl,term3sl,term3m,term3e,term3hpeth,term3hpepr,term3aweth,term3awepr FROM %s WHERE id = ?", tableName)
	err := DB.Get(&student, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &student, err
}
