package moment_test

import (
	"fmt"
	"time"

	"github.com/moment/moment"
	_ "github.com/moment/moment/locale"
)

func Example() {
	time.Local = time.UTC
	m, err := moment.New("2010-10-20 4:30", "YYYY-MM-DD HH:mm")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m.Format("LLLL"))
	fmt.Println(m.Clone().SetLocale("fr").Format("LLLL"))
	fmt.Println(moment.CreateDuration(90.0, "m").Humanize(true, nil))
	// Output:
	// Wednesday, October 20, 2010 4:30 AM
	// mercredi 20 octobre 2010 04:30
	// in 2 hours
}
