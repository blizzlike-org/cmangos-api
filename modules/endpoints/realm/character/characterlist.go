package character

import (
  "encoding/json"
  "fmt"
  "os"
  "net/http"
  "strconv"
  "github.com/gorilla/mux"

  e_account "metagit.org/blizzlike/cmangos-api/modules/endpoints/account"
  cmangos_realm "metagit.org/blizzlike/cmangos-api/cmangos/realmd/realm"
  cmangos_character "metagit.org/blizzlike/cmangos-api/cmangos/mangosd/character"
)

func DoCharacterlistByAccount(w http.ResponseWriter, r *http.Request) {
  id, err := e_account.AuthenticateByToken(w, r)
  if err != nil {
    return
  }

  var characterlist []cmangos_character.CharacterInfo
  realmlist := cmangos_realm.GetRealms()
  vars := mux.Vars(r)
  realmid, _ := strconv.Atoi(vars["realm"])
  accountid, _ := strconv.Atoi(vars["account"])
  for _, v := range realmlist {
    if realmid == v.Id && accountid == id {
      characterlist, err = v.CharacterInstance.GetCharacterByAccountId(accountid)
      if err != nil {
        fmt.Fprintf(os.Stderr, "Cannot fetch characterlist (%v)\n", err)
        w.WriteHeader(http.StatusInternalServerError)
	return
      }
    }
  }

  if len(characterlist) == 0 {
    w.WriteHeader(http.StatusNotFound)
    return
  }

  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(characterlist)
  return
}
