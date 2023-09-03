package main

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type ReadOnlyEntry struct {
	widget.Entry
}

func (e *ReadOnlyEntry) TypedKey(key *fyne.KeyEvent) {
	// no-op
}

func (e *ReadOnlyEntry) TypedRune(r rune) {
	// no-op
}

func NewReadOnlyEntry() *ReadOnlyEntry {
	entry := &ReadOnlyEntry{}
	entry.ExtendBaseWidget(entry)

	return entry
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Double Side Printing Page Order Generator")
	myWindow.Resize(fyne.NewSize(800, 600))

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter value of N")
	input.Validator = func(text string) error {
		_, err := strconv.Atoi(text)
		if err != nil {
			return err
		}
		return nil
	}
	input.MultiLine = false

	output := NewReadOnlyEntry()
	output.MultiLine = true

	generateFunc := func() {
		n, err := strconv.Atoi(input.Text)
		if err != nil || n <= 0 || n%4 != 0 {
			output.SetText("Invalid input: N must be a number greater than 0 and divisible by 4")
			return
		}

		mixedSeries := generateMixedSeries(n)
		chunks := splitToChunks(mixedSeries, 16)

		var printStr string
		for _, chunk := range chunks {
			printStr += fmt.Sprintf("%s\n", Join(chunk, ", "))
		}

		output.SetText(printStr)
	}

	input.OnSubmitted = func(_ string) {
		generateFunc()
	}
	inputContainer := container.New(layout.NewGridLayout(1), input)

	btn := widget.NewButton("Generate", generateFunc)
	btnContainer := container.New(layout.NewVBoxLayout(), btn)

	// Create a horizontal layout for the input and button
	inputAndButton := container.New(layout.NewVBoxLayout(), inputContainer, btnContainer)

	// Make the output box vertically scrollable
	outputContainer := container.New(layout.NewAdaptiveGridLayout(1), output)
	outputContainer.Resize(fyne.NewSize(800, 400))

	// Combine inputAndButton, outputScroll, and apply layouts
	content := container.New(layout.NewBorderLayout(inputAndButton, nil, nil, nil), inputAndButton, outputContainer)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func Join[T any](items []T, sep string) string {
	var builder strings.Builder

	for i, item := range items {
		if i > 0 {
			builder.WriteString(sep)
		}
		builder.WriteString(fmt.Sprintf("%v", item))
	}

	return builder.String()
}

func generateMixedSeries(N int) []int {
	mN := N
	nN := N / 2
	mM := 1
	nM := N/2 - 1

	if N%4 != 0 {
		return []int{}
	}

	pN := generatePN(mN, nN)
	pM := generatePM(mM, nM)

	fullLen := len(pN) + len(pM)
	mixedSeries := make([]int, 0, fullLen)

	type nextPT struct {
		p   []int
		idx int
	}

	pNS := &nextPT{
		p:   pN,
		idx: -1,
	}

	pMS := &nextPT{
		p:   pM,
		idx: -1,
	}

	nextP := pNS
	for i := 0; i < fullLen; i++ {
		nextP.idx++
		mixedSeries = append(mixedSeries, nextP.p[nextP.idx])

		if i%2 == 0 {
			if nextP == pNS {
				nextP = pMS
			} else {
				nextP = pNS
			}
		}
	}

	return mixedSeries
}

func generatePN(m, n int) []int {
	series := make([]int, 0)
	for i := m; i > m-n; i-- {
		series = append(series, i)
	}
	return series
}

func generatePM(m, n int) []int {
	series := make([]int, 0)
	for i := m; i <= m+n; i++ {
		series = append(series, i)
	}
	return series
}

func splitToChunks(slice []int, chunkSize int) (divided [][]int) {
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		divided = append(divided, slice[i:end])
	}

	return divided
}
