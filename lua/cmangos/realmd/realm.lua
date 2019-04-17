local _M = {}

function _M.get_realmlist()
  local result, err = sql.realmd.query(
    "SELECT * FROM realmlist ORDER BY id ASC;")
  if err ~= nil then return nil, err end

  return result
end

return _M
