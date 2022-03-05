package entities

type Profile struct {
	Address           string
	ENSName           string
	DisplayConfig     *DisplayConfig // DisplayConfig can be nil
	NonFungibleTokens *[]NonFungibleToken
	FungibleTokens    *[]FungibleToken
	Statistics        *[]Statistic
	Interactions      *[]Interaction
}
