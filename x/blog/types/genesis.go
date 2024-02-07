package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PostList:         []Post{},
		LoanList:         []Loan{},
		SentPostList:     []SentPost{},
		TimedoutPostList: []TimedoutPost{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in post
	postIdMap := make(map[uint64]bool)
	postCount := gs.GetPostCount()
	for _, elem := range gs.PostList {
		if _, ok := postIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for post")
		}
		if elem.Id >= postCount {
			return fmt.Errorf("post id should be lower or equal than the last id")
		}
		postIdMap[elem.Id] = true
	}
	// Check for duplicated ID in loan
	loanIdMap := make(map[uint64]bool)
	loanCount := gs.GetLoanCount()
	for _, elem := range gs.LoanList {
		if _, ok := loanIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for loan")
		}
		if elem.Id >= loanCount {
			return fmt.Errorf("loan id should be lower or equal than the last id")
		}
		loanIdMap[elem.Id] = true
	}
	// Check for duplicated ID in sentPost
	sentPostIdMap := make(map[uint64]bool)
	sentPostCount := gs.GetSentPostCount()
	for _, elem := range gs.SentPostList {
		if _, ok := sentPostIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for sentPost")
		}
		if elem.Id >= sentPostCount {
			return fmt.Errorf("sentPost id should be lower or equal than the last id")
		}
		sentPostIdMap[elem.Id] = true
	}
	// Check for duplicated ID in timedoutPost
	timedoutPostIdMap := make(map[uint64]bool)
	timedoutPostCount := gs.GetTimedoutPostCount()
	for _, elem := range gs.TimedoutPostList {
		if _, ok := timedoutPostIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for timedoutPost")
		}
		if elem.Id >= timedoutPostCount {
			return fmt.Errorf("timedoutPost id should be lower or equal than the last id")
		}
		timedoutPostIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
