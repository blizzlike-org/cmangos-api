local cmangos_invite = require("cmangos.realmd.account.invite")
local http_auth = require("http.auth")
local json = require("json")

local _M = {
  get = {},
  post = {}
}

function _M.get.render(w, r)
  local account, err = http_auth:authenticate(w, r)
  if err then return end

  local invites, err = cmangos_invite:get_tokens(account.id)
  if err then
    w.set_status(500)
    return
  end

  local resp, err = json.encode(invites)
  if err then
    w.set_status(500)
    return
  end

  w.add_header("Content-Type", "application/json")
  w.set_status(200)
  w.write(resp)
  return
end

function _M.post.render(w, r)
  local account, err = http_auth:authenticate(w, r)
  if err then return end

  local body, err = r.get_body()
  if err then
    w.set_status(500)
    return
  end

  local req, err = json.decode(body or "")
  if err then
    w.set_status(400)
    return
  end

  if not req.gmlevel or req.gmlevel > account.gmlevel then
    w.set_status(403)
    return
  end

  local invite, err = cmangos_invite:create_token(account.id, req.gmlevel, req.info)
  if err then
    w.set_status(500)
    return
  end

  w.add_header("X-Invite-Token", invite)
  w.set_status(201)
  return
end

return _M
