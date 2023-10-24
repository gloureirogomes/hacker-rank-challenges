package comparethetriplets

func CompareTriplets(a []int32, b []int32) []int32 {
	var alice int32 = 0
	var bob int32 = 0

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}

		if a[i] > b[i] {
			alice++
			continue
		}

		bob++
	}
	return []int32{alice, bob}
}
