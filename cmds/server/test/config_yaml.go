/*
 Copyright (c) Facebook, Inc. and its affiliates.

 This source code is licensed under the MIT license found in the
 LICENSE file in the root directory of this source tree.
*/

package test

import "os"

// SampleConfig ...
const SampleConfig = `
# resuable references
authenticator_type_bcrypt: &authenticator_type_bcrypt 1

privlvl_root: &privlvl_root 15

action_deny: &action_deny 1
action_permit: &action_permit 2

logger_type_stderr: &logger_type_stderr 1
logger_type_syslog: &logger_type_syslog 2
logger_type_file: &logger_type_file 3

logger_stderr: &logger_stderr
  name: stderr
  type: *logger_type_stderr

logger_syslog: &logger_syslog
  name: syslog
  type: *logger_type_syslog

logger_file: &logger_file
  name: file
  type: *logger_type_file

bcrypt: &bcrypt
  type: *authenticator_type_bcrypt
  options:
    # password
    hash: 24326124313024614d6761663134486e35366b6a734b2f79564a384b2e577678754c6b34314364586a4d727a6276794a7844304c4371757345765171

# services
cmd: &cmd
  name: cmd
  is_optional: true
  set_values:
    - name: priv-lvl
      values: [15]
      is_optional: true

cisco_avp: &cisco_avp
  name: cisco-av-pair
  is_optional: true
  set_values:
    - name: shell:roles
      values: [ admin ]
      is_optional: true
    - name: shell:roles
      values: [ network-admin vdc-admin ]
      is_optional: true

# commands
conf_t: &conf_t
  name: configure
  match: [terminal, exclusive]
  action: *action_permit

conf_b: &conf_b
  name: configure
  match: [batch]
  action: *action_permit

# groups
noc: &noc
  name: noc
  services: [*cisco_avp, *cmd]
  commands: [*conf_t, *conf_b]
  authenticator: *bcrypt
  accounter: *logger_file


# finally, declare users
users:
  - name: mr_uses_group
    scopes: ["localhost"]
    groups: [*noc]
  - name: mr_no_group
    scopes: ["localhost"]
    services: [*cisco_avp]
    commands: [*conf_t]
    authenticator: *bcrypt
    accounter: *logger_file
  - name: ms_commands_only
    scopes: ["localhost"]
    commands: [*conf_t]


handler_type_start: &handler_type_start 1
handler_type_span: &handler_type_span 2

provider_type_prefix: &provider_type_prefix 1
provider_type_dns: &provider_type_dns 2

secrets:
  - name: localhost
    secret:
      group: tacquito
      key: fooman
    handler:
      type: *handler_type_start
    type: *provider_type_prefix
    options:
      prefixes: |
        [
          "::0/0"
        ]
`

// WriteSampleConfig write config to current root
func WriteSampleConfig(path string) error {
	return os.WriteFile(path, []byte(SampleConfig), 0644)
}
