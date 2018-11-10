//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt and sort packages
import (
	"fmt"
	"sort"
)

//OSProcess class
type OSProcess struct {
	ProcessID     int
	ArrivalTime   int
	BurstTime     int
	Priority      int
	TimePeriod    int
	LeftBurstTime int
	Wait          int
	TAT           int
}

// GetLeftBurstTime method
func (process *OSProcess) GetLeftBurstTime() int {
	return process.LeftBurstTime
}

//OS Process Step method
type OSProcessStep struct {
	Process OSProcess
	IsNil   bool
}

// SortFCFS method
func SortFCFS(processes [3]OSProcess) []OSProcessStep {

	var steps int
	steps = 0
	var process OSProcess
	for _, process = range processes {
		steps += int(process.BurstTime)
	}

	var processSteps = []OSProcessStep{}

	//fmt.Println(processSteps)
	var sProcesses []OSProcess
	sProcesses = getSortedProcesses(processes, steps)
	//fmt.Println(sProcesses)
	var proc OSProcess
	for _, proc = range sProcesses {
		//fmt.Println(proc)
		if proc.ProcessID != 0 {
			processSteps = append(processSteps, OSProcessStep{Process: proc, IsNil: false})
		}

	}

	return processSteps
}

// getSortedProcesses
func getSortedProcesses(processes [3]OSProcess, timePeriod int) []OSProcess {
	var cProcesses []OSProcess
	cProcesses = []OSProcess{}

	var process OSProcess
	for _, process = range processes {
		if int(process.ArrivalTime) <= timePeriod && process.GetLeftBurstTime() > 0 {
			cProcesses = append(cProcesses, process)
		}
	}
	sort.Slice(cProcesses, func(i int, j int) bool {
		if cProcesses[i].ArrivalTime < cProcesses[j].ArrivalTime {
			return true
		} else if cProcesses[i].Priority == cProcesses[j].Priority {
			return cProcesses[i].ProcessID < cProcesses[j].ProcessID
		}
		return false
	})
	return cProcesses
}

func getProcess(processes []OSProcess, ProcessID int) (*OSProcess, int) {
	var i int
	for i = 0; i < len(processes); i++ {
		if int(processes[i].ProcessID) == ProcessID {
			return &processes[i], i
		}
	}
	return nil, 0
}

// main method
func main() {

	var processes = [3]OSProcess{}

	processes[0] = OSProcess{4, 1, 3, 3, 0, 3, 0, 0}
	processes[1] = OSProcess{3, 2, 9, 5, 0, 9, 0, 0}
	processes[2] = OSProcess{1, 4, 3, 4, 0, 3, 0, 0}

	var processSteps []OSProcessStep

	processSteps = SortFCFS(processes)

	var processStep OSProcessStep
	for _, processStep = range processSteps {

		fmt.Println(processStep.Process.ProcessID)

	}

}
