local cmangos_account = require("cmangos.realmd.account")
local cmangos_auth = require("cmangos.realmd.account.auth")
local cmangos_invite = require("cmangos.realmd.account.invite")
local http_auth = require("http.auth")
local json = require("json")
local helper = require("helper")
local logger = require("logger")

local _CMANGOS_USERNAME_MAXLEN = 16
local _CMANGOS_PASSWORD_MAXLEN = 16

local _M = {
  get = {},
  post = {},
  helper = {}
}

function _M.helper.validate_username(self, username)
  if not username or #username > _CMANGOS_USERNAME_MAXLEN or
     cmangos_account:username_exists(username) then
    return nil
  end
  return true
end

function _M.helper.validate_password(self, password)
  if not password or #password > _CMANGOS_PASSWORD_MAXLEN then return nil end
  return true
end

function _M.helper.validate_email(self, email)
  if config.signup.email.required then
    if not email or not helper:check_email_format(email) then return nil end
    if config.signup.email.unique then
      if cmangos_account:email_exists(email) then return nil end
    end
  else
    if email and not helper:check_email_format(email) then return nil end
  end
  return true
end

function _M.helper.validate_invite(self, token)
  local invite = nil
  if config.signup.invite.required then
    invite = cmangos_invite:validate_token(token)
    if not token or not invite then return nil end
    return invite
  else
    return true
  end
end

function _M.get.render(w, r)
  local account, err = http_auth:authenticate(w, r)
  if err then return end

  local account, err = cmangos_account:get_account_by_id(account.id)
  if err then
    w.set_status(500)
    return
  end

  local resp, err = json.encode(account)
  if err then
    w.set_status(500)
    return
  end

  w.add_header("Content-Type", "application/json")
  w.write(resp)
  return
end

function _M.post.render(w, r)
  local header = r.get_header("Content-Type")
  if header ~= "application/json" then
    w.set_status(400)
    return
  end

  local body, err = r.get_body()
  if err then
    w.set_status(400)
    return
  end

  local req, err = json.decode(body)
  if err then
    logger.error(err)
    w.set_status(500)
    return
  end

  local resp = {
    username = true,
    password = true,
    validate = true,
    email = true,
    invite = true
  }
  if not _M.helper:validate_username(req.username) then resp.username = false end
  if not _M.helper:validate_password(req.password) then resp.password = false end
  if not req.validate or req.password ~= req.validate then resp.validate = false end
  if not _M.helper:validate_email(req.email) then resp.email = false end
  local invite = _M.helper:validate_invite(req.invite)
  if not invite then resp.invite = false end

  if not resp.username or not resp.password or not resp.validate or not resp.email then
    w.set_status(400)
    w.write(json.encode(resp))
    return
  end

  if not resp.invite then
    w.set_status(401)
    w.write(json.encode(resp))
    return
  end

  local _, err = cmangos_account:create(req.username, req.password, req.email, (invite or {}).gmlevel or 0)
  if err then
    logger.error(err)
    w.set_status(500)
    return
  end

  if config.signup.invite.required then
    local account, err = cmangos_auth:authenticate_by_password(req.username, req.password)
    print(err or "-")
    local _, err = cmangos_invite:update_token(req.invite, account.id)
    print(err or "-")
  end

  w.set_status(201)
  w.write(json.encode(resp))
  return
end

function _M.post.change_account(w, r)
  local account, err = http_auth:authenticate(w, r)
  if err then return end

  local vars = r.parse_vars()
  if account.id ~= tonumber(vars.account) then
    w.set_status(401)
    return
  end

  local header = r.get_header("Content-Type")
  if not header or header ~= "application/json" then
    w.set_status(400)
    return
  end

  local body, err = r.get_body()
  if err then
    w.set_status(400)
    return
  end
  local req = json.decode(body)

  if not req.username and not req.password then
    w.set_status(400)
    return
  end

  local update, err = cmangos_account:change_username(account.id, req.username, req.password)
  if err then
    w.set_status(400)
    return
  end

  w.set_status(200)
  return
end

return _M
