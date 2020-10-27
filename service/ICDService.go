package service

import (
	"api-icd-migration-service/dao"
	"api-icd-migration-service/transformer"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var d dao.ICDDao

type ICDService struct {
}

func (is ICDService) Migrate() {
	totaldoc, err := d.GetCount()
	if err != nil {
		log.Fatal(err)
	}
	perpage := os.Getenv("N_PER_PAGE")
	nperpage, err := strconv.ParseInt(perpage, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(totaldoc)
	var i int64
	i=0
	var noofpages = totaldoc/nperpage
	log.Println(noofpages)
	for i < noofpages {
		micds, err := d.Paginate(i*nperpage, nperpage)
		if err != nil {
			log.Fatal(err)
		}
		icds := transformer.Transform(micds)

		err = d.BulkInsert(icds, nperpage)
		if err != nil {
			log.Fatal(err)
		}
		i++
	}
	micds, err := d.Paginate(i*nperpage, totaldoc - (nperpage*(i)))
	if err != nil {
		log.Fatal(err)
	}
	icds := transformer.Transform(micds)
	err = d.BulkInsert(icds, nperpage)
	if err != nil {
		log.Fatal(err)
	}
}
