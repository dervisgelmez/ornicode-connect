package initializers

import core "ornicode-connect/initializers/packages"

func Load() {
	core.LoadEnv()
	core.ConnectDatabase()
	core.SyncDatabase()
}
