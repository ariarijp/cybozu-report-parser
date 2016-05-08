package cybozureport

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Schedule struct {
	StartTime time.Time
	EndTime   time.Time
	Title     string
}

func Parse(filename string) []Schedule {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	r1 := regexp.MustCompile(`^(\d+) 月 (\d+) 日`)
	r2 := regexp.MustCompile(`^(\d+):(\d+)-(\d+):(\d+) (.+)`)
	now := time.Now()

	var dateStr string
	var schedules []Schedule

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		} else if line == "-----" {
			break
		}

		grp := r1.FindStringSubmatch(string(line))
		if len(grp) == 3 {
			month, _ := strconv.Atoi(grp[1])
			day, _ := strconv.Atoi(grp[2])
			dateStr = fmt.Sprintf("%04d-%02d-%02d", now.Year(), month, day)
			continue
		}

		grp = r2.FindStringSubmatch(string(line))
		if len(grp) == 6 {
			sHour, _ := strconv.Atoi(grp[1])
			sMinute, _ := strconv.Atoi(grp[2])
			eHour, _ := strconv.Atoi(grp[3])
			eMinute, _ := strconv.Atoi(grp[4])

			var s Schedule
			s.Title = grp[5]
			s.StartTime, _ = time.Parse("2006-01-02T15:04:05 MST", fmt.Sprintf("%sT%02d:%02d:00 JST", dateStr, sHour, sMinute))
			s.EndTime, _ = time.Parse("2006-01-02T15:04:05 MST", fmt.Sprintf("%sT%02d:%02d:00 JST", dateStr, eHour, eMinute))

			schedules = append(schedules, s)
		} else {
			log.Fatal(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return schedules
}
