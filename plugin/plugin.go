// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import (
	"os"

	"github.com/drone-plugins/drone-plugin-lib/drone"
)

// Plugin implements drone.Plugin to provide the plugin implementation.
type Plugin struct {
	settings Settings
	pipeline drone.Pipeline
	network  drone.Network
	owner    string
	repo     string
}

// New initializes a plugin from the given Settings, Pipeline, and Network.
func New(settings Settings, pipeline drone.Pipeline, network drone.Network) drone.Plugin {
	// Retrieve owner and repo from environment variables
	owner := os.Getenv("DRONE_REPO_NAMESPACE")
	repo := os.Getenv("DRONE_REPO_NAME")

	return &Plugin{
		settings: settings,
		pipeline: pipeline,
		network:  network,
		owner:    owner,
		repo:     repo,
	}
}
