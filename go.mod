module github.com/hardw01f/Vulnerability-goapp

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/hardw01f/vulnerability-goapp v0.0.0-20200608095627-9a1291f40edf // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

replace github.com/hardw01f/Vulnerability-goapp/pkg/admin => ./pkg/admin

replace github.com/hardw01f/Vulnerability-goapp/pkg/cookei => ./pkg/cookie

replace github.com/hardw01f/Vulnerability-goapp/pkg/user => ./pkg/user

replace github.com/hardw01f/Vulnerability-goapp/pkg/image => ./pkg/image

replace github.com/hardw01f/Vulnerability-goapp/pkg/login => ./pkg/login

replace github.com/hardw01f/Vulnerability-goapp/pkg/logout => ./pkg/logout

replace github.com/hardw01f/Vulnerability-goapp/pkg/post => ./pkg/post

replace github.com/hardw01f/Vulnerability-goapp/pkg/register => ./pkg/register

replace github.com/hardw01f/Vulnerability-goapp/pkg/search => ./pkg/search
