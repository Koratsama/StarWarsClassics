# StarWarsClassics
<!-- First Created Feb 1st, 2022
     _______.___________.    ___      .______         ____    __    ____  ___      .______          _______.
    /       |           |   /   \     |   _  \        \   \  /  \  /   / /   \     |   _  \        /       |
   |   (----`---|  |----`  /  ^  \    |  |_)  |        \   \/    \/   / /  ^  \    |  |_)  |      |   (----`
    \   \       |  |      /  /_\  \   |      /          \            / /  /_\  \   |      /        \   \    
.----)   |      |  |     /  _____  \  |  |\  \----.      \    /\    / /  _____  \  |  |\  \----.----)   |   
|_______/       |__|    /__/     \__\ | _| `._____|       \__/  \__/ /__/     \__\ | _| `._____|_______/    
           ______  __          ___           _______.     _______. __    ______     _______.                
          /      ||  |        /   \         /       |    /       ||  |  /      |   /       |                
         |  ,----'|  |       /  ^  \       |   (----`   |   (----`|  | |  ,----'  |   (----`                
         |  |     |  |      /  /_\  \       \   \        \   \    |  | |  |        \   \                    
         |  `----.|  `----./  _____  \  .----)   |   .----)   |   |  | |  `----.----)   |                   
          \______||_______/__/     \__\ |_______/    |_______/    |__|  \______|_______/                                   
-->
<!-- First Created Feb 1st, 2022 -->
this is a simple go application that let's you play 3 different types of games from the Star Wars universe!
Choose between Sabacc, Corellian Spike, and Coruscant Shift.
Right now the goal is to get a basic menu working. The games are not available to play yet as they still need to be worked on.

<details><summary>Sabacc (classic 76-card variant)</summary>
<p>

no content yet.

</p>
</details>

---

<details><summary>Corellian Spike (Black Spire Outpost Rules)</summary>
<p>

### The Deck

1. 30 green cards with positive values 1 through 10, three each ●, ■, ▲
2. 30 red cards with negative values -1 through -10, three each ●, ■, ▲
3. 2 zero-value cards known as sylops (Old Corellian for "idiots")
    
The three suits, known as "Staves", are shown by the shape of the pips on the cards: ● (Circle), ■ (Square), ▲ (Triangle).
These do not have any bearing on the gameplay. These are used in Coruscant Shift. 

### Setup

Choose a dealer for the first game. Hand the deck and dice to the dealer. Each subsequent game, the persont to the left of
the dealer will be the new dealer.

### Ante

Players must pay in to play the roun, 1 credit into the game pot and 2 credit into the sabacc pot. If a player cannot afford 
to pay in, that player is out of the game.

### Dealer

1. Shuffle the deck, then deal two cards fave down to each player. Players can look at their cards, but must not show other players.
2. Place the remaining deck face down in the center of play. This is the draw pile.
3. Place the top card of the draw pile face-up on the table. This is the discard pile.

### Gameplay

Each game is played in three rounds. Each round consists of a turn phase, a betting phase, and a spike dice phase.
The goal of the game is to have the best hand with a total value of zero, which is called Sabacc, or if no player 
gets Sabacc, then the hand with the closest total value to zero wins, which is called Nulrhek. There is also a 
hierarchy of special named Sabacc hands that can be obtained. When scoring, a positive value beats an equal 
negative value.

### 1. Turn phase

Play starts with the player to the dealer's left and continues going left around the table.
One your turn, you can stand, gain, or swap. You may only choose one.
    
1. **Stand** - Retain your current hand for this round and end your turn.
2. **Gain** - To gain a card from the draw pile, you have two options:
    - Take the top card from the draw pile.
    - Discard a card from your hand, then take the top card from the draw pile.
3. **Swap** - Take the top card from the discard pile and add it to your hand, then discard a different card from your hand.

After you stand, gain, or swap your turn is over. Once all players have had a turn, betting begins.

### 2. Betting phase

Players calculate their current hand value and place their bets, beginning with the player to  the dealer's left. Players can check, bet, call, raise, or junk. The cycle continues until all bets are equal.

1. **Check**: The player stays in the game, but wagers no credits. This can only be done if no bet has been made yet.
Id any player places a bet, all players must call or raise in order to stay in.
2. **Bet**: The player makes a wager and adds it to the game pot. All other players must pay the same amount into 
the game pot (call) to stay in the round. Otherwise, they can raise the bet or junk. 
3. **Call**: The player matches the highest bet placed so far by paying that amount into the game pot. When 
verbally declared it is often said as "I'll see your bet."
4. **Raise**: The player raises the highest bet so far by betting a higher amount and paying that amount into the game pot.
All players must now pay this amount to stay in the round, and those who bey prior to the raise must pay the difference
between their current bet and the current highest raise when the betting cycle comes back to them in order to stay in.
Betting goes in cycle until all players bet the same amount. A player can raise the bet only if they were not the last one
to raise in this betting phase and can't raise if they started the betting and no one else raise.
5. **Junk**: The player shuffles his hand (to randomize card order) and puts it on the discard pile. This player 
forfeits all winnings for the round and cannot play until the next round.

Once all players have equaled the highest bet or junked (or all players checked), move on to the spike dice phase.

### 3. Spike Dice Phase

The dealer rolls both dice. If the symbols are different, the round ends. If the symbols match, everyone must note the number 
of cards in their hand, reveal them, then place all the cards onto the discard pile, The dealer then deals each player the same
number of cards they discarded. The player to the left of the dealer is always first to receive new cards. Deal each player all 
the needed cards before dealing to the next player. Then place the rest of the deck face down on the table to form a new draw 
pile and flip over the top card onto the discard pile to start it with a new random value. If at any point the draw pile is 
depleted, shuffle the discard pile to create a new one, then turn over the top card to start a new discard pile as usual.

Repeat phases 1-3 for two more rounds to complete a single game.

### Showdown

At the end of the third round, players reveal their cards in order of play starting at the dealer's left.
Hands are compated with respect to the scoring hierarchy and a winner is determined.
- The winning player collects the game pot.
- If a player wins with Sabacc (any hand totaling zero) that player also collects the Sabacc pot. Otherwise, the credits in the Sabacc pot for the next game.

To begin a new game, the player to the left of the dealer collects the deck and dice, shuffles the cards, and is the new 
game's dealers. Players pay the ante for the new game and games begins.

</p>
</details>

---

<details><summary>Coruscant Shift</summary>
<p>

no content yet.

</p>
</details>

---