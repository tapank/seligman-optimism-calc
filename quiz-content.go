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
		// question 21
		q: "You are penalized for not returning your income-tax forms on time.",
		options: []Option{
			{
				txt: "I always put off doing my taxes.",
				agg: PmB,
			},
			{
				txt: "I was lazy about getting my taxes done this year.",
				agg: NIL,
			},
		},
	},
	{
		// question 22
		q: "You ask a person out on a date and he/she says no.",
		options: []Option{
			{
				txt: "I was a wreck that day.",
				agg: PvB,
			},
			{
				txt: "I got tongue-tied when I asked him/her on the date.",
				agg: NIL,
			},
		},
	},
	{
		// question 23
		q: "A game-show host picks you out of the audience to participate in the show.",
		options: []Option{
			{
				txt: "I was sitting in the right seat.",
				agg: NIL,
			},
			{
				txt: "I looked the most enthusiastic.",
				agg: PsG,
			},
		},
	},
	{
		// question 24
		q: "You are frequently asked to dance at a party.",
		options: []Option{
			{
				txt: "I am outgoing at parties.",
				agg: PmG,
			},
			{
				txt: "I was in perfect form that night.",
				agg: NIL,
			},
		},
	},
	{
		// question 25
		q: "You buy your spouse a gift and he/she doesn't like it.",
		options: []Option{
			{
				txt: "I don't put enough thought into things like that.",
				agg: PsB,
			},
			{
				txt: "He/she has very picky tastes.",
				agg: NIL,
			},
		},
	},
	{
		// question 26
		q: "You do exceptionally well in a job interview.",
		options: []Option{
			{
				txt: "I felt extremely confident during the interview.",
				agg: NIL,
			},
			{
				txt: "I interview well.",
				agg: PmG,
			},
		},
	},
	{
		// question 27
		q: "You tell a joke and everyone laughs.",
		options: []Option{
			{
				txt: "The joke was funny.",
				agg: NIL,
			},
			{
				txt: "My timing was perfect.",
				agg: PsG,
			},
		},
	},
	{
		// question 28
		q: "Your boss gives you too little time in which to finish a project,\n\tbut you get finished anyway.",
		options: []Option{
			{
				txt: "I am good at my job.",
				agg: NIL,
			},
			{
				txt: "I am an efficient person.",
				agg: PvG,
			},
		},
	},
	{
		// question 29
		q: "You've been feeling run-down lately.",
		options: []Option{
			{
				txt: "I never get a chance to relax.",
				agg: PmB,
			},
			{
				txt: "I was exceptionally busy this week.",
				agg: NIL,
			},
		},
	},
	{
		// question 30
		q: "You ask someone to dance and he/she says no.",
		options: []Option{
			{
				txt: "I am not a good enough dancer.",
				agg: PsB,
			},
			{
				txt: "He/she doesn't like to dance.",
				agg: NIL,
			},
		},
	},
	{
		// question 31
		q: "You save a person from choking to death.",
		options: []Option{
			{
				txt: "I know a technique to stop someone from choking.",
				agg: NIL,
			},
			{
				txt: "I know what to do in crisis situations.",
				agg: PvG,
			},
		},
	},
	{
		// question 32
		q: "Your romantic partner wants to cool things off for a while.",
		options: []Option{
			{
				txt: "I'm too self-centered.",
				agg: PvB,
			},
			{
				txt: "I don't spend enough time with him/her.",
				agg: NIL,
			},
		},
	},
	{
		// question 33
		q: "A friend says something that hurts your feelings.",
		options: []Option{
			{
				txt: "She always blurts things out without thinking of others.",
				agg: PmB,
			},
			{
				txt: "My friend was in a bad mood and took it out on me.",
				agg: NIL,
			},
		},
	},
	{
		// question 34
		q: "Your employer comes to you for advice.",
		options: []Option{
			{
				txt: "I am an expert in the area about which I was asked.",
				agg: NIL,
			},
			{
				txt: "I am good at giving useful advice.",
				agg: PvG,
			},
		},
	},
	{
		// question 35
		q: "A friend thanks you for helping him/her get through a bad time.",
		options: []Option{
			{
				txt: "I enjoy helping him/her through tough times.",
				agg: NIL,
			},
			{
				txt: "I care about people.",
				agg: PvG,
			},
		},
	},
	{
		// question 36
		q: "You have a wonderful time at a party.",
		options: []Option{
			{
				txt: "Everyone was friendly.",
				agg: NIL,
			},
			{
				txt: "I was friendly.",
				agg: PsG,
			},
		},
	},
	{
		// question 37
		q: "Your doctor tells you that you are in good physical shape.",
		options: []Option{
			{
				txt: "I make sure I exercise frequently.",
				agg: NIL,
			},
			{
				txt: "I am very health-conscious.",
				agg: PvG,
			},
		},
	},
	{
		// question 38
		q: "Your spouse takes you away for a romantic weekend.",
		options: []Option{
			{
				txt: "He/she needed to get away for a few days.",
				agg: NIL,
			},
			{
				txt: "He/she likes to explore new areas.",
				agg: PmG,
			},
		},
	},
	{
		// question 39
		q: "Your doctor tells you that you eat too much sugar.",
		options: []Option{
			{
				txt: "I don't pay much attention to my diet.",
				agg: PsB,
			},
			{
				txt: "You can't avoid sugar, it's in everything.",
				agg: NIL,
			},
		},
	},
	{
		// question 40
		q: "You are asked to head an important project.",
		options: []Option{
			{
				txt: "I just successfully completed a similar project.",
				agg: NIL,
			},
			{
				txt: "I am a good supervisor.",
				agg: PmG,
			},
		},
	},
	{
		// question 41
		q: "You and your spouse have been fighting a great deal.",
		options: []Option{
			{
				txt: "I have been feeling cranky and pressured lately.",
				agg: PsB,
			},
			{
				txt: "He/she has been hostile lately.",
				agg: NIL,
			},
		},
	},
	{
		// question 42
		q: "You fall down a great deal while skiing.",
		options: []Option{
			{
				txt: "Skiing is difficult.",
				agg: PmB,
			},
			{
				txt: "The trails were icy.",
				agg: NIL,
			},
		},
	},
	{
		// question 43
		q: "You win a prestigious award.",
		options: []Option{
			{
				txt: "I solved an important problem.",
				agg: NIL,
			},
			{
				txt: "I was the best employee.",
				agg: PvG,
			},
		},
	},
	{
		// question 44
		q: "Your stocks are at an all-time low.",
		options: []Option{
			{
				txt: "I didn't know much about the business climate at the time.",
				agg: PvB,
			},
			{
				txt: "I made a poor choice of stocks.",
				agg: NIL,
			},
		},
	},
	{
		// question 45
		q: "You win the lottery.",
		options: []Option{
			{
				txt: "It was pure chance.",
				agg: NIL,
			},
			{
				txt: "I picked the right numbers.",
				agg: PsG,
			},
		},
	},
	{
		// question 46
		q: "You gain weight over the holidays and you can't lose it.",
		options: []Option{
			{
				txt: "Diets don't work in the long run.",
				agg: PmB,
			},
			{
				txt: "The diet I tried didn't work.",
				agg: NIL,
			},
		},
	},
	{
		// question 47
		q: "You are in the hospital and few people come to visit.",
		options: []Option{
			{
				txt: "I'm irritable when I am sick.",
				agg: PsB,
			},
			{
				txt: "My friends are negligent about things like that.",
				agg: NIL,
			},
		},
	},
	{
		// question 48
		q: "They won't honor your credit card at a store.",
		options: []Option{
			{
				txt: "I sometimes overestimate how much money I have.",
				agg: PvB,
			},
			{
				txt: "I sometimes forget to pay my credit-card bill.",
				agg: NIL,
			},
		},
	},
}

var childQs = []Question{
	{
		// question 1
		q: "You get an A on a test.",
		options: []Option{
			{
				txt: "I am smart.",
				agg: PvG,
			},
			{
				txt: "I am good in the subject that the test was in.",
				agg: NIL,
			},
		},
	},
	{
		// question 2
		q: "You play a game with some friends and you win.",
		options: []Option{
			{
				txt: "The people that I played with did not play the game well.",
				agg: NIL,
			},
			{
				txt: "I play that game well.",
				agg: PsG,
			},
		},
	},
	{
		// question 3
		q: "You spend a night at a friend's house and you have a good time.",
		options: []Option{
			{
				txt: "My friend was in a friendly mood that night.",
				agg: NIL,
			},
			{
				txt: "Everyone in my friend's family was in a friendly mood that night.",
				agg: PvG,
			},
		},
	},
	{
		// question 4
		q: "You go on a vacation with a group of people and you have fun.",
		options: []Option{
			{
				txt: "I was in a good mood.",
				agg: PsG,
			},
			{
				txt: "The people I was with were in good moods.",
				agg: NIL,
			},
		},
	},
	{
		// question 5
		q: "All of your friends catch a cold except you.",
		options: []Option{
			{
				txt: "I have been healthy lately.",
				agg: NIL,
			},
			{
				txt: "I am a healthy person.",
				agg: PmG,
			},
		},
	},
	{
		// question 6
		q: "Your pet gets run over by a car.",
		options: []Option{
			{
				txt: "I don't take good care of my pets.",
				agg: PsB,
			},
			{
				txt: "Drivers are not cautious enough.",
				agg: NIL,
			},
		},
	},
	{
		// question 7
		q: "Some kids you know say that they don't like you.",
		options: []Option{
			{
				txt: "Once in a while people are mean to me.",
				agg: NIL,
			},
			{
				txt: "Once in a while I am mean to other people.",
				agg: PsB,
			},
		},
	},
	{
		// question 8
		q: "You get very good grades.",
		options: []Option{
			{
				txt: "Schoolwork is simple.",
				agg: NIL,
			},
			{
				txt: "I am a hard worker.",
				agg: PsG,
			},
		},
	},
	{
		// question 9
		q: "You meet a friend and your friend tells you that you look nice.",
		options: []Option{
			{
				txt: "My friend felt like praising the way people looked that day.",
				agg: NIL,
			},
			{
				txt: "Usually my friend praises the way people look.",
				agg: PmG,
			},
		},
	},
	{
		// question 10
		q: "A good friend tells you that he hates you.",
		options: []Option{
			{
				txt: "My friend was in a bad mood that day.",
				agg: NIL,
			},
			{
				txt: "I wasn't nice to my friend that day.",
				agg: PsB,
			},
		},
	},
	{
		// question 11
		q: "You tell a joke and no one laughs.",
		options: []Option{
			{
				txt: "I don't tell jokes well.",
				agg: PsB,
			},
			{
				txt: "The joke is so well known that it is no longer funny.",
				agg: NIL,
			},
		},
	},
	{
		// question 12
		q: "Your teacher gives a lesson and you don't understand it.",
		options: []Option{
			{
				txt: "I didn't pay attention to anything that day.",
				agg: PvB,
			},
			{
				txt: "I didn't pay attention when my teacher was talking.",
				agg: NIL,
			},
		},
	},
	{
		// question 13
		q: "You fail a test.",
		options: []Option{
			{
				txt: "My teacher makes hard tests.",
				agg: PmB,
			},
			{
				txt: "The past few weeks, my teacher has made hard tests.",
				agg: NIL,
			},
		},
	},
	{
		// question 14
		q: "You gain a lot of weight and start to look fat.",
		options: []Option{
			{
				txt: "The food I have to eat is fattening.",
				agg: NIL,
			},
			{
				txt: "I like fattening foods.",
				agg: PsB,
			},
		},
	},
	{
		// question 15
		q: "A person steals money from you.",
		options: []Option{
			{
				txt: "That person is dishonest.",
				agg: NIL,
			},
			{
				txt: "People are dishonest.",
				agg: PvB,
			},
		},
	},
	{
		// question 16
		q: "Your parents praise something you make.",
		options: []Option{
			{
				txt: "I am good at making some things.",
				agg: PsG,
			},
			{
				txt: "My parents like some things I make.",
				agg: NIL,
			},
		},
	},
	{
		// question 17
		q: "You play a game and you win money.",
		options: []Option{
			{
				txt: "I am a lucky person.",
				agg: PvG,
			},
			{
				txt: "I am lucky when I play games.",
				agg: NIL,
			},
		},
	},
	{
		// question 18
		q: "You almost drown when swimming in a river.",
		options: []Option{
			{
				txt: "I am not a very cautious person.",
				agg: PmB,
			},
			{
				txt: "Some days I am not a cautious person.",
				agg: NIL,
			},
		},
	},
	{
		// question 19
		q: "You are invited to a lot of parties.",
		options: []Option{
			{
				txt: "A lot of people have been acting friendly toward me lately.",
				agg: NIL,
			},
			{
				txt: "I have been acting friendly toward a lot of people lately.",
				agg: PsG,
			},
		},
	},
	{
		// question 20
		q: "A grown-up yells at you.",
		options: []Option{
			{
				txt: "That person yelled at the first person he saw.",
				agg: NIL,
			},
			{
				txt: "That person yelled at a lot of people he saw that day.",
				agg: PvB,
			},
		},
	},
	{
		// question 21
		q: "You do a project with a group of kids and it turns out badly.",
		options: []Option{
			{
				txt: "I don't work well with the people in the group.",
				agg: NIL,
			},
			{
				txt: "I never work well with a group.",
				agg: PvB,
			},
		},
	},
	{
		// question 22
		q: "You make a new friend.",
		options: []Option{
			{
				txt: "I am a nice person.",
				agg: PsG,
			},
			{
				txt: "The people that I meet are nice.",
				agg: NIL,
			},
		},
	},
	{
		// question 23
		q: "You have been getting along well with your family.",
		options: []Option{
			{
				txt: "I am easy to get along with when I am with my family.",
				agg: PmG,
			},
			{
				txt: "Once in a while I am easy to get along with when I am with my family.",
				agg: NIL,
			},
		},
	},
	{
		// question 24
		q: "You try to sell candy, but no one will buy any.",
		options: []Option{
			{
				txt: "Lately a lot of children are selling things, so people don't want to buy anything else from children.",
				agg: NIL,
			},
			{
				txt: "People don't like to buy things from children.",
				agg: PmB,
			},
		},
	},
	{
		// question 25
		q: "You play a game and you win.",
		options: []Option{
			{
				txt: "Sometimes I try as hard as I can at games.",
				agg: NIL,
			},
			{
				txt: "Sometimes I try as hard as I can.",
				agg: PvG,
			},
		},
	},
	{
		// question 26
		q: "You get a bad grade in school.",
		options: []Option{
			{
				txt: "I am stupid.",
				agg: PsB,
			},
			{
				txt: "Teachers are unfair graders.",
				agg: NIL,
			},
		},
	},
	{
		// question 27
		q: "You walk into a door and you get a bloody nose.",
		options: []Option{
			{
				txt: "I wasn't looking where I was going.",
				agg: NIL,
			},
			{
				txt: "I have been careless lately.",
				agg: PvB,
			},
		},
	},
	{
		// question 28
		q: "You miss the ball and your team loses the game.",
		options: []Option{
			{
				txt: "I didn't try hard while playing ball that day.",
				agg: NIL,
			},
			{
				txt: "I usually don't try hard when I am playing ball.",
				agg: PmB,
			},
		},
	},
	{
		// question 29
		q: "You twist your ankle in gym class.",
		options: []Option{
			{
				txt: "The past few weeks, the sports we played in gym class have been dangerous.",
				agg: NIL,
			},
			{
				txt: "The past few weeks I have been clumsy in gym class.",
				agg: PsB,
			},
		},
	},
	{
		// question 30
		q: "Your parents take you to the beach and you have a good time.",
		options: []Option{
			{
				txt: "Everything at the beach was nice that day.",
				agg: PvG,
			},
			{
				txt: "The weather at the beach was nice that day.",
				agg: NIL,
			},
		},
	},
	{
		// question 31
		q: "You take a train which arrives so late that you miss a movie.",
		options: []Option{
			{
				txt: "The past few days there have been problems with the train being on time.",
				agg: NIL,
			},
			{
				txt: "The trains are almost never on time.",
				agg: PmB,
			},
		},
	},
	{
		// question 32
		q: "Your mother makes your favorite dinner for you.",
		options: []Option{
			{
				txt: "There are a few things that my mother does to please me.",
				agg: NIL,
			},
			{
				txt: "My mother likes to please me.",
				agg: PvG,
			},
		},
	},
	{
		// question 33
		q: "The team that you are on loses a game.",
		options: []Option{
			{
				txt: "The team members don't play well together.",
				agg: PmB,
			},
			{
				txt: "That day the team members didn't play well together.",
				agg: NIL,
			},
		},
	},
	{
		// question 34
		q: "You finish your homework quickly.",
		options: []Option{
			{
				txt: "Lately I have been doing everything quickly.",
				agg: PvG,
			},
			{
				txt: "Lately I have been doing schoolwork quickly.",
				agg: NIL,
			},
		},
	},
	{
		// question 35
		q: "Your teacher asks you a question and you give the wrong answer.",
		options: []Option{
			{
				txt: "I get nervous when I have to answer questions.",
				agg: PmB,
			},
			{
				txt: "That day I got nervous when I had to answer questions.",
				agg: NIL,
			},
		},
	},
	{
		// question 36
		q: "You get on the wrong bus and you get lost.",
		options: []Option{
			{
				txt: "That day I wasn't paying attention to what was going on.",
				agg: NIL,
			},
			{
				txt: "I usually don't pay attention to what's going on.",
				agg: PmB,
			},
		},
	},
	{
		// question 37
		q: "You go to an amusement park and you have a good time.",
		options: []Option{
			{
				txt: "I usually enjoy myself at amusement parks.",
				agg: NIL,
			},
			{
				txt: "I usually enjoy myself.",
				agg: PvG,
			},
		},
	},
	{
		// question 38
		q: "An older kid slaps you in the face.",
		options: []Option{
			{
				txt: "I teased his younger brother.",
				agg: PsB,
			},
			{
				txt: "His younger brother told him I had teased him.",
				agg: NIL,
			},
		},
	},
	{
		// question 39
		q: "You get all the toys you want on your birthday.",
		options: []Option{
			{
				txt: "People always guess what toys to buy me for my birthday.",
				agg: PmG,
			},
			{
				txt: "This birthday people guessed right as to what toys I wanted.",
				agg: NIL,
			},
		},
	},
	{
		// question 40
		q: "You take a vacation in the country and you have a wonderful time.",
		options: []Option{
			{
				txt: "The country is a beautiful place to be.",
				agg: PmG,
			},
			{
				txt: "The time of the year that we went was beautiful.",
				agg: NIL,
			},
		},
	},
	{
		// question 41
		q: "Your neighbors ask you over for dinner.",
		options: []Option{
			{
				txt: "Sometimes people are in kind moods.",
				agg: NIL,
			},
			{
				txt: "People are kind.",
				agg: PmG,
			},
		},
	},
	{
		// question 42
		q: "You have a substitute teacher and she likes you.",
		options: []Option{
			{
				txt: "I was well behaved during class that day.",
				agg: NIL,
			},
			{
				txt: "I am almost always well behaved during class.",
				agg: PmG,
			},
		},
	},
	{
		// question 43
		q: "You make your friends happy.",
		options: []Option{
			{
				txt: "I am a fun person to be with.",
				agg: PmG,
			},
			{
				txt: "Sometimes I am a fun person to be with.",
				agg: NIL,
			},
		},
	},
	{
		// question 44
		q: "You get a free ice-cream cone.",
		options: []Option{
			{
				txt: "I was friendly to the ice-cream man that day.",
				agg: PsG,
			},
			{
				txt: "The ice-cream man was feeling friendly that day.",
				agg: NIL,
			},
		},
	},
	{
		// question 45
		q: "At your friend's party the magician asks you to help him out.",
		options: []Option{
			{
				txt: "It was just luck that I got picked.",
				agg: NIL,
			},
			{
				txt: "I looked really interested in what was gonig on.",
				agg: PsG,
			},
		},
	},
	{
		// question 46
		q: "You try to convince a kid to go to the movies with you, but he won't go.",
		options: []Option{
			{
				txt: "That day he did not feel like doing anything.",
				agg: PvB,
			},
			{
				txt: "That day he did not feel like going to the movies.",
				agg: NIL,
			},
		},
	},
	{
		// question 47
		q: "Your parents get a divorce.",
		options: []Option{
			{
				txt: "It is hard for people to get along well when they are married.",
				agg: PvB,
			},
			{
				txt: "It is hard for my parents to get along well when they are married.",
				agg: NIL,
			},
		},
	},
	{
		// question 48
		q: "You have been trying to get into a club and you don't get in.",
		options: []Option{
			{
				txt: "I don't get along well with other people.",
				agg: PvB,
			},
			{
				txt: "I don't get along well with the people in the club.",
				agg: NIL,
			},
		},
	},
}
