package utils

import (
	"encoding/json"
	"nutrimama/internal/entity"
)

type TrackingPayload struct {
	ProteinAnimal          bool `json:"protein_animal"`
	WaterGlasses           int  `json:"water_glasses"` // Operates on an int multiplier
	IronPill               bool `json:"iron_pill"`
	FolicAcidVitD3         bool `json:"folic_acid_vit_d3"`
	ClinicalSymptoms       bool `json:"clinical_symptoms"`
	ExclusiveBreastfeeding bool `json:"exclusive_breastfeeding"`
	MpasiHighProtein       bool `json:"mpasi_high_protein"`
	ExtraFat               bool `json:"extra_fat"`
	CarbsPlantProtein      bool `json:"carbs_plant_protein"`
	VegetablesFruits       bool `json:"vegetables_fruits"`
	IodizedSalt            bool `json:"iodized_salt"`
	VitD3IronDrops         bool `json:"vit_d3_iron_drops"`
	BowelBladderPatterns   bool `json:"bowel_bladder_patterns"`
	Omega3                 bool `json:"omega_3"`
	SelfWeightCheck        bool `json:"self_weight_check"`
	PhysicalActivity       bool `json:"physical_activity"`
	FoodTextureEval        bool `json:"food_texture_eval"`
	MotorCognitiveStim     bool `json:"motor_cognitive_stim"`
	CalciumSupplement      bool `json:"calcium_supplement"`
	LilaMeasurementCheck   bool `json:"lila_measurement_check"`
	HemoglobinCheck        bool `json:"hemoglobin_check"`
	AncCheck               bool `json:"anc_check"`
	BloodPressureCheck     bool `json:"blood_pressure_check"`
	VitA                   bool `json:"vit_a"`
	ChildWeightCheck       bool `json:"child_weight_check"`
	ChildHeightCheck       bool `json:"child_height_check"`
	HeadCircumferenceCheck bool `json:"head_circumference_check"`
	Immunization           bool `json:"immunization"`
}

func evalBool(chk bool) float64 {
	if chk {
		return 100.0
	}
	return 0.0
}

func CalculateNutrition(answersData string, trackingType string) *entity.NutritionTracking {
	var payload TrackingPayload
	_ = json.Unmarshal([]byte(answersData), &payload)

	scores := make(map[string]float64)

	hydration := float64(payload.WaterGlasses) * 10.0
	if hydration > 100 {
		hydration = 100
	}
	scores["water_glasses"] = hydration

	scores["protein_animal"] = evalBool(payload.ProteinAnimal)
	scores["iron_pill"] = evalBool(payload.IronPill)
	scores["folic_acid_vit_d3"] = evalBool(payload.FolicAcidVitD3)
	scores["clinical_symptoms"] = evalBool(payload.ClinicalSymptoms)
	scores["exclusive_breastfeeding"] = evalBool(payload.ExclusiveBreastfeeding)
	scores["mpasi_high_protein"] = evalBool(payload.MpasiHighProtein)
	scores["extra_fat"] = evalBool(payload.ExtraFat)
	scores["carbs_plant_protein"] = evalBool(payload.CarbsPlantProtein)
	scores["vegetables_fruits"] = evalBool(payload.VegetablesFruits)
	scores["iodized_salt"] = evalBool(payload.IodizedSalt)
	scores["vit_d3_iron_drops"] = evalBool(payload.VitD3IronDrops)
	scores["bowel_bladder_patterns"] = evalBool(payload.BowelBladderPatterns)
	scores["omega_3"] = evalBool(payload.Omega3)
	scores["self_weight_check"] = evalBool(payload.SelfWeightCheck)
	scores["physical_activity"] = evalBool(payload.PhysicalActivity)
	scores["food_texture_eval"] = evalBool(payload.FoodTextureEval)
	scores["motor_cognitive_stim"] = evalBool(payload.MotorCognitiveStim)
	scores["calcium_supplement"] = evalBool(payload.CalciumSupplement)
	scores["lila_measurement_check"] = evalBool(payload.LilaMeasurementCheck)
	scores["hemoglobin_check"] = evalBool(payload.HemoglobinCheck)
	scores["anc_check"] = evalBool(payload.AncCheck)
	scores["blood_pressure_check"] = evalBool(payload.BloodPressureCheck)
	scores["vit_a"] = evalBool(payload.VitA)
	scores["child_weight_check"] = evalBool(payload.ChildWeightCheck)
	scores["child_height_check"] = evalBool(payload.ChildHeightCheck)
	scores["head_circumference_check"] = evalBool(payload.HeadCircumferenceCheck)
	scores["immunization"] = evalBool(payload.Immunization)

	scoresBytes, _ := json.Marshal(scores)

	var tracking entity.NutritionTracking
	tracking.TrackingType = &trackingType
	tracking.ScoresData = string(scoresBytes)

	return &tracking
}
