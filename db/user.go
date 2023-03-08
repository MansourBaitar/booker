package db

import (
	"app/internal"
	"app/internal/model"
	"database/sql"
	"errors"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (i *UserRepo) Save(user model.User) (model.User, error) {
	if err := i.db.Ping(); err != nil {
		return user, ErrNoConn
	}

	if user.Id == 0 {
		return i.createNewUser(user)
	} else {
		// TODO: perform UPDATE
		return user, errors.New("not implemented")
	}
}

func (i *UserRepo) FindAll() ([]model.User, error) {
	if err := i.db.Ping(); err != nil {
		return nil, ErrNoConn
	}

	const query = `SELECT * FROM public.user`
	rows, err := i.db.Query(query)
	if err != nil {
		return nil, ErrInvalidQuery
	}

	defer rows.Close()
	return usersFromRows(rows)
}

func (r *UserRepo) GetUserByEmail(email string, pwd string) (model.User, error) {
	user := model.User{}
	if err := r.db.Ping(); err != nil {
		return user, ErrNoConn
	}

	const query = `SELECT * FROM public.user p WHERE p.email = $1 AND p.password = $2`
	row := r.db.QueryRow(query, email, pwd)
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.BirthDate, &user.CreatedAt)

	return user, err
}

func (i *UserRepo) createNewUser(user model.User) (model.User, error) {
	logger := internal.NewDefaultLogger()
	const query = "INSERT INTO public.user (first_name, last_name, email,password, birth_date, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING public.user.id"
	row := i.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password, user.BirthDate, user.CreatedAt)
	err := row.Scan(&user.Id)
	if err != nil {
		logger.Printf("error occurred during save of user: %v", err)
	}
	logger.Printf("Successfully added new user to the database with id: %d", user.Id)
	return user, err
}

func usersFromRows(r *sql.Rows) ([]model.User, error) {
	users := make([]model.User, 0)

	for r.Next() {
		u := model.User{}
		err := r.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email,
			&u.BirthDate, &u.CreatedAt)
		if err != nil {
			return nil, ErrScanFault
		}

		users = append(users, u)
	}

	return users, nil
}
