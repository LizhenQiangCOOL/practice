
# 显示有特殊需求，使用 case when then
select employee_id,
       ( case when mod(employee_id, 2) != 0 AND left(name, 1) != 'M' then salary
              when mod(employee_id, 2)=0 or left(name, 1) = 'M' then 0
           end) as  bonus
from Employees
order by employee_id asc;
