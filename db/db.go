package db

import (
	"github.com/golang/glog"
)

func saveword(nickName, word string) error {
	sql := newstatement().insert("word").columns("nickName", "word").values(nickName, word).toString()
	if suc, err := Exec(sql); !suc {
		glog.Warning("DB saveword:", err)
		return err
	}

	return nil
}

func topword(start, len int) ([]map[string]string, error) {
	glog.Info("topword", start, len)
	sql := newstatement().
		selects("nickName", "word").
		from("word").
		orderBy("id", false).
		limit(start, len).
		toString()
	rows, err := Query(sql)
	if nil != err {
		return nil, err
	}
	defer rows.Close()

	var nickName string
	var word string
	infos := []map[string]string{}
	for rows.Next() {
		err = rows.Scan(&nickName, &word)
		if nil != err {
			return nil, err
		}
		infos = append(infos, map[string]string{"nickName": nickName, "word": word})
	}

	return infos, err
}
