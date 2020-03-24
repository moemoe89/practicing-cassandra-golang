//
//  Practicing Cassandra
//
//  Copyright © 2020. All rights reserved.
//

package main

import (
	ap "github.com/moemoe89/practicing-cassandra-golang/api"
	conf "github.com/moemoe89/practicing-cassandra-golang/config"
	"github.com/moemoe89/practicing-cassandra-golang/routers"

	"fmt"

	"github.com/DeanThompson/ginpprof"
)

func main() {
	sess, err := conf.InitDB()
	if err != nil {
		panic(err)
	}

	defer sess.Close()

	log := conf.InitLog()

	repo := ap.NewCassandraRepository(sess)
	svc := ap.NewService(log, repo)

	app := routers.GetRouter(log, svc)
	ginpprof.Wrap(app)
	err = app.Run(":" + conf.Configuration.Port)
	if err != nil {
		panic(fmt.Sprintf("Can't start the app: %s", err.Error()))
	}
}
