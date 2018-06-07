package main

import (
	"github.com/graphql-go/graphql"
	"encoding/json"
	"fmt"
)

var GrootType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Groot",
	Fields: graphql.Fields{
		"time": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if g, ok := p.Source.(Groot); ok {
					return g.Time, nil
				}
				return nil, nil
			},
		},
		"bank": &graphql.Field{
			Type: GrootBankType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if g, ok := p.Source.(Groot); ok {
					return g.Bank, nil
				}
				return nil, nil
			},
		},
	},
})

var GrootBankType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Groot_Bank",
	Fields: graphql.Fields{
		"kind": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if b, ok := p.Source.(GrootBank); ok {
					return b.Kind, nil
				}
				return nil, nil
			},
		},
		"balance": &graphql.Field{
			Type: graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if b, ok := p.Source.(GrootBank); ok {
					return b.Balance, nil
				}
				return nil, nil
			},
		},
		"status": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if b, ok := p.Source.(GrootBank); ok {
					return b.Status, nil
				}
				return nil, nil
			},
		},
	},
})

var QueryGroot = &graphql.Field{
	Type: graphql.NewList(GrootType),
	Args: graphql.FieldConfigArgument{
		"month": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		month, _ := p.Args["month"].(string)

		if month == "" {
			return FindAllGroot()
		}

		if gs, err := FindSpecifyGroot(month); err != nil {
			return nil, err
		} else {
			return []Groot{gs}, nil
		}
	},
}

var AddGroot = &graphql.Field{
	Type: GrootType,
	Args: graphql.FieldConfigArgument{
		"month": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"bank": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		month, _ := p.Args["month"].(string)
		bank, _ := p.Args["bank"].(string)

		var gb GrootBank

		err := json.Unmarshal([]byte(bank), &gb)
		if err != nil {
			return nil, z.Error(fmt.Sprintf("Json Unmarshal Error [%s]", err))
		}

		g := Groot{
			Time: month,
			Bank: gb,
		}

		if err = AddNewGroot(g); err != nil {
			return nil, z.Error(fmt.Sprintf("Add Groot Error [%s]", err))
		}

		return g, nil
	},
}

var UpdateGroot = &graphql.Field{
	Type: GrootType,
	Args: graphql.FieldConfigArgument{
		"month": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"bank": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		month, _ := p.Args["month"].(string)
		bank, _ := p.Args["bank"].(string)

		var gb GrootBank

		err := json.Unmarshal([]byte(bank), &gb)
		if err != nil {
			return nil, z.Error(fmt.Sprintf("Json Unmarshal Error [%s]", err))
		}

		g := Groot{
			Time: month,
			Bank: gb,
		}

		if err = UpdateGrootByMonth(g); err != nil {
			return nil, z.Error(fmt.Sprintf("Add Groot Error [%s]", err))
		}

		return g, nil
	},
}
