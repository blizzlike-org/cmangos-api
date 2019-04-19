local _M = {}

function _M.check_email_format(self, email)
-- http://lua-users.org/wiki/StringRecipes
if email:match("[A-Za-z0-9%.%%%+%-]+@[A-Za-z0-9%.%%%+%-]+%.%w%w%w?%w?") then
  return true
end
return nil
end

return _M
