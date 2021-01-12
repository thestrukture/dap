package globals

import "github.com/thestrukture/dap/types"

var OpenSource = []string{
	"Agility",
}

var LessRisk = []string{
	"Pace",
}

var Diversity = []string{
	"Diversity",
}

var Groups = []types.Group{
	types.Group{
		Text:   "Part 1",
		Resume: "Answer the following questions as directly as possible. There are no wrong or right answers. The package referred to in the following questions can refer to any package, written in any language.",
		Questions: []types.Question{

			types.Question{
				Text: "Given the task of writing a PDF exporter for a web server would you?",
				Choices: []types.Choice{

					types.Choice{
						Text:      "Find a package capable of exporting a PDF.",
						Weight:    2,
						Interests: OpenSource,
					},
					types.Choice{
						Text:      "Write a package from scratch. (to reduce technical debt)",
						Weight:    1,
						Interests: LessRisk,
					},
					types.Choice{
						Text:      "Use a terminal command to handle export.",
						Weight:    3,
						Interests: OpenSource,
					},
					types.Choice{
						Text:      "Bridge code from a lower level language",
						Weight:    2,
						Interests: Diversity,
					},
					types.Choice{
						Text:      "N/A",
						Weight:    0,
						Interests: Diversity,
					},
				},
			},

			types.Question{
				Text: "How do you choose a programming language?",
				Choices: []types.Choice{

					types.Choice{
						Text:      "Based on level of comfort",
						Weight:    1,
						Interests: LessRisk,
					},
					types.Choice{
						Text:      "Based on utility of language with relation to the problem at hand",
						Weight:    1,
						Interests: Diversity,
					},
					types.Choice{
						Text:      "Based on level of support (Books, online resources & documentation)",
						Weight:    2,
						Interests: OpenSource,
					},
					types.Choice{
						Text:      "N/A",
						Weight:    0,
						Interests: Diversity,
					},
				},
			},

			types.Question{
				Text: "How do you prefer to write your code?",
				Choices: []types.Choice{

					types.Choice{
						Text:      "With clean architecture in mind. This includes adding comments.",
						Weight:    1,
						Interests: LessRisk,
					},
					types.Choice{
						Text:      "With multiple revisions. This includes having multiple drafts.",
						Weight:    2,
						Interests: Diversity,
					},
					types.Choice{
						Text:      "Complete first, then clean it up.",
						Weight:    3,
						Interests: OpenSource,
					},
					types.Choice{
						Text:      "N/A",
						Weight:    0,
						Interests: Diversity,
					},
				},
			},
		},
	},
	types.Group{
		Text:   "Part 2",
		Resume: "Answer the following questions as directly as possible. There are no wrong or right answers.",
		Questions: []types.Question{

			types.Question{
				Text: "What is your favorite communication medium to get new objectives?",
				Choices: []types.Choice{

					types.Choice{
						Text:      "Video conference",
						Weight:    1,
						Interests: LessRisk,
					},
					types.Choice{
						Text:      "Voice conference",
						Weight:    1,
						Interests: Diversity,
					},
					types.Choice{
						Text:      "Documentation",
						Weight:    1,
						Interests: OpenSource,
					},
					types.Choice{
						Text:      "N/A",
						Weight:    0,
						Interests: Diversity,
					},
				},
			},

			types.Question{
				Text: "What is your favorite API documentation format?",
				Choices: []types.Choice{

					types.Choice{
						Text:      "Word/PDF Document",
						Weight:    1,
						Interests: LessRisk,
					},
					types.Choice{
						Text:      "Website or Blog",
						Weight:    1,
						Interests: Diversity,
					},
					types.Choice{
						Text:      "Markdown",
						Weight:    1,
						Interests: OpenSource,
					},
					types.Choice{
						Text:      "N/A",
						Weight:    0,
						Interests: Diversity,
					},
				},
			},

			types.Question{
				Text: "How often would you like to have meetings?",
				Choices: []types.Choice{

					types.Choice{
						Text:      "Bi-weekly",
						Weight:    1,
						Interests: LessRisk,
					},
					types.Choice{
						Text:      "Monthly",
						Weight:    0,
						Interests: Diversity,
					},
					types.Choice{
						Text:      "Weekly",
						Weight:    1,
						Interests: OpenSource,
					},
					types.Choice{
						Text:      "N/A",
						Weight:    0,
						Interests: Diversity,
					},
				},
			},

			types.Question{
				Text: "How often do you go on StackOverflow or other forum websites?",
				Choices: []types.Choice{

					types.Choice{
						Text:      "I don't use any forums.",
						Weight:    0,
						Interests: LessRisk,
					},
					types.Choice{
						Text:      "Sometimes",
						Weight:    0,
						Interests: OpenSource,
					},
					types.Choice{
						Text:      "N/A",
						Weight:    0,
						Interests: Diversity,
					},
				},
			},
		},
	},
}
