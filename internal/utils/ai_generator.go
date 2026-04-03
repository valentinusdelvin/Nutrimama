package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// GenerateMealPlanIDs pushes DB context into Gemini and forces strict integer array extraction natively!
func GenerateMealPlanIDs(scoresJSON string, foodsJSON string, days int) ([]int, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("GEMINI_API_KEY mathematically missing from .env explicitly blocking generation")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.5-flash")
	model.SetTemperature(0.2) // Low temperature for highly determined JSON validation

	totalMeals := days * 3

	prompt := fmt.Sprintf(`You are a Medical Nutritionist AI building a dynamic backend pipeline.
I will provide you with a structured Mother's condition (nutrition tracker score) and an exact JSON list of available system Food Options.

Mother's Condition JSON:
%s

Available Foods JSON:
%s

TASK:
Generate a perfectly balanced %d-day Meal Plan (Breakfast, Lunch, Dinner every day).
This requires EXACTLY %d total dietary selections physically mapping ONLY the provided food_id integers.
Optimize her nutritional deficiencies rigorously.

CRITICAL JSON STRUCTURAL CONSTRAINT:
You MUST ONLY return a raw JSON array containing exactly %d integers.
DO NOT return ANY markdown formatting, no backticks, no text. Just the array.
Example Output:
[1, 5, 2, 8, 3, 10, ... ]`, scoresJSON, foodsJSON, days, totalMeals, totalMeals)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("gemini structural failure: %v", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, errors.New("gemini violently failed to map content seamlessly")
	}

	// Safely extracting literal response payload cleanly!
	rawOutput := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])
	
	// Scrubbing aggressive ghost markdown backticks seamlessly ensuring native cast compatibility!
	rawOutput = strings.TrimSpace(rawOutput)
	rawOutput = strings.TrimPrefix(rawOutput, "```json")
	rawOutput = strings.TrimPrefix(rawOutput, "```")
	rawOutput = strings.TrimSuffix(rawOutput, "```")
	rawOutput = strings.TrimSpace(rawOutput)

	var extractedIDs []int
	if err := json.Unmarshal([]byte(rawOutput), &extractedIDs); err != nil {
		return nil, fmt.Errorf("gemini output malformed unmarshal crash natively: %v\nRaw output: %s", err, rawOutput)
	}

	if len(extractedIDs) != totalMeals {
		return nil, fmt.Errorf("gemini hallucinated mapping securely returning %d items natively instead of %d", len(extractedIDs), totalMeals)
	}

	return extractedIDs, nil
}
