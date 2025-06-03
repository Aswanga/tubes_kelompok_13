package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Expense struct {
	Category string
	Amount   float64
}

type BudgetApp struct {
	Expenses []Expense
	Budget   float64
}

func (b *BudgetApp) AddExpense(category string, amount float64) {
	b.Expenses = append(b.Expenses, Expense{Category: category, Amount: amount})
}

func (b *BudgetApp) UbahBudgetInteractive() {
	var input string
	var budgetBaru float64

	for {
		fmt.Print("Masukkan anggaran baru: ")
		fmt.Scanln(&input)

		nilai, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Input Tidak Valid!")
			continue
		}

		budgetBaru = nilai
		break
	}

	b.Budget = budgetBaru
	fmt.Println("✅ Anggaran berhasil diatur ulang.")
}

func (b *BudgetApp) AddExpenseInteractive() {
	var input string
	var choice int
	var amount float64

	categories := []string{"Transportasi", "Akomodasi", "Makanan", "Hiburan"}

	for {
		fmt.Println("\nPilih kategori pengeluaran:")
		for i, cat := range categories {
			fmt.Printf("%d. %s\n", i+1, cat)
		}
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&input)

		parsedChoice, err := strconv.Atoi(input)
		if err != nil || parsedChoice < 1 || parsedChoice > len(categories) {
			fmt.Println("Pilihan tidak valid")
			continue
		}

		choice = parsedChoice
		break
	}

	category := categories[choice-1]

	for {
		fmt.Printf("Masukkan jumlah pengeluaran untuk kategori %s: ", category)
		fmt.Scanln(&input)

		parsedAmount, err := strconv.ParseFloat(input, 64)
		if err != nil || parsedAmount < 0 {
			fmt.Println("Jumlah tidak valid.")
			continue
		}

		amount = parsedAmount
		break
	}

	b.AddExpense(category, amount)
	fmt.Println("✅ Pengeluaran berhasil ditambahkan.")
}

func (b *BudgetApp) EditExpenseInteractive() {
	if len(b.Expenses) == 0 {
		fmt.Println("Belum ada data pengeluaran.")
		return
	}

	b.SuggestSaving()
	b.PrintExpensesWithIndex()

	var input string
	var index int

	for {
		fmt.Print("\nMasukkan index pengeluaran yang ingin diedit: ")
		fmt.Scanln(&input)

		parsedIndex, err := strconv.Atoi(input)
		if err != nil || parsedIndex < 0 || parsedIndex >= len(b.Expenses) {
			fmt.Println("Index tidak valid!")
			continue
		}

		index = parsedIndex
		break
	}

	categories := []string{"Transportasi", "Akomodasi", "Makanan", "Hiburan", "Lainnya"}
	var choice int
	for {
		fmt.Println("\nPilih kategori baru untuk pengeluaran:")
		for i, cat := range categories {
			fmt.Printf("%d. %s\n", i+1, cat)
		}
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&input)

		parsedChoice, err := strconv.Atoi(input)
		if err != nil || parsedChoice < 1 || parsedChoice > len(categories) {
			fmt.Println("Pilihan tidak valid!")
			continue
		}

		choice = parsedChoice
		break
	}

	newCategory := categories[choice-1]

	var newAmount float64
	for {
		fmt.Print("Masukkan jumlah baru: ")
		fmt.Scanln(&input)

		if _, err := strconv.Atoi(input); err == nil {
			parsedAmount, err := strconv.ParseFloat(input, 64)
			if err == nil && parsedAmount >= 0 {
				newAmount = parsedAmount
				break
			}
		}

		fmt.Println("Jumlah tidak valid!")
	}

	b.Expenses[index] = Expense{Category: newCategory, Amount: newAmount}
	fmt.Println("✅ Pengeluaran berhasil diubah.")
}

func (b *BudgetApp) DeleteExpenseInteractive() {
	if len(b.Expenses) == 0 {
		fmt.Println("Belum ada data pengeluaran.")
		return
	}

	b.PrintExpensesWithIndex()

	var input string
	var index int
	for {
		fmt.Print("Masukkan index pengeluaran yang ingin dihapus: ")
		fmt.Scanln(&input)

		parsedIndex, err := strconv.Atoi(input)
		if err != nil || parsedIndex < 0 || parsedIndex >= len(b.Expenses) {
			fmt.Println("Index tidak valid!")
			continue
		}

		index = parsedIndex
		break
	}

	b.Expenses = append(b.Expenses[:index], b.Expenses[index+1:]...)
	fmt.Println("✅ Pengeluaran berhasil dihapus.")
}

func (b *BudgetApp) PrintExpensesWithIndex() {
	fmt.Println("\nDaftar Pengeluaran:")
	for i, e := range b.Expenses {
		fmt.Printf("[%d] %s - %.2f\n", i, e.Category, e.Amount)
	}
}

func (b *BudgetApp) TotalExpenses() float64 {
	total := 0.0
	for _, e := range b.Expenses {
		total += e.Amount
	}
	return total
}

func (b *BudgetApp) SuggestSaving() {
	total := b.TotalExpenses()
	if total > b.Budget {
		fmt.Println("Saran: Pengeluaran melebihi anggaran. Kurangi biaya di salah satu kategori.")
	} else if total == b.Budget {
		fmt.Println("Saran: Pengeluaran Anda pas dengan anggaran.")
	} else {
		fmt.Println("Saran: Pengeluaran Anda lebih hemat dari anggaran. Pertahankan!")
	}
}
func (b *BudgetApp) SearchInteractive() {
	for {
		fmt.Println("Pilih metode pencarian:")
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search")
		fmt.Print("Pilihan Anda: ")

		var searchChoice int
		fmt.Scanln(&searchChoice)

		if searchChoice != 1 && searchChoice != 2 {
			fmt.Println("Pilihan tidak valid!")
			continue
		}

		fmt.Print("Masukkan kategori: ")
		var category string
		fmt.Scanln(&category)

		switch searchChoice {
		case 1:
			results := b.SearchSequential(category)
			if len(results) > 0 {
				for _, e := range results {
					fmt.Printf("Hasil (sequential) - %s: %.2f\n", e.Category, e.Amount)
				}
			} else {
				fmt.Println("Tidak ditemukan dengan sequential search.")
			}
		case 2:
			results := b.SearchBinary(category)
			if len(results) > 0 {
				for _, e := range results {
					fmt.Printf("Hasil (binary) - %s: %.2f\n", e.Category, e.Amount)
				}
			} else {
				fmt.Println("Tidak ditemukan dengan binary search.")
			}
		}

		break
	}
}

func (b *BudgetApp) SearchSequential(category string) []Expense {
	var result []Expense
	for _, e := range b.Expenses {
		if strings.EqualFold(e.Category, category) {
			result = append(result, e)
		}
	}
	return result
}

func (b *BudgetApp) SearchBinary(category string) []Expense {
	b.SortByCategory()
	var results []Expense
	low, high := 0, len(b.Expenses)-1

	for low <= high {
		mid := (low + high) / 2
		cmp := strings.Compare(strings.ToLower(b.Expenses[mid].Category), strings.ToLower(category))

		if cmp == 0 {
			results = append(results, b.Expenses[mid])

			left := mid - 1
			for left >= 0 && strings.EqualFold(b.Expenses[left].Category, category) {
				results = append(results, b.Expenses[left])
				left--
			}

			right := mid + 1
			for right < len(b.Expenses) && strings.EqualFold(b.Expenses[right].Category, category) {
				results = append(results, b.Expenses[right])
				right++
			}
			break
		} else if cmp < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return results
}
func (b *BudgetApp) TampilkanExpenses() {
	var input string
	var sortOpt int

	b.PrintExpensesWithIndex()

	for {
		fmt.Println("\nUrutkan data pengeluaran?")
		fmt.Println("1. Urutkan Berdasarkan Kategori")
		fmt.Println("2. Urutkan Berdasarkan Jumlah")
		fmt.Println("3. Tidak")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&input)

		parsedOpt, err := strconv.Atoi(input)
		if err != nil || parsedOpt < 1 || parsedOpt > 3 {
			fmt.Println("Pilihan tidak valid!")
			continue
		}

		sortOpt = parsedOpt
		break
	}

	switch sortOpt {
	case 1:
		b.SortByCategory()
	case 2:
		b.SelectionSortByAmount()
	}

	b.PrintExpensesWithIndex()
}

// Selection Sort
func (b *BudgetApp) SelectionSortByAmount() {
	n := len(b.Expenses)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if b.Expenses[j].Amount < b.Expenses[minIdx].Amount {
				minIdx = j
			}
		}
		b.Expenses[i], b.Expenses[minIdx] = b.Expenses[minIdx], b.Expenses[i]
	}
}

// Insertion Sort
func (b *BudgetApp) SortByCategory() {
	n := len(b.Expenses)
	for i := 1; i < n; i++ {
		key := b.Expenses[i]
		j := i - 1

		for j >= 0 && strings.ToLower(b.Expenses[j].Category) > strings.ToLower(key.Category) {
			b.Expenses[j+1] = b.Expenses[j]
			j--
		}
		b.Expenses[j+1] = key
	}
}

func main() {
	app := BudgetApp{
		Budget: 5000000,
	}

	app.AddExpense("Transportasi", 150000)
	app.AddExpense("Transportasi", 200000)
	app.AddExpense("Transportasi", 100000)

	app.AddExpense("Akomodasi", 500000)
	app.AddExpense("Akomodasi", 600000)
	app.AddExpense("Akomodasi", 550000)

	app.AddExpense("Makanan", 75000)
	app.AddExpense("Makanan", 100000)
	app.AddExpense("Makanan", 90000)

	app.AddExpense("Hiburan", 120000)
	app.AddExpense("Hiburan", 130000)
	app.AddExpense("Hiburan", 110000)

	for {
		fmt.Println("\n===== APLIKASI ANGGARAN =====")
		fmt.Printf("Budget Anda         : %.2f\n", app.Budget)
		fmt.Printf("Total Pengeluaran   : %.2f\n", app.TotalExpenses())
		fmt.Printf("Selisih Pengeluaran : %.2f\n", app.Budget-app.TotalExpenses())
		app.SuggestSaving()

		fmt.Println("=============================")
		fmt.Println("1. Atur ulang anggaran")
		fmt.Println("2. Tambah pengeluaran")
		fmt.Println("3. Edit pengeluaran")
		fmt.Println("4. Hapus pengeluaran")
		fmt.Println("5. Tampilkan laporan ")
		fmt.Println("6. Pencarian pengeluaran")
		fmt.Println("7. Keluar")
		fmt.Print("Pilihan Anda: ")

		var menu int
		fmt.Scanln(&menu)

		switch menu {
		case 1:
			app.UbahBudgetInteractive()

		case 2:
			app.AddExpenseInteractive()

		case 3:
			app.EditExpenseInteractive()

		case 4:
			app.DeleteExpenseInteractive()

		case 5:
			app.TampilkanExpenses()
		case 6:
			app.SearchInteractive()
		case 7:
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
