// +build !windows

package main

import (
	"github.com/docker/docker/pkg/integration/checker"
	"github.com/go-check/check"
)

func (s *DockerSuite) TestInspectOomKilledTrue(c *check.C) {
	testRequires(c, DaemonIsLinux, memoryLimitSupport)

	name := "testoomkilled"
	_, exitCode, _ := dockerCmdWithError("run", "--name", name, "--memory", "32MB", "busybox", "sh", "-c", "x=a; while true; do x=$x$x$x$x; done")

	c.Assert(exitCode, checker.Equals, 137, check.Commentf("OOM exit should be 137"))

	oomKilled, err := inspectField(name, "State.OOMKilled")
	c.Assert(oomKilled, checker.Equals, "true")
	c.Assert(err, checker.IsNil)
}

func (s *DockerSuite) TestInspectOomKilledFalse(c *check.C) {
	testRequires(c, DaemonIsLinux, memoryLimitSupport)

	name := "testoomkilled"
	dockerCmd(c, "run", "--name", name, "--memory", "32MB", "busybox", "sh", "-c", "echo hello world")

	oomKilled, err := inspectField(name, "State.OOMKilled")
	c.Assert(oomKilled, checker.Equals, "false")
	c.Assert(err, checker.IsNil)
}
