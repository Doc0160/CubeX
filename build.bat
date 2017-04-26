@echo off
ibt -begin CubeX.ibt
go-bindata -debug -prefix web web
go build
set LastError=%ERRORLEVEL%
ibt -end CubeX.ibt %LastError%
