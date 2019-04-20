local cmangos_character = require("cmangos.mangosd.character")
local http_auth = require("http.auth")
local json = require("json")

local _M = {
  get = {}
}

function _M.get.render(w, r)
  local account, err = http_auth:authenticate(w, r)
  if err then return end

  local vars = r.parse_vars()
  if not sql.mangosd[tonumber(vars.realm or -1)] then
    w.set_status(404)
    return
  end

  local characterlist, err = cmangos_character.list(
    { realm = tonumber(vars.realm) }, account.id)
  if err then
	  print(err)
    w.set_status(500)
    return
  end

  for k, v in pairs(characterlist) do
    characterlist[k].sha_pass_hash = nil
  end

  local list = json.encode(characterlist)
  w.set_status(200)
  w.write(list)
  return
end

return _M
