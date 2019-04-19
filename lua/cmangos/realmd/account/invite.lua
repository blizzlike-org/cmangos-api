local _M = {}

function _M.update_token(self, token, id)
  local result, err = sql.api.query(
    "UPDATE invitetoken SET account = ? WHERE token = ?;",
    id, token)
  if err ~= nil then return nil, err end

  return true
end

function _M.validate_token(self, token)
  local result, err = sql.api.query(
    "SELECT token FROM invitetoken " ..
    " WHERE token = ? AND account IS NULL;", token)
  if err ~= nil then return nil, err end
  if #result == 0 then return nil end

  return result[1]
end

return _M
