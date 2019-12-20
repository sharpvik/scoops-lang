package assembly

import (
  //  "fmt"
    "bufio"
    "io"
    "os"
    "regexp"
)

func Read(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    emptyLine := regexp.MustCompile(`^\s*\n$`)
    comment := regexp.MustCompile(`^\s*#.*\n$`)
    lastLine := regexp.MustCompile(`^\s*$`)

    rdr := bufio.NewReader(file)
    var assemblyСode []string
    line, err := rdr.ReadString('\n')

    for err == nil {
        // skip empty lines and comments
        if !emptyLine.MatchString(line) && !comment.MatchString(line) {
            assemblyСode = append(assemblyСode, line)
        }

        line, err = rdr.ReadString('\n')
    }

    if err != io.EOF {
        return nil, err
    }

    if !lastLine.MatchString(line) && !comment.MatchString(line) {
        assemblyСode = append(assemblyСode, line)
    }

    return assemblyСode, nil
}
