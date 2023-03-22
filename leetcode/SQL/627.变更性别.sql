
# 一个update语句需要根据不同条件变换
# mysql case when else 相当于 golang中的 select case default
update Salary
set sex = (case when sex = 'm' then 'f' else 'm' end);