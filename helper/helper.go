package helper

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

const CSV_PATH = "database/todo.csv"

func HandleAdd() {
	if len(os.Args) == 3 {
		task := os.Args[2]
		file, err := os.OpenFile(CSV_PATH, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			log.Fatal(err)
		}
		csvRows := openFileRead()
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		t := time.Now()
		data := [][]string{
			{strconv.Itoa(len(csvRows)), task, t.Format("02-Jan-2006 15:04:05"), strconv.FormatBool(false)},
		}
		err = writer.WriteAll(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("New task added to csv file successfully.")
	} else {
		panic("Add command must have a task to add: retro-todo add <task>")
	}
}

func HandleList() {
	if len(os.Args) != 2 {
		log.Fatal("Allowed just list arg")
	}
	printTasks()
}

func HandleComplete() {
	if len(os.Args) != 3 {
		log.Fatal("Number of args must be 3! main.go complete <id>")
	}
	rows := openFileRead()
	fOut, err := os.Create("database/todo.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fOut.Close()
	w := csv.NewWriter(fOut)
	selected_id := os.Args[2]
	for _, row := range rows {
		id := row[0]
		if id == selected_id {
			new_row := row
			status, err := strconv.ParseBool(row[3])
			if err != nil {
				log.Fatal(err)
			}
			new_row[3] = strconv.FormatBool(!status)
			w.Write(new_row)
		} else {
			w.Write(row)
		}
	}
	w.Flush()
	fmt.Println("Csv update completed")
}

func HandleDelete() {
	if len(os.Args) != 3 {
		log.Fatal("Number of args must be 3! main.go complete <id>")
	}
	rows := openFileRead()
	fOut, err := os.Create("database/todo.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fOut.Close()
	w := csv.NewWriter(fOut)
	selected_id := os.Args[2]
	for _, row := range rows {
		id := row[0]
		if id == selected_id {
			fmt.Println("Task " + id + " deleted")
		} else {
			w.Write(row)
		}
	}
	w.Flush()
}

func openFileRead() [][]string {
	fileRead, err := os.Open(CSV_PATH)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(fileRead)
	reader.FieldsPerRecord = -1
	csvRows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return csvRows
}

func printTasks() {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	rows := openFileRead()
	for _, row := range rows {
		var formatRow string
		for _, col := range row {
			formatRow += (col + "\t" + "|" + "\t")
		}
		fmt.Fprintln(writer, formatRow)
	}
	writer.Flush()
}
