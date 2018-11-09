package questionnaire

import (
    "github.com/kun-lun/digester/pkg/artifactgen"
    "github.com/kun-lun/digester/pkg/common"
    "github.com/kun-lun/digester/pkg/detector"
    "github.com/kun-lun/digester/pkg/vmgroupcalc"
    "bufio"
    "fmt"
    "log"
    "os"
    "reflect"
    "strconv"
    "strings"
)

var (
    scanner = bufio.NewScanner(os.Stdin)
)

func Run() artifactgen.Blueprint {
    fmt.Println("Project path?")
    scanner.Scan()
    path := scanner.Text()
    d, err := detector.New(path)
    if err != nil {
        log.Fatal(err)
    }

    possiblePackageManagers := d.DetectPackageManager()
    fmt.Printf("What is the package manager for the project?")
    for i, pm := range possiblePackageManagers {
        if i == 0 {
            fmt.Printf(" %s \\", strings.ToUpper(string(pm)))
        } else {
            fmt.Printf(" %s \\", string(pm))
        }
    }
    if len(possiblePackageManagers) > 0 {
        fmt.Printf(" other?\n")
    } else {
        fmt.Printf("\n")
    }

    scanner.Scan()
    inputPackageManager := scanner.Text()
    if inputPackageManager == "" {
        inputPackageManager = string(possiblePackageManagers[0])
    }
    d.ConfirmPackageManager(inputPackageManager)

    possibleFrameworks := d.DetectFramework()
    fmt.Printf("What is the framework of the project?")
    for i, fw := range possibleFrameworks {
        if i == 0 {
            fmt.Printf(" %s \\", strings.ToUpper(string(fw)))
        } else {
            fmt.Printf(" %s \\", string(fw))
        }
    }
    if len(possibleFrameworks) > 0 {
        fmt.Printf(" other?\n")
    } else {
        fmt.Printf(" NONE?\n")
    }

    scanner.Scan()
    inputFramework := scanner.Text()
    if inputFramework == "" {
        if len(possibleFrameworks) > 0 {
            inputFramework = string(possibleFrameworks[0])
        }
    }
    d.ConfirmFramework(inputFramework)

    d.DetectConfig()

    bp := d.ExposeKnownInfo()

    // Ask for the empty fields
    nisp := &bp.NonIaaS
    if nisp.ProgrammingLanguage == "" {
        fmt.Println("What's the programming language? Allowed values: {php}")
        scanner.Scan()
        input := scanner.Text()
        pl, err := common.ParseProgrammingLanguage(input)
        if err != nil {
            log.Fatal(err)
        }
        nisp.ProgrammingLanguage = pl
    }
    if len(nisp.Databases) > 0 {
        needExtraInfo := false
        fmt.Println("Here is the database(s) we found:" )
        for i, db := range nisp.Databases {
            fmt.Printf("No.%d: {\n", i+1)
            s := reflect.ValueOf(&db).Elem()
            for j := 0; j < s.NumField(); j++ {
                valField := s.Field(j)
                typeField := s.Type().Field(j)
                tag := typeField.Tag
                val := valField.Interface().(string)
                fmt.Printf("  %s: %s\n", tag.Get("name"), val)
                if (val == "") {
                    needExtraInfo = true
                }
            }
            fmt.Println("}")
        }
        if (needExtraInfo) {
            fmt.Println("Please help fill the blank field(s).")
            for i, _ := range nisp.Databases {
                db := &nisp.Databases[i]
                askForDbEmptyFields(i+1, db)
            }
            fmt.Println("Done.")
        }
    }

    if len(nisp.Databases) > 0 {
        fmt.Println("Do you use any more databases? How many? Allowed values: {n | n >= 0}")
    } else {
        fmt.Println("Do you use any databases? How many? Allowed values: {n | n >= 0}")
    }
    scanner.Scan()
    input := scanner.Text()
    extraDatabasesNum, err := strconv.Atoi(input)
    if err != nil {
		log.Fatal(err)
	}
    for i := 1; i <= extraDatabasesNum; i++ {
        newDb := common.Database{}
        askForDbEmptyFields(len(nisp.Databases) + 1, &newDb)
        nisp.Databases = append(nisp.Databases, newDb)
    }

    fmt.Println("What's your expected number of concurrent users?")
    scanner.Scan()
    concurrentUserNumberStr := scanner.Text()
    var concurrentUserNumber int
    concurrentUserNumber, err = strconv.Atoi(concurrentUserNumberStr)
    if err != nil {
        log.Fatal(err)
    }
    bp.IaaS.VMGroup = vmgroupcalc.Calc(vmgroupcalc.Requirment{
        ConcurrentUserNumber: concurrentUserNumber,
    })


    return bp
}

func askForDbEmptyFields(num int, db *common.Database) {
    s := reflect.ValueOf(db).Elem()
    for i := 0; i < s.NumField(); i++ {
        valField := s.Field(i)
        typeField := s.Type().Field(i)
        tag := typeField.Tag
        val := valField.Interface().(string)
        allowedSentence:= ""
        if tag.Get("allow") != "" {
            allowedSentence = fmt.Sprintf(" Allowed values: %s", tag.Get("allow"))
        }
        if (val == "") {
            fmt.Printf(
                "For the database No.%d: %s%s\n",
                num,
                tag.Get("question"),
                allowedSentence,
            )
            scanner.Scan()
            input := scanner.Text()
            if err := db.ValidateField(typeField.Name, input); err != nil {
                log.Fatal(err)
            }
            valField.SetString(input)
        }
    }
}
