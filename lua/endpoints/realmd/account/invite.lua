local cmangos_invite = require("cmangos.realmd.account.invite")
local http_auth = require("http.auth")
local json = require("json")

local _M = {
  post = {}
}

function _M.post.render(w, r)
  print("asd")
  local account, err = http_auth:authenticate(w, r)
  if err then
    print(err)
    return
  end

  local body, err = r.get_body()
  if err then
    print(err)
    w.set_status(500)
    return
  end

  local req, err = json.decode(body or "")
  if err then
    print(err)
    w.set_status(400)
    return
  end

  if not req.gmlevel or req.gmlevel > account.gmlevel then
    print(err)
    w.set_status(403)
    return
  end

  print("asd")
  local invite, err = cmangos_invite:create_token(account.id, req.gmlevel, req.info)
  if err then
    print(err)
    w.set_status(500)
    return
  end

  w.add_header("X-Invite-Token", invite)
  w.set_status(201)
  return
end

return _M
