local cmangos_social = require("cmangos.mangosd.character.social")
local json = require("json")
local http_auth = require("http.auth")

local _M = {}

function _M.get_list(w, r)
  local account, err = http_auth:authenticate(w, r)
  if err then return end

  local vars = r.parse_vars()
  local realm = tonumber(vars.realm)
  if not sql.mangosd[realm or -1] then
    w.set_status(404)
    return
  end

  local list, err = cmangos_social:list(realm, tonumber(vars.character), account.id)
  if err then
    print(err)
    w.set_status(500)
    return
  end

  local list = json.encode(list)
  w.set_status(200)
  w.write(list)
  return
end

return _M
