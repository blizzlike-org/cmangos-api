local cmangos_guild = require("cmangos.mangosd.character.guild")
local http_auth = require("http.auth")
local json = require("json")

local _M = {
  get = {}
}

function _M.get.render(w, r)
  local header = r.get_header("Authorization")
  local account, err = nil, nil
  if header then
    account, err = http_auth:authenticate(w, r)
    if err then return end
  end

  local vars = r.parse_vars()
  local guild, err = cmangos_guild.get_info(
    { realm = tonumber(vars.realm) }, tonumber(vars.character), account)
  if err then
    print(err)
    w.set_status(500)
    return
  end

  local resp = json.encode(guild)
  w.set_status(200)
  w.write(resp)
  return
end

return _M
