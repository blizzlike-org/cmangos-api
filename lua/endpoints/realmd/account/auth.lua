local auth = require("http.auth")
local json = require("json")
local realm = require("cmangos.realmd.realm")

local _M = {}

local _doc = {
  context = "/realmd/account/auth",
  request = {
    get = {
      header = { ["Authorization"] = {
        Basic = { comment = "username:password base64 encoded" },
        Token = { comment = "verify auth token (uuid form)" }
      } }
    }
  },
  response = {
    get = {
      [200] = {
        header = { ["X-Auth-Token"] = "uuid string" }
      }
    }
  }
}

function _M.render(w, r)
  local a, err = auth:authenticate(r)
  if err ~= nil then
    w.setStatus(500)
    return
  end

  local resp, err = json.encode(a)
  if err ~= nil then
    w.setStatus(500)
    return
  end

  for k, v in pairs(a) do
    print("k " .. k)
    if type(v) == "table" then
      for i, j in pairs(v) do
        print("i " .. i)
      end
    end
  end

  w.addHeader("Content-Type", "application/json")
  w.setStatus(200)
  w.write(resp)
end

return _M
