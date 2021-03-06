local auth = require("endpoints.realmd.account.auth")
local account = require("endpoints.realmd.account")
local character = require("endpoints.mangosd.character")
local guild = require("endpoints.mangosd.character.guild")
local invite = require("endpoints.realmd.account.invite")
local realm = require("endpoints.realmd.realm")
local social = require("endpoints.mangosd.character.social")

return {
  { method = "DELETE", context = "/realmd/account/invite/{token}", callback = invite.delete.render },

  { method = "GET", context = "/realmd/account", callback = account.get.render },
  { method = "GET", context = "/realmd/account/auth", callback = auth.render },
  { method = "GET", context = "/realmd/account/invite", callback = invite.get.render },
  { method = "GET", context = "/realmd/realmlist", callback = realm.get.render },

  { method = "GET", context = "/mangosd/{realm}/character", callback = character.get.render },
  { method = "GET", context = "/mangosd/{realm}/character/{character}/social", callback = social.get_list },
  { method = "GET", context = "/mangosd/{realm}/character/{character}/guild", callback = guild.get.render },

  { method = "POST", context = "/realmd/account", callback = account.post.render },
  { method = "POST", context = "/realmd/account/{account}", callback = account.post.change_account },
  { method = "POST", context = "/realmd/account/invite", callback = invite.post.render }
}
