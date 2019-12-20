package main

import "testing"

func TestTypeOfArg(t *testing.T) {
    cases := []string{
        // valid flags
        "-a", "-c", "-e",
        // invalid flags
        "-ace", "-w", "-www",

        // valid filenames
        "a.scp", "a.scpa", "a.scpb", "ab.scp", "ab.scpa", "ab.scpb",
        // invalid filenames
        ".scp", ".scpa", ".scpb", "a.", "ab.", "a", "ab", "a.a", "ab.ab",
        "a.scpw", "a.scpaa", "a.scpbb",
    }
    answers := []interface{}{
        // flags
        "flag", "flag", "flag",
        "", "", "",

        // filenames
        "filename", "filename", "filename", "filename", "filename", "filename",
        "", "", "", "", "", "", "", "", "",
        "", "", "",
    }
    for i, c := range cases {
        argType, _ := TypeOfArg(c)
        if argType != answers[i] {
            t.Errorf("For '%s' received '%s' but expected '%s'.",
                     c, argType, answers[i])
        }
    }
}


func TestParseArgs(t *testing.T) {
    cases := [][]string{
        // valid cases
        {"-a", "a.scp"},
        {"-a", "a.scpa"},
        {"-a", "a.scpb"},
        {"-c", "a.scp"},
        {"-c", "a.scpa"},
        {"-c", "a.scpb"},
        {"-e", "a.scp"},
        {"-e", "a.scpa"},
        {"-e", "a.scpb"},
        {"a.scp"},
        {"a.scpa"},
        {"a.scpb"},

        // invalid cases
        {},
        {"-a"},
        {"-c"},
        {"-e"},
        {"-a", "-c"},
        {"-a", "-e"},
        {"-a", "a.scp", "wrong"},
        {"-c", "a.scpw"},
        {"a.scpw"},
        {""},
        {"", ""},
        {"-f", "a.scpw"},
    }
    answers := []struct{
        flag byte
        filename string
    }{
        // valid cases
        {'a', "a.scp"},
        {'a', "a.scpa"},
        {'a', "a.scpb"},
        {'c', "a.scp"},
        {'c', "a.scpa"},
        {'c', "a.scpb"},
        {'e', "a.scp"},
        {'e', "a.scpa"},
        {'e', "a.scpb"},
        {0, "a.scp"},
        {0, "a.scpa"},
        {0, "a.scpb"},

        // invalid cases
        {0, ""},
        {0, ""},
        {0, ""},
        {0, ""},
        {0, ""},
        {0, ""},
        {0, ""},
        {0, ""},
        {0, ""},
        {0, ""},
        {0, ""},
        {0, ""},
    }
    for i, c := range cases {
        flag, filename, _ := ParseArgs(c)
        eflag, efilename := answers[i].flag, answers[i].filename
        if flag != eflag || filename != efilename {
            t.Errorf("For '%s %s' received '%c, %s'. Expected: '%c, %s'.",
                     c[0], c[1], flag, filename, eflag, efilename)
        }
    }
}


func TestReverseString(t *testing.T) {
    cases := "Hello, World!"
    answer := "!dlroW ,olleH"
    reply := ReverseString(cases) 
    if reply != answer {
        t.Errorf(
            "Cannot properly reverse string. Got %s for %s", reply, cases,
        )
    }
}


func TestGetFileExtention(t *testing.T) {
    cases := []string{
        "a.scp", "a.scpa", "a.scpb", "a.exe", "a",
    }
    answers := []string{
        "scp", "scpa", "scpb", "exe", "",
    }
    for i, c := range cases {
        extention := GetFileExtention(c)
        if extention != answers[i] {
            t.Errorf("For '%s' received '%s' byt expected '%s'.",
                     c, extention, answers[i])
        }
    }
}
