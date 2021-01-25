package repo

var (
	portInserSQL = `
insert into ports (id,name,city,country,alias,regions,coordinates,province,timezone,unlocs,code)
values (?,?,?,?,?,?,?,?,?,?,?);
`
	portUpdateSQL = `
update ports set
name = ?,
city = ?,
country = ?,
alias = ?,
regions = ?,
coordinates = ?,
province = ?,
timezone = ?,
unlocs = ?
WHERE code = ?
`
	postSelectSQL = `
SELECT 
name,
city,
country,
alias,
regions,
coordinates,
province,
timezone,
unlocs,
code,
id
FROM ports WHERE code = ? LIMIT 1;
`
)
