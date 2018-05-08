# wikipedia
Reads the paragraph its words as tokens and relates them to questions to match them to answers based on weighted average.

# Approach
* Parses the paragraph and identifies words as tokens (named entities, verbs, pronouns etc). Tokens are typically the special characters, strings, plurals and numbers.
* Parses the questions and identifies the words or the named entities, verbs, pronouns etc as tokens.
* Parses the answers jumbled and separated by delimited ";"
* Compares the question tokens to possible answer tokens to find any occurance of question token in answers. Pluralality both in questions
and answers are considered while its expected for the program to atleast know the list of plural and its equivalent singular words. Dictionary
with token.go to be updated with possible plural named entities within the paragraph for the intelligence to functiona properly
* Finds the offset of answers in paragraph, identifies the question token offset in paragraph and determines the closest proximity of question token offset
with the answer.
* Derives the weighted average of both minimum offset difference of questions to answer and maximum occurance of question tokens in answer text.

# Given
* A paragraph, list of questions and possible answers in jumbled fashion separated by ";"

# Assumption
* Program to be aware of list of plurals used in paragraph upfront in case questions and answers used them orthogonally on same context.
* Set of tokens are stemmed/ignored as part of parsing/answer identifying intelligence to enchance performance and avoid complexity.

