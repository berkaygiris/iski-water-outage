package internal

func ParseOutages(apiResponse ApiResponse) ([]Outage, error) {
	type grp struct {
		outage   Outage
		zonesMap map[string]map[string]bool
	}
	grouped := make(map[string]*grp)

	for _, feat := range apiResponse.Features {
		p := feat.Properties
		if _, ok := grouped[p.IDInfo]; !ok {
			grouped[p.IDInfo] = &grp{
				outage: Outage{
					Info:      p.Description,
					ID:        p.IDInfo,
					StartDate: p.StartDate,
					EndDate:   p.EndDate,
					Zones:     []Zone{},
				},
				zonesMap: make(map[string]map[string]bool),
			}
		}
		group := grouped[p.IDInfo]
		if group.zonesMap[p.District] == nil {
			group.zonesMap[p.District] = make(map[string]bool)
		}
		group.zonesMap[p.District][p.Neighborhood] = true
	}

	outages := make([]Outage, 0, len(grouped))
	for _, g := range grouped {
		for district, neighs := range g.zonesMap {
			ns := make([]string, 0, len(neighs))
			for n := range neighs {
				ns = append(ns, n)
			}
			g.outage.Zones = append(g.outage.Zones, Zone{
				District:      district,
				Neighborhoods: ns,
			})
		}
		outages = append(outages, g.outage)
	}
	return outages, nil
}
