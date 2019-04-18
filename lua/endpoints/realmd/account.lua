local cmangos_account = require("cmangos.realmd.account")
local http_auth = require("http.auth")
local json = require("json")

local _M = {
  get = {}
}

function _M.get.render(w, r)
  local account, err = http_auth:authenticate(w, r)
  if err ~= nil then return end

  local account, err = cmangos_account:get_account(account.id)
  if err ~= nil then
    w.set_status(500)
    return
  end

  local resp, err = json.encode(account)
  if err ~= nil then
    w.set_status(500)
    return
  end

  w.add_header("Content-Type", "application/json")
  w.write(resp)
  return
end

return _M
