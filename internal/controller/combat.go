package controller

import (
	"fmt"
	"kimond/gladiatext/internal/actor"
	"kimond/gladiatext/internal/combat"
	"math/rand"
)

const (
	attackCost     = 10
	dodgeCost      = 10
	parryCost      = 10
	breathRecovery = 10
)

var costPerAction = map[combat.Action]int{
	combat.ActionAttack: attackCost,
	combat.ActionDodge:  dodgeCost,
	combat.ActionParry:  parryCost,
}

type CombatController struct {
	Player *actor.Character
	NPC    *actor.NPC
}

func NewCombatController(player *actor.Character, npc *actor.NPC) *CombatController {
	return &CombatController{
		Player: player,
		NPC:    npc,
	}
}

func (cc *CombatController) ResolveRound(playerAction combat.Action) []string {
	roundLog := []string{}
	npcAction := cc.NPC.ChooseAction()

	if playerAction != combat.ActionBreath && cc.Player.Stamina < costPerAction[playerAction] {
		roundLog = append(roundLog, fmt.Sprintf("%s is too exhausted to perform this action...", cc.Player.Name))
		return roundLog
	}

	switch playerAction {
	case combat.ActionAttack:
		roundLog = append(roundLog, cc.handlePlayerAttack(npcAction)...)
	case combat.ActionDodge:
		roundLog = append(roundLog, cc.handlePlayerDodge(npcAction)...)
	case combat.ActionParry:
		roundLog = append(roundLog, cc.handlePlayerParry(npcAction)...)
	case combat.ActionBreath:
		roundLog = append(roundLog, cc.handlePlayerBreath(npcAction)...)
	}

	cc.NPC.UpdateMood()
	return roundLog
}

func (cc *CombatController) handlePlayerAttack(npcAction combat.Action) []string {
	roundLog := []string{}

	playerDamage := cc.Player.Stats.Strength + 1

	switch npcAction {
	case combat.ActionAttack:
		roundLog = append(roundLog, fmt.Sprintf("%s and %s attack each other!", cc.Player.Name, cc.NPC.Name))
		// Player attacks first because has higher dexterity TODO: this is bad
		if cc.Player.Stats.Dexterity > cc.NPC.Character.Stats.Dexterity { // TODO: refine the logic
			roundLog = append(roundLog, fmt.Sprintf("%s manages to hit %s first", cc.Player.Name, cc.NPC.Name))
			cc.NPC.Character.HP -= playerDamage
		} else {
			roundLog = append(roundLog, fmt.Sprintf("%s lands a hit before %s!", cc.NPC.Name, cc.Player.Name))
			cc.Player.HP -= cc.NPC.Character.Stats.Strength + 1
		}
		cc.NPC.Character.Stamina -= attackCost
	case combat.ActionDodge:
		dodgeChance := 25 + (cc.NPC.Character.Stats.Dexterity-cc.Player.Stats.Dexterity)*5
		dodgeChance = min(max(dodgeChance, 0), 100)

		if rollProbability(dodgeChance) {
			roundLog = append(roundLog, fmt.Sprintf("%s dodges %s's attack!", cc.NPC.Name, cc.Player.Name))
		} else {
			roundLog = append(roundLog, fmt.Sprintf("%s fails to dodge %s's attack!", cc.NPC.Name, cc.Player.Name))
			roundLog = append(roundLog, fmt.Sprintf("and receive an attack from %s", cc.Player.Name))
			cc.NPC.Character.HP -= playerDamage
		}
		cc.NPC.Character.Stamina -= dodgeCost
	case combat.ActionParry:
		parryChance := 25 + (cc.NPC.Character.Stats.Strength-cc.Player.Stats.Strength)*5
		parryChance = min(max(parryChance, 0), 100)

		if rollProbability(parryChance) {
			roundLog = append(roundLog, fmt.Sprintf("%s parries %s's attack!", cc.NPC.Name, cc.Player.Name))
			// TODO: handle counter attack logic
		} else {
			roundLog = append(roundLog, fmt.Sprintf("%s fails to parry %s's attack!", cc.NPC.Name, cc.Player.Name))
			roundLog = append(roundLog, fmt.Sprintf("and receive an attack from %s", cc.Player.Name))
			cc.NPC.Character.HP -= playerDamage
		}
		cc.NPC.Character.Stamina -= parryCost
	case combat.ActionBreath:
		roundLog = append(roundLog, fmt.Sprintf("%s takes a deep breath.", cc.NPC.Name))
		roundLog = append(roundLog, fmt.Sprintf("and receive an attack from %s", cc.Player.Name))
		cc.NPC.Character.HP -= playerDamage
		cc.NPC.Character.Stamina += breathRecovery
	}

	cc.Player.Stamina -= attackCost

	return roundLog
}

func (cc *CombatController) handlePlayerDodge(npcAction combat.Action) []string {
	roundLog := []string{}

	switch npcAction {
	case combat.ActionAttack:
		dodgeChance := 25 + (cc.Player.Stats.Dexterity-cc.NPC.Character.Stats.Dexterity)*5
		dodgeChance = min(max(dodgeChance, 0), 100)

		if rollProbability(dodgeChance) {
			roundLog = append(roundLog, fmt.Sprintf("%s dodges %s's attack!", cc.Player.Name, cc.NPC.Name))
		} else {
			roundLog = append(roundLog, fmt.Sprintf("%s fails to dodge %s's attack!", cc.Player.Name, cc.NPC.Name))
			roundLog = append(roundLog, fmt.Sprintf("and receives an attack from %s", cc.NPC.Name))
			cc.Player.HP -= cc.NPC.Character.Stats.Strength + 1
		}
		cc.NPC.Character.Stamina -= attackCost
	case combat.ActionDodge:
		roundLog = append(roundLog, fmt.Sprintf("%s and %s dodge each other!", cc.Player.Name, cc.NPC.Name))
		cc.NPC.Character.Stamina -= dodgeCost
	case combat.ActionParry:
		roundLog = append(roundLog, fmt.Sprintf("%s dodges while %s parries... nothing...", cc.Player.Name, cc.NPC.Name))
		cc.NPC.Character.Stamina -= parryCost
	case combat.ActionBreath:
		roundLog = append(roundLog, fmt.Sprintf("%s takes a deep breath.", cc.Player.Name))
		cc.NPC.Character.Stamina += breathRecovery
	}
	cc.Player.Stamina -= dodgeCost

	return roundLog
}

func (cc *CombatController) handlePlayerParry(npcAction combat.Action) []string {
	roundLog := []string{}

	switch npcAction {
	case combat.ActionAttack:
		parryChance := 25 + (cc.Player.Stats.Strength-cc.NPC.Character.Stats.Strength)*5
		parryChance = min(max(parryChance, 0), 100)

		if rollProbability(parryChance) {
			roundLog = append(roundLog, fmt.Sprintf("%s parries %s's attack!", cc.Player.Name, cc.NPC.Name))
		} else {
			roundLog = append(roundLog, fmt.Sprintf("%s fails to parry %s's attack!", cc.Player.Name, cc.NPC.Name))
			roundLog = append(roundLog, fmt.Sprintf("and receives an attack from %s", cc.NPC.Name))
			cc.Player.HP -= cc.NPC.Character.Stats.Strength + 1
		}
		cc.NPC.Character.Stamina -= attackCost
	case combat.ActionDodge:
		roundLog = append(roundLog, fmt.Sprintf("%s dodges while %s parries... nothing...", cc.NPC.Name, cc.Player.Name))
		cc.NPC.Character.Stamina -= dodgeCost
	case combat.ActionParry:
		roundLog = append(roundLog, fmt.Sprintf("%s and %s parry each other!", cc.Player.Name, cc.NPC.Name))
		cc.NPC.Character.Stamina -= parryCost
	case combat.ActionBreath:
		roundLog = append(roundLog, fmt.Sprintf("%s takes a deep breath.", cc.Player.Name))
		cc.NPC.Character.Stamina += breathRecovery
	}

	cc.Player.Stamina -= parryCost
	return roundLog
}

func (cc *CombatController) handlePlayerBreath(npcAction combat.Action) []string {
	roundLog := []string{}

	switch npcAction {
	case combat.ActionAttack:
		roundLog = append(roundLog, fmt.Sprintf("%s takes a deep breath.", cc.Player.Name))
		roundLog = append(roundLog, fmt.Sprintf("and receives an attack from %s", cc.NPC.Name))
		cc.Player.HP -= cc.NPC.Character.Stats.Strength + 1
		cc.NPC.Character.Stamina -= attackCost
	case combat.ActionDodge:
		roundLog = append(roundLog, fmt.Sprintf("%s takes a deep breath.", cc.Player.Name))
		roundLog = append(roundLog, fmt.Sprintf("%s dodges a deep breath...", cc.NPC.Name))
		cc.NPC.Character.Stamina -= dodgeCost
	case combat.ActionParry:
		roundLog = append(roundLog, fmt.Sprintf("%s takes a deep breath.", cc.Player.Name))
		roundLog = append(roundLog, fmt.Sprintf("%s parries a deep breath...", cc.NPC.Name))
		cc.NPC.Character.Stamina -= parryCost
	case combat.ActionBreath:
		roundLog = append(roundLog, fmt.Sprintf("%s and %s take a deep breath...", cc.Player.Name, cc.NPC.Name))
		cc.NPC.Character.Stamina += breathRecovery
	}
	cc.Player.Stamina += breathRecovery

	return roundLog
}

func rollProbability(chance int) bool {
	return rand.Intn(100) < chance
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
