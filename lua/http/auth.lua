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

function _M.authenticate(self, w, r)
  header = r.get_header("Authorization")
  if not header then
    w.set_status(400)
    return nil, "missing authorization header"
  end

  method, token = split(header, " ")
  if not method or not token then
    w.set_status(400)
    return nil, "bad authorization header"
  end

  if method:lower() == "basic" then
    username, password = split(base64.decode(token), ":")
    if not username or not password then
      w.set_status(400)
      return nil, "malformed authorization header"
    end
    account, err = auth:authenticate_by_password(username, password)
    if err then
      w.set_status(500)
      return nil, err
    end

    if not account then
      w.set_status(401)
      return nil, "cannot authenticate"
    end

    return account, nil
  end

  if method:lower() == "token" then
    account, err = auth:authenticate_by_token(token)
    if err then return nil, err end
    if not account then
      w.set_status(401)
      return nil, "cannot authenticate token"
    end
    account.token = token
    return account, nil
  end

  return nil, "auth method not supported " .. (method:lower() or "-")
end

return _M
