
# 编号不为2 (存在null的情况)
select name
from customer
where referee_id != 2 or referee_id is null;