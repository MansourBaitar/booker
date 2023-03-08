package db

import (
	"app/internal"
	"app/internal/model"
	"database/sql"
	"errors"
)

type HolidayRepo struct {
	db *sql.DB
}

func NewHolidayRepo(db *sql.DB) *HolidayRepo {
	return &HolidayRepo{db}
}

func (r *HolidayRepo) Save(holiday model.Holiday) (model.Holiday, error) {
	if err := r.db.Ping(); err != nil {
		return holiday, ErrNoConn
	}
	if holiday.Id == 0 {
		return r.createNewHoliday(holiday)
	} else {
		// TODO: perform UPDATE
		return holiday, errors.New("not implemented")
	}
}

func (r *HolidayRepo) FindById(id uint) (model.Holiday, error) {
	holiday := model.Holiday{}

	if err := r.db.Ping(); err != nil {
		return holiday, ErrNoConn
	}

	const query = `SELECT * FROM public.holiday h WHERE h.id = $1`
	row := r.db.QueryRow(query, id)
	err := row.Scan(&holiday.Id, &holiday.Type, &holiday.UserId, &holiday.Name, &holiday.Status, &holiday.FromDate, &holiday.ToDate, &holiday.Creation)
	return holiday, err
}

func (r *HolidayRepo) FindAll() ([]model.Holiday, error) {
	if err := r.db.Ping(); err != nil {
		return nil, ErrNoConn
	}

	const query = `SELECT * FROM public.holiday`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, ErrInvalidQuery
	}

	defer rows.Close()
	return holidayFromRows(rows)
}

func (r *HolidayRepo) DeleteHoliday(id uint) (int64, error) {
	if err := r.db.Ping(); err != nil {
		return 0, ErrNoConn
	}

	const query = `DELETE FROM holiday l WHERE l.id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return 0, ErrInvalidQuery
	}

	return result.RowsAffected()
}

func (r *HolidayRepo) createNewHoliday(holiday model.Holiday) (model.Holiday, error) {
	logger := internal.NewDefaultLogger()
	const query = "INSERT INTO public.holiday (type, userId, name, status, fromDate, toDate, creation) VALUES ($1, $2, $3, $4, $5, $6, $7)  RETURNING public.holiday.id"
	row := r.db.QueryRow(query, holiday.Type, holiday.Name, holiday.Name, holiday.Status, holiday.FromDate, holiday.ToDate, holiday.Creation)
	err := row.Scan(&holiday.Id)
	if err != nil {
		logger.Printf("Error occurred during save of holiday: %v", err)
		return holiday, err
	}
	logger.Printf("Successfully added new holiday to the database with id: %d", holiday.Id)
	return holiday, err
}

func holidayFromRows(r *sql.Rows) ([]model.Holiday, error) {
	holiday := make([]model.Holiday, 0)

	for r.Next() {
		o := model.Holiday{}
		err := r.Scan(&o.Id, &o.Type, &o.UserId, &o.Name,
			&o.Status, &o.FromDate, &o.ToDate, &o.Creation)
		if err != nil {
			return nil, ErrScanFault
		}

		holiday = append(holiday, o)
	}

	return holiday, nil
}
