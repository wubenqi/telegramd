/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mysql_dao

import (
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	do "github.com/nebulaim/telegramd/biz_model/dal/dataobject"
)

type AuthUsersDAO struct {
	db *sqlx.DB
}

func NewAuthUsersDAO(db *sqlx.DB) *AuthUsersDAO {
	return &AuthUsersDAO{db}
}

// insert into auth_users(auth_id, user_id) values (:auth_id, :user_id)
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) Insert(do *do.AuthUsersDO) (id int64, err error) {
	var query = "insert into auth_users(auth_id, user_id) values (:auth_id, :user_id)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		glog.Error("AuthUsersDAO/Insert error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("AuthUsersDAO/LastInsertId error: ", err)
	}
	return
}

// select id, user_id from auth_users where auth_id = :auth_id
// TODO(@benqi): sqlmap
func (dao *AuthUsersDAO) SelectByAuthId(auth_id int64) (*do.AuthUsersDO, error) {
	var query = "select id, user_id from auth_users where auth_id = ?"
	rows, err := dao.db.Queryx(query, auth_id)

	if err != nil {
		glog.Error("AuthUsersDAO/SelectByAuthId error: ", err)
		return nil, err
	}

	defer rows.Close()

	do := &do.AuthUsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			glog.Error("AuthUsersDAO/SelectByAuthId error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}
