package main

import (
	"Assignment1/Employee"
	"Assignment1/Gym"
	"Assignment1/Hotel"
	"Assignment1/Wallet"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("\nMain menu")
		fmt.Println("1. Hotel ")
		fmt.Println("2. Employees ")
		fmt.Println("3. Gym ")
		fmt.Println("4. Wallet ")
		fmt.Println("5. Exit")
		fmt.Print("Your choice: ")

		var choice string
		fmt.Scan(&choice)

		switch choice {
		case "1":
			runHotel()
		case "2":
			runEmployees()
		case "3":
			runGym()
		case "4":
			runWallet()
		case "5":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Error! Choose 1-5")
		}
	}
}

// ===== HOTEL =====
func runHotel() {
	hotel := Hotel.Hotel{
		Rooms: make(map[string]Hotel.Room),
	}

	for {
		fmt.Println("\nHOTEL")
		fmt.Println("1. Add room")
		fmt.Println("2. Check in guest")
		fmt.Println("3. Check out guest")
		fmt.Println("4. Show vacant rooms")
		fmt.Println("5. Back to main menu")
		fmt.Print("Choice: ")

		var cmd string
		fmt.Scan(&cmd)

		if cmd == "1" {
			var room Hotel.Room
			fmt.Print("Room number: ")
			fmt.Scan(&room.RoomNumber)

			fmt.Print("Room type (luxury/standard): ")
			fmt.Scan(&room.Type)

			fmt.Print("Price per night: ")
			fmt.Scan(&room.PricePerNight)

			room.IsOccupied = false
			hotel.AddRoom(room)

		} else if cmd == "2" {
			var num string
			fmt.Print("Room number to check in: ")
			fmt.Scan(&num)
			hotel.CheckIn(num)

		} else if cmd == "3" {
			var num string
			fmt.Print("Room number to check out: ")
			fmt.Scan(&num)
			hotel.CheckOut(num)

		} else if cmd == "4" {
			fmt.Println("\nVacant rooms:")
			hotel.ListVacantRooms()

		} else if cmd == "5" {
			break
		} else {
			fmt.Println("Wrong command")
		}
	}
}

// ===== EMPLOYEES =====
func runEmployees() {
	var employees []Employee.Employee

	for {
		fmt.Println("\nEMPLOYEE")
		fmt.Println("1. Add full-time employee")
		fmt.Println("2. Add part-time employee")
		fmt.Println("3. Add contractor")
		fmt.Println("4. Add intern")
		fmt.Println("5. Calculate all salaries")
		fmt.Println("6. Back")
		fmt.Print("Choice: ")

		var cmd string
		fmt.Scan(&cmd)

		if cmd == "1" {
			var emp Employee.FullTime
			fmt.Print("Monthly salary: ")
			fmt.Scan(&emp.MonthlySalary)

			fmt.Print("Bonus rate (e.g., 0.1 for 10%): ")
			fmt.Scan(&emp.BonusRate)

			employees = append(employees, emp)
			fmt.Println("Full-time employee added")

		} else if cmd == "2" {
			var emp Employee.PartTime
			fmt.Print("Hourly rate: ")
			fmt.Scan(&emp.HourlyRate)

			fmt.Print("Hours worked: ")
			fmt.Scan(&emp.HoursWorked)

			employees = append(employees, emp)
			fmt.Println("Part-time employee added")

		} else if cmd == "3" {
			var emp Employee.Contractor
			fmt.Print("Project rate: ")
			fmt.Scan(&emp.ProjectRate)

			fmt.Print("Projects completed: ")
			fmt.Scan(&emp.ProjectsCompleted)

			employees = append(employees, emp)
			fmt.Println("Contractor added")

		} else if cmd == "4" {
			var emp Employee.Intern
			fmt.Print("Daily rate: ")
			fmt.Scan(&emp.DailyRate)

			fmt.Print("Days worked: ")
			fmt.Scan(&emp.DaysWorked)

			employees = append(employees, emp)
			fmt.Println("Intern added")

		} else if cmd == "5" {
			fmt.Println("\n=== SALARY CALCULATION ===")
			total := 0.0
			for i, emp := range employees {
				salary := emp.CalculateSalary()
				total += salary
				fmt.Printf("Employee %d: salary = %.2f, hours = %d\n",
					i+1, salary, emp.GetWorkHours())
			}
			fmt.Printf("TOTAL: %.2f\n", total)

		} else if cmd == "6" {
			break
		} else {
			fmt.Println("Input error")
		}
	}
}

// ===== GYM =====
func runGym() {
	gym := Gym.Gym{
		Members: make(map[int]Gym.Member),
	}
	nextID := 1

	for {
		fmt.Println("\nGYM")
		fmt.Println("1. Add basic member")
		fmt.Println("2. Add premium member")
		fmt.Println("3. List all members")
		fmt.Println("4. Back")
		fmt.Print("Choice: ")

		var cmd string
		fmt.Scan(&cmd)

		if cmd == "1" {
			var member Gym.BasicMember
			fmt.Print("Name: ")
			fmt.Scan(&member.Name)

			fmt.Print("Age: ")
			fmt.Scan(&member.Age)

			gym.AddMember(nextID, member)
			fmt.Printf("Basic member added, ID: %d\n", nextID)
			nextID++

		} else if cmd == "2" {
			var member Gym.PremiumMember
			fmt.Print("Name: ")
			fmt.Scan(&member.Name)

			fmt.Print("Age: ")
			fmt.Scan(&member.Age)

			fmt.Print("Has trainer? (1 - yes, 0 - no): ")
			var trainer int
			fmt.Scan(&trainer)
			member.HasTrainer = trainer == 1

			gym.AddMember(nextID, member)
			fmt.Printf("Premium member added, ID: %d\n", nextID)
			nextID++

		} else if cmd == "3" {
			fmt.Println("\nMEMBER LIST:")
			gym.ListMembers()

		} else if cmd == "4" {
			break
		} else {
			fmt.Println("Wrong choice")
		}
	}
}

// ===== WALLET =====
func runWallet() {
	wallet := Wallet.Wallet{
		Balance:      0,
		Transactions: []Wallet.Transaction{},
	}

	for {
		fmt.Println("\nWALLET")
		fmt.Printf("Balance: %.2f\n", wallet.GetBalance())
		fmt.Println("1. Add money")
		fmt.Println("2. Spend money")
		fmt.Println("3. Transaction history")
		fmt.Println("4. Back")
		fmt.Print("Choice: ")

		var cmd string
		fmt.Scan(&cmd)

		if cmd == "1" {
			var amount float64
			fmt.Print("Amount to add: ")
			fmt.Scan(&amount)
			wallet.AddMoney(amount)

		} else if cmd == "2" {
			var amount float64
			fmt.Print("Amount to spend: ")
			fmt.Scan(&amount)
			wallet.SpendMoney(amount)

		} else if cmd == "3" {
			fmt.Println("\nTRANSACTION HISTORY:")
			if len(wallet.Transactions) == 0 {
				fmt.Println("No transactions")
			} else {
				for i, t := range wallet.Transactions {
					operation := "Add"
					if t.Type == "spend" {
						operation = "Spend"
					}
					fmt.Printf("%d. %s: %.2f\n", i+1, operation, t.Amount)
				}
			}

		} else if cmd == "4" {
			break
		} else {
			fmt.Println("Input error")
		}
	}
}
