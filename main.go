package main

import (
	"fmt"
	"math/rand"
	"strings"

)

//dadu only
type Dice struct{
	AtasDadu int
}

func (d *Dice) getAtasDadu() int {
	return d.AtasDadu
}

func (d *Dice) roll() *Dice {
	d.AtasDadu = rand.Intn(6) + 1
	return d
}

func (d *Dice) setAtasDadu(AtasDadu int) *Dice {
	d.AtasDadu = AtasDadu
	return d
}
//Player Only
type Player struct {
	Jumlahdadu []Dice
	Name string
	Position int
	Point int
}
	
func (p *Player) getJumlahdadu() []Dice {
	return p.Jumlahdadu
}
	
func (p *Player) getName() string {
	return p.Name
}
	
func (p *Player) getPosition() int {
	return p.Position
}
func (p *Player) getPoint() int {
	return p.Point
}

//newplayae dengan jumlah dadu

//Constructor
// constructor menambahkan jumlah dadu
func NewPlayer1(Jumlahdadu int, position int, name string) *Player {
	player := Player{
	Position: position,
	Name: name,
	Point: 0,
	}
	for i := 0; i < Jumlahdadu; i++ {
		player.Jumlahdadu = append(player.Jumlahdadu, Dice{} )
	}
	fmt.Println(player)
	
	return &player
}


//Addpoint
func (p *Player) AddPoint(Point int) {
	p.Point += Point
}
	
func (p *Player) GetPoint() int {
	return p.Point
}
	
func (p *Player) Play() {
	for _, dice := range p.Jumlahdadu {
	dice.roll()
	}
	
}
//hapus atau insert dadu
func (p *Player) ApusDadu(key int) {
	p.Jumlahdadu = append(p.Jumlahdadu[:key], p.Jumlahdadu[key+1:]...)
}
	
func (p *Player) InsertDice(dice *Dice) {
	p.Jumlahdadu = append(p.Jumlahdadu, *dice)
}

type Game struct {
	players []Player
	round, numberOfPlayer, numberOfDicePerPlayer int
}

	
const REMOVED_WHEN_DICE_TOP = 6
const MOVE_WHEN_DICE_TOP = 1
//sukses
func NewGame(numberOfPlayer, numberOfDicePerPlayer int) *Game {
	g := Game{
	round: 0,
	numberOfPlayer: numberOfPlayer,
	numberOfDicePerPlayer: numberOfDicePerPlayer,}
	for i := 0; i < g.numberOfPlayer; i++ {
		g.players = append(g.players, *NewPlayer1(g.numberOfDicePerPlayer, i, string(rune(65+i))))
	}

	return &g
}
//sukses
func (g *Game) displayRound() *Game {
    fmt.Printf("Giliran %d\r\n", g.round)
    return g
}

func (g *Game) displayTopSideDice(title string) *Game {
    fmt.Printf("%s", title)
    for _, player := range g.players {
        fmt.Printf("Pemain #%s: ", player.getName())
        diceTopSide := ""
        for _, dice := range player.getJumlahdadu() {
            diceTopSide += fmt.Sprintf("%d, ", dice.getAtasDadu())
        }
        fmt.Printf("%s\r\n", strings.TrimSuffix(diceTopSide, ", "))
    }
   
    return g
}

func (g *Game) displayWinner(player *Player) *Game {
	fmt.Println("Pemenang")
	fmt.Printf("Pemain %s\r\n", player.getName())
	return g
}

//Start the game
func (g *Game) Start() {
    fmt.Printf("Pemain = %d, Dadu = %d\n\n", g.numberOfPlayer, g.numberOfDicePerPlayer)

    // Loop until found the winner
    for {
        g.round++
        diceCarryForward := make(map[int][]*Dice)
		
        for _, player := range g.players {
            player.Play()
        }
	 // Display before moved/removed
		g.displayRound().displayTopSideDice("Sebelum")
		// Check player the top side
        for index, player := range g.players {
            tempDiceArray := []*Dice{}
		
			
			for diceIndex, dice := range player.getJumlahdadu() {
                // Check for any occurrence of 6
                if dice.getAtasDadu() == REMOVED_WHEN_DICE_TOP  {
                    player.AddPoint(1)
                    player.ApusDadu(diceIndex)
                }
				 // Check for occurrence of 1
				if dice.getAtasDadu() == MOVE_WHEN_DICE_TOP {
                    // Determine player position
                    // Max player is right most side.
                    // So move the dice to left most side.
                    if player.getPosition() == g.numberOfPlayer-1 {
                        g.players[0].InsertDice(&dice)
                        player.ApusDadu(diceIndex)
                    } else {
                        tempDiceArray = append(tempDiceArray, &dice)
                        player.ApusDadu(diceIndex)
                    }
					
				}

			}

			
            diceCarryForward[index+1] = tempDiceArray

            if diceArray, ok := diceCarryForward[index]; ok && len(diceArray) > 0 {
                // Insert the dice
                for _, dice := range diceArray {
                    player.InsertDice(dice)
                }

                // Reset
                delete(diceCarryForward, index)
            }
        }

        // Display after moved/removed
      
		g.displayRound().displayTopSideDice("Sesudah")
        // Set number player who have dice.
        playerHasDice := g.numberOfPlayer

        for _, player := range g.players {
            if len(player.getJumlahdadu()) <= 0 {
                playerHasDice--
            }
        }

        // Check if player has dice only one
        if playerHasDice == 1 {
            g.displayWinner(g.getWinner())

            // Exit the loop
            break
        }
	
	}
}


func (g *Game) getWinner() *Player {
	var winner *Player
	highscore := 0
	for _, player := range g.players {
	if player.GetPoint() > highscore {
	highscore = player.GetPoint()
	winner = &player
	}
	}
	return winner
}


func main(){
	//memaukakan Input Player dan Input dadu
	Game:=NewGame(2, 2)
	Game.Start()
}

