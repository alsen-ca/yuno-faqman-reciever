// Entrypoint for main app's Algorithm

// Listens to /ask/:query,:amount
const query = query
const amount = amount

// 'query' refers to a 'question' from model Qa that the clients wants to search for.
// Instead of putting the whole question word for word, this endpoint simply expects some keywords.

// Uses those keywords and uses some algorithm.
// If the word is present on 'tag', the weight assigned to that word is *4
// If the word is present on 'thema', the weight assigned to that word is *5
// If it finds the word in 'quesrion' of Qa, it also takes the weight of that word into consideration
// Uses some kind of algorithm to determine to search list
// The list must not only take the "total weight obtained" into consideration - also in relation to the total amount of 
// words that it found and how importat they are

// List of 10 most probable answers the algorithm found
let answers = nil

// Save only the first x amount of answers from that list. With x being the 'amount' given by the /ask
answers = filter_answers(answers, amount)

return answers

func filter_answers(answers: list, amount: int) {
	answers = answers[0,amount-1]
	
	// Returns the new list of answers
	return answers
}