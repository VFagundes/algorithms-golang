package main

import "fmt"

/*
influencers = [
 {"id": 1, "sales": 15000, "engagement": 0.07, "reach": 500000},
 {"id": 2, "sales": 12000, "engagement": 0.05, "reach": 300000},
 {"id": 3, "sales": 20000, "engagement": 0.08, "reach": 1000000},
]
conditions = [
 {"metric": "sales",      "operator": ">=", "value": 15000},
 {"metric": "engagement", "operator": ">=", "value": 0.06}
]
The number of influencers can be as large as 10^6.
Each influencer will have at least the metrics: `sales`, `engagement`, and `reach`.
The number of conditions will not exceed 10.
You may use any programming language.
Please talk through your thought process as you code.
Focus on writing clean and efficient code.
Once done, explain the time and space complexity of your solution. Challenge yourself to optimize the solution if possible.
influencers = [
{"id": 1, "sales": 17000, "engagement": 0.04, "reach": 200000},
{"id": 2, "sales": 21000, "engagement": 0.05, "reach": 300000},
{"id": 3, "sales": 22000, "engagement": 0.06, "reach": 800000},
{"id": 4, "sales": 15000, "engagement": 0.05, "reach": 600000}
]

conditions = [
{"metric": "sales", "operator": ">", "value": 20000},
{"metric": "reach", "operator": "<=", "value": 500000}
]
*/
// Influencer structure
type Influencer struct {
	ID         int
	Sales      int
	Engagement float64
	Reach      int
}

// Condition structures
type Condition struct {
	Metric   string
	Operator string
	Value    interface{}
}

// Function to filter influencers based on conditions
func filterInfluencers(influencers []Influencer, conditions []Condition) []int {
	var result []int

	// Loop through each influencer
	for _, influencer := range influencers {
		meetsAllConditions := true

		// Check each condition
		for _, condition := range conditions {
			var influencerValue interface{}
			switch condition.Metric {
			case "sales":
				influencerValue = influencer.Sales
			case "engagement":
				influencerValue = influencer.Engagement
			case "reach":
				influencerValue = influencer.Reach
			}

			// Evaluate the condition
			if !evaluateCondition(influencerValue, condition.Operator, condition.Value) {
				meetsAllConditions = false
				break
			}
		}

		// If all conditions are met, add the influencer ID to the result
		if meetsAllConditions {
			result = append(result, influencer.ID)
		}
	}

	return result
}

// Simplified function to evaluate conditions for both int and float64
func evaluateCondition(value interface{}, operator string, conditionValue interface{}) bool {
	switch v := value.(type) {
	case int:
		cv := conditionValue.(int)
		return compare(v, cv, operator)
	case float64:
		cv := conditionValue.(float64)
		return compare(v, cv, operator)
	}
	return false
}

// Generic comparison function for int and float64
func compare[T int | float64](a, b T, operator string) bool {
	switch operator {
	case ">=":
		return a >= b
	case "<=":
		return a <= b
	case ">":
		return a > b
	case "<":
		return a < b
	}
	return false
}

func main() {
	// Example influencers
	influencers := []Influencer{
		{ID: 1, Sales: 17000, Engagement: 0.04, Reach: 200000},
		{ID: 2, Sales: 21000, Engagement: 0.05, Reach: 300000},
		{ID: 3, Sales: 22000, Engagement: 0.06, Reach: 800000},
		{ID: 4, Sales: 15000, Engagement: 0.05, Reach: 600000},
	}

	// Example conditions
	conditions := []Condition{
		{Metric: "sales", Operator: ">", Value: 20000},
		{Metric: "reach", Operator: "<=", Value: 500000},
	}

	// Call filter function
	result := filterInfluencers(influencers, conditions)
	fmt.Println(result) // Expected output: [2]
}
