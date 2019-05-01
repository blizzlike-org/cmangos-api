local seclevel = require("cmangos.seclevel")

local _M = {}

function _M.list(self, realm, character, account)
  if not sql.mangosd[realm] then return nil, "no database configuration defined for realmid" end

  local result, err = sql.mangosd[realm].chars.query(
    "SELECT " ..
    " f.name AS name, " ..
    " f.online AS online, " ..
    " s.friend AS friend, " ..
    " s.flags AS flags " ..
    "FROM character_social AS s " ..
    "INNER JOIN characters AS c ON (s.guid = c.guid) " ..
    "INNER JOIN characters AS f ON (s.friend = f.guid) " ..
    "WHERE s.guid = ? AND c.account = ?;", character, account)
  if err then return nil, err end
  if #result == 0 then return nil, "no social information available" end

  return result
end

return _M
