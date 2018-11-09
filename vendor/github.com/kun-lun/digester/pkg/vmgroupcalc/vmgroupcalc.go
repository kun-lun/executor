package vmgroupcalc

import (
    "github.com/kun-lun/digester/pkg/common"
)

type Requirment struct {
    ConcurrentUserNumber int
}

func Calc(r Requirment) common.VMGroup {
    res := common.VMGroup{
        Count: 1,
        Size: "Standard_DS1_v2",
    }
    x := r.ConcurrentUserNumber
    if x >= 1000 {
        res.Count += 1
    }
    if x >= 2000 {
        res.Count += 2
    }

    return res
}
