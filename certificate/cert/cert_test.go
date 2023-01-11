package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2023-01-08")
	if err != nil {
		t.Errorf("Cert data should be valid, err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference, got=nil")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid, expected 'GOLANG COURSE', got=%v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2023-01-08")
	if err == nil {
		t.Error("Error must return an empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "azertyazertyazertyazertyazertyazertyazertyazertyazertyazertyazertyazertyazertyazertyazertyazertyazertyazertyazertyazertyazerty"
	_, err := New(course, "Bob", "2023-01-08")
	if err == nil {
		t.Errorf("Must return an error, string is too long, course=%s", course)
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "", "2023-01-11")
	if err == nil {
		t.Error("Name can't be empty")
	}
}

func TestNameTooLong(t *testing.T) {
	name := "Bobobobobo bobob"
	_, err := New("Golang", name, "2023-01-11")
	if err == nil {
		t.Errorf("Name of student is too long, given='%v', len=%d", name, len(name))
	}
}
