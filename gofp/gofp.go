// Generates functional code locally for user defined data type.
/*
```
1. Install "gofp" to generate code
   go get github.com/logic-building/functional-go/gofp
   go get -u github.com/logic-building/functional-go/gofp
   go install github.com/logic-building/functional-go/gofp

2. Add this line in a file where user defined data type exists
   //go:generate gofp -destination <file> -pkg <pkg> -type <Types separated by comma>

example:
    package employee

   //go:generate gofp -destination fp.go -pkg employee -type "Employee, Teacher"
   type Employee struct {
   	id     int
   	name   string
   	salary float64
   }

   type Teacher struct {
   	id     int
   	name   string
   	salary float64
   }

Note:
   A. fp.go               :  generated code
   B. employee            :  package name
   C. "Employee, Teacher" :  User defined data types

3. Generate functional code
   go generate ./...

4. Now write your code

    emp1 := employee.Employee{1, "A", 1000}
   	emp2 := employee.Employee{1, "A", 1000}
   	emp3 := employee.Employee{1, "A", 1000}

   	empList := []employee.Employee{emp1, emp2, emp3}

   	newEmpList := employee.Map(incrementSalary, empList) //  Returns: [{1 A 1500} {1 A 1500} {1 A 1500}]

   func incrementSalary(emp employee.Employee) employee.Employee {
        emp.Salary = emp.Salary + 500
        return emp
   }

```
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/logic-building/functional-go/fp"
	template2 "github.com/logic-building/functional-go/internal/template"
	"github.com/logic-building/functional-go/internal/template/basic"
	"math/rand"
	"runtime"
	"time"
)

var (
	destination = flag.String("destination", "", "functional code for user defined data type")
	pkgName     = flag.String("pkg", "", "package name for generated files")
	types       = flag.String("type", "", "user defined type")
	imports     = flag.String("imports", "", "import statements for user defined types when structs are in different package")
)

func main() {
	isAlreadyRun := runWithin(time.Second * 15)
	defer func() {
		if !isAlreadyRun {
			fmt.Println("\n\t\t\"" + quoteForTheDay() + "\"\n")
		}
	}()

	flag.Parse()

	if *destination == "" || *pkgName == "" || *types == "" {
		fmt.Println("either of these fields : (destination, package, types) - is not provided")
		usage()
		return
	}

	if len(*destination) > 0 {
		if err := os.MkdirAll(filepath.Dir(*destination), os.ModePerm); err != nil {
			log.Fatalf("Unable to create destination directory: %v", err)
			usage()
		}
		f, err := os.Create(*destination)
		if err != nil {
			log.Fatalf("Failed opening destination file: %v", err)
		}
		generatedCode, err := generateFPCode(*pkgName, *types, *imports)
		if err != nil {
			usage()
			log.Fatalf("Failed code generation: %v", err)
		}

		generatedCodeIO, err := generateFPCodeIO(*pkgName, *types)
		if err != nil {
			usage()
			log.Fatalf("Failed code generation for different IO combination: %v", err)
		}

		generatedCodeII, err := generateFPCodeII(*pkgName, *types)
		if err != nil {
			usage()
			log.Fatalf("Failed code generation for different IO combination: %v", err)
		}

		f.Write([]byte(generatedCode + "\n" + generatedCodeIO + "\n" + generatedCodeII))
		defer f.Close()
	}

	if !isAlreadyRun {
		fmt.Println("Functional code generation is successful.")
	}

}

// When imports are passed, Remove 1st part of "." in <FOUTPUT_TYPE> and <FINPUT_TYPE>
func removeFirstPartOfDot(str string) string {
	if strings.Contains(str, ".") {
		return strings.Split(str, ".")[1]
	}
	return str
}

func generateFPCode(pkg, dataTypes, imports string) (string, error) {
	basicTypes := "int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, string, bool"
	conditionalType := ""
	types := strings.Split(dataTypes, ",")
	types = fp.DistinctStrIgnoreCase(types)

	template := "// Code generated by 'gofp'. DO NOT EDIT.\n"
	template += "package <PACKAGE>\n"
	template += "import \"sync\" \n"

	if imports != "" {
		importList := strings.Split(imports, ",")
		importList = fp.DistinctStr(importList)
		for _, v := range importList {
			template += fmt.Sprintf("import \"%s\" \n", strings.TrimSpace(v))
		}
	}

	for _, t := range types {

		if strings.TrimSpace(strings.ToLower(t)) != strings.ToLower(pkg) {
			conditionalType = strings.TrimSpace(t)
		}
		t = strings.TrimSpace(t)

		if strings.Contains(basicTypes, t) {
			continue
		}
		r := strings.NewReplacer("<PACKAGE>", pkg, "<TYPE>", t, "<CONDITIONAL_TYPE>", removeFirstPartOfDot(conditionalType))

		template = r.Replace(template)

		template += template2.Map()
		template = r.Replace(template)

		template += template2.Filter()
		template = r.Replace(template)

		template += template2.Remove()
		template = r.Replace(template)

		template += template2.Some()
		template = r.Replace(template)

		template += template2.Every()
		template = r.Replace(template)

		template += template2.DropWhile()
		template = r.Replace(template)

		template += template2.TakeWhile()
		template = r.Replace(template)

		template += template2.Pmap()
		template = r.Replace(template)

		template += template2.FilterMap()
		template = r.Replace(template)

		template += template2.Rest()
		template = r.Replace(template)

		template += template2.Reduce()
		template = r.Replace(template)

		template += template2.DropLast()
		template = r.Replace(template)
	}
	return template, nil
}

func generateFPCodeIO(pkg, dataTypes string) (string, error) {
	basicTypes := "int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, string, bool"
	template := ""
	types := strings.Split(dataTypes, ",")

	types = fp.DistinctStrIgnoreCase(types)

	// For different input output combination
	for _, inputType := range types {
		for _, outputType := range types {

			inputType = strings.TrimSpace(inputType)
			outputType = strings.TrimSpace(outputType)

			if inputType == outputType || (strings.Contains(basicTypes, inputType) && strings.Contains(basicTypes, outputType)) {
				continue
			}

			if strings.Contains(basicTypes, strings.ToLower(inputType)) {
				inputType = strings.ToLower(inputType)
			}

			if strings.Contains(basicTypes, strings.ToLower(outputType)) {
				outputType = strings.ToLower(outputType)
			}

			fInputType := strings.Title(inputType)
			fOutputType := strings.Title(outputType)

			if fInputType == "String" {
				fInputType = "Str"
			}

			if fOutputType == "String" {
				fOutputType = "Str"
			}

			r := strings.NewReplacer("<FINPUT_TYPE>", removeFirstPartOfDot(fInputType), "<FOUTPUT_TYPE>", removeFirstPartOfDot(fOutputType), "<INPUT_TYPE>", inputType, "<OUTPUT_TYPE>", outputType)
			template += basic.MapIO()
			template = r.Replace(template)

			template += basic.PMapIO()
			template = r.Replace(template)

			template += basic.FilterMapIO()
			template = r.Replace(template)
		}
	}

	return template, nil
}

func generateFPCodeII(pkg, dataTypes string) (string, error) {
	basicTypes := "int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, string, bool"
	template := ""
	types := strings.Split(dataTypes, ",")

	types = fp.DistinctStrIgnoreCase(types)

	// For different input output combination
	for _, inputType1 := range types {
		for _, inputType2 := range types {

			inputType1 = strings.TrimSpace(inputType1)
			inputType2 = strings.TrimSpace(inputType2)

			// Skip same basic data type
			if strings.Contains(basicTypes, inputType1) && strings.Contains(basicTypes, inputType2) {
				continue
			}

			if strings.Contains(basicTypes, strings.ToLower(inputType1)) {
				inputType1 = strings.ToLower(inputType1)
			}

			if strings.Contains(basicTypes, strings.ToLower(inputType2)) {
				inputType2 = strings.ToLower(inputType2)
			}

			fInputType1 := strings.Title(inputType1)
			fInputType2 := strings.Title(inputType2)

			if fInputType1 == "String" {
				fInputType1 = "Str"
			}

			if fInputType2 == "String" {
				fInputType2 = "Str"
			}

			// Standard function name for same user defined type
			if inputType1 == inputType2 && strings.ToLower(inputType1) == strings.ToLower(pkg) {
				fInputType1 = ""
				fInputType2 = ""
			}

			// Standard function name for same user defined type
			if inputType1 == inputType2 {
				fInputType2 = ""
			}

			r := strings.NewReplacer("<FINPUT_TYPE1>", removeFirstPartOfDot(fInputType1), "<FINPUT_TYPE2>", removeFirstPartOfDot(fInputType2), "<INPUT_TYPE1>", inputType1, "<INPUT_TYPE2>", inputType2)

			template += basic.Merge()
			template = r.Replace(template)

			template += basic.Zip()
			template = r.Replace(template)
		}
	}

	return template, nil
}

func usage() {
	fmt.Println("\nUsage:")
	fmt.Println("go:generate -destination fp.go -source employee.go -pkg Employee")
}
func quoteForTheDay() string {
	quotes := []string{
		"Time spent in love is never waste",
		"Enjoy every moment",
		"Wherever there is love, there is God",
		"The real way to loving is giving not demanding",
		"No one is greater or smaller than other. Everyone in this world is unique. Love everyone",
		"The person who has heart full of love, has always something to give",
		"Hell has three gates lust, anger, greed",
		"Be happy with nothing and you will be happy with everything",
		"Detachment is not that you should own nothing, but that nothing should own you",
		"Devotion has the power to burn down any karma",
		"Love is the greatest power on earth",
		"When you wish good for others, good things come back to you. This is the Law of Nature",
		"If you can win over your mind, you can win over the whole world",
		"Darkness cannot drive out darkness, only light can do that. Hate cannot drive out hate. only love cna do that",
		"Silence says so mcuh. Just listen",
		"The greatest gift human can give to himself and others are tolerance and forgiveness",
		"The practice of devotion involves replacing desires for the world with the desires for God",
		"The wealth of divine love is the only true wealth. Every other form of wealth simply enhances one's pride",
		"Speak only when you feel your words are better than the silence",
		"For our spiritual growth, negative people are often placed in our path, so we may learn selfless love, forgiveness & surrender to the will of God",
		"The happiest people are givers not takers",
		"Why do we close our eyes when we pray, cry, kiss or dream? Because the most beautiful things in life are not seen, but felt by the heart",
		"If you have to choose between being kind and being right choose being kind and you will always be right",
		"Silence & Smile are two powerful tools.Smile is the way to solve many problems & Silence is the way to avoid many problems",
		"Don't get upset with people and situations, because both are powerless without your reaction",
		"Most of the people are in lack of knowledge.Don't hate people.Love people and understand people are under influence of ignorance. Always do righteously.",
		"Every way and means that leads our mind to God is Devotion",
		"The Only Purpose of Our Human Life is to Attain God",
		"Meditation. Because some questions can't be answered by Google!",
		"This is the nature of existence - if you do the right things, the right things will happen to you",
		"Devotion is the only way to be liberated from material attachment. It is only then that we become free from lust, anger and greed",
		"I belong to no religion. My religion is love. Every heart is my temple",
		"Don't focus too much on the defects, be aware of them, but our endeavor should be towards positive",
		"To purify the mind, we must engage in devotion to the lord, When our mind is purified, out attitude and our experience will change towards the outer world",
		"The reason that we are in a state of suffering and we are enveloped in the darkness of material existence, is our forgetfulness of God",
		"If you can establish your relationship with God, that ultimate satisfaction that you have been searching for since innumerable lifetimes, will eventually be attained",
		"The Joy of the mind is the measure of its strength",
		"When you come to a point where you have no need to impress anybody, your freedom will begin",
		"Ritualistic worship, chanting and meditation are done with the body, voice and the mind: they excel each other in the ascending order",
		"Uttering the sacred word, either in a loud or low tone is preferable to chants in praise of the Supreme. Mental contemplation is superior to both",
		"When one learns to turn the mind away from material allurements and renounces the desires of the senses, such a person comes in touch with the inner bliss of the soul",
		"When we decide that God is ours and the whole world is His, then our consciousness transforms from seeking self-enjoyment to serving the Lord with everything that we have",
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	ind := r.Intn(len(quotes))

	return quotes[ind]
}

func runWithin(duration time.Duration) bool {

	runWithin := func(file string, duration time.Duration) bool {
		writeToFile := func(file string) {
			f, _ := os.Create(file)
			defer f.Close()
			f.WriteString("Functional go")
		}
		runWithin := true
		if f, err := os.Stat(file); err == nil {
			modificationTime := f.ModTime()

			currentTime := time.Now()
			diffSeconds := currentTime.Sub(modificationTime).Seconds()
			if diffSeconds > duration.Seconds() {
				runWithin = false
				writeToFile(file)
			}

		} else {
			writeToFile(file)
		}
		return runWithin
	}

	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}

		return runWithin(home+"\functional-go.txt", duration)
	}
	return runWithin("/tmp/functional-go.txt", duration)
}
