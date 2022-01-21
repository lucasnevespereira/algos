## Binary Search

A binary search looks for an element in a sorted list. If this element is present it returns it's position, else it returns null.

### Pratical example

Let's say you need to guess a number between <b>0</b> and <b>100</b>.

#### ❌ Wrong way to guess this or Simple Search

You could just ask for every number untill finding the right one

<i>eg: Is it 1? Is it 2? Is it 3? ... Is it 99?</i>

The cost of this would be very time consumming, the complexity of this is <b>O(n)</b> meaning that in the worst case for <b>100</b> numbers you need <b>100</b> guesses. This is <b>O(100)</b>.


#### ✔️ Better way to guess this or Binary Search

Better way to do this is start at the middle. The answer would be either <b>too high</b> or <b>too low</b>.

<i>eg: Is it 50? No, to low.. </i>
<i> -> This means that the numbers <b>1</b>
 to <b>50</b> are not my guess.

This eliminates half of the possibilities. The cost of this is less time consuming. You are no longer looking at O(n) but at <b>log2(n)</b>.In the the worst case for <b>100</b> numbers you need <b>log2(100)</b> guesses which is <b>30</b>.

<i>Reminder: log2(100) === 30 because 2^30 === 100.</i>

### Pseudo code
```
myList
itemToFind

lowest = 0
highest = size of myList

WHILE lowest is SMALLER or EQUAL to highest
(meaning WHILE we have more than one number)

// start by checking middle position
middle = (lowest + highest) / 2

myGuess = list[middle]

IF myGuess EQUALS itemToFind
    return middle
IF myGuess BIGGER THAN itemToFind
    highest = middle - 1
ELSE
    lowest = middle + 1
```