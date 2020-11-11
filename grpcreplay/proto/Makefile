# Copyright 2019 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Makefile for building Go files from protos.

# Change these to match your environment.
PROTOC=/usr/local/bin/protoc
PROTOC_GO_PLUGIN_DIR=$(HOME)/work/backend/platform/bin

.PHONY: gen-protos

gen-protos:
	for d in *; do \
            if [ -d $$d ]; then \
	        PATH="$(PATH):$(PROTOC_GO_PLUGIN_DIR)" \
					$(PROTOC) --gogo_out=\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,plugins=grpc:$$d \
            -I $(HOME)/go/src/github.com/gogo/protobuf/protobuf \
	        -I $$d $$d/*.proto; \
            fi \
	done

