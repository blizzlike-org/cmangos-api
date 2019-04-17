local cmangos_auth = require("cmangos.realmd.account.auth")
local http_auth = require("http.auth")
local realm = require("cmangos.realmd.realm")

local _M = {}

local _doc = {
  context = "/realmd/account/auth",
  request = {
    get = {
      header = { ["Authorization"] = {
        Basic = { comment = "username:password base64 encoded" },
        Token = { comment = "verify auth token (uuid form)" }
      } }
    }
  },
  response = {
    get = {
      [200] = {
        header = { ["X-Auth-Token"] = "uuid string" }
      },
      [400] = true,
      [401] = true,
      [500] = true
    }
  }
}

function _M.render(w, r)
  local account, err = http_auth:authenticate(w, r)
  if err ~= nil then return end

  if not account.token then
    account.token, err = cmangos_auth:create_token(account.id)
    if err ~= nil then
      w.set_status(500)
      return
    end
  end

  w.add_header("X-Auth-Token", account.token)
  w.set_status(200)
end

return _M
