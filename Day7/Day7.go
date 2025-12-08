package Day7

type Beam struct {
	loc        int
	numOfPaths int
}

// if beam just hits . then it just changes to | and continues because we dont know if will share a parent yet
// if beam enoucnters a ^, then two new beams are created the value of the current beam being iterated over numOfPaths needs to be added to createdBeams.
// however the beam to the left could already be created, if it exists, then just add numofPaths to existing beam.
// to check if it exists, then it should be the last appended beam in beams to be created, just increment that beams numofpaths directly

func CountNumOfAvailPaths(starting_loc int, input [][]string) int {
	beamLocations := []Beam{{loc: starting_loc, numOfPaths: 1}}
	for i := 0; i < len(input); i++ {
		newBeamLocations := []Beam{}
		for _, currBeam := range beamLocations {
			if input[i][currBeam.loc] == "." {
				input[i][currBeam.loc] = "|"
				newBeamLocations = append(newBeamLocations, currBeam)
				continue
			} else if input[i][currBeam.loc] == "|" {
				newBeamLocations[len(newBeamLocations)-1].numOfPaths += currBeam.numOfPaths
			} else if input[i][currBeam.loc] == "^" {
				if input[i][currBeam.loc-1] == "|" {
					newBeamLocations[len(newBeamLocations)-1].numOfPaths += currBeam.numOfPaths
				} else {
					newBeamLocations = append(newBeamLocations, Beam{loc: currBeam.loc - 1, numOfPaths: currBeam.numOfPaths})
				}

				if currBeam.loc+1 < len(input[i]) {
					input[i][currBeam.loc+1] = "|"
					newBeamLocations = append(newBeamLocations, Beam{loc: currBeam.loc + 1, numOfPaths: currBeam.numOfPaths})
				}
			}
		}
		beamLocations = newBeamLocations
	}

	total := 0
	for _, currBeam := range beamLocations {
		total += currBeam.numOfPaths
	}
	return total
}

func BeamMagicCountSplits(starting_loc int, input [][]string) int {
	total := 0
	beamLocations := []int{starting_loc}
	for i := 0; i < len(input); i++ {
		newBeamLocations := []int{}
		for _, loc := range beamLocations {
			if input[i][loc] == "." {
				input[i][loc] = "|"
				newBeamLocations = append(newBeamLocations, loc)
			}
			if input[i][loc] == "^" { // beam ends, new beams created
				total++
				if loc-1 > -1 {
					if input[i][loc-1] != "|" {
						input[i][loc-1] = "|"
						newBeamLocations = append(newBeamLocations, loc-1)
					}
				}

				if loc+1 < len(input[i]) {
					if input[i][loc+1] != "|" {
						input[i][loc+1] = "|"
						newBeamLocations = append(newBeamLocations, loc+1)
					}
				}
			}
		}
		beamLocations = newBeamLocations
	}
	return total
}
