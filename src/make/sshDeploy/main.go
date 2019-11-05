package main

import (
	"fmt"
	"github.com/tachyon-protocol/udw/udwConsole"
	"github.com/tachyon-protocol/udw/udwGoSource/udwGoBuild"
	"github.com/tachyon-protocol/udw/udwSsh"
	"path/filepath"
)

func main() {
	udwConsole.MustRunCommandLineFromFuncV2(func(req struct {
		Ip      string
		PkgPath string
		Command string
	}) {
		const (
			_os  = `linux`
			arch = `amd64`
		)
		resp := udwGoBuild.MustBuild(udwGoBuild.BuildRequest{
			PkgPath:       req.PkgPath,
			TargetOs:      _os,
			TargetCpuArch: arch,
			EnableRace:    false,
		})
		pkgName := filepath.Base(req.PkgPath)
		fmt.Println("build successfully", pkgName, _os, "/", arch)
		udwSsh.MustScpToRemoteDefault(req.Ip, resp.GetOutputExeFilePath(), "/tmp/"+pkgName)
		//udwSsh.MustRpcSshDefault(serverIp, "mv /tmp/"+pkgName+" /usr/local/bin/"+pkgName+";killall "+pkgName+";setsid "+pkgName+" >> /tmp/server.log 2>&1")
		udwSsh.MustRpcSshDefault(req.Ip, "mv /tmp/"+pkgName+" /usr/local/bin/"+pkgName+";killall "+pkgName+";"+req.Command)
	})
}
