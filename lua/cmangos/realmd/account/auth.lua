local uuid = require("uuid")

local _M = {}

function _M.authenticate_by_password(self, username, password)
  local result, err = sql.realmd.query(
    "SELECT id FROM account " ..
    " WHERE " ..
    "  UPPER(username) = UPPER(?) AND " ..
    "  sha_pass_hash = SHA1(CONCAT(UPPER(?), ':', UPPER(?)));",
    username, username, password)
  if err ~= nil then return nil, err end
  if #result == 0 then return nil end

  return result[1]
end

function _M.authenticate_by_token(self, token)
  local result, err = sql.api.query(
    "SELECT owner FROM authtoken " ..
    " WHERE token = ?;", token)
  if err ~= nil then return nil, err end
  if #result == 0 then return nil end

  result[1].id = result[1].owner
  result[1].owner = nil

  local expiry = os.time() + config.expiry
  local _r, err = sql.api.query(
    "UPDATE authtoken SET expiry = ? WHERE token = ?;",
    expiry, token)
  if err ~= nil then return nil, err end

  return result[1]
end

function _M.create_token(self, accountid)
  local token, err = uuid.create()
  local expiry = os.time() + config.expiry
  if err ~= nil then return nil, err end

  local result, err = sql.api.query(
    "INSERT INTO authtoken (token, owner, expiry) VALUES (?, ?, ?);",
    token, accountid, expiry)
  if err ~= nil then return nil, err end
  return token
end

return _M
