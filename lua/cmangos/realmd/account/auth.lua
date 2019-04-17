local _M = {}

function _M.authByPassword(self, username, password)
  local result, err = sql.realmd.query(
    "SELECT id FROM account " ..
    " WHERE " ..
    "  UPPER(username) = UPPER(?) AND " ..
    "  sha_pass_hash = SHA1(CONCAT(UPPER(?), ':', UPPER(?)));",
    username, username, password)
  if err ~= nil then return nil, err end

  return result[1]
end

function _M.authByToken(self, token)
  local result, err = sql.api.query(
    "SELECT owner FROM authtoken " ..
    " WHERE token = ?;", token)
  if err ~= nil then return nil, err end

  return result[1]
end

return _M
