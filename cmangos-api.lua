local cron = require("cron")
local http = require("http")
local mariadb = require("mariadb")
local realmstatus = require("scheduler.realmstatus")
local routes = require("routes")
local logger = require("logger")

local _M = {}

config = {}
jobs = {}
sql = {
  mangosd = {}
}

function _M.configure(self, f)
  fd = io.open(f, "r")
  if fd == nil then
    print(string.format("No such file or directory %s", f))
    os.exit(2)
  end

  fd:close()
  return dofile(f)
end

function _M.open_database(self, cfg)
  local db, err = mariadb.open(
    cfg.username, cfg.password,
    cfg.address .. ":" .. cfg.port,
    cfg.database)
  if err ~= nil then
    logger.error(err)
    os.exit(3)
  end
  return db
end

function _M.main(self)
  if #arg ~= 3 then self:print_usage() end
  config = self:configure(arg[3])
  logger.set_level(config.loglvl)

  sql.api = self:open_database(config.db.api)
  sql.realmd = self:open_database(config.db.realmd)

  for k, v in pairs(config.db.mangosd) do
    sql.mangosd[k] = {
      chars = self:open_database(config.db.mangosd[k].chars),
      world = self:open_database(config.db.mangosd[k].world)
    }
  end

  jobs.realmstatus = cron.run(config.mangosd.check_interval, realmstatus.check)

  local listen = config.address .. ":" .. tostring(config.port)
  http.serve(listen, routes, nil)

  sql.api.close()
  sql.realmd.close()

  for k, v in pairs(sql.mangosd) do
    sql.mangosd[k].chars.close()
    sql.mangosd[k].world.close()
  end

  jobs.realmstatus.stop()
end

function _M.print_usage(self)
  print(string.format("USAGE: %s %s <config>", arg[1], arg[2]))
  os.exit(1)
end

_M:main()
