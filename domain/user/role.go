package domain

APPLICATION := map[string] uint64 {
"axul" : 1,
"fileit" : 1 << 1
}

read_role := 1
write_role := 1 << 1
update_role := 1 << 2
delete_role := 1 << 3
admin_r_role := 1 << 4 + read_role + write_role + update_role + delete_role
admin_rw_role := 1 << 5 + admin_rw_role
owner_role := 1 << 6 + owner_role

ROLE := map[string] uint64 {
"read" : read_role,
"write" : write_role,
"update" : update_role,
"delete" : delete_role,
"admin_r" : admin_r_role,
"admin_rw" : admin_rw_role,
"owner" : owner_role,
}

