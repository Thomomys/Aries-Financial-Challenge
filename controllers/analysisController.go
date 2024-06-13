package controllers

import (
	"aries-financial-challenge/model"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"slices"
)

// AnalysisResponse represents the data structure of the analysis result
type AnalysisResponse struct {
	XYValues        []XYValue `json:"xy_values"`
	MaxProfit       float64   `json:"max_profit"`
	MaxLoss         float64   `json:"max_loss"`
	BreakEvenPoints []float64 `json:"break_even_points"`
}

// XYValue represents a pair of X and Y values
type XYValue struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func AnalysisHandler(w http.ResponseWriter, r *http.Request) {
	var contracts []model.OptionsContract

	if err := json.NewDecoder(r.Body).Decode(&contracts); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AnalysisResponse{
		XYValues:        calculateXYValues(contracts),
		MaxProfit:       calculateMaxProfit(contracts),
		MaxLoss:         calculateMaxLoss(contracts),
		BreakEvenPoints: calculateBreakEvenPoints(contracts),
	})
}

func calculateXYValues(contracts []model.OptionsContract) []XYValue {
	xyValues := []XYValue{}
	risks := []float64{}

	maxStrikePrice := 0.0
	minStrikePrice := math.Inf(1)
	for _, contract := range contracts {
		risks = append(risks, contract.StrikePrice)
		minStrikePrice = min(minStrikePrice, contract.StrikePrice)
		maxStrikePrice = max(maxStrikePrice, contract.StrikePrice)
	}
	for i := 1.0; i < 10.0; i += 1.0 {
		risks = append(risks, minStrikePrice-(minStrikePrice*i/20.0))
		risks = append(risks, maxStrikePrice+(maxStrikePrice*i/20.0))
	}

	slices.SortFunc(risks, func(a float64, b float64) int {
		return int(math.Ceil(a - b))
	})

	fmt.Println(risks)

	for _, risk := range risks {
		totalReward := 0.0
		for _, contract := range contracts {
			reward := 0.0
			if contract.Type == "Call" && contract.LongShort == "long" {
				reward = max(0, risk-contract.StrikePrice) - contract.Ask
			} else if contract.Type == "Call" && contract.LongShort == "short" {
				reward = contract.Bid - max(0, risk-contract.StrikePrice)
			} else if contract.Type == "Put" && contract.LongShort == "long" {
				reward = max(0, contract.StrikePrice-risk) - contract.Ask
			} else if contract.Type == "Put" && contract.LongShort == "short" {
				if risk > contract.StrikePrice {
					reward = contract.Bid
				} else {
					reward = contract.StrikePrice - risk - contract.Bid
				}
			}
			fmt.Print(reward, " ")
			totalReward += reward
		}
		fmt.Println()
		xyValues = append(xyValues, XYValue{
			X: risk,
			Y: totalReward,
		})
	}

	return xyValues
}

func calculateMaxProfit(contracts []model.OptionsContract) float64 {
	maxProfit := 0.0

	for _, option := range contracts {
		if option.Type == "Call" && option.LongShort == "long" {
			maxProfit = max(maxProfit, 0.0)
		} else if option.Type == "Call" && option.LongShort == "short" {
			maxProfit = max(maxProfit, option.Bid)
		} else if option.Type == "Put" && option.LongShort == "short" {
			maxProfit = max(maxProfit, option.Bid)
		} else if option.Type == "Put" && option.LongShort == "long" {
			maxProfit = max(maxProfit, option.StrikePrice-option.Ask)
		}
	}

	return maxProfit
}

func calculateMaxLoss(contracts []model.OptionsContract) float64 {
	maxLoss := 0.0
	for _, contract := range contracts {
		if contract.Type == "Call" && contract.LongShort == "long" {
			maxLoss = max(maxLoss, contract.Ask)
		} else if contract.Type == "Call" && contract.LongShort == "short" {
			maxLoss = max(maxLoss, contract.StrikePrice-contract.Bid)
		} else if contract.Type == "Put" && contract.LongShort == "short" {
			maxLoss = max(maxLoss, float64(0))
		} else if contract.Type == "Put" && contract.LongShort == "long" {
			maxLoss = max(maxLoss, contract.Ask)
		}
	}

	return maxLoss
}

func calculateBreakEvenPoints(contracts []model.OptionsContract) []float64 {
	breakEvenPoint := make([]float64, len(contracts))

	for i, contract := range contracts {
		if contract.Type == "Call" && contract.LongShort == "long" {
			breakEvenPoint[i] = contract.StrikePrice + contract.Ask
		} else if contract.Type == "Call" && contract.LongShort == "short" {
			breakEvenPoint[i] = contract.StrikePrice - contract.Bid
		} else if contract.Type == "Put" && contract.LongShort == "long" {
			breakEvenPoint[i] = contract.StrikePrice - contract.Ask
		} else if contract.Type == "Put" && contract.LongShort == "short" {
			breakEvenPoint[i] = contract.StrikePrice + contract.Bid
		}
	}

	return breakEvenPoint
}
