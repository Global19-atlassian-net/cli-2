echo - Active state: {{.Project.Owner}}/{{.Project.Name}}

{{range $K, $V := .Env}}
set -xg {{$K}} "{{$V}}"
{{end}}

{{range $K, $CMD := .Scripts}}
alias {{$K}}='state run {{$CMD}}'
{{end}}

cd "{{.WD}}"

{{.UserScripts}}