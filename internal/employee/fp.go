// Code generated by 'gofp'. DO NOT EDIT.
package employee 

func Map(f func(Employee) Employee, list []Employee) []Employee {
	if f == nil {
		return []Employee{}
	}
	newList := make([]Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
func Filter(f func(Employee) bool, list []Employee) []Employee {
	if f == nil {
		return []Employee{}
	}
	var newList []Employee
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
func Remove(f func(Employee) bool, list []Employee) []Employee {
	if f == nil {
		return []Employee{}
	}
	var newList []Employee
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
func Some(f func(Employee) bool, list []Employee) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}
func Every(f func(Employee) bool, list []Employee) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}
func MapTeacher(f func(Teacher) Teacher, list []Teacher) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	newList := make([]Teacher, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
func FilterTeacher(f func(Teacher) bool, list []Teacher) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	var newList []Teacher
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
func RemoveTeacher(f func(Teacher) bool, list []Teacher) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	var newList []Teacher
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
func SomeTeacher(f func(Teacher) bool, list []Teacher) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}
func EveryTeacher(f func(Teacher) bool, list []Teacher) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}