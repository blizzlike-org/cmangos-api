local auth = require("endpoints.realmd.account.auth")
local realm = require("endpoints.realmd.realm")

return {
  { method = "GET", context = "/realmd/account/auth", callback = auth.render },

  { method = "GET", context = "/realmd/realmlist", callback = realm.get.render }
}
