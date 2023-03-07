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
