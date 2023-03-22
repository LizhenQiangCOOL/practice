
#过滤面积或人口
select name, population, area
from World
where area >= 3000000 or population >= 25000000;