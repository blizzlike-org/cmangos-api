local _M = {
  realm = nil
}

function _M.set_realm(self, realmid)
  local M = _M
  M.realm = realmid
  return M
end

function _M.get_info(self)
  if not sql.mangosd[self.realm] then return nil, "no database configuration defined for realmid" end

  local result, err = sql.mangosd[self.realm].chars.query(
    "SELECT SUM(CASE WHEN race IN (2, 5, 6, 8, 10) THEN 1 ELSE 0 END) as horde, SUM(CASE WHEN race IN (1, 3, 4, 7, 11) THEN 1 ELSE 0 END) as alliance from characters;")
  if err then return nil, err end

  return result
end

function _M.list(self, accountid)
  if not sql.mangosd[self.realm] then return nil, "no database configuration defined for realmid" end

  local result, err = sql.mangosd[self.realm].chars.query(
    "SELECT * FROM characters WHERE account = ?;", accountid)
  if err then return nil, err end

  return result
end

return _M
