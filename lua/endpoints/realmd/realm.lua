local json = require("json")
local realm = require("cmangos.realmd.realm")

local _M = {
  get = {}
}

local _doc = {
  context = "/realmd/realmlist",
  request = {
    get = true
  },
  response = {
    get = {
      [200] = {
        header = { ["Content-Type"] = "application/json" },
        body = {{
          { key = "id", type = "int", comment = "realm id" },
          { key = "name", type = "string", comment = "realm name" },
          { key = "address", type = "string", comment = "realm ip address" },
          { key = "port", type = "int", comment = "realm tcp port" }
        }}
      }
    }
  }
}

function _M.get.render(w, r)
  local realmlist, err = realm:getRealmlist()
  if err ~= nil then
    print(err)
    w.setStatus(500)
    return
  end

  local resp, err = json.encode(realmlist)
  if err ~= nil then
    print(err)
    w.setStatus(500)
    return
  end

  w.addHeader("Content-Type", "application/json")
  w.setStatus(200)
  w.write(resp)
end

return _M
