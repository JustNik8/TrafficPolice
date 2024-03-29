package repository

import (
	"TrafficPolice/internal/domain"
	"TrafficPolice/internal/errs"
	"TrafficPolice/internal/repository"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"time"
)

type caseRepoPostgres struct {
	conn *pgx.Conn
}

func NewCaseRepoPostgres(conn *pgx.Conn) repository.CaseRepo {
	return &caseRepoPostgres{conn: conn}
}

const insertCaseQuery = `INSERT INTO cases (case_id, transport_id, camera_id, 
                   violation_id, violation_value, required_skill, case_date, is_solved, fine_decision, solved_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, false, false, NULL) RETURNING case_id`

func (r *caseRepoPostgres) InsertCase(c domain.Case) (string, error) {
	var caseID string

	err := r.conn.QueryRow(context.Background(), insertCaseQuery,
		c.ID, c.Transport.ID, c.Camera.ID, c.Violation.ID, c.ViolationValue, c.RequiredSkill, c.Date,
	).Scan(&caseID)

	if err != nil {
		return "", errs.ErrInvalidRelevantParams
	}

	return caseID, err
}

const getCaseByIDQuery = `SELECT c.case_id, t.transport_id, t.transport_chars, 
       t.transport_nums, t.region, t.person_id, cam.camera_type_id,
       cam.camera_latitude, cam.camera_longitude, cam.short_desc, v.violation_name, v.fine_amount,
       c.violation_value, c.required_skill, c.case_date,
       c.is_solved, c.fine_decision
FROM cases as c
JOIN transports AS t ON c.transport_id = t.transport_id
JOIN violations AS v ON c.violation_id = v.violation_id
JOIN cameras AS cam ON c.camera_id = cam.camera_id
WHERE c.case_id = $1
LIMIT 1`

func (r *caseRepoPostgres) GetCaseByID(caseID string) (domain.Case, error) {
	c := domain.Case{Transport: domain.Transport{Person: &domain.Person{}}, Camera: domain.Camera{}, Violation: domain.Violation{}}

	row := r.conn.QueryRow(context.Background(), getCaseByIDQuery, caseID)

	err := row.Scan(&c.ID, &c.Transport.ID, &c.Transport.Chars, &c.Transport.Num, &c.Transport.Region,
		&c.Transport.Person.ID, &c.Camera.CameraType.ID, &c.Camera.Latitude, &c.Camera.Longitude,
		&c.Camera.ShortDesc, &c.Violation.Name, &c.Violation.FineAmount, &c.ViolationValue,
		&c.RequiredSkill, &c.Date, &c.IsSolved, &c.FineDecision)

	if errors.Is(err, pgx.ErrNoRows) {
		return domain.Case{}, errs.ErrNoCase
	}
	if err != nil {
		return domain.Case{}, err
	}

	return c, nil
}

const setCaseFineDecisionQuery = `UPDATE cases
SET fine_decision = $1, is_solved = true, solved_at = $2
WHERE case_id = $3`

func (r *caseRepoPostgres) SetCaseFineDecision(caseID string, fineDecision bool, solvedAt time.Time) error {
	_, err := r.conn.Exec(context.Background(), setCaseFineDecisionQuery, fineDecision, solvedAt, caseID)
	return err
}

const updateCaseRequiredSkillQuery = `UPDATE cases
SET required_skill = $1
WHERE case_id = $2`

func (r *caseRepoPostgres) UpdateCaseRequiredSkill(caseID string, requiredSkill int) error {
	_, err := r.conn.Exec(context.Background(), updateCaseRequiredSkillQuery, requiredSkill, caseID)
	return err
}

const getCaseWithPersonInfoQuery = `SELECT c.case_id, t.transport_id, t.transport_chars, 
       t.transport_nums, t.region, p.id, p.phone_num, p.email, p.vk_id, p.tg_id, 
       cam.camera_id ,cam.camera_type_id, cam.camera_latitude, cam.camera_longitude, cam.short_desc, 
       v.violation_id, v.violation_name, v.fine_amount, 
       c.violation_value, c.required_skill, c.case_date, c.is_solved, c.fine_decision
FROM cases as c
JOIN transports AS t ON c.transport_id = t.transport_id
JOIN persons AS p ON t.person_id = p.id
JOIN violations AS v ON c.violation_id = v.violation_id
JOIN cameras AS cam ON c.camera_id = cam.camera_id
WHERE c.case_id = $1
LIMIT 1`

func (r *caseRepoPostgres) GetCaseWithPersonInfo(caseID string) (domain.Case, error) {
	c := domain.Case{Transport: domain.Transport{Person: &domain.Person{}}, Camera: domain.Camera{}, Violation: domain.Violation{}}

	row := r.conn.QueryRow(context.Background(), getCaseWithPersonInfoQuery, caseID)

	err := row.Scan(&c.ID, &c.Transport.ID, &c.Transport.Chars, &c.Transport.Num,
		&c.Transport.Region, &c.Transport.Person.ID, &c.Transport.Person.PhoneNum,
		&c.Transport.Person.Email, &c.Transport.Person.VkID, &c.Transport.Person.TgID,
		&c.Camera.ID, &c.Camera.CameraType.ID, &c.Camera.Latitude, &c.Camera.Longitude,
		&c.Camera.ShortDesc, &c.Violation.ID, &c.Violation.Name, &c.Violation.FineAmount,
		&c.ViolationValue, &c.RequiredSkill, &c.Date, &c.IsSolved, &c.FineDecision)

	if err != nil {
		return domain.Case{}, err
	}
	return c, nil
}
