local auth = require("cmangos.realmd.account.auth")
local base64 = require("base64")

local _M = {}

local function split(s, delimeter)
  i, j = string.find(s, delimeter)
  m = string.sub(s, 0, j - 1)
  t = string.sub(s, j + 1, s:len())
  return m, t
end

function _M.authenticate(self, r)
  header = r.getHeader("Authorization")
  if not auth then
    r.setStatus(400)
    return nil, "missing authorization header"
  end

  method, token = split(header, " ")
  if method == nil or token == nil then
    r.setStatus(400)
    return nil, "bad authorization header"
  end

  if method:lower() == "basic" then
    username, password = split(base64.decode(token), ":")
    accountid, err = auth:authByPassword(username, password)
    if err ~= nil then return nil, err end
    return accountid, nil
  end

  if method:lower() == "token" then
    accountid, err = auth:authByToken(token)
    if err ~= nil then return nil, err end
    return accountid, nil
  end

  return nil, "auth method not supported " .. (method:lower() or "-")
end

return _M
