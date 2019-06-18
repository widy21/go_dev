package main

import (
	"testing"
	"time"
)

func Test_save(t *testing.T) {
	stu := &student{
		Name: "stu01",
		Age:  18,
	}
	err := stu.save()
	if err != nil {
		t.Fatalf("save student false, error: %v", err)
		return
	}
	t.Logf("save student success.")
}

func Test_save_and_load(t *testing.T) {
	stu := &student{
		Name: "stu01",
		Age:  18,
	}
	err := stu.save()
	if err != nil {
		t.Fatalf("save student false, error: %v", err)
		return
	}
	//t.Logf("save student success.")

	time.Sleep(5 * time.Second)
	var newStu student
	newStu.load()
	//t.Logf("new stu is: %v", newStu)
	if newStu.Name != stu.Name {
		t.Fatalf("name error, expect stu01")
	}
	if newStu.Age != stu.Age {
		t.Fatalf("age error, expect 18")
	}
	t.Logf("save&load student success.")

}
