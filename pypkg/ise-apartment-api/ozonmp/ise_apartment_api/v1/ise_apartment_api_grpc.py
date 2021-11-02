# Generated by the Protocol Buffers compiler. DO NOT EDIT!
# source: ozonmp/ise_apartment_api/v1/ise_apartment_api.proto
# plugin: grpclib.plugin.main
import abc
import typing

import grpclib.const
import grpclib.client
if typing.TYPE_CHECKING:
    import grpclib.server

import validate.validate_pb2
import google.api.annotations_pb2
import google.protobuf.empty_pb2
import google.protobuf.timestamp_pb2
import ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2


class IseApartmentApiServiceBase(abc.ABC):

    @abc.abstractmethod
    async def DescribeApartmentV1(self, stream: 'grpclib.server.Stream[ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.DescribeApartmentV1Request, ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.DescribeApartmentV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def CreateApartmentV1(self, stream: 'grpclib.server.Stream[ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.CreateApartmentV1Request, ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.CreateApartmentV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def ListApartmentV1(self, stream: 'grpclib.server.Stream[ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.ListApartmentV1Request, ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.ListApartmentV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def RemoveApartmentV1(self, stream: 'grpclib.server.Stream[ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.RemoveApartmentV1Request, google.protobuf.empty_pb2.Empty]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/ozonmp.ise_apartment_api.v1.IseApartmentApiService/DescribeApartmentV1': grpclib.const.Handler(
                self.DescribeApartmentV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.DescribeApartmentV1Request,
                ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.DescribeApartmentV1Response,
            ),
            '/ozonmp.ise_apartment_api.v1.IseApartmentApiService/CreateApartmentV1': grpclib.const.Handler(
                self.CreateApartmentV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.CreateApartmentV1Request,
                ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.CreateApartmentV1Response,
            ),
            '/ozonmp.ise_apartment_api.v1.IseApartmentApiService/ListApartmentV1': grpclib.const.Handler(
                self.ListApartmentV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.ListApartmentV1Request,
                ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.ListApartmentV1Response,
            ),
            '/ozonmp.ise_apartment_api.v1.IseApartmentApiService/RemoveApartmentV1': grpclib.const.Handler(
                self.RemoveApartmentV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.RemoveApartmentV1Request,
                google.protobuf.empty_pb2.Empty,
            ),
        }


class IseApartmentApiServiceStub:

    def __init__(self, channel: grpclib.client.Channel) -> None:
        self.DescribeApartmentV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.ise_apartment_api.v1.IseApartmentApiService/DescribeApartmentV1',
            ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.DescribeApartmentV1Request,
            ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.DescribeApartmentV1Response,
        )
        self.CreateApartmentV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.ise_apartment_api.v1.IseApartmentApiService/CreateApartmentV1',
            ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.CreateApartmentV1Request,
            ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.CreateApartmentV1Response,
        )
        self.ListApartmentV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.ise_apartment_api.v1.IseApartmentApiService/ListApartmentV1',
            ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.ListApartmentV1Request,
            ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.ListApartmentV1Response,
        )
        self.RemoveApartmentV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.ise_apartment_api.v1.IseApartmentApiService/RemoveApartmentV1',
            ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2.RemoveApartmentV1Request,
            google.protobuf.empty_pb2.Empty,
        )
