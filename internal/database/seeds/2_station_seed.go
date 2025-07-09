package database

type OverpassElement struct {
	Type  string            `json:"type"`
	ID    int64             `json:"id"`
	Lat   float64           `json:"lat,omitempty"`
	Lon   float64           `json:"lon,omitempty"`
	Nodes []int64           `json:"nodes,omitempty"`
	Tags  map[string]string `json:"tags,omitempty"`
}

type OverpassData struct {
	Elements []OverpassElement `json:"elements"`
}

func init() {
	//RegisterSeeder(Seeder{
	//	Name: "stationSeed",
	//	Run: func(tx *gorm.DB) error {
	//		raw, err := os.ReadFile("data/daftar_stasiun_overpass.json")
	//		if err != nil {
	//			log.Fatal("read file error:", err)
	//		}
	//
	//		var parsed OverpassData
	//		if err := json.Unmarshal(raw, &parsed); err != nil {
	//			log.Fatal("parse json error:", err)
	//		}
	//
	//		stationMap := map[int64]uuid.UUID{}
	//
	//		// --- Insert Stations ---
	//		for _, el := range parsed.Elements {
	//			if el.Type == "node" && (el.Tags["railway"] == "station" || el.Tags["public_transport"] == "station") {
	//				name := el.Tags["name"]
	//				ref := el.Tags["ref"]
	//				railRef := el.Tags["railway:ref"]
	//
	//				station := model.Station{
	//					Lat:        el.Lat,
	//					Lon:        el.Lon,
	//					Name:       name,
	//					OverpassID: &el.ID,
	//				}
	//
	//				if ref != "" {
	//					station.Ref = ref
	//				} else if railRef != "" {
	//					station.Ref = railRef
	//				}
	//
	//				if err := tx.Create(&station).Error; err != nil {
	//					return fmt.Errorf("insert station error: %w", err)
	//				}
	//
	//				stationMap[el.ID] = station.ID
	//			}
	//		}
	//
	//		file, err := os.Open("data/daftar_stasiun_overpass_update.json")
	//		if err != nil {
	//			return fmt.Errorf("gagal buka file JSON: %w", err)
	//		}
	//		defer func(file *os.File) {
	//			err := file.Close()
	//			if err != nil {
	//			}
	//		}(file)
	//
	//		// --- Update Station Refs dari daftar_stasiun_overpass_update.json ---
	//		file, err = os.Open("data/daftar_stasiun_overpass_update.json")
	//		if err != nil {
	//			return fmt.Errorf("gagal buka file JSON: %w", err)
	//		}
	//		defer func(file *os.File) {
	//			err := file.Close()
	//			if err != nil {
	//			}
	//		}(file)
	//
	//		var updatedStationsOverpass []model.StationOverpass
	//		if err := json.NewDecoder(file).Decode(&updatedStationsOverpass); err != nil {
	//			return fmt.Errorf("gagal decode JSON: %w", err)
	//		}
	//
	//		for _, station := range updatedStationsOverpass {
	//			err = tx.Table("stations").Where("overpass_id = ?", station.ID).
	//				Update("ref", station.Ref).Error
	//			if err != nil {
	//				return fmt.Errorf("gagal update stasiun %s: %w", station.ID, err)
	//			}
	//		}
	//
	//		file, err = os.Open("data/daftar_stasiun_wiki.json")
	//		if err != nil {
	//			return fmt.Errorf("gagal buka file JSON: %w", err)
	//		}
	//		defer func(file *os.File) {
	//			err := file.Close()
	//			if err != nil {
	//			}
	//		}(file)
	//
	//		var updatedStationsWiki []model.StationOverpass
	//		if err := json.NewDecoder(file).Decode(&updatedStationsWiki); err != nil {
	//			return fmt.Errorf("gagal decode JSON: %w", err)
	//		}
	//
	//		for _, station := range updatedStationsWiki {
	//			var existing model.Station
	//			err := tx.Where("ref = ?", station.Ref).First(&existing).Error
	//			if err != nil {
	//				if errors.Is(err, gorm.ErrRecordNotFound) {
	//					newStation := model.Station{
	//						Name: station.Name,
	//						Ref:  station.Ref,
	//						Lat:  station.Lat,
	//						Lon:  station.Lon,
	//					}
	//					if err := tx.Create(&newStation).Error; err != nil {
	//						return fmt.Errorf("gagal insert stasiun %s: %w", station.Ref, err)
	//					}
	//				} else {
	//					return fmt.Errorf("gagal cek stasiun %s: %w", station.Ref, err)
	//				}
	//			}
	//		}
	//
	//		log.Println("✅ rail seeded successfully")
	//		return nil
	//	},
	//})
}
