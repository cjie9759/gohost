SELECT DISTINCT sid,host_name,sys_info,ip,sip,mem,cpu,disk,date 
from host_infos 
WHERE date BETWEEN ${__from:date:seconds} AND ${__to:date:seconds} 
GROUP by sid order by date desc 



SELECT * from host_infos
GROUP by sid order by date desc 


SELECT $__unixEpochGroupSeconds(date, 500) as window, avg(value)
FROM host_infos
GROUP BY 1
ORDER BY 1 ASC


SELECT count(*) from host_infos;

SELECT json_extract(mem, '$.usedPercent') from host_infos

SELECT $__unixEpochGroupSeconds(date, 5) as time, avg(json_extract(mem, '$.usedPercent')) ,host_name
FROM host_infos
GROUP BY time,host_name
ORDER BY time,host_name ASC
limit 10000

SELECT json_extract(cpu, '$.Percent') as a from host_infos ORDER by a;

SELECT CAST(strftime('%s', 'now', '-1 minute') as INTEGER) as time, 4 as value 
WHERE time >= $__from / 1000 and time < $__to / 1000

SELECT  host_name,date,json_extract(disk, '$.free') as dp , json_extract(disk, '$.used') as di
from host_infos
WHERE date >= $__from / 1000 and date < $__to / 1000
GROUP by host_name
ORDER by date desc

SELECT  host_name, json_extract(disk, '$.used') /1024/1024/1024 as used,json_extract(disk, '$.free') /1024/1024/1024 as free 
from host_infos
WHERE date >= $__from / 1000 and date < $__to / 1000
GROUP by host_name
ORDER by date desc


SELECT $__unixEpochGroupSeconds(date, ${epoch}, null)  as time,host_name,1 as up
from host_infos 
-- WHERE date BETWEEN ${__from:date:seconds} AND ${__to:date:seconds}
-- WHERE date >= 1702364444312 / 1000 and date < 1702386044312 / 1000 
-- WHERE date BETWEEN '2023-12-12T07:00:44.312Z' AND '2023-12-12T13:00:44.312Z'
-- order by time asc 
GROUP BY 1,2
ORDER BY 1,2 ASC
-- limit 1000;