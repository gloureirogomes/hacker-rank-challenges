package comparethetriplets

func CompareTriplets() []int32 {
	aliceTriplets := []int32{5, 6, 7}
	bobTriplets := []int32{3, 6, 10}
	var alice int32 = 0
	var bob int32 = 0

	for i := 0; i < len(aliceTriplets); i++ {
		if aliceTriplets[i] == bobTriplets[i] {
			continue
		}

		if aliceTriplets[i] > bobTriplets[i] {
			alice++
			continue
		}

		bob++
	}
	return []int32{alice, bob}
}
