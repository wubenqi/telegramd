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

type UserDialogsDAO struct {
	db *sqlx.DB
}

func NewUserDialogsDAO(db *sqlx.DB) *UserDialogsDAO {
	return &UserDialogsDAO{db}
}

// insert into user_dialogs(user_id, peer_type, peer_id, created_at) values (:user_id, :peer_type, :peer_id, :created_at)
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) Insert(do *do.UserDialogsDO) (id int64, err error) {
	var query = "insert into user_dialogs(user_id, peer_type, peer_id, created_at) values (:user_id, :peer_type, :peer_id, :created_at)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		glog.Error("UserDialogsDAO/Insert error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("UserDialogsDAO/LastInsertId error: ", err)
	}
	return
}

// select peer_type, peer_id from user_dialogs where user_id = :user_id and is_pinned = 1
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) SelectPinnedDialogs(user_id int32) ([]do.UserDialogsDO, error) {
	var query = "select peer_type, peer_id from user_dialogs where user_id = ? and is_pinned = 1"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		glog.Errorf("UserDialogsDAO/SelectPinnedDialogs error: ", err)
		return nil, err
	}

	defer rows.Close()

	var values []do.UserDialogsDO
	for rows.Next() {
		v := do.UserDialogsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			glog.Errorf("UserDialogsDAO/SelectPinnedDialogs error: %s", err)
			return nil, err
		}
		values = append(values, v)
	}

	return values, nil
}

// select id from user_dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) CheckExists(user_id int32, peer_type int32, peer_id int32) (*do.UserDialogsDO, error) {
	var query = "select id from user_dialogs where user_id = ? and peer_type = ? and peer_id = ?"
	rows, err := dao.db.Queryx(query, user_id, peer_type, peer_id)

	if err != nil {
		glog.Error("UserDialogsDAO/CheckExists error: ", err)
		return nil, err
	}

	defer rows.Close()

	do := &do.UserDialogsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			glog.Error("UserDialogsDAO/CheckExists error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}

// select peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count from user_dialogs where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) SelectDialogsByUserID(user_id int32) ([]do.UserDialogsDO, error) {
	var query = "select peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count from user_dialogs where user_id = ?"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		glog.Errorf("UserDialogsDAO/SelectDialogsByUserID error: ", err)
		return nil, err
	}

	defer rows.Close()

	var values []do.UserDialogsDO
	for rows.Next() {
		v := do.UserDialogsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			glog.Errorf("UserDialogsDAO/SelectDialogsByUserID error: %s", err)
			return nil, err
		}
		values = append(values, v)
	}

	return values, nil
}

// select peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count from user_dialogs where user_id = :user_id and peer_type = :peer_type
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) SelectDialogsByPeerType(user_id int32, peer_type int32) ([]do.UserDialogsDO, error) {
	var query = "select peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count from user_dialogs where user_id = ? and peer_type = ?"
	rows, err := dao.db.Queryx(query, user_id, peer_type)

	if err != nil {
		glog.Errorf("UserDialogsDAO/SelectDialogsByPeerType error: ", err)
		return nil, err
	}

	defer rows.Close()

	var values []do.UserDialogsDO
	for rows.Next() {
		v := do.UserDialogsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			glog.Errorf("UserDialogsDAO/SelectDialogsByPeerType error: %s", err)
			return nil, err
		}
		values = append(values, v)
	}

	return values, nil
}
