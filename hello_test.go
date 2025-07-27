package main

import (
	"testing"
)

func TestEnglish(test *testing.T) {
	name := "Bane"
	language := "english"
	expected := "Hello " + name
	actual,err := Hello(name , language)
	if err != nil {
		test.Errorf("!!! Should not produce an error: %s" , err)
	}
	if expected != actual {
		test.Errorf("Result was incorrect , got: %s ,want: %s " , actual , expected)
	}
}

func TestNorway(test *testing.T) {
	name := "Kirill"
	language := "norway"
	expected := "Hei " + name
	actual , err := Hello(name , language)
	if err != nil {
		test.Errorf("!!! Should not produce an error: %s" , err)
	}
	if expected != actual {
		test.Errorf("Result was incorrect , got: %s ,want: %s " , actual , expected)
	}
}

func TestItalian(test *testing.T) {
	language := "italian"
	expected := "Ciao Marco"
	actual , err := Hello("" , language)
	if err != nil {
			test.Errorf("!!! Should not produce an error: %s" , err)
	}
	if expected != actual {
		test.Errorf("Result was incorrect , got: %s ,want: %s " , actual , expected)
	}
}

/* func TestIndividual(test *testing.T) {
	oldStdin := os.Stdin
	defer func() {os.Stdin = oldStdin}()
	r,w , _ := os.Pipe()
	go func() {
        defer w.Close()
        _, err := w.WriteString("Kirill\nenglish\n")
        if err != nil {
            test.Errorf("Write failed: %v", err)
        }
    }()
	os.Stdin = r
	time.Sleep(100 * time.Millisecond)
	expected := "Hello Kirill"
	actual ,err := Hello("","")
	if err != nil {
		test.Errorf("!!! Should not produce an error: %s" , err)
	}
	if expected != actual {
		test.Errorf("Result was incorrect , got: %s ,want: %s " , actual , expected)
	}
} */