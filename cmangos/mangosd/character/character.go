package character

import (
  "database/sql"
)

type CharacterInfo struct {
  Guid int64 `json:"guid"`
  Account int64 `json:"account"`
  Name string `json:"name"`
  Race int `json:"race"`
  Class int `json:"class"`
  Gender int `json:"gender"`
  Level int `json:"level"`
  Xp int64 `json:"xp"`
  Money int64 `json:"money"`
  PlayerBytes int64 `json:"playerBytes"`
  PlayerBytes2 int64 `json:"playerBytes2"`
  PlayerFlags int64 `json:"playerFlags"`
  Position_x float64 `json:"position_x"`
  Position_y float64 `json:"position_y"`
  Position_z float64 `json:"position_z"`
  Map int64 `json:"map"`
  Orientation float64 `json:"orientation"`
  Taximask string `json:"taximask"`
  Online int `json:"online"`
  Cinematic int `json:"cinematic"`
  Totaltime int64 `json:"totaltime"`
  Leveltime int64 `json:"leveltime"`
  Logout_time int64 `json:"logout_time"`
  Is_logout_resting int `json:"is_logout_resting"`
  Rest_bonus float64 `json:"rest_bonus"`
  Resettalents_cost int64 `json:"resettalents_cost"`
  Resettalents_time int64 `json:"resettalents_time"`
  Trans_x float64 `json:"trans_x"`
  Trans_y float64 `json:"trans_y"`
  Trans_z float64 `json:"trans_z"`
  Trans_o float64 `json:"trans_o"`
  Transguid int64 `json:"transguid"`
  Extra_flags int64 `json:"extra_flags"`
  Stable_slots int `json:"stable_slots"`
  At_login int64 `json:"at_login"`
  Zone int64 `json:"zone"`
  Death_expire_time int64 `json:"death_expire_time"`
  Taxi_path string `json:"taxi_path"`
  Honor_highest_rank int64 `json:"honor_highest_rank"`
  Honor_standing int64 `json:"honor_standing"`
  Stored_honor_rating float64 `json:"stored_honor_rating"`
  Stored_dishonorable_kills int64 `json:"stored_dishonorable_kills"`
  Stored_honorable_kills int64 `json:"stored_honorable_kills"`
  WatchedFaction int64 `json:"watchedFaction"`
  Drunk int `json:"drunk"`
  Health int64 `json:"health"`
  Power1 int64 `json:"power1"`
  Power2 int64 `json:"power2"`
  Power3 int64 `json:"power3"`
  Power4 int64 `json:"power4"`
  Power5 int64 `json:"power5"`
  ExploredZones string `json:"exploredZones"`
  EquipmentCache string `json:"equipmentCache"`
  AmmoId int64 `json:"ammoId"`
  ActionBars int `json:"actionBars"`
  DeleteInfos_Account sql.NullInt64 `json:"deleteInfos_Account"`
  DeleteInfos_Name sql.NullString `json:"deleteInfos_Name"`
  DeleteDate sql.NullInt64 `json:"deleteDate"`
}

func (c *CharacterInstanceInfo) GetCharacterByAccountId(id int) ([]CharacterInfo, error) {
  var ci []CharacterInfo
  stmt, err := c.Db.Prepare(
    `SELECT
       guid, account, name, race, class, gender,
       level, xp, money, online, totaltime, leveltime,
       logout_time, is_logout_resting, rest_bonus, drunk, health,
       deleteInfos_Account, deleteInfos_Name, deleteDate
     FROM characters
     WHERE account = ?;`)
  if err != nil {
    return ci, err
  }
  defer stmt.Close()

  var rows *sql.Rows
  rows, err = stmt.Query(id)
  for rows.Next() {
    var c CharacterInfo
    err = rows.Scan(
      &c.Guid, &c.Account, &c.Name, &c.Race, &c.Class,
      &c.Gender, &c.Level, &c.Xp, &c.Money, &c.Online, &c.Totaltime,
      &c.Leveltime, &c.Logout_time, &c.Is_logout_resting, &c.Rest_bonus,
      &c.Drunk, &c.Health, &c.DeleteInfos_Account, &c.DeleteInfos_Name, &c.DeleteDate)
    if err != nil {
      return ci, err
    }

    ci = append(ci, c)
  }

  return ci, nil
}
