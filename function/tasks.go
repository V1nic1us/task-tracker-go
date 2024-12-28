package function

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func findAvailableID(tasks Tasks) int {
	usedIDs := make(map[int]bool)
	for _, task := range tasks.Tasks {
		usedIDs[task.ID] = true
	}

	// Procura o menor ID disponível a partir de 1
	for id := 1; ; id++ {
		if !usedIDs[id] {
			return id
		}
	}
}

func AddTask(content []byte) {
	if len(os.Args) < 3 {
		fmt.Println("task is required")
		os.Exit(1)
	}

	task := os.Args[2]
	var tasks Tasks
	err := json.Unmarshal(content, &tasks)
	if err != nil {
		tasks.Tasks = append(tasks.Tasks, Task{ID: 1, Description: task, Status: "to-do", CreatedAt: time.Now().String(), UpdatedAt: time.Now().String()})
		dat, err := json.MarshalIndent(tasks, "", "    ")
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		err = os.WriteFile("tasks.json", dat, 0666)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	} else {
		newId := findAvailableID(tasks)
		tasks.Tasks = append(tasks.Tasks, Task{ID: newId, Description: task, Status: "to-do", CreatedAt: time.Now().String(), UpdatedAt: time.Now().String()})
		dat, err := json.MarshalIndent(tasks, "", "    ")
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		err = os.WriteFile("tasks.json", dat, 0666)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	}
}

func UpdateTask(content []byte) {
	if len(os.Args) < 4 {
		fmt.Println("task is required")
		os.Exit(1)
	}
	taskiD, e := strconv.Atoi(os.Args[2])
	if e != nil {
		fmt.Println("error:", e)
		os.Exit(1)
	}
	taskDescription := os.Args[3]

	var tasks Tasks
	err := json.Unmarshal(content, &tasks)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	} else {
		for i, task := range tasks.Tasks {
			if taskiD == task.ID {
				tasks.Tasks[i].Description = taskDescription
				tasks.Tasks[i].UpdatedAt = time.Now().String()
				dat, err := json.MarshalIndent(tasks, "", "    ")
				if err != nil {
					fmt.Println("error:", err)
					os.Exit(1)
				}
				err = os.WriteFile("tasks.json", dat, 0666)
				if err != nil {
					fmt.Println("error:", err)
					os.Exit(1)
				}
				return
			}
		}
		fmt.Println("Task not found")
		os.Exit(1)
	}

}

func DeleteTask(content []byte) {
	if len(os.Args) < 3 {
		fmt.Println("task is required")
		os.Exit(1)
	}

	taskiD, e := strconv.Atoi(os.Args[2])
	if e != nil {
		fmt.Println("error:", e)
		os.Exit(1)
	}

	var tasks Tasks
	err := json.Unmarshal(content, &tasks)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	} else {
		for i, task := range tasks.Tasks {
			if taskiD == task.ID {
				tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
				dat, err := json.MarshalIndent(tasks, "", "    ")
				if err != nil {
					fmt.Println("error:", err)
					os.Exit(1)
				}
				err = os.WriteFile("tasks.json", dat, 0666)
				if err != nil {
					fmt.Println("error:", err)
					os.Exit(1)
				}
				return
			}
		}
		fmt.Println("Task not found")
		os.Exit(1)
	}
}

func MarkInProgressTask(content []byte) {
	if len(os.Args) < 3 {
		fmt.Println("task is required")
		os.Exit(1)
	}
	taskiD, e := strconv.Atoi(os.Args[2])
	if e != nil {
		fmt.Println("error:", e)
		os.Exit(1)
	}

	var tasks Tasks
	err := json.Unmarshal(content, &tasks)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	} else {
		for i, v := range tasks.Tasks {
			if v.ID ==  taskiD{
				tasks.Tasks[i].Status = "in-progress"
				tasks.Tasks[i].UpdatedAt = time.Now().String()
				dat, err := json.MarshalIndent(tasks, "", "    ")
				if err != nil {
					fmt.Println("error:", err)
					os.Exit(1)
				}
				err = os.WriteFile("tasks.json", dat, 0666)
				if err != nil {
					fmt.Println("error:", err)
					os.Exit(1)
				}
				return
			}
		}
	}
}

func MarkDoneTask(content []byte) {
	if len(os.Args) < 3 {
		fmt.Println("task is required")
		os.Exit(1)
	}
	taskiD, e := strconv.Atoi(os.Args[2])
	if e != nil {
		fmt.Println("error:", e)
		os.Exit(1)
	}

	var tasks Tasks
	err := json.Unmarshal(content, &tasks)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	} else {
		for i, v := range tasks.Tasks {
			if v.ID ==  taskiD{
				tasks.Tasks[i].Status = "done"
				tasks.Tasks[i].UpdatedAt = time.Now().String()
				dat, err := json.MarshalIndent(tasks, "", "    ")
				if err != nil {
					fmt.Println("error:", err)
					os.Exit(1)
				}
				err = os.WriteFile("tasks.json", dat, 0666)
				if err != nil {
					fmt.Println("error:", err)
					os.Exit(1)
				}
				return
			}
		}
	}
}


func ListTasks(content []byte) {
	var tasks Tasks
	err := json.Unmarshal(content, &tasks)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("Tasks:")
	fmt.Println()
	if len(os.Args) < 3 {
		for _, task := range tasks.Tasks {
			fmt.Printf("ID: %d\n", task.ID)
			fmt.Printf("Description: %s\n", task.Description)
			fmt.Printf("Status: %s\n", task.Status)
			fmt.Printf("Created At: %s\n", task.CreatedAt)
			fmt.Printf("Updated At: %s\n", task.UpdatedAt)
			fmt.Println()
		}
		return
	}

	listOption:= os.Args[2]

	if listOption == "todo" {
		for _, task := range tasks.Tasks {
			if task.Status == "to-do" {
				fmt.Printf("ID: %d\n", task.ID)
				fmt.Printf("Description: %s\n", task.Description)
				fmt.Printf("Status: %s\n", task.Status)
				fmt.Printf("Created At: %s\n", task.CreatedAt)
				fmt.Printf("Updated At: %s\n", task.UpdatedAt)
				fmt.Println()
			}
		}
		return
	}

	if listOption == "in-progress" {
		for _, task := range tasks.Tasks {
			if task.Status == "in-progress" {
				fmt.Printf("ID: %d\n", task.ID)
				fmt.Printf("Description: %s\n", task.Description)
				fmt.Printf("Status: %s\n", task.Status)
				fmt.Printf("Created At: %s\n", task.CreatedAt)
				fmt.Printf("Updated At: %s\n", task.UpdatedAt)
				fmt.Println()
			}
		}
		return
	}

	if listOption == "done" {
		for _, task := range tasks.Tasks {
			if task.Status == "done" {
				fmt.Printf("ID: %d\n", task.ID)
				fmt.Printf("Description: %s\n", task.Description)
				fmt.Printf("Status: %s\n", task.Status)
				fmt.Printf("Created At: %s\n", task.CreatedAt)
				fmt.Printf("Updated At: %s\n", task.UpdatedAt)
				fmt.Println()
			}
		}
		return
	}
}
