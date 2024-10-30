// SMS/sms/model.go
package sms

import (
	"time"
)
const SchoolName string = "Name of the School"
const (
	Term1atcm  = 10
	Term1atco  = 10
	Term1aips  = 10
	Term1aimpc = 10
	Term2atcm  = 10
	Term2atco  = 10
	Term2aips  = 10
	Term2aimpc = 10
	Term3atcm  = 30
	Term3atco  = 30
	Term3aips  = 30
	Term3aimpc = 30
)
const (
	Term1fl    = 10
	Term1sl    = 10
	Term1M     = 10
	Term1E     = 10
	Term1hpeth = 0
	Term1hpepr = 5
	Term1aweth = 0
	Term1awepr = 5
	Term2fl    = 20
	Term2sl    = 20
	Term2M     = 20
	Term2E     = 20
	Term2hpeth = 0
	Term2hpepr = 10
	Term2aweth = 0
	Term2awepr = 10
	Term3fl    = 70
	Term3sl    = 70
	Term3M     = 70
	Term3E     = 70
	Term3hpeth = 15
	Term3hpepr = 10
	Term3aweth = 15
	Term3awepr = 10
)

type ClassInfo struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Year      int       `db:"year" json:"year"`
	TableName string    `db:"tablename" json:"TableName"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
type School struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
/* SQL for Student12 table:
CREATE TABLE student12 (
				id BIGINT PRIMARY KEY AUTO_INCREMENT,
				name VARCHAR(255) NOT NULL,
				roll BIGINT NOT NULL,
				student_id VARCHAR(50) NOT NULL,
				mother_name VARCHAR(255),
				father_name VARCHAR(255),
				guardian_name VARCHAR(255),
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				term1atcm BIGINT DEFAULT 0,
				term1atco BIGINT DEFAULT 0,
				term1aips BIGINT DEFAULT 0,
				term1aimpc BIGINT DEFAULT 0,
				term2atcm BIGINT DEFAULT 0,
				term2atco BIGINT DEFAULT 0,
				term2aips BIGINT DEFAULT 0,
				term2aimpc BIGINT DEFAULT 0,
				term3atcm BIGINT DEFAULT 0,
				term3atco BIGINT DEFAULT 0,
				term3aips BIGINT DEFAULT 0,
				term3aimpc BIGINT DEFAULT 0
);
*/

type Student12 struct {
	ID           int64     `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Roll         int64     `db:"roll" json:"roll"`
	StudentId    string    `db:"student_id" json:"studentId"`
	MotherName   string    `db:"mother_name" json:"motherName"`
	FatherName   string    `db:"father_name" json:"fatherName"`
	GuardianName string    `db:"guardian_name" json:"guardianName"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt"`

	Term1atcm  int64 `db:"term1atcm" json:"term1atcm"`
	Term1atco  int64 `db:"term1atco" json:"term1atco"`
	Term1aips  int64 `db:"term1aips" json:"term1aips"`
	Term1aimpc int64 `db:"term1aimpc" json:"term1aimpc"`
	Term2atcm  int64 `db:"term2atcm" json:"term2atcm"`
	Term2atco  int64 `db:"term2atco" json:"term2atco"`
	Term2aips  int64 `db:"term2aips" json:"term2aips"`
	Term2aimpc int64 `db:"term2aimpc" json:"term2aimpc"`
	Term3atcm  int64 `db:"term3atcm" json:"term3atcm"`
	Term3atco  int64 `db:"term3atco" json:"term3atco"`
	Term3aips  int64 `db:"term3aips" json:"term3aips"`
	Term3aimpc int64 `db:"term3aimpc" json:"term3aimpc"`
}

/* SQL for Student345 table:
CREATE TABLE student345 (
				id BIGINT PRIMARY KEY AUTO_INCREMENT,
				name VARCHAR(255) NOT NULL,
				roll BIGINT NOT NULL,
				student_id VARCHAR(50) NOT NULL,
				mother_name VARCHAR(255),
				father_name VARCHAR(255),
				guardian_name VARCHAR(255),
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				term1fl BIGINT DEFAULT 0,
				term1sl BIGINT DEFAULT 0,
				term1m BIGINT DEFAULT 0,
				term1e BIGINT DEFAULT 0,
				term1hpeth BIGINT DEFAULT 0,
				term1hpepr BIGINT DEFAULT 0,
				term1aweth BIGINT DEFAULT 0,
				term1awepr BIGINT DEFAULT 0,
				term2fl BIGINT DEFAULT 0,
				term2sl BIGINT DEFAULT 0,
				term2m BIGINT DEFAULT 0,
				term2e BIGINT DEFAULT 0,
				term2hpeth BIGINT DEFAULT 0,
				term2hpepr BIGINT DEFAULT 0,
				term2aweth BIGINT DEFAULT 0,
				term2awepr BIGINT DEFAULT 0,
				term3fl BIGINT DEFAULT 0,
				term3sl BIGINT DEFAULT 0,
				term3m BIGINT DEFAULT 0,
				term3e BIGINT DEFAULT 0,
				term3hpeth BIGINT DEFAULT 0,
				term3hpepr BIGINT DEFAULT 0,
				term3aweth BIGINT DEFAULT 0,
				term3awepr BIGINT DEFAULT 0
);
*/

type Student345 struct {
	ID           int64     `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Roll         int64     `db:"roll" json:"roll"`
	StudentId    string    `db:"student_id" json:"studentId"`
	MotherName   string    `db:"mother_name" json:"motherName"`
	FatherName   string    `db:"father_name" json:"fatherName"`
	GuardianName string    `db:"guardian_name" json:"guardianName"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt"`

	Term1fl    int64 `db:"term1fl" json:"term1fl"`
	Term1sl    int64 `db:"term1sl" json:"term1sl"`
	Term1M     int64 `db:"term1m" json:"term1m"`
	Term1E     int64 `db:"term1e" json:"term1e"`
	Term1hpeth int64 `db:"term1hpeth" json:"term1hpeth"`
	Term1hpepr int64 `db:"term1hpepr" json:"term1hpepr"`
	Term1aweth int64 `db:"term1aweth" json:"term1aweth"`
	Term1awepr int64 `db:"term1awepr" json:"term1awepr"`
	Term2fl    int64 `db:"term2fl" json:"term2fl"`
	Term2sl    int64 `db:"term2sl" json:"term2sl"`
	Term2M     int64 `db:"term2m" json:"term2m"`
	Term2E     int64 `db:"term2e" json:"term2e"`
	Term2hpeth int64 `db:"term2hpeth" json:"term2hpeth"`
	Term2hpepr int64 `db:"term2hpepr" json:"term2hpepr"`
	Term2aweth int64 `db:"term2aweth" json:"term2aweth"`
	Term2awepr int64 `db:"term2awepr" json:"term2awepr"`
	Term3fl    int64 `db:"term3fl" json:"term3fl"`
	Term3sl    int64 `db:"term3sl" json:"term3sl"`
	Term3M     int64 `db:"term3m" json:"term3m"`
	Term3E     int64 `db:"term3e" json:"term3e"`
	Term3hpeth int64 `db:"term3hpeth" json:"term3hpeth"`
	Term3hpepr int64 `db:"term3hpepr" json:"term3hpepr"`
	Term3aweth int64 `db:"term3aweth" json:"term3aweth"`
	Term3awepr int64 `db:"term3awepr" json:"term3awepr"`
}
