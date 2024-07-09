package habit

type Habit struct {
	Name      string
	Completed bool
	Count     int64
}

type User struct {
	Habits map[string]*Habit
}

var Users map[int64]*User
