
# 正则表达式 regexp
select  patient_id, patient_name, conditions
from Patients
where conditions regexp "(\\s)*DIAB1"