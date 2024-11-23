package structs

import (
	"cinema/helpers/common"
	"errors"
	"fmt"
	"regexp"
	"time"
)

type Customer struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Fullname   string    `json:"fullname"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	ModifiedAt time.Time `gorm:"autoUpdateTime"`
}

type Booking struct {
	ID          int              `json:"id"`
	CustomerID  int              `json:"customer_id"`
	ShowtimeID  int              `json:"showtime_id"`
	BookingDate time.Time        `json:"booking_date"`
	SeatNumber  int              `json:"seat_number"`
	Status      string           `json:"status"`
	CreatedAt   time.Time        `gorm:"autoCreateTime"`
	ModifiedAt  time.Time        `gorm:"autoUpdateTime"`
	Customers   CustomerResponse `json:"customer" gorm:"foreignKey:CustomerID;references:ID"`
	Showtimes   Showtime         `json:"showtime" gorm:"foreignKey:ShowtimeID;references:ID"`
}

type Movie struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Genre       string     `json:"genre"`
	Duration    string     `json:"duration"`
	Rating      string     `json:"rating"`
	ReleaseDate time.Time  `json:"release_date"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	ModifiedAt  time.Time  `gorm:"autoUpdateTime"`
	Showtimes   []Showtime `json:"showtimes" gorm:"foreignKey:MovieID;references:ID"`
}

type Cinema_hall struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Location   string     `json:"location"`
	Capacity   int        `json:"capacity"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
	ModifiedAt time.Time  `gorm:"autoUpdateTime"`
	Showtimes  []Showtime `json:"showtimes" gorm:"foreignKey:CinemaHallID;references:ID"`
}

type Showtime struct {
	ID           int       `json:"id"`
	CinemaHallID int       `json:"cinema_hall_id"`
	MovieID      int       `json:"movie_id"`
	ShowtimeDate time.Time `json:"showtime_date"`
	Price        int       `json:"price"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	ModifiedAt   time.Time `gorm:"autoUpdateTime"`
}

type LoginRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SignUpRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	ReTypePassword string `json:"re_type_password"`
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
}

type CustomerResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (CustomerResponse) TableName() string {
	return "customers"
}

func (l *LoginRequest) ValidateLogin() (err error) {
	if common.IsEmptyField(l.Username) {
		return errors.New("username required")
	}

	if common.IsEmptyField(l.Password) {
		return errors.New("password required")
	}

	return
}

func (s *SignUpRequest) ValidateSignUp() (err error) {
	if common.IsEmptyField(s.Username) {
		return errors.New("username required")
	}

	if common.IsEmptyField(s.Password) {

		return errors.New("password required")
	}

	if common.IsEmptyField(s.ReTypePassword) {
		return errors.New("retype password required")
	}

	if common.IsEmptyField(s.Fullname) {
		return errors.New("fullname required")
	}

	if common.IsEmptyField(s.Email) {
		return errors.New("email required")
	}

	if common.IsEmptyField(s.Phone) {
		return errors.New("phone required")
	}

	if s.ReTypePassword != s.Password {
		return errors.New("password mismatch!")
	}

	re := regexp.MustCompile(fmt.Sprintf(`^(.{8,})$`))
	if !re.MatchString(s.Password) {
		return errors.New("please make sure that the password contains at least 8 character")
	}

	return nil
}

func (s *SignUpRequest) ConvertToModelForSignUp() (customer Customer, err error) {
	hashedPassword, err := common.HashPassword(s.Password)
	if err != nil {
		err = errors.New("hashing password failed")
		return
	}

	return Customer{
		Username: s.Username,
		Password: hashedPassword,
		Fullname: s.Fullname,
		Email:    s.Email,
		Phone:    s.Phone,
	}, nil
}

func (c *Customer) ValidateUpdate() (err error) {
	if common.IsEmptyField(c.Fullname) {
		return errors.New("fullname required")
	}

	if common.IsEmptyField(c.Email) {

		return errors.New("email required")
	}

	if common.IsEmptyField(c.Phone) {
		return errors.New("phone required")
	}

	return nil
}

func (c *Cinema_hall) ValidateCreate() (err error) {
	if common.IsEmptyField(c.Name) {
		return errors.New("Name required")
	}

	if common.IsEmptyField(c.Capacity) {
		return errors.New("Capacity required")
	}

	if common.IsEmptyField(c.Location) {
		return errors.New("Location required")
	}

	return
}

func (c *Cinema_hall) ValidateUpdate() (err error) {
	if common.IsEmptyField(c.Name) {
		return errors.New("Name required")
	}

	if common.IsEmptyField(c.Capacity) {

		return errors.New("Capacity required")
	}

	if common.IsEmptyField(c.Location) {
		return errors.New("Location required")
	}

	return nil
}

func (m *Movie) ValidateCreate() (err error) {
	if common.IsEmptyField(m.Title) {
		return errors.New("title required")
	}

	if common.IsEmptyField(m.Genre) {
		return errors.New("genre required")
	}

	if common.IsEmptyField(m.Duration) {
		return errors.New("Duration required")
	}

	if common.IsEmptyField(m.Rating) {
		return errors.New("rating required")
	}

	if common.IsEmptyField(m.ReleaseDate) {
		return errors.New("Release date required")
	}

	if common.IsEmptyField(m.Description) {
		return errors.New("Description required")
	}

	return
}

func (m *Movie) ValidateUpdate() (err error) {
	if common.IsEmptyField(m.Title) {
		return errors.New("title required")
	}

	if common.IsEmptyField(m.Genre) {
		return errors.New("genre required")
	}

	if common.IsEmptyField(m.Duration) {
		return errors.New("Duration required")
	}

	if common.IsEmptyField(m.Rating) {
		return errors.New("rating required")
	}

	if common.IsEmptyField(m.ReleaseDate) {
		return errors.New("Release date required")
	}

	if common.IsEmptyField(m.Description) {
		return errors.New("Description required")
	}

	return nil
}

func (s *Showtime) ValidateCreate() (err error) {
	if common.IsEmptyField(s.MovieID) {
		return errors.New("Movie Id required")
	}

	if common.IsEmptyField(s.CinemaHallID) {
		return errors.New("Cinema Hall required")
	}

	if common.IsEmptyField(s.ShowtimeDate) {
		return errors.New("Showtime Date required")
	}

	if common.IsEmptyField(s.Price) {
		return errors.New("Price required")
	}

	return
}

func (s *Showtime) ValidateUpdate() (err error) {
	if common.IsEmptyField(s.MovieID) {
		return errors.New("Movie Id required")
	}

	if common.IsEmptyField(s.CinemaHallID) {
		return errors.New("Cinema Hall required")
	}

	if common.IsEmptyField(s.ShowtimeDate) {
		return errors.New("Showtime Date required")
	}

	if common.IsEmptyField(s.Price) {
		return errors.New("Pice required")
	}

	return nil
}

func (b *Booking) ValidateCreate() (err error) {
	if common.IsEmptyField(b.CustomerID) {
		return errors.New("Customer Id required")
	}
	if common.IsEmptyField(b.ShowtimeID) {
		return errors.New("Showtime Id required")
	}
	if common.IsEmptyField(b.SeatNumber) {
		return errors.New("Seat Number required")
	}
	return
}

func (b *Booking) ValidateUpdate() (err error) {
	if common.IsEmptyField(b.ShowtimeID) {
		return errors.New("Showtime Id required")
	}
	if common.IsEmptyField(b.Status) {
		return errors.New("Status required")
	}

	return nil
}
