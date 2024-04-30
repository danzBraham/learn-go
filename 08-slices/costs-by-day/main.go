package main

func main() {
	costs := []cost{
		{0, 1.8},
		{1, 2.8},
		{5, 4.8},
		{3, 2.8},
		{1, 1.3},
	}

	getCostsByDay(costs)
}

type cost struct {
	day   int
	value float64
}

/*
1st Approach
Pros:
  - More efficient for large input slices since it allocates the slice with the correct length upfront.
  - Easier to understand and maintain.

Cons:
  - Requires an extra loop to find the maximum day number.
*/
func getCostsByDay(costs []cost) []float64 {
	lastDay := 0
	for _, c := range costs {
		if c.day > lastDay {
			lastDay = c.day
		}
	}

	costsByDay := make([]float64, lastDay+1)

	for _, c := range costs {
		costsByDay[c.day] += c.value
	}

	return costsByDay
}

/*
2nd Approach
Pros:
- Doesn't require an extra loop to find the maximum day number.
- Handles the case where days are not in sequential order more naturally.

Cons:
- Less efficient for large input slices due to the repeated use of append,
	which involves reallocating and copying the underlying array.
- Slightly more complex and harder to understand at a glance.
*/
// func getCostsByDay(costs []cost) []float64 {
// 	costsByDay := []float64{}
// 	for i := 0; i < len(costs); i++ {
// 		day := costs[i].day
// 		value := costs[i].value
// 		for day >= len(costsByDay) {
// 			costsByDay = append(costsByDay, 0.0)
// 		}
// 		costsByDay[day] += value
// 	}
// 	return costsByDay
// }
