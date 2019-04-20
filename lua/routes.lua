local auth = require("endpoints.realmd.account.auth")
local account = require("endpoints.realmd.account")
local character = require("endpoints.mangosd.character")
local invite = require("endpoints.realmd.account.invite")
local realm = require("endpoints.realmd.realm")

return {
  { method = "DELETE", context = "/realmd/account/invite/{token}", callback = invite.delete.render },

  { method = "GET", context = "/realmd/account", callback = account.get.render },
  { method = "GET", context = "/realmd/account/auth", callback = auth.render },
  { method = "GET", context = "/realmd/account/invite", callback = invite.get.render },
  { method = "GET", context = "/realmd/realmlist", callback = realm.get.render },

  { method = "GET", context = "/mangosd/{realm}/character", callback = character.get.render },

  { method = "POST", context = "/realmd/account", callback = account.post.render },
  { method = "POST", context = "/realmd/account/invite", callback = invite.post.render }
}
