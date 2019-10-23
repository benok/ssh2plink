package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func trimArgs(argToTrim string, args []string) []string {
	var result []string
	for i := 0; i < len(args); i = i + 1 {
		if args[i] != argToTrim {
			result = append(result, args[i])
		}
	}
	return result
}

const command = "plink.exe"

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Fprintf(os.Stderr, "ssh2plink")
	} else if len(args) > 1 {
		cmd := args[0]
		cmdPath := filepath.Dir(cmd)
		if args[1] == "-V" {
			fmt.Fprint(os.Stderr, "OpenSSH_for_Windows_7.7p1, LibreSSL 2.6.5")
		} else {
			fullCmd := filepath.Join(cmdPath, command)
			args = trimArgs("-T", args[1:])

			sshCmd := exec.Command(fullCmd, args...)
			sshCmd.Stdout = os.Stdout
			sshCmd.Stderr = os.Stderr
			sshCmd.Stdin = os.Stdin

			err := sshCmd.Start()
			if err != nil {
				fmt.Fprint(os.Stderr, "failed to start sshCmd:", fullCmd, args)
				fmt.Fprintf(os.Stderr, "sshCmd.Start faild with %s", err)
			}

			err = sshCmd.Wait()
			if err != nil {
				fmt.Fprint(os.Stderr, "failed to wait sshCmd:", fullCmd, args)
				fmt.Fprintf(os.Stderr, "sshCmd.Wait faild with %s", err)
			}
		}
	}
}
