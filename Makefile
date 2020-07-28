.PHONY: none
none:

.PHONY: testdata
testdata:
	rm -f dstest/*.gen.go
	go run ./internal/generate-testdata -latitude 34.0219 -longitude -118.4814 -time 1556668800 > dstest/santamonica_20190501.gen.go
	go run ./internal/generate-testdata -latitude 34.0219 -longitude -118.4814 -extend hourly -units si > dstest/santamonica_hourly_si.gen.go
	go run ./internal/generate-testdata -latitude 34.0219 -longitude -118.4814 -exclude alerts,currently,daily,flags,minutely -units si > dstest/santamonica_exclude_si.gen.go
	go run ./internal/generate-testdata -latitude 34.0219 -longitude -118.4814 -exclude alerts,currently,daily,flags,minutely -extend hourly -units si > dstest/santamonica_exclude_hourly_si.gen.go
	go run ./internal/generate-testdata -latitude 34.0219 -longitude -118.4814 -lang fr > dstest/santamonica_fr.gen.go
