package server

import mgo "gopkg.in/mgo.v2"

func (a *API) SetupDatabase() (*mgo.Session, error) {
	session, err := mgo.Dial(a.Config.GetString("mongo_host"))
	if err != nil {
		return nil, err
	}

	a.Database = session.DB(a.Config.GetString("mongo_database"))

	return session, nil
}
