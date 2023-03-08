package db

import (
	"app/internal"
	"app/internal/model"
	"database/sql"
	"errors"
)

type AbsenceRepo struct {
	db *sql.DB
}

func NewAbsenceRepo(db *sql.DB) *AbsenceRepo {
	return &AbsenceRepo{db}
}

func (r *AbsenceRepo) Save(absence model.Absence) (model.Absence, error) {
	if err := r.db.Ping(); err != nil {
		return absence, ErrNoConn
	}

	if absence.Id == 0 {
		return r.createNewAbsence(absence)
	} else {
		// TODO: perform UPDATE
		return absence, errors.New("not implemented")
	}
}

func (r *AbsenceRepo) FindById(id uint) (model.Absence, error) {
	absence := model.Absence{}

	if err := r.db.Ping(); err != nil {
		return absence, ErrNoConn
	}

	const query = `SELECT * FROM public.absence h WHERE h.id = $1`
	row := r.db.QueryRow(query, id)
	err := row.Scan(&absence.Id, &absence.Type, &absence.UserId, &absence.Name, &absence.Status, &absence.FromDate, &absence.ToDate, &absence.Creation)
	return absence, err
}

func (r *AbsenceRepo) FindAll() ([]model.Absence, error) {
	if err := r.db.Ping(); err != nil {
		return nil, ErrNoConn
	}

	const query = `SELECT * FROM public.absence`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, ErrInvalidQuery
	}

	defer rows.Close()
	return absenceFromRows(rows)
}

func (r *AbsenceRepo) DeleteAbsence(id uint) (int64, error) {
	if err := r.db.Ping(); err != nil {
		return 0, ErrNoConn
	}

	const query = `DELETE FROM absence a WHERE a.id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return 0, ErrInvalidQuery
	}

	return result.RowsAffected()
}

func (r *AbsenceRepo) createNewAbsence(absence model.Absence) (model.Absence, error) {
	logger := internal.NewDefaultLogger()
	const query = "INSERT INTO public.absence (type, userId, name, status, fromDate, toDate, creation) VALUES ($1, $2, $3, $4, $5, $6, $7)  RETURNING public.absence.id"
	row := r.db.QueryRow(query, absence.Type, absence.UserId, absence.Name, absence.Status, absence.FromDate, absence.ToDate, absence.Creation)
	err := row.Scan(&absence.Id)
	if err != nil {
		logger.Printf("Error occurred during save of absence: %v", err)
		return absence, err
	}
	logger.Printf("Successfully added new absence to the database with id: %d", absence.Id)
	return absence, err
}

func absenceFromRows(r *sql.Rows) ([]model.Absence, error) {
	absence := make([]model.Absence, 0)

	for r.Next() {
		o := model.Absence{}
		err := r.Scan(&o.Id, &o.Type, &o.UserId, &o.Name,
			&o.Status, &o.FromDate, &o.ToDate, &o.Creation)
		if err != nil {
			return nil, ErrScanFault
		}

		absence = append(absence, o)
	}

	return absence, nil
}
