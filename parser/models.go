package main

import "encoding/xml"

// FeatureCollection was generated 2023-11-19 10:18:16 by bram on PC.
type FeatureCollection struct {
	XMLName        xml.Name `xml:"FeatureCollection"`
	Text           string   `xml:",chardata"`
	TnRo           string   `xml:"tn-ro,attr"`
	Gml            string   `xml:"gml,attr"`
	TnRa           string   `xml:"tn-ra,attr"`
	Tn             string   `xml:"tn,attr"`
	Net            string   `xml:"net,attr"`
	Gn             string   `xml:"gn,attr"`
	Base           string   `xml:"base,attr"`
	Gmd            string   `xml:"gmd,attr"`
	Xlink          string   `xml:"xlink,attr"`
	Gco            string   `xml:"gco,attr"`
	Ns1            string   `xml:"ns1,attr"`
	Gss            string   `xml:"gss,attr"`
	Gsr            string   `xml:"gsr,attr"`
	Gts            string   `xml:"gts,attr"`
	Hfp            string   `xml:"hfp,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	FeatureMember  []struct {
		Text                string `xml:",chardata"`
		ConditionOfFacility struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"id,attr"`
			NetworkRef struct {
				Text          string `xml:",chardata"`
				LinkReference struct {
					Text    string `xml:",chardata"`
					Element struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"element"`
					ApplicableDirection struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
					} `xml:"applicableDirection"`
				} `xml:"LinkReference"`
			} `xml:"networkRef"`
			InspireId struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text      string `xml:",chardata"`
					LocalId   string `xml:"localId"`   // cof-117911911, cof-117911...
					Namespace string `xml:"namespace"` // nl-top10nl-tn, nl-top10nl...
				} `xml:"Identifier"`
			} `xml:"inspireId"`
			BeginLifespanVersion string `xml:"beginLifespanVersion"` // 2019-02-01T00:00:00Z, 201...
			EndLifespanVersion   struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"endLifespanVersion"`
			ValidFrom struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"validFrom"`
			CurrentStatus struct {
				Text  string `xml:",chardata"`
				Href  string `xml:"href,attr"`
				Title string `xml:"title,attr"`
			} `xml:"currentStatus"`
		} `xml:"ConditionOfFacility"`
		MarkerPost struct {
			Text                 string `xml:",chardata"`
			ID                   string `xml:"id,attr"`
			BeginLifespanVersion string `xml:"beginLifespanVersion"` // 2014-03-29T00:00:00Z, 201...
			InspireId            struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text      string `xml:",chardata"`
					LocalId   string `xml:"localId"`   // 127315625, 129204299, 129...
					Namespace string `xml:"namespace"` // nl-top10nl-tn, nl-top10nl...
				} `xml:"Identifier"`
			} `xml:"inspireId"`
			EndLifespanVersion struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"endLifespanVersion"`
			InNetwork struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"inNetwork"`
			Geometry struct {
				Text  string `xml:",chardata"`
				Point struct {
					Text         string `xml:",chardata"`
					SrsName      string `xml:"srsName,attr"`
					SrsDimension string `xml:"srsDimension,attr"`
					Pos          string `xml:"pos"` // 52.71538974 5.02795098, 5...
				} `xml:"Point"`
			} `xml:"geometry"`
			ValidFrom struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"validFrom"`
			Location struct {
				Text string `xml:",chardata"` // 1.0, 1.0, 1.0, 1.0, 1.0, ...
				Uom  string `xml:"uom,attr"`
			} `xml:"location"`
			Route struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"route"`
		} `xml:"MarkerPost"`
		NominalTrackGauge struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"id,attr"`
			NetworkRef struct {
				Text          string `xml:",chardata"`
				LinkReference struct {
					Text    string `xml:",chardata"`
					Element struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"element"`
					ApplicableDirection struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
					} `xml:"applicableDirection"`
				} `xml:"LinkReference"`
			} `xml:"networkRef"`
			InspireId struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text      string `xml:",chardata"`
					LocalId   string `xml:"localId"`   // ntg-117911911, ntg-117911...
					Namespace string `xml:"namespace"` // nl-top10nl-tn, nl-top10nl...
				} `xml:"Identifier"`
			} `xml:"inspireId"`
			BeginLifespanVersion string `xml:"beginLifespanVersion"` // 2019-02-01T00:00:00Z, 201...
			EndLifespanVersion   struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"endLifespanVersion"`
			ValidFrom struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"validFrom"`
			NominalGauge struct {
				Text      string `xml:",chardata"`
				Uom       string `xml:"uom,attr"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"nominalGauge"`
			NominalGaugeCategory string `xml:"nominalGaugeCategory"` // standard, standard, stand...
		} `xml:"NominalTrackGauge"`
		NumberOfTracks struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"id,attr"`
			NetworkRef struct {
				Text          string `xml:",chardata"`
				LinkReference struct {
					Text    string `xml:",chardata"`
					Element struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"element"`
					ApplicableDirection struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
					} `xml:"applicableDirection"`
				} `xml:"LinkReference"`
			} `xml:"networkRef"`
			InspireId struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text      string `xml:",chardata"`
					LocalId   string `xml:"localId"`   // not-117911911, not-117911...
					Namespace string `xml:"namespace"` // nl-top10nl-tn, nl-top10nl...
				} `xml:"Identifier"`
			} `xml:"inspireId"`
			BeginLifespanVersion string `xml:"beginLifespanVersion"` // 2019-02-01T00:00:00Z, 201...
			EndLifespanVersion   struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"endLifespanVersion"`
			ValidFrom struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"validFrom"`
			MinMaxNumberOfTracks struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"minMaxNumberOfTracks"`
			NumberOfTracks string `xml:"numberOfTracks"` // 1, 1, 2, 2, 1, 2, 2, 2, 1...
		} `xml:"NumberOfTracks"`
		RailwayArea struct {
			Text                 string `xml:",chardata"`
			ID                   string `xml:"id,attr"`
			BeginLifespanVersion string `xml:"beginLifespanVersion"` // 2015-05-31T00:00:00Z, 201...
			InspireId            struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text      string `xml:",chardata"`
					LocalId   string `xml:"localId"`   // 123122948, 127079277, 106...
					Namespace string `xml:"namespace"` // nl-top10nl-tn, nl-top10nl...
				} `xml:"Identifier"`
			} `xml:"inspireId"`
			EndLifespanVersion struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"endLifespanVersion"`
			InNetwork struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"inNetwork"`
			Geometry struct {
				Text    string `xml:",chardata"`
				Polygon struct {
					Text         string `xml:",chardata"`
					SrsName      string `xml:"srsName,attr"`
					SrsDimension string `xml:"srsDimension,attr"`
					Exterior     struct {
						Text       string `xml:",chardata"`
						LinearRing struct {
							Text    string `xml:",chardata"`
							PosList string `xml:"posList"` // 51.50161090 3.76397274 51...
						} `xml:"LinearRing"`
					} `xml:"exterior"`
					Interior []struct {
						Text       string `xml:",chardata"`
						LinearRing struct {
							Text    string `xml:",chardata"`
							PosList string `xml:"posList"` // 51.00509679 5.86034360 51...
						} `xml:"LinearRing"`
					} `xml:"interior"`
				} `xml:"Polygon"`
			} `xml:"geometry"`
			GeographicalName struct {
				Text             string `xml:",chardata"`
				NilReason        string `xml:"nilReason,attr"`
				Nil              string `xml:"nil,attr"`
				GeographicalName struct {
					Text       string `xml:",chardata"`
					Language   string `xml:"language"` // nld, nld, nld, nld, nld, ...
					Nativeness struct {
						Text  string `xml:",chardata"`
						Href  string `xml:"href,attr"`
						Title string `xml:"title,attr"`
					} `xml:"nativeness"`
					NameStatus struct {
						Text  string `xml:",chardata"`
						Href  string `xml:"href,attr"`
						Title string `xml:"title,attr"`
					} `xml:"nameStatus"`
					SourceOfName  string `xml:"sourceOfName"` // Gazetteer, Gazetteer, Gaz...
					Pronunciation struct {
						Text      string `xml:",chardata"`
						NilReason string `xml:"nilReason,attr"`
						Nil       string `xml:"nil,attr"`
					} `xml:"pronunciation"`
					Spelling struct {
						Text           string `xml:",chardata"`
						SpellingOfName struct {
							Chardata              string `xml:",chardata"`
							Text                  string `xml:"text"`   // Reefbrug, Deelbrug, Deelb...
							Script                string `xml:"script"` // Latn, Latn, Latn, Latn, L...
							TransliterationScheme struct {
								Text      string `xml:",chardata"`
								NilReason string `xml:"nilReason,attr"`
								Nil       string `xml:"nil,attr"`
							} `xml:"transliterationScheme"`
						} `xml:"SpellingOfName"`
					} `xml:"spelling"`
				} `xml:"GeographicalName"`
			} `xml:"geographicalName"`
			ValidFrom struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"validFrom"`
		} `xml:"RailwayArea"`
		RailwayElectrification struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"id,attr"`
			NetworkRef struct {
				Text          string `xml:",chardata"`
				LinkReference struct {
					Text    string `xml:",chardata"`
					Element struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"element"`
					ApplicableDirection struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
					} `xml:"applicableDirection"`
				} `xml:"LinkReference"`
			} `xml:"networkRef"`
			InspireId struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text      string `xml:",chardata"`
					LocalId   string `xml:"localId"`   // re-117911911, re-11791199...
					Namespace string `xml:"namespace"` // nl-top10nl-tn, nl-top10nl...
				} `xml:"Identifier"`
			} `xml:"inspireId"`
			BeginLifespanVersion string `xml:"beginLifespanVersion"` // 2019-02-01T00:00:00Z, 201...
			EndLifespanVersion   struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"endLifespanVersion"`
			ValidFrom struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"validFrom"`
			Electrified string `xml:"electrified"` // false, false, true, true,...
		} `xml:"RailwayElectrification"`
		RailwayLink struct {
			Text                 string `xml:",chardata"`
			ID                   string `xml:"id,attr"`
			BeginLifespanVersion string `xml:"beginLifespanVersion"` // 2019-02-01T00:00:00Z, 201...
			InspireId            struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text      string `xml:",chardata"`
					LocalId   string `xml:"localId"`   // rl-117911911, rl-11791199...
					Namespace string `xml:"namespace"` // nl-top10nl-tn, nl-top10nl...
				} `xml:"Identifier"`
			} `xml:"inspireId"`
			EndLifespanVersion struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"endLifespanVersion"`
			InNetwork struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"inNetwork"`
			CentrelineGeometry struct {
				Text       string `xml:",chardata"`
				LineString struct {
					Text         string `xml:",chardata"`
					SrsName      string `xml:"srsName,attr"`
					SrsDimension string `xml:"srsDimension,attr"`
					PosList      string `xml:"posList"` // 53.41789240 6.75796325 53...
				} `xml:"LineString"`
			} `xml:"centrelineGeometry"`
			Fictitious       string `xml:"fictitious"` // false, false, false, fals...
			GeographicalName struct {
				Text             string `xml:",chardata"`
				NilReason        string `xml:"nilReason,attr"`
				Nil              string `xml:"nil,attr"`
				GeographicalName struct {
					Text       string `xml:",chardata"`
					Language   string `xml:"language"` // nld, nld, nld, nld, nld, ...
					Nativeness struct {
						Text  string `xml:",chardata"`
						Href  string `xml:"href,attr"`
						Title string `xml:"title,attr"`
					} `xml:"nativeness"`
					NameStatus struct {
						Text  string `xml:",chardata"`
						Href  string `xml:"href,attr"`
						Title string `xml:"title,attr"`
					} `xml:"nameStatus"`
					SourceOfName  string `xml:"sourceOfName"` // Gazetteer, Gazetteer, Gaz...
					Pronunciation struct {
						Text      string `xml:",chardata"`
						NilReason string `xml:"nilReason,attr"`
						Nil       string `xml:"nil,attr"`
					} `xml:"pronunciation"`
					Spelling struct {
						Text           string `xml:",chardata"`
						SpellingOfName struct {
							Chardata              string `xml:",chardata"`
							Text                  string `xml:"text"`   // Zwijndrechtse Brug, Zwijn...
							Script                string `xml:"script"` // Latn, Latn, Latn, Latn, L...
							TransliterationScheme struct {
								Text      string `xml:",chardata"`
								NilReason string `xml:"nilReason,attr"`
								Nil       string `xml:"nil,attr"`
							} `xml:"transliterationScheme"`
						} `xml:"SpellingOfName"`
					} `xml:"spelling"`
				} `xml:"GeographicalName"`
			} `xml:"geographicalName"`
			ValidFrom struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"validFrom"`
		} `xml:"RailwayLink"`
		RailwayNode struct {
			Text                 string `xml:",chardata"`
			ID                   string `xml:"id,attr"`
			BeginLifespanVersion string `xml:"beginLifespanVersion"` // 2023-04-01T00:00:00Z, 202...
			InspireId            struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text      string `xml:",chardata"`
					LocalId   string `xml:"localId"`   // 131953575, 131953574, 129...
					Namespace string `xml:"namespace"` // nl-top10nl-tn, nl-top10nl...
				} `xml:"Identifier"`
			} `xml:"inspireId"`
			EndLifespanVersion struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"endLifespanVersion"`
			InNetwork struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"inNetwork"`
			Geometry struct {
				Text  string `xml:",chardata"`
				Point struct {
					Text         string `xml:",chardata"`
					SrsName      string `xml:"srsName,attr"`
					SrsDimension string `xml:"srsDimension,attr"`
					Pos          string `xml:"pos"` // 53.02885619 5.64446635, 5...
				} `xml:"Point"`
			} `xml:"geometry"`
			GeographicalName struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"geographicalName"`
			ValidFrom struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"validFrom"`
			FormOfNode struct {
				Text  string `xml:",chardata"`
				Href  string `xml:"href,attr"`
				Title string `xml:"title,attr"`
			} `xml:"formOfNode"`
		} `xml:"RailwayNode"`
		RailwayStationNode struct {
			Text                 string `xml:",chardata"`
			ID                   string `xml:"id,attr"`
			BeginLifespanVersion string `xml:"beginLifespanVersion"` // 2021-06-01T00:00:00Z, 202...
			InspireId            struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text      string `xml:",chardata"`
					LocalId   string `xml:"localId"`   // 131435898, 128889210, 131...
					Namespace string `xml:"namespace"` // nl-top10nl-tn, nl-top10nl...
				} `xml:"Identifier"`
			} `xml:"inspireId"`
			EndLifespanVersion struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"endLifespanVersion"`
			InNetwork struct {
				Text string `xml:",chardata"`
				Nil  string `xml:"nil,attr"`
			} `xml:"inNetwork"`
			Geometry struct {
				Text  string `xml:",chardata"`
				Point struct {
					Text         string `xml:",chardata"`
					SrsName      string `xml:"srsName,attr"`
					SrsDimension string `xml:"srsDimension,attr"`
					Pos          string `xml:"pos"` // 53.33382006 6.92387040, 5...
				} `xml:"Point"`
			} `xml:"geometry"`
			GeographicalName struct {
				Text             string `xml:",chardata"`
				GeographicalName struct {
					Text       string `xml:",chardata"`
					Language   string `xml:"language"` // nld, nld, nld, nld, nld, ...
					Nativeness struct {
						Text  string `xml:",chardata"`
						Href  string `xml:"href,attr"`
						Title string `xml:"title,attr"`
					} `xml:"nativeness"`
					NameStatus struct {
						Text  string `xml:",chardata"`
						Href  string `xml:"href,attr"`
						Title string `xml:"title,attr"`
					} `xml:"nameStatus"`
					SourceOfName  string `xml:"sourceOfName"` // Gazetteer, Gazetteer, Gaz...
					Pronunciation struct {
						Text      string `xml:",chardata"`
						NilReason string `xml:"nilReason,attr"`
						Nil       string `xml:"nil,attr"`
					} `xml:"pronunciation"`
					Spelling struct {
						Text           string `xml:",chardata"`
						SpellingOfName struct {
							Chardata              string `xml:",chardata"`
							Text                  string `xml:"text"`   // Delfzijl, Amersfoort Cent...
							Script                string `xml:"script"` // Latn, Latn, Latn, Latn, L...
							TransliterationScheme struct {
								Text      string `xml:",chardata"`
								NilReason string `xml:"nilReason,attr"`
								Nil       string `xml:"nil,attr"`
							} `xml:"transliterationScheme"`
						} `xml:"SpellingOfName"`
					} `xml:"spelling"`
				} `xml:"GeographicalName"`
			} `xml:"geographicalName"`
			ValidFrom struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"validFrom"`
			FormOfNode struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"formOfNode"`
			NumberOfPlatforms struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"numberOfPlatforms"`
		} `xml:"RailwayStationNode"`
		RailwayType struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"id,attr"`
			NetworkRef struct {
				Text          string `xml:",chardata"`
				LinkReference struct {
					Text    string `xml:",chardata"`
					Element struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"element"`
					ApplicableDirection struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
					} `xml:"applicableDirection"`
				} `xml:"LinkReference"`
			} `xml:"networkRef"`
			InspireId struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text      string `xml:",chardata"`
					LocalId   string `xml:"localId"`   // rt-117911911, rt-11791199...
					Namespace string `xml:"namespace"` // nl-top10nl-tn, nl-top10nl...
				} `xml:"Identifier"`
			} `xml:"inspireId"`
			BeginLifespanVersion string `xml:"beginLifespanVersion"` // 2019-02-01T00:00:00Z, 201...
			EndLifespanVersion   struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"endLifespanVersion"`
			ValidFrom struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"validFrom"`
			Type struct {
				Text  string `xml:",chardata"`
				Href  string `xml:"href,attr"`
				Title string `xml:"title,attr"`
			} `xml:"type"`
		} `xml:"RailwayType"`
		RailwayUse struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"id,attr"`
			NetworkRef struct {
				Text          string `xml:",chardata"`
				LinkReference struct {
					Text    string `xml:",chardata"`
					Element struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"element"`
					ApplicableDirection struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
					} `xml:"applicableDirection"`
				} `xml:"LinkReference"`
			} `xml:"networkRef"`
			InspireId struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text      string `xml:",chardata"`
					LocalId   string `xml:"localId"`   // ru-117911911, ru-11791199...
					Namespace string `xml:"namespace"` // nl-top10nl-tn, nl-top10nl...
				} `xml:"Identifier"`
			} `xml:"inspireId"`
			BeginLifespanVersion string `xml:"beginLifespanVersion"` // 2019-02-01T00:00:00Z, 201...
			EndLifespanVersion   struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"endLifespanVersion"`
			ValidFrom struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"validFrom"`
			Use struct {
				Text  string `xml:",chardata"`
				Href  string `xml:"href,attr"`
				Title string `xml:"title,attr"`
			} `xml:"use"`
		} `xml:"RailwayUse"`
		VerticalPosition struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"id,attr"`
			NetworkRef struct {
				Text          string `xml:",chardata"`
				LinkReference struct {
					Text    string `xml:",chardata"`
					Element struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"element"`
					ApplicableDirection struct {
						Text string `xml:",chardata"`
						Nil  string `xml:"nil,attr"`
					} `xml:"applicableDirection"`
				} `xml:"LinkReference"`
			} `xml:"networkRef"`
			InspireId struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text      string `xml:",chardata"`
					LocalId   string `xml:"localId"`   // vp-117911911, vp-11791199...
					Namespace string `xml:"namespace"` // nl-top10nl-tn, nl-top10nl...
				} `xml:"Identifier"`
			} `xml:"inspireId"`
			BeginLifespanVersion string `xml:"beginLifespanVersion"` // 2019-02-01T00:00:00Z, 201...
			EndLifespanVersion   struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"endLifespanVersion"`
			ValidFrom struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"validFrom"`
			VerticalPosition string `xml:"verticalPosition"` // onGroundSurface, onGround...
		} `xml:"VerticalPosition"`
	} `xml:"featureMember"`
} 


type RailwayLink struct {
			Text                 string `xml:",chardata"`
			ID                   string `xml:"id,attr"`
			BeginLifespanVersion string `xml:"beginLifespanVersion"` // 2019-02-01T00:00:00Z, 201...
			InspireId            struct {
				Text       string `xml:",chardata"`
				Identifier struct {
					Text      string `xml:",chardata"`
					LocalId   string `xml:"localId"`   // rl-117911911, rl-11791199...
					Namespace string `xml:"namespace"` // nl-top10nl-tn, nl-top10nl...
				} `xml:"Identifier"`
			} `xml:"inspireId"`
			EndLifespanVersion struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"endLifespanVersion"`
			InNetwork struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"inNetwork"`
			CentrelineGeometry struct {
				Text       string `xml:",chardata"`
				LineString struct {
					Text         string `xml:",chardata"`
					SrsName      string `xml:"srsName,attr"`
					SrsDimension string `xml:"srsDimension,attr"`
					PosList      string `xml:"posList"` // 53.41789240 6.75796325 53...
				} `xml:"LineString"`
			} `xml:"centrelineGeometry"`
			Fictitious       string `xml:"fictitious"` // false, false, false, fals...
			GeographicalName struct {
				Text             string `xml:",chardata"`
				NilReason        string `xml:"nilReason,attr"`
				Nil              string `xml:"nil,attr"`
				GeographicalName struct {
					Text       string `xml:",chardata"`
					Language   string `xml:"language"` // nld, nld, nld, nld, nld, ...
					Nativeness struct {
						Text  string `xml:",chardata"`
						Href  string `xml:"href,attr"`
						Title string `xml:"title,attr"`
					} `xml:"nativeness"`
					NameStatus struct {
						Text  string `xml:",chardata"`
						Href  string `xml:"href,attr"`
						Title string `xml:"title,attr"`
					} `xml:"nameStatus"`
					SourceOfName  string `xml:"sourceOfName"` // Gazetteer, Gazetteer, Gaz...
					Pronunciation struct {
						Text      string `xml:",chardata"`
						NilReason string `xml:"nilReason,attr"`
						Nil       string `xml:"nil,attr"`
					} `xml:"pronunciation"`
					Spelling struct {
						Text           string `xml:",chardata"`
						SpellingOfName struct {
							Chardata              string `xml:",chardata"`
							Text                  string `xml:"text"`   // Zwijndrechtse Brug, Zwijn...
							Script                string `xml:"script"` // Latn, Latn, Latn, Latn, L...
							TransliterationScheme struct {
								Text      string `xml:",chardata"`
								NilReason string `xml:"nilReason,attr"`
								Nil       string `xml:"nil,attr"`
							} `xml:"transliterationScheme"`
						} `xml:"SpellingOfName"`
					} `xml:"spelling"`
				} `xml:"GeographicalName"`
			} `xml:"geographicalName"`
			ValidFrom struct {
				Text      string `xml:",chardata"`
				NilReason string `xml:"nilReason,attr"`
				Nil       string `xml:"nil,attr"`
			} `xml:"validFrom"`
		}