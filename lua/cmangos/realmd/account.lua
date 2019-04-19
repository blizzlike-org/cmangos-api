local _M = {}

function _M.create(self, username, password, email, gmlevel)
  local result, err = sql.realmd.query(
    "INSERT INTO account (username, sha_pass_hash, gmlevel, email, joindate) " ..
    " VALUES (UPPER(?), SHA1(CONCAT(UPPER(?), ':', UPPER(?))), ?, LOWER(?), NOW());",
    username, username, password, gmlevel, email or "")
  if err then return nil, err end

  return true
end

function _M.get_account_by_id(self, id)
  local result, err = sql.realmd.query(
    "SELECT " ..
    " id, username, gmlevel, email, " ..
    " joindate, last_ip, failed_logins, " ..
    " locked, last_login, expansion, locale " ..
    "FROM account " ..
    "WHERE id = ?;", id)
  if err ~= nil then return nil, err end
  if #result == 0 then return nil end

  return result[1]
end

function _M.username_exists(self, username)
  local result, err = sql.realmd.query(
    "SELECT " ..
    " id, username, gmlevel, email, " ..
    " joindate, last_ip, failed_logins, " ..
    " locked, last_login, expansion, locale " ..
    "FROM account " ..
    "WHERE username = ?;", username)
  if err ~= nil then return nil, err end
  if #result == 0 then return nil end

  return result[1]
end

function _M.email_exists(self, email)
  local result, err = sql.realmd.query(
    "SELECT " ..
    " id, username, gmlevel, email, " ..
    " joindate, last_ip, failed_logins, " ..
    " locked, last_login, expansion, locale " ..
    "FROM account " ..
    "WHERE email = ?;", email)
  if err ~= nil then return nil, err end
  if #result == 0 then return nil end

  return result[1]
end

return _M
