# Luca Comba

pdf:
	pandoc -V colorlinks=true \
	-V linkcolor=blue \
	-V urlcolor=blue \
	-V toccolor=gray ./report.md -o report.pdf
