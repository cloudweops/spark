package impl_test

import (
	"context"

	"github.com/cloudweops/phoenix/app"
	"github.com/cloudweops/spark/apps/user"
	"github.com/cloudweops/spark/test/tools"
)

var impl user.RPCServer
var ctx = context.Background()

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(user.AppName).(user.RPCServer)
}
