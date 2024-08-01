package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]map[string]float64)

	for i := range equations {
		if _, ok := m[equations[i][0]]; !ok {
			m[equations[i][0]] = make(map[string]float64)
		}
		if _, ok := m[equations[i][1]]; !ok {
			m[equations[i][1]] = make(map[string]float64)
		}
		m[equations[i][0]][equations[i][0]] = 1.0
		m[equations[i][0]][equations[i][1]] = values[i]
		m[equations[i][1]][equations[i][0]] = 1 / values[i]
	}
	for k := range m {
		for i := range m {
			for j := range m {
				_, ok1 := m[i][k]
				_, ok2 := m[k][j]

				if ok1 && ok2 {
					m[i][j] = m[i][k] * m[k][j]
				}
			}
		}
	}

	result := make([]float64, len(queries))

	for i, q := range queries {
		_, ok1 := m[q[0]]
		if !ok1 {
			result[i] = -1.0
			continue
		}
		_, ok2 := m[q[0]][q[1]]
		if ok2 {
			result[i] = m[q[0]][q[1]]
		} else {
			result[i] = -1.0
		}
	}
	return result
}
