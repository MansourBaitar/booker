package db

import (
	"app/internal"
	"app/internal/model"
	"database/sql"
	"errors"
)

type ExpenseRepo struct {
	db *sql.DB
}

func NewExpenseRepo(db *sql.DB) *ExpenseRepo {
	return &ExpenseRepo{db}
}

func (r *ExpenseRepo) Save(expense model.Expense) (model.Expense, error) {
	if err := r.db.Ping(); err != nil {
		return expense, ErrNoConn
	}

	if expense.Id == 0 {
		return r.createNewExpense(expense)
	} else {
		// TODO: perform UPDATE
		return expense, errors.New("not implemented")
	}
}

func (r *ExpenseRepo) FindById(id uint) (model.Expense, error) {
	expense := model.Expense{}

	if err := r.db.Ping(); err != nil {
		return expense, ErrNoConn
	}

	const query = `SELECT * FROM expense e WHERE e.id = $1`
	row := r.db.QueryRow(query, id)
	err := row.Scan(&expense.Id, &expense.Type, &expense.UserId, &expense.Name, &expense.Amount, &expense.Status, &expense.Category, &expense.Group, &expense.Comment, &expense.Date, &expense.Creation)
	return expense, err
}

func (r *ExpenseRepo) FindAll() ([]model.Expense, error) {
	if err := r.db.Ping(); err != nil {
		return nil, ErrNoConn
	}

	const query = `SELECT * FROM public.expense`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, ErrInvalidQuery
	}

	defer rows.Close()
	return expensesFromRows(rows)
}

func (r *ExpenseRepo) DeleteExpense(id uint) (int64, error) {
	if err := r.db.Ping(); err != nil {
		return 0, ErrNoConn
	}

	const query = `DELETE FROM public.expense e WHERE e.id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return 0, ErrInvalidQuery
	}

	return result.RowsAffected()
}

func (r *ExpenseRepo) createNewExpense(expense model.Expense) (model.Expense, error) {
	logger := internal.NewDefaultLogger()
	const query = "INSERT INTO public.expense (type, userId, name, amount, status, category, group, comment, date, creation) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING public.expense.id"
	row := r.db.QueryRow(query, expense.Type, expense.UserId, expense.Name, expense.Amount, expense.Status, expense.Category, expense.Group, expense.Comment, expense.Date, expense.Creation)
	err := row.Scan(&expense.Id)
	if err != nil {
		logger.Printf("Error occurred during save of location: %v", err)
		return expense, err
	}
	logger.Printf("Successfully added new location to the database with id: %d", expense.Id)
	return expense, err
}

func expensesFromRows(r *sql.Rows) ([]model.Expense, error) {
	locations := make([]model.Expense, 0)

	for r.Next() {
		o := model.Expense{}
		err := r.Scan(&o.Id, &o.Type, &o.UserId, &o.Name, &o.Amount, &o.Status, &o.Category, &o.Group, &o.Comment, &o.Date, &o.Creation)
		if err != nil {
			return nil, ErrScanFault
		}
		locations = append(locations, o)
	}

	return locations, nil
}
