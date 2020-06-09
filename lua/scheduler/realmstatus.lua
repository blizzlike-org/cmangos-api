local realm = require("cmangos.realmd.realm")
local socket = require("socket")

local _M = {}

function _M.check()
  local rl, err = realm:get_realmlist()
  if err ~= nil then return end

  for k, v in pairs(rl or {}) do
    local check = _M:check_host(v.address, v.port)
    rl[k].check = check.last_check
    rl[k].state = check.state
  end
  realmlist = rl
  return
end

function _M.check_host(self, address, port)
  local s, err = socket.open("tcp", address .. ":" .. tostring(port), config.mangosd.check_timeout or nil)
  local state = 0
  if not err then
    s.close()
    state = 1
  end

  return {
    state = state,
    last_check = os.time()
  }
end

return _M
