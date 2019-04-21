local seclevel = require("cmangos.seclevel")

local _M = {
  realm = nil
}

function _M.get_info(self, characterid, account)
  if not sql.mangosd[self.realm] then return nil, "no database configuration defined for realmid" end

  local result, err = sql.mangosd[self.realm].chars.query(
    "SELECT " ..
    " c.account AS account, " ..
    " c.guid AS guid, " ..
    " c.name AS charactername, " ..
    " gr.rname AS rank, " ..
    " gr.rights AS rights, " ..
    " gm.pnote AS playernote, " ..
    " gm.offnote AS officernote, " ..
    " g.guildid AS guildid, " ..
    " g.name AS guildname, " ..
    " g.leaderguid AS leader, " ..
    " g.EmblemStyle AS emblemstyle, " ..
    " g.EmblemColor AS emblemcolor, " ..
    " g.BorderStyle AS borderstyle, " ..
    " g.BorderColor AS bordercolor, " ..
    " g.BackgroundColor AS backgroundcolor, " ..
    " g.info AS info, " ..
    " g.motd AS motd, " ..
    " g.createdate AS createdate " ..
    "FROM characters AS c " ..
    "INNER JOIN guild_member AS gm ON (c.guid = gm.guid) " ..
    "INNER JOIN guild AS g ON (gm.guildid = g.guildid) " ..
    "INNER JOIN guild_rank AS gr ON (gm.rank = gr.rid) " ..
    "WHERE c.guid = ?;", characterid)
  if err then return nil, err end
  if #result == 0 then return nil, "no guild membership found" end

  if not account then
    account = { gmlevel = seclevel._VISITOR }
  end

  if account.gmlevel < seclevel._GAMEMASTER then
    result[1].playernote = nil
    result[1].officernote = nil
    result[1].leader = nil
    result[1].info = nil
    result[1].motd = nil
    result[1].rights = nil
  end

  return result[1]
end

return _M
