local uuid = require("uuid")

local _M = {}

function _M.create_token(self, id, gmlevel, info)
  local token = uuid.create()
  local result, err = sql.api.query(
    "INSERT INTO invitetoken (token, friend, gmlevel, info) VALUES (?, ?, ?, ?)",
    token, id, gmlevel, info or "")
  if err then return nil, err end

  return token
end

function _M.delete_token(self, token)
  local _, err = sql.api.query(
    "DELETE FROM invitetoken WHERE token = ? AND account IS NULL;", token)
  if err then return nil, err end

  return true
end

function _M.get_tokens(self, id)
  local result, err = sql.api.query(
    "SELECT * FROM invitetoken WHERE friend = ? AND account IS NULL;", id)
  if err then return nil, err end
  return result
end

function _M.update_token(self, token, id)
  local result, err = sql.api.query(
    "UPDATE invitetoken SET account = ? WHERE token = ?;",
    id, token)
  if err ~= nil then return nil, err end

  return true
end

function _M.validate_token(self, token)
  local result, err = sql.api.query(
    "SELECT friend, token, gmlevel FROM invitetoken " ..
    " WHERE token = ? AND account IS NULL;", token)
  if err then return nil, err end
  if #result == 0 then return nil end

  return result[1]
end

return _M
