local auth = require("endpoints.realmd.account.auth")
local account = require("endpoints.realmd.account")
local realm = require("endpoints.realmd.realm")

return {
  { method = "GET", context = "/realmd/account", callback = account.get.render },
  { method = "GET", context = "/realmd/account/auth", callback = auth.render },

  { method = "POST", context = "/realmd/account", callback = account.post.render },

  { method = "GET", context = "/realmd/realmlist", callback = realm.get.render }
}
