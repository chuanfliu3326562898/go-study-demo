/*
 * @author: onno
 * @date: 2020/4/27
 */
package user

type People struct {
    name     string
    sex      bool
    birthday int64
}

type Student struct {
    school string
    *People
    address Address
}

type Address struct {
    province string
    city     string
}

func NewPeople1(name string, sex bool, birthday int64) People {
    p := People{name, sex, birthday}
    return p
}

func NewPeople2(name string, sex bool, birthday int64) *People {
    p := new(People)
    p.name = name
    p.sex = sex
    p.birthday = birthday
    return p
}

func NewPeople3(name string, sex bool, birthday int64) *People {
    p := &People{name: name, sex: sex, birthday: birthday}
    return p
}

func newStudent1(name string, sex bool, birthday int64, school string, province string, city string) *Student {
    s := &Student{
        school,
        NewPeople2(name, sex, birthday),
        Address{province, city},
    }
    return s
}

func (s Student) GetSex() bool {
    return s.sex
}

func (s Student) GetProvince() string {
    return s.address.province;
}

func (s *Student) GetName() string {
    return s.name
}
