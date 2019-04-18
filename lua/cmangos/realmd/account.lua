local _M = {}

function _M.get_account(self, id)
  local result, err = sql.realmd.query(
    "SELECT " ..
    "  id, username, gmlevel, email, " ..
    "  joindate, last_ip, failed_logins, " ..
    "  locked, last_login, expansion, locale " ..
    "FROM account " ..
    "WHERE id = ?;", id)
  if err ~= nil then return nil, err end
  if #result == 0 then return nil end

  return result[1]
end

return _M
