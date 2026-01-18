package day10

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-25/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day10/input")
	return part1(input), part2(input)
}

type Button []int

type Machine struct {
	target   []bool
	buttons  []Button
	joltages []int
}

func loadMachines(input []string) []Machine {
	machines := make([]Machine, 0)
	for _, line := range input {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")

		targetStr := lineSplit[0]
		target := make([]bool, 0)
		for _, char := range targetStr {
			if char == '.' {
				target = append(target, false)
			} else if char == '#' {
				target = append(target, true)
			}
		}

		buttons := make([]Button, 0)
		for _, buttonStr := range lineSplit[1 : len(lineSplit)-1] {
			button := make(Button, 0)
			for _, buttonChar := range strings.Split(buttonStr[1:len(buttonStr)-1], ",") {
				toggle, _ := strconv.Atoi(string(buttonChar))
				button = append(button, toggle)
			}
			buttons = append(buttons, button)

		}

		joltagesStr := lineSplit[len(lineSplit)-1]
		joltagesStr = strings.Trim(joltagesStr, "{")
		joltagesStr = strings.Trim(joltagesStr, "}")
		joltagesChars := strings.Split(joltagesStr, ",")
		joltages := make([]int, 0)
		for _, char := range joltagesChars {
			jolt, _ := strconv.Atoi(string(char))
			joltages = append(joltages, jolt)
		}

		machines = append(machines, Machine{
			target:   target,
			joltages: joltages,
			buttons:  buttons,
		})

	}
	return machines
}

func (b Button) push(lights []bool) []bool {
	newLights := make([]bool, len(lights))
	for i, light := range lights {
		newLights[i] = light
	}
	for _, light := range b {
		newLights[light] = !newLights[light]
	}
	return newLights
}

func (m Machine) configure(lights []bool, buttons []Button) (bool, int) {
	// fmt.Println(m.target, lights, buttons)
	if slices.Equal(lights, m.target) {
		return true, 0
	}
	if len(buttons) == 0 {
		return false, 0
	}

	button := buttons[0]
	newLights := button.push(lights)
	newButtons := buttons[1:]

	resWithBtn, countWithBtn := m.configure(newLights, newButtons)
	countWithBtn += 1
	resWithoutBtn, countWithoutBtn := m.configure(lights, newButtons)

	if resWithBtn {
		if (!resWithoutBtn) || (resWithoutBtn && countWithBtn < countWithoutBtn) {
			return true, countWithBtn
		}
	}

	return resWithoutBtn, countWithoutBtn
}

func part1(input []string) string {
	machines := loadMachines(input)
	total := 0
	for _, machine := range machines {
		_, res := machine.configure(make([]bool, len(machine.target)), machine.buttons)
		total += res
	}
	return fmt.Sprint(total)
}

func part2(input []string) string {
	return fmt.Sprint(0)
}
