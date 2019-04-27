local _M = {}

_M._CMANGOS_USERNAME_MAXLEN = 16
_M._CMANGOS_PASSWORD_MAXLEN = 16

function _M.change_username(self, id, username, password)
  if #username > _M._CMANGOS_USERNAME_MAXLEN then
    return nil, "username longer then expected"
  end
  if #password > _M._CMANGOS_PASSWORD_MAXLEN then
    return nil, "password longer then expected"
  end

  if self:username_exists(username) then
    return nil, "username already exists"
  end

  local result, err = sql.realmd.query(
    "UPDATE account SET " ..
    " username = ?, " ..
    " sha_pass_hash = SHA1(CONCAT(UPPER(?), ':', UPPER(?))) " ..
    "WHERE id = ? AND " ..
    " sha_pass_hash = SHA1(CONCAT(UPPER(username), ':', UPPER(?)));",
    username, username, password, id, password)
  if err then
    return nil, err
  end

  if result.affected_rows ~= 1 then
    return nil, "cannot change username"
  end

  return true
end

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
