/*
 *    Copyright (C) 2015      Dominik Schmidt
 *                  2015-2017 Christian Muehlhaeuser
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Dominik Schmidt <domme@tomahawk-player.org>
 *      Christian Muehlhaeuser <muesli@gmail.com>
 */

package execbee

import (
	"github.com/muesli/beehive/bees"
)

// ExecBeeFactory is a factory for ExecBees.
type ExecBeeFactory struct {
	bees.BeeFactory
}

// New returns a new Bee instance configured with the supplied options.
func (factory *ExecBeeFactory) New(name, description string, options bees.BeeOptions) bees.BeeInterface {
	bee := ExecBee{
		Bee: bees.NewBee(name, factory.Name(), description, options),
	}
	bee.ReloadOptions(options)

	return &bee
}

// Name returns the name of this Bee.
func (factory *ExecBeeFactory) Name() string {
	return "execbee"
}

// Description returns the description of this Bee.
func (factory *ExecBeeFactory) Description() string {
	return "A bee that allows executing commands"
}

// Image returns the filename of an image for this Bee.
// func (factory *ExecBeeFactory) Image() string {
// 	return factory.Name() + ".png"
// }

// Events describes the available events provided by this Bee.
func (factory *ExecBeeFactory) Events() []bees.EventDescriptor {
	events := []bees.EventDescriptor{
		{
			Namespace:   factory.Name(),
			Name:        "cmd",
			Description: "A command was executed",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "stdout",
					Description: "std output of the executed command",
					Type:        "string",
				},
				{
					Name:        "stderr",
					Description: "stderr output of the executed command",
					Type:        "string",
				},
			},
		},
	}
	return events
}

// Actions describes the available actions provided by this Bee.
func (factory *ExecBeeFactory) Actions() []bees.ActionDescriptor {
	actions := []bees.ActionDescriptor{
		{
			Namespace:   factory.Name(),
			Name:        "localCommand",
			Description: "Executes a command on the local host",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "command",
					Description: "command to be executed",
					Type:        "string",
				},
			},
		},
	}
	return actions
}

func init() {
	f := ExecBeeFactory{}
	bees.RegisterFactory(&f)
}
