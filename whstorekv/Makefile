#
# Copyright 2018 gRPC authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

HOST_SYSTEM = $(shell uname | cut -f 1 -d_)
SYSTEM ?= $(HOST_SYSTEM)
CXX = g++
CPPFLAGS += `pkg-config --cflags protobuf grpc`
CXXFLAGS += -std=c++11
LDFLAGS +=  librocksdb.a -L/usr/local/lib `pkg-config --libs protobuf grpc++ grpc`\
           -Wl,--no-as-needed -lgrpc++_reflection -Wl,--as-needed\
           -ldl -lpthread -lrt -lsnappy -lgflags -lz -lbz2 -llz4 -lzstd 
PROTOC = protoc
PROTOS_PATH = proto 
GEN_PATH = ./genfiles
GRPC_CPP_PLUGIN = grpc_cpp_plugin
GRPC_CPP_PLUGIN_PATH ?= `which $(GRPC_CPP_PLUGIN)`
PB_NAME = whstorekv
PB_OBJS = $(GEN_PATH)/$(PB_NAME).pb.o $(GEN_PATH)/$(PB_NAME).grpc.pb.o
BIN_PATH = /home/ubuntu/GoPath/bin

$(shell test -d $(GEN_PATH) || mkdir -p $(GEN_PATH))

vpath %.proto $(PROTOS_PATH)

all: system-check $(BIN_PATH)/whstorekv_client $(BIN_PATH)/whstorekv_server

$(BIN_PATH)/whstorekv_client: $(PB_OBJS) $(GEN_PATH)/whstorekv_client.o
	$(CXX) $^ $(LDFLAGS) -o $@

$(BIN_PATH)/whstorekv_server: $(PB_OBJS) $(GEN_PATH)/whstorekv_server.o
	$(CXX) $^ $(LDFLAGS) -o $@

$(GEN_PATH)/%.o:%.cpp
	g++ $^ -o $@ -c $(LDFLAGS)

.PRECIOUS: $(GEN_PATH)/$(PB_NAME).grpc.pb.cc
#$(GEN_PATH)/$(PB_NAME).grpc.pb.cc: $(PROTOS_PATH)/$(PB_NAME).proto
$(GEN_PATH)/$(PB_NAME).grpc.pb.cc: ./proto/$(PB_NAME).proto
	$(PROTOC) -I $(PROTOS_PATH) --grpc_out=$(GEN_PATH) --plugin=protoc-gen-grpc=$(GRPC_CPP_PLUGIN_PATH) $<

.PRECIOUS: $(GEN_PATH)/$(PB_NAME).pb.cc
#$(GEN_PATH)/$(PB_NAME).pb.cc: $(PROTOS_PATH)/$(PB_NAME).proto
$(GEN_PATH)/$(PB_NAME).pb.cc: ./proto/$(PB_NAME).proto
	$(PROTOC) -I $(PROTOS_PATH) --cpp_out=$(GEN_PATH)  $<

clean:
	rm -f *.o *.pb.cc *.pb.h whstorekv_client whstorekv_server
	rm -rdf $(GEN_PATH)


# The following is to test your system and ensure a smoother experience.
# They are by no means necessary to actually compile a grpc-enabled software.

PROTOC_CMD = which $(PROTOC)
PROTOC_CHECK_CMD = $(PROTOC) --version | grep -q libprotoc.3
PLUGIN_CHECK_CMD = which $(GRPC_CPP_PLUGIN)
HAS_PROTOC = $(shell $(PROTOC_CMD) > /dev/null && echo true || echo false)
ifeq ($(HAS_PROTOC),true)
HAS_VALID_PROTOC = $(shell $(PROTOC_CHECK_CMD) 2> /dev/null && echo true || echo false)
endif
HAS_PLUGIN = $(shell $(PLUGIN_CHECK_CMD) > /dev/null && echo true || echo false)

SYSTEM_OK = false
ifeq ($(HAS_VALID_PROTOC),true)
ifeq ($(HAS_PLUGIN),true)
SYSTEM_OK = true
endif
endif

system-check:
ifneq ($(HAS_VALID_PROTOC),true)
	@echo " DEPENDENCY ERROR"
	@echo
	@echo "You don't have protoc 3.0.0 installed in your path."
	@echo "Please install Google protocol buffers 3.0.0 and its compiler."
	@echo "You can find it here:"
	@echo
	@echo "   https://github.com/google/protobuf/releases/tag/v3.0.0"
	@echo
	@echo "Here is what I get when trying to evaluate your version of protoc:"
	@echo
	-$(PROTOC) --version
	@echo
	@echo
endif
ifneq ($(HAS_PLUGIN),true)
	@echo " DEPENDENCY ERROR"
	@echo
	@echo "You don't have the grpc c++ protobuf plugin installed in your path."
	@echo "Please install grpc. You can find it here:"
	@echo
	@echo "   https://github.com/grpc/grpc"
	@echo
	@echo "Here is what I get when trying to detect if you have the plugin:"
	@echo
	-which $(GRPC_CPP_PLUGIN)
	@echo
	@echo
endif
ifneq ($(SYSTEM_OK),true)
	@false
endif
