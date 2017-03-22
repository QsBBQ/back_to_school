package main
import "fmt"

// As a school administrator you need to create a school system for managing the school
// School
// school registration
// Departments
// majors
// sports
// graduation


type School struct {
  Name string
  Departments []Department
  Students []Student
}

func (s *School) Register() {

}

type Student struct {
  student_id int
}

type MusicDepartment struct {

}

type MusicStudent struct {
  Student
}

type EngineerStudent struct {
  Student
}

type CivilEngineeringStudent struct {
  Student
  EngineerStudent
}

func main() {
  school := School{}

}
