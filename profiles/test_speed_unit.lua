print('hello world')

local tags = require('lib/tags')

local way = {
    speed_unit = 'M'
}

function way:get_value_by_key(k)
    return self[k]
end

local speed_unit = tags.get_speed_unit_by_key(way, 'speed_unit')

-- If speed_unit == 'M', then it's mile per hour
-- If speed_unit == 'K', then it's kilometer per hour
local adjust_speed_by_unit = 1
if speed_unit == 'M' then
    adjust_speed_by_unit = 1.60934
end

print(speed_unit, adjust_speed_by_unit)