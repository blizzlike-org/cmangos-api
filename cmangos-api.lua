local http = require("http")
local mariadb = require("mariadb")
local routes = require("routes")

local _M = {}

config = {}
sql = {}

function _M.configure(self, f)
  fd = io.open(f, "r")
  if fd == nil then
    print(string.format("No such file or directory %s", f))
    os.exit(2)
  end

  fd:close()
  return dofile(f)
end

function _M.database(self, cfg)
  local db, err = mariadb.open(
    cfg.username, cfg.password,
    cfg.address .. ":" .. cfg.port,
    cfg.database)
  if err ~= nil then
    print(err)
    os.exit(3)
  end
  return db
end

function _M.main(self)
  if #arg ~= 3 then self:usage() end
  config = self:configure(arg[3])

  sql.api = self:database(config.db.api)
  sql.realmd = self:database(config.db.realmd)

  local listen = config.address .. ":" .. tostring(config.port)
  http.serve(listen, routes, nil)

  sql.api.close()
  sql.realmd.close()
end

function _M.usage(self)
  print(string.format("USAGE: %s %s <config>", arg[1], arg[2]))
  os.exit(1)
end

_M:main()
