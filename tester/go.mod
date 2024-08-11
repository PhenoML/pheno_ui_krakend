module pheno_ui/tester

go 1.22.5

toolchain go1.22.6

require (
	github.com/jmespath-community/go-jmespath v1.1.2-0.20240117150817-e430401a2172
	pheno_ui/maps v0.0.0-00010101000000-000000000000
)

require golang.org/x/exp v0.0.0-20240808152545-0cdaa3abc0fa // indirect

replace pheno_ui/maps => ../maps
