package main

import (
	"fmt"
	zt "github.com/andy-zhangtao/gogather/time"
	"time"
)

func addGroot(g Groot) (err error) {
	return bw.Save(g)
}

func findSpecifyGroot(g *Groot) (err error) {
	return bw.FindOne(g)
}

func findAllGroot(g Groot, allGroot interface{}, sort []string) (err error) {
	return bw.FindAllWithSort(g, allGroot, sort)
}

func updateGroot(g Groot, fields []string) (err error) {
	_, err = bw.Update(&g, fields)
	return
}

func AddNewGroot(g Groot) (err error) {
	if g.Time == "" {
		t := zt.Ztime{}
		g.Time, err = t.Now().Format("YYYY-MM")
		if err != nil {
			return z.Error(fmt.Sprintf("Time Format Error [%s]", err))
		}
	}

	g.Total = g.B_oth+g.B_ali+g.B_ccb+g.B_citic+g.B_cmb+g.B_bocom+g.B_abc+g.B_icbc
	g.Timestamp = time.Now().Unix()
	return addGroot(g)
}

func FindSpecifyGroot(month string) (g Groot, err error) {
	g.Time = month
	err = findSpecifyGroot(&g)
	return
}

func FindAllGroot() (allGroot []Groot, err error) {
	err = findAllGroot(Groot{}, &allGroot, []string{"+time"})
	return
}

func UpdateGrootByMonth(g Groot) (err error) {
	g.Total = g.B_oth+g.B_ali+g.B_ccb+g.B_citic+g.B_cmb+g.B_bocom+g.B_abc+g.B_icbc
	g.Timestamp = time.Now().Unix()
	err = updateGroot(g, []string{"time"})
	return
}
