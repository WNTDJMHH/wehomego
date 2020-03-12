/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

#include <iostream>
#include <memory>
#include <string>
#include <cstdio>

#include "rocksdb/db.h"
#include "rocksdb/slice.h"
#include "rocksdb/options.h"

#include <grpcpp/grpcpp.h>
#include <grpcpp/health_check_service_interface.h>
#include <grpcpp/ext/proto_server_reflection_plugin.h>

#ifdef BAZEL_BUILD
#include "examples/protos/whstorekv.grpc.pb.h"
#else
#include "genfiles/whstorekv.grpc.pb.h"
#endif

using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;
using namespace  whstorekv;
using namespace  std;

// Logic and data behind the server's behavior.
class WhStoreKvServiceImpl final : public WhStoreKv::Service {
  Status SayHello(ServerContext* context, const HelloRequest* ptrReq,
                  HelloReply* ptrRsp) override {
    std::string prefix("Hello ");
    ptrRsp->set_message(prefix + ptrReq->name());
    return Status::OK;
  }

  Status Get(ServerContext* context, const GetReq * ptrReq,
                  GetRsp * ptrRsp) override {
    std::string prefix("Hello ");
		cout << "ReqPb" << ptrReq->ShortDebugString() <<endl;
		// get value
		std::string strValue;
		auto  s = m_ptrDb->Get(rocksdb::ReadOptions(), ptrReq->key(), &strValue);
		ptrRsp->mutable_data()->ParseFromString(strValue);
		cout << "RsqPb" << ptrRsp->ShortDebugString() <<endl;
		return Status::OK;
	}

	Status Set(ServerContext* context, const SetReq * ptrReq,
					SetRsp * ptrRsp) override {
			std::string prefix("Hello ");
			cout << "ReqPb" << ptrReq->ShortDebugString() <<endl;
			// Put key-value
			auto  s = m_ptrDb->Put(rocksdb::WriteOptions(), ptrReq->key(), ptrReq->data().SerializeAsString());
			cout << "Status:" << s.ok() << endl;
			return Status::OK;
	}

  Status BatchGet(ServerContext* context, const BatchGetReq * ptrReq,
                  BatchGetRsp * ptrRsp) override {
    std::string prefix("Hello ");
		cout << "ReqPb" << ptrReq->ShortDebugString() <<endl;
    return Status::OK;
  }

  Status SearchKeys(ServerContext* context, const SearchKeysReq * ptrReq,
                  SearchKeysRsp * ptrRsp) override {
    std::string prefix("Hello ");
		cout << "ReqPb" << ptrReq->ShortDebugString() <<endl;
    return Status::OK;
  }

	public:
		rocksdb::DB* m_ptrDb; 
};

void RunServer() {
  std::string server_address("0.0.0.0:50050");
  WhStoreKvServiceImpl service;

	rocksdb::DB * ptrDb;
	std::string kDBPath = "/tmp/rocksdb_simple_example";
  rocksdb::Options options;
  options.IncreaseParallelism();
  options.OptimizeLevelStyleCompaction();
  options.create_if_missing = true;
  rocksdb::Status s = rocksdb::DB::Open(options, kDBPath, &ptrDb);
	assert(s.ok());
  std::cout << "DbInitOk" << kDBPath << std::endl;
	service.m_ptrDb = ptrDb;

  grpc::EnableDefaultHealthCheckService(true);
  grpc::reflection::InitProtoReflectionServerBuilderPlugin();
  ServerBuilder builder;
  builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());
  builder.RegisterService(&service);
  std::unique_ptr<Server> server(builder.BuildAndStart());
  std::cout << "Server listening on " << server_address << std::endl;

  server->Wait();
}

int main(int argc, char** argv) {
  RunServer();

  return 0;
}
