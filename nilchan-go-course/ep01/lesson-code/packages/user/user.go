package user

type User struct {
	name string
	age  int
}

func (u *User) NewUser(name string, age int) User {
	if name == "" {
		return User{}
	}

	if age <= 0 || age >= 150 {
		return User{}
	}

	return User{
		name: name,
		age:  age,
	}
}

func (u *User) ChangeName(name string) {
	if name != "" {
		u.name = name
	}
}

func (u *User) ChangeAge(age int) {
	if age > 0 && age < 150 {
		u.age = age
	}
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetAge() int {
	return u.age
}
