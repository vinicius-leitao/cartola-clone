package entity

type Team struct {
	ID      string
	Name    string
	Players []*Player
}

func NewTeam(id, name string) *Team {
	return &Team{
		ID:   id,
		Name: name,
	}
}

func (team *Team) AddPlayer(player *Player) {
	team.Players = append(team.Players, player)
}

func (team *Team) RemovePlayer(player *Player) {
	for indice, playerValue := range team.Players {
		if playerValue.ID == player.ID {
			team.Players = append(team.Players[:indice], team.Players[indice+1:]...)
			return
		}
	}
}
