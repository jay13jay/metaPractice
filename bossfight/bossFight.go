package bossfight

// var N int32 = 3													// # of fighters
// var H []int32 = []int32{2, 1, 4} 	// Health
// var D []int32 = []int32{3, 1, 2} 		// Fighter Damage / sec
// var B int32 = 4 												// Boss Damage / sec


	
func getMaxDamageDealt(N int32, H []int32, D []int32, B int32) float64 {
  // Write your code here
	var maxTotalDamage float64 = 0

	// Maps to memoize the seconds and damages each warrior can last and deal
	secondsMap := make(map[int]float64)
	damageMap := make(map[int]float64)

	// Calculate damage for each combination of frontline and rearguard positions
	for i := 0; i < int(N); i++ {
		frontlineSeconds, exists := secondsMap[i]
		if !exists {
			frontlineSeconds = float64(H[i]) / float64(B)
			secondsMap[i] = frontlineSeconds
		}
		frontlineDamage, exists := damageMap[i]
		if !exists {
			frontlineDamage = float64(D[i]) * frontlineSeconds
			damageMap[i] = frontlineDamage
		}
		for j := 0; j < int(N); j++ {
			if i != j {
				rearguardSeconds, exists := secondsMap[j]
				if !exists {
					rearguardSeconds = float64(H[j]) / float64(B)
					secondsMap[j] = rearguardSeconds
				}
				rearguardDamage := float64(D[j]) * (frontlineSeconds + rearguardSeconds)

				// Total damage
				totalDamage := frontlineDamage + rearguardDamage

				// Update maximum total damage if the current combination deals more damage
				if totalDamage > maxTotalDamage {
					maxTotalDamage = totalDamage
				}
			}
		}
	}

	return maxTotalDamage
}

func GetDamage(N int32, H []int32, D []int32, B int32) float64 {
	return getMaxDamageDealt(N, H, D, B)
}