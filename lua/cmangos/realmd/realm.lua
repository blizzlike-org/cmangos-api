local cmangos_character = require("cmangos.mangosd.character")

local _M = {}

function _M.get_realmlist()
  local result, err = sql.realmd.query(
    "SELECT r.*, COALESCE(rc.accounts, 0) as accounts, COALESCE(rc.characters, 0) as characters, COALESCE(u.starttime, 0) as starttime, COALESCE(a.online, 0) AS online FROM realmlist AS r LEFT JOIN (SELECT realmid, MAX(starttime) AS starttime FROM uptime GROUP BY realmid) AS u ON (r.id = u.realmid) LEFT JOIN (SELECT realmid, COUNT(acctid) AS accounts, SUM(numchars) AS characters FROM realmcharacters GROUP BY realmid) AS rc ON (r.id = rc.realmid) LEFT JOIN (SELECT active_realm_id, COUNT(active_realm_id) as online FROM account GROUP BY active_realm_id) AS a ON (r.id = a.active_realm_id);")
  if err ~= nil then return nil, err end

  for k, v in pairs(result) do
    local info, err = cmangos_character.get_info({ realm = tonumber(v.id) })
    if not err then
      result[k].alliance = info[1].alliance
      result[k].horde = info[1].horde
    else
      result[k].alliance = 0
      result[k].horde = 0
    end
  end

  return result
end

return _M
