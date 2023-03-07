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
		// question 11
		q: "You stop a crime by calling the police.",
		options: []Option{
			{
				txt: "A strange noise caught my attention.",
				agg: NIL,
			},
			{
				txt: "I was alert that day.",
				agg: PsG,
			},
		},
	},
	{
		// question 12
		q: "You were extremely healthy all year.",
		options: []Option{
			{
				txt: "Few people around me were sick, so I wasn't exposed.",
				agg: NIL,
			},
			{
				txt: "I made sure I ate well and got enough rest.",
				agg: PsG,
			},
		},
	},
	{
		// question 13
		q: "You owe the library ten dollars for an overdue book.",
		options: []Option{
			{
				txt: "When I am really involved in what I am reading, I often forget when it's due.",
				agg: PmB,
			},
			{
				txt: "I was so involved in writing the report that I forgot to return the book.",
				agg: NIL,
			},
		},
	},
	{
		// question 14
		q: "Your stocks make you a lot of money.",
		options: []Option{
			{
				txt: "My broker decided to take on something new.",
				agg: NIL,
			},
			{
				txt: "My broker is a top-notch investor.",
				agg: PmG,
			},
		},
	},
	{
		// question 15
		q: "You win an athletic contest.",
		options: []Option{
			{
				txt: "I was feeling unbeatable.",
				agg: NIL,
			},
			{
				txt: "I train hard.",
				agg: PmG,
			},
		},
	},
	{
		// question 16
		q: "You fail an important examination.",
		options: []Option{
			{
				txt: "I wasn't as smart as the other people taking the exam.",
				agg: PvB,
			},
			{
				txt: "I didn't prepare for it well.",
				agg: NIL,
			},
		},
	},
	{
		// question 17
		q: "You prepared a special meal for a friend and he/she barely touched the food.",
		options: []Option{
			{
				txt: "I wasn't a good cook.",
				agg: PvB,
			},
			{
				txt: "I made the meal in a rush.",
				agg: NIL,
			},
		},
	},
	{
		// question 18
		q: "You lose a sporting event for which you have been training for a long time.",
		options: []Option{
			{
				txt: "I'm not very athletic.",
				agg: PvB,
			},
			{
				txt: "I'm not good at that sport.",
				agg: NIL,
			},
		},
	},
	{
		// question 19
		q: "Your car runs out of gas on a dark street late at night.",
		options: []Option{
			{
				txt: "I din't check to see how much gas was in the tank.",
				agg: PsB,
			},
			{
				txt: "The gas gauge was broken.",
				agg: NIL,
			},
		},
	},
	{
		// question 20
		q: "You lose your temper with a friend.",
		options: []Option{
			{
				txt: "He/she is always nagging me.",
				agg: PmB,
			},
			{
				txt: "He/she was in a hostile mood.",
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
