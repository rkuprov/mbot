// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: service.proto

package mbotpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6d, 0x62, 0x6f, 0x74, 0x1a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xaf, 0x05, 0x0a, 0x11, 0x4d, 0x42, 0x6f, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1b,
	0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x62,
	0x6f, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x6d, 0x62, 0x6f,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x50, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x41, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x59, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x62, 0x6f, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x41, 0x6c, 0x6c, 0x12, 0x20, 0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x70, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x62, 0x6f, 0x74, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x6b, 0x75, 0x70, 0x72, 0x6f, 0x76, 0x2f, 0x6d, 0x62, 0x6f, 0x74, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x62, 0x6f, 0x74, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x4d,
	0x58, 0x58, 0xaa, 0x02, 0x04, 0x4d, 0x62, 0x6f, 0x74, 0xca, 0x02, 0x04, 0x4d, 0x62, 0x6f, 0x74,
	0xe2, 0x02, 0x10, 0x4d, 0x62, 0x6f, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x4d, 0x62, 0x6f, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_service_proto_goTypes = []any{
	(*CreateCustomerRequest)(nil),       // 0: mbot.CreateCustomerRequest
	(*GetCustomerRequest)(nil),          // 1: mbot.GetCustomerRequest
	(*GetCustomersAllRequest)(nil),      // 2: mbot.GetCustomersAllRequest
	(*UpdateCustomerRequest)(nil),       // 3: mbot.UpdateCustomerRequest
	(*CreateSubscriptionRequest)(nil),   // 4: mbot.CreateSubscriptionRequest
	(*GetSubscriptionRequest)(nil),      // 5: mbot.GetSubscriptionRequest
	(*GetSubscriptionsAllRequest)(nil),  // 6: mbot.GetSubscriptionsAllRequest
	(*UpdateSubscriptionRequest)(nil),   // 7: mbot.UpdateSubscriptionRequest
	(*CreateCustomerResponse)(nil),      // 8: mbot.CreateCustomerResponse
	(*GetCustomerResponse)(nil),         // 9: mbot.GetCustomerResponse
	(*GetCustomersAllResponse)(nil),     // 10: mbot.GetCustomersAllResponse
	(*UpdateCustomerResponse)(nil),      // 11: mbot.UpdateCustomerResponse
	(*CreateSubscriptionResponse)(nil),  // 12: mbot.CreateSubscriptionResponse
	(*GetSubscriptionResponse)(nil),     // 13: mbot.GetSubscriptionResponse
	(*GetSubscriptionsAllResponse)(nil), // 14: mbot.GetSubscriptionsAllResponse
	(*UpdateSubscriptionResponse)(nil),  // 15: mbot.UpdateSubscriptionResponse
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: mbot.MBotServerService.CreateCustomer:input_type -> mbot.CreateCustomerRequest
	1,  // 1: mbot.MBotServerService.GetCustomer:input_type -> mbot.GetCustomerRequest
	2,  // 2: mbot.MBotServerService.GetCustomersAll:input_type -> mbot.GetCustomersAllRequest
	3,  // 3: mbot.MBotServerService.UpdateCustomer:input_type -> mbot.UpdateCustomerRequest
	4,  // 4: mbot.MBotServerService.CreateSubscription:input_type -> mbot.CreateSubscriptionRequest
	5,  // 5: mbot.MBotServerService.GetSubscription:input_type -> mbot.GetSubscriptionRequest
	6,  // 6: mbot.MBotServerService.GetSubscriptionsAll:input_type -> mbot.GetSubscriptionsAllRequest
	7,  // 7: mbot.MBotServerService.UpdateSubscription:input_type -> mbot.UpdateSubscriptionRequest
	8,  // 8: mbot.MBotServerService.CreateCustomer:output_type -> mbot.CreateCustomerResponse
	9,  // 9: mbot.MBotServerService.GetCustomer:output_type -> mbot.GetCustomerResponse
	10, // 10: mbot.MBotServerService.GetCustomersAll:output_type -> mbot.GetCustomersAllResponse
	11, // 11: mbot.MBotServerService.UpdateCustomer:output_type -> mbot.UpdateCustomerResponse
	12, // 12: mbot.MBotServerService.CreateSubscription:output_type -> mbot.CreateSubscriptionResponse
	13, // 13: mbot.MBotServerService.GetSubscription:output_type -> mbot.GetSubscriptionResponse
	14, // 14: mbot.MBotServerService.GetSubscriptionsAll:output_type -> mbot.GetSubscriptionsAllResponse
	15, // 15: mbot.MBotServerService.UpdateSubscription:output_type -> mbot.UpdateSubscriptionResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_customer_proto_init()
	file_stats_proto_init()
	file_subscription_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
