import asyncio

from grpclib.client import Channel

from ozonmp.ise_apartment_api.v1.ise_apartment_api_grpc import IseApartmentApiServiceStub
from ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2 import DescribeApartmentV1Request

async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = IseApartmentApiServiceStub(channel)

        req = DescribeApartmentV1Request(apartment_id=1)
        reply = await client.DescribeApartmentV1(req)
        print(reply.message)


if __name__ == '__main__':
    asyncio.run(main())
