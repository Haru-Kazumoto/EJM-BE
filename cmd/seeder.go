package cmd

import (
	"EJM/config"
	db "EJM/init"
	"EJM/scripts"
	"EJM/utils"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

type Cases interface {
	Execute() error
	GetName() string
}

type SeederStatus struct {
	Index        int
	WorkerIndex  int
	Case1        Cases
	ErrorMessage error
}

func SeederJobs(args []string, cases ...Cases) <-chan SeederStatus {
	chanOut := make(chan SeederStatus)

	go func() {
		for i, case1 := range cases {
			if args[0] != utils.ALL {
				for _, arg := range args {
					if arg == reflect.TypeOf(case1).Elem().Name() {
						chanOut <- SeederStatus{
							Index: i,
							Case1: case1,
						}
					}
				}
			} else {
				chanOut <- SeederStatus{
					Index: i,
					Case1: case1,
				}
			}
		}
		close(chanOut)
	}()

	return chanOut
}

func InsertData(chanIn <-chan SeederStatus, numWorkers int) <-chan SeederStatus {
	// wait group to control the workers
	wg := new(sync.WaitGroup)
	chanOut := make(chan SeederStatus)

	// allocate N of workers
	wg.Add(numWorkers)

	go func() {
		// dispatch N workers
		for workerIndex := 0; workerIndex < numWorkers; workerIndex++ {
			go func(workerIndex int) {
				// listen to `chanIn` channel for incoming jobs
				for job := range chanIn {
					err := InsertToDB(job.Case1, workerIndex)
					chanOut <- SeederStatus{
						WorkerIndex:  workerIndex,
						ErrorMessage: err,
					}
				}
				wg.Done()
			}(workerIndex)
		}
	}()

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut

}

func InsertToDB(case1 Cases, workerIndex int) error {
	log.Println("Executing " + case1.GetName() + " Worker Index : " + strconv.Itoa(workerIndex))
	// execute insert
	err := case1.Execute()
	if err != nil {
		return err
	}

	return nil
}

func initSeeder(cfg *config.Config, args []string) {
	//name := reflect.TypeOf(*str).Name()
	start := time.Now()
	dbObject := db.DBInitialize(cfg)

	chanIn := SeederJobs(
		args,
		&scripts.MenuSeeder{
			Name: "Menu Seeder",
			DB:   dbObject,
		},
		&scripts.RoleSeeder{
			Name: "Role Seeder",
			DB:   dbObject,
		},
		&scripts.UserSeeder{
			Name: "User Seeder",
			DB:   dbObject,
		},
		&scripts.MappingCodeSeeder{
			Name: "Mapping Code Seeder",
			DB:   dbObject,
		},
		&scripts.TermSeeder{
			Name: "Term Seeder",
			DB:   dbObject,
		},
	)

	chanResults := InsertData(chanIn, 10)

	// track and print output
	counterTotal := 0
	counterSuccess := 0
	for index := range chanResults {
		if index.ErrorMessage != nil {
			fmt.Println("Ada error di index " + strconv.Itoa(index.Index))
		} else {
			counterSuccess++
		}
		counterTotal++
	}

	log.Printf("%d/%d of total seeder inserted", counterSuccess, counterTotal)
	log.Println("It takes", time.Since(start))
}

func NewSeeder(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "seeder [SEEDER STRUCT]",
		Short: "Running seeder",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			initSeeder(cfg, args)
		},
	}
}
