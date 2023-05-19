dev: database
	watchexec -c -r -e go -- go run .

database: 
	systemctl is-active --quiet postgresql || systemctl start postgresql

.PHONY: dev database
