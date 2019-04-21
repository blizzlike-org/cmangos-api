local auth = require("cmangos.realmd.account.auth")
local base64 = require("base64")

local _M = {}

local function split(s, delimeter)
  i, j = string.find(s, delimeter)
  if not i or not j then return nil end
  m = string.sub(s, 0, j - 1)
  t = string.sub(s, j + 1, s:len())
  return m, t
end

function _M.parse_auth_header(self, r)
  local header = r.get_header("Authorization")
  if not header then return nil, nil, "missing authorization header" end

  local method, token = split(header, " ")
  if not method or not token then return nil, nil, "malformed authorization header" end

  return method, token
end

function _M.authenticate(self, w, r)
  local method, token, err = self:parse_auth_header(r)
  if err then
    if w then w.set_status(400) end
    return nil, err
  end

  local account = {}
  if method:lower() == "basic" then
    local username, password = split(base64.decode(token), ":")
    if not username or not password then
      if w then w.set_status(400) end
      return nil, "malformed basic auth"
    end
    account, err = auth:authenticate_by_password(username, password)
    if err then
      if w then w.set_status(500) end
      return nil, err
    end

    if not account then
      if w then w.set_status(401) end
      return nil, "cannot authenticate"
    end

    return account, nil
  end

  if method:lower() == "token" then
    account, err = auth:authenticate_by_token(token)
    if err then return nil, err end
    if not account then
      if w then w.set_status(401) end
      return nil, "cannot authenticate token"
    end
    account.token = token
    return account, nil
  end

  if w then w.set_status(501) end
  return nil, "auth method not supported " .. (method:lower() or "-")
end

return _M
