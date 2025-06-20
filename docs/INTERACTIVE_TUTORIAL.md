# Interactive Chorlang Tutorial

Follow along by creating and running each example!

## Lesson 1: Your First Steps 👋

**Goal**: Print a greeting

Create `lesson1.chore`:
```chorelang
spin print("Welcome to Chorlang!")
spin print("Let's dance with code! 💃")
```

**Run it**: `./chorelang -r lesson1.chore`

**Try it yourself**: 
- Change the messages
- Add more print statements
- Try printing numbers: `spin print(42)`

---

## Lesson 2: Variables are Dancers 🕺

**Goal**: Work with variables

Create `lesson2.chore`:
```chorelang
// Introduce our dancers
dance firstName = "Grace"
dance lastName = "Hopper"
dance age = 85

// They can change!
age = age + 1

spin print(firstName, lastName, "is", age, "years old")
```

**Run it**: `./chorelang -r lesson2.chore`

**Exercises**:
1. Add a `dance city = "New York"` and print it
2. Calculate `dance birthYear = 2024 - age`
3. Try reassigning firstName (notice: no 'dance' needed!)

---

## Lesson 3: Loops are Choreography 🔄

**Goal**: Master the `sway` loop

Create `lesson3.chore`:
```chorelang
spin print("5... 4... 3... 2... 1... DANCE!")

// Countdown
sway count from 5 to 1 {
    spin print(count, "...")
}
spin print("Let's dance! 🎉")

// Sum numbers
dance total = 0
sway num from 1 to 10 {
    total = total + num
}
spin print("Sum of 1-10 is:", total)
```

**Run it**: `./chorelang -r lesson3.chore`

**Challenges**:
1. Print even numbers from 2 to 20
2. Calculate factorial of 5
3. Create a multiplication table for 7

---

## Lesson 4: Decisions on the Dance Floor 🤔

**Goal**: Use conditionals

Create `lesson4.chore`:
```chorelang
dance score = 85
dance grade = ""

if score >= 90 {
    grade = "A"
} else if score >= 80 {
    grade = "B"
} else if score >= 70 {
    grade = "C"
} else {
    grade = "F"
}

spin print("Score:", score, "Grade:", grade)

// Nested conditions
dance age = 25
dance hasTicket = true

if age >= 18 {
    if hasTicket {
        spin print("Welcome to the dance!")
    } else {
        spin print("Please buy a ticket")
    }
} else {
    spin print("Sorry, 18+ only")
}
```

**Run it**: `./chorelang -r lesson4.chore`

**Tasks**:
1. Add more grade levels (A+, A-, B+, etc.)
2. Create a "FizzBuzz" using sway and if
3. Check if a number is positive, negative, or zero

---

## Lesson 5: Concurrent Dancers 🎭

**Goal**: Launch goroutines

Create `lesson5.chore`:
```chorelang
// Solo performance
spin print("Main dancer starts")

// Backup dancers join
start {
    sway i from 1 to 3 {
        spin print("Backup dancer A - move", i)
    }
}

start {
    sway i from 1 to 3 {
        spin print("Backup dancer B - move", i)
    }
}

// Main continues
spin print("Main dancer finishes")
```

**Run it**: `./chorelang -r lesson5.chore`

**Note**: Output may vary due to concurrency!

**Experiments**:
1. Add more backup dancers
2. Make dancers count at different speeds
3. Add delays between prints (simulate work)

---

## Lesson 6: Channel Communications 📡

**Goal**: Coordinate with channels

Create `lesson6.chore`:
```chorelang
// Create a communication channel
flow messages = flow channel<string>

// Sender routine
start {
    send messages <- "Hello"
    send messages <- "from"
    send messages <- "Chorlang!"
}

// Receiver (main)
dance msg1 = <-messages
dance msg2 = <-messages  
dance msg3 = <-messages

spin print(msg1, msg2, msg3)
```

**Run it**: `./chorelang -r lesson6.chore`

**Projects**:
1. Send numbers and calculate their sum
2. Create a ping-pong pattern between two goroutines
3. Build a simple work queue

---

## Lesson 7: Pattern Matching Elegance 🎯

**Goal**: Use match expressions

Create `lesson7.chore`:
```chorelang
// Day of week checker
dance day = "Monday"
dance activity = match day {
    when "Monday": flow "Start fresh!"
    when "Friday": flow "TGIF!"
    when "Saturday": flow "Party time!"
    when "Sunday": flow "Rest day"
}
spin print("It's", day, "-", activity)

// Number classifier
dance num = 15
dance numType = match num {
    when 0: flow "Zero"
    when 1: flow "Unity"
    when 2: flow "Pair"
    when 3: flow "Trinity"
}
spin print(num, "is:", numType)
```

**Run it**: `./chorelang -r lesson7.chore`

**Ideas**:
1. Create a simple calculator with match
2. Build a menu system
3. Classify grades using match instead of if/else

---

## Final Project: Dance Competition 🏆

**Goal**: Combine everything!

Create `competition.chore`:
```chorelang
// Dance competition with concurrent judges
flow scores = flow channel<int>

// Three judges score concurrently
start {
    dance score = 8
    spin print("Judge 1 scores:", score)
    send scores <- score
}

start {
    dance score = 9
    spin print("Judge 2 scores:", score)
    send scores <- score
}

start {
    dance score = 7
    spin print("Judge 3 scores:", score)
    send scores <- score
}

// Calculate average
dance total = 0
sway i from 1 to 3 {
    dance score = <-scores
    total = total + score
}

dance average = total / 3
spin print("Average score:", average)

// Determine placement
dance placement = match average {
    when 10: flow "Perfect! Gold medal! 🥇"
    when 9: flow "Excellent! Gold medal! 🥇"
    when 8: flow "Great! Silver medal! 🥈"
    when 7: flow "Good! Bronze medal! 🥉"
}

spin print("Result:", placement)
```

**Run it**: `./chorelang -r competition.chore`

**Extensions**:
1. Add more judges
2. Handle different dance styles
3. Create a tournament bracket
4. Add dancer names and track multiple performers

---

## Congratulations! 🎉

You've learned:
- ✅ Variables with `dance`
- ✅ Loops with `sway`
- ✅ Function calls with `spin`
- ✅ Conditionals with `if/else`
- ✅ Concurrency with `start`
- ✅ Channels with `flow` and `send`
- ✅ Pattern matching with `match/when`

**What's next?**
- Build a real project
- Explore the examples directory
- Share your Chorlang creations
- Keep dancing with code! 💃🕺

---

*Remember: In Chorlang, every bug is just a misstep in the dance, and every successful compilation is a standing ovation!*