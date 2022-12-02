/*
Copyright ApeCloud Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package engine

type mysql struct{}

const (
	mysqlEngineName    = "mysql"
	mysqlClient        = "mysql"
	mysqlContainerName = "mysql"
)

var _ Interface = &mysql{}

func (m *mysql) ConnectCommand(database string) []string {
	if len(database) > 0 {
		return []string{mysqlClient, "-D", database}
	}
	return []string{mysqlClient}
}

func (m *mysql) EngineName() string {
	return mysqlEngineName
}

func (m *mysql) EngineContainer() string {
	return mysqlContainerName
}