package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
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
			Type: graphql.NewList(GrootBankType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				//if g, ok := p.Source.(Groot); ok {
				//	return g.Bank, nil
				//}
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
				//if b, ok := p.Source.(GrootBank); ok {
				//	return b.Kind, nil
				//}
				return nil, nil
			},
		},
		"balance": &graphql.Field{
			Type: graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				//if b, ok := p.Source.(GrootBank); ok {
				//	return b.Balance, nil
				//}
				return nil, nil
			},
		},
		"status": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				//if b, ok := p.Source.(GrootBank); ok {
				//	return b.Status, nil
				//}
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

		if len(month) == 4 {
			gs, err := FindAllGroot()
			if err != nil {
				return nil, err
			}

			var tg []Groot
			for _, g := range gs {
				if g.Time[:4] == month {
					tg = append(tg, g)
				}
			}

			return tg, nil
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

		var gb []GrootBank
		err := json.Unmarshal([]byte(bank), &gb)
		if err != nil {
			return nil, z.Error(fmt.Sprintf("Json Unmarshal Error [%s]", err))
		}

		g := Groot{
			Time: month,
		}

		for _, bank := range gb {
			switch bank.Kind {
			case BANK_TYPE_ICBC:
				g.B_icbc = bank.Balance
			case BANK_TYPE_ABCHINA:
				g.B_abc = bank.Balance
			case BANK_TYPE_BOCOM:
				g.B_bocom = bank.Balance
			case BANK_TYPE_CMBCHINA:
				g.B_cmb = bank.Balance
			case BANK_TYPE_CITIC:
				g.B_citic = bank.Balance
			case BANK_TYPE_CCB:
				g.B_ccb = bank.Balance
			case BANK_TYPE_BJ:
				g.B_bj = bank.Balance
			case BANK_TYPE_ALIPAY:
				g.B_ali = bank.Balance
			case BANK_TYPE_OTHER:
				g.B_oth = bank.Balance
			}
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

		var gb []GrootBank

		err := json.Unmarshal([]byte(bank), &gb)
		if err != nil {
			return nil, z.Error(fmt.Sprintf("Json Unmarshal Error [%s]", err))
		}

		g := Groot{
			Time: month,
		}

		for _, bank := range gb {
			switch bank.Kind {
			case BANK_TYPE_ICBC:
				g.B_icbc = bank.Balance
			case BANK_TYPE_ABCHINA:
				g.B_abc = bank.Balance
			case BANK_TYPE_BOCOM:
				g.B_bocom = bank.Balance
			case BANK_TYPE_CMBCHINA:
				g.B_cmb = bank.Balance
			case BANK_TYPE_CITIC:
				g.B_citic = bank.Balance
			case BANK_TYPE_CCB:
				g.B_ccb = bank.Balance
			case BANK_TYPE_BJ:
				g.B_bj = bank.Balance
			case BANK_TYPE_ALIPAY:
				g.B_ali = bank.Balance
			case BANK_TYPE_OTHER:
				g.B_oth = bank.Balance
			}
		}
		
		oldGroot, err := FindSpecifyGroot(month)
		if err != nil {
			return nil, err
		}

		if oldGroot.B_icbc == 0 && oldGroot.B_oth == 0 && oldGroot.B_ali == 0 && oldGroot.B_bj == 0 && oldGroot.B_ccb == 0 && oldGroot.B_citic == 0 && oldGroot.B_cmb == 0 && oldGroot.B_bocom == 0 && oldGroot.B_abc == 0 {
			if err = AddNewGroot(g); err != nil {
				return nil, z.Error(fmt.Sprintf("Add Groot Error [%s]", err))
			}
		}

		if err = UpdateGrootByMonth(g); err != nil {
			return nil, z.Error(fmt.Sprintf("Update Groot Error [%s]", err))
		}

		return g, nil
	},
}
