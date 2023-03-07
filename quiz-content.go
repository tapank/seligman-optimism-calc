package main

const INSTRUCTIONS = `
Take as much time as you need to answer each of the questions. On average the
test takes about fifteen minutes. There are no right or wrong answers. Read the
description of each situation and vividly imagine it happening to you. You have
probably not experienced some of the situations, but that doesn't matter.
Perhaps neither reponse will seem to fit; go ahead anyway and select A or B
(type a or b on the keyboard), choosing the cause likelier to apply to you. You
may not like the way some of the responses sound, but don't choose what you
think you *should* say or what would sound right to other people; choose the
response your'd be likelier to have.

Press any key to start.`

var adultQs = []Question{
	{
		// question 1
		q: "The project you are in charge of is a great success.",
		options: []Option{
			{
				txt: "I kept a close watch over everyone's work.",
				agg: PsG,
			},
			{
				txt: "Everyone devoted a lot of time and energy to it.",
				agg: NIL,
			},
		},
	},
	{
		// question 2
		q: "You and your spouse (boyfriend/girlfriend) make up after a fight.",
		options: []Option{
			{
				txt: "I forgave him/her.",
				agg: NIL,
			},
			{
				txt: "I'm usually forgiving.",
				agg: PmG,
			},
		},
	},
	{
		// question 3
		q: "You get lost driving to a friend's house.",
		options: []Option{
			{
				txt: "I missed a turn.",
				agg: PsB,
			},
			{
				txt: "My friend gave me bad directions.",
				agg: NIL,
			},
		},
	},
	{
		// question 4
		q: "Your spouse surprises you with a gift.",
		options: []Option{
			{
				txt: "He/she just got a raise at work.",
				agg: NIL,
			},
			{
				txt: "I took him/her out to a special dinner the night before.",
				agg: PsG,
			},
		},
	},
	{
		// question 5
		q: "You forget your spouse's birthday.",
		options: []Option{
			{
				txt: "I'm not good at remembering birthdays.",
				agg: PmB,
			},
			{
				txt: "I was preoccupied with other things.",
				agg: NIL,
			},
		},
	},
	{
		// question 6
		q: "You get a flower from a secret admirer.",
		options: []Option{
			{
				txt: "I am attractive to him/her.",
				agg: NIL,
			},
			{
				txt: "I am a popular person.",
				agg: PvG,
			},
		},
	},
	{
		// question 7
		q: "You run for a community office position and you win.",
		options: []Option{
			{
				txt: "I devote a lot of time and energy to campaigning.",
				agg: NIL,
			},
			{
				txt: "I work very hard at everything I do.",
				agg: PvG,
			},
		},
	},
	{
		// question 8
		q: "You miss an important engagement.",
		options: []Option{
			{
				txt: "Sometimes my memory fails me.",
				agg: PvB,
			},
			{
				txt: "I sometimes forget to check my appointment book.",
				agg: NIL,
			},
		},
	},
	{
		// question 9
		q: "You run for a community office position and you lose.",
		options: []Option{
			{
				txt: "I didn't campaign hard enough.",
				agg: PsB,
			},
			{
				txt: "The person who won knew more people.",
				agg: NIL,
			},
		},
	},
	{
		// question 10
		q: "You host a successful dinner.",
		options: []Option{
			{
				txt: "I was particularly charming that night.",
				agg: NIL,
			},
			{
				txt: "I am a good host.",
				agg: PmG,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
	{
		// question x
		q: "",
		options: []Option{
			{
				txt: "",
				agg: NIL,
			},
			{
				txt: "",
				agg: NIL,
			},
		},
	},
}

var childQs = []Question{
	{
		q: "CH The project you are in charge of is a great success.",
		options: []Option{
			{
				txt: "I kept a close watch over everyone's work.",
				agg: PsG,
			},
			{
				txt: "Everyone devoted a lot of time and energy to it.",
				agg: NIL,
			},
		},
	},
	{
		q: "CH You and your spouse (boyfriend/girlfriend) make up after a fight.",
		options: []Option{
			{
				txt: "I forgave him/her.",
				agg: NIL,
			},
			{
				txt: "I'm usually forgiving.",
				agg: PmG,
			},
		},
	},
	{
		q: "CH You get lost driving to a friend's house.",
		options: []Option{
			{
				txt: "I missed a turn.",
				agg: PsB,
			},
			{
				txt: "My friend gave me bad directions.",
				agg: NIL,
			},
		},
	},
	// {
	// 	q: "",
	// 	options: []Option{
	// 		{
	// 			txt: "",
	// 			agg: NIL,
	// 		},
	// 		{
	// 			txt: "",
	// 			agg: NIL,
	// 		},
	// 	},
	// },
}
