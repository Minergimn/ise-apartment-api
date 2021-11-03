# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ozonmp/ise_apartment_api/v1/ise_apartment_api.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from validate import validate_pb2 as validate_dot_validate__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='ozonmp/ise_apartment_api/v1/ise_apartment_api.proto',
  package='ozonmp.ise_apartment_api.v1',
  syntax='proto3',
  serialized_options=_b('ZKgithub.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api;ise_apartment_api'),
  serialized_pb=_b('\n3ozonmp/ise_apartment_api/v1/ise_apartment_api.proto\x12\x1bozonmp.ise_apartment_api.v1\x1a\x17validate/validate.proto\x1a\x1cgoogle/api/annotations.proto\"j\n\tApartment\x12\x17\n\x02id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x02id\x12\"\n\x06object\x18\x02 \x01(\tB\n\xfa\x42\x07r\x05\x10\x01\x18\xe8\x07R\x06object\x12 \n\x05owner\x18\x03 \x01(\tB\n\xfa\x42\x07r\x05\x10\x01\x18\xe8\x07R\x05owner\"H\n\x1a\x44\x65scribeApartmentV1Request\x12*\n\x0c\x61partment_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x0b\x61partmentId\"[\n\x1b\x44\x65scribeApartmentV1Response\x12<\n\x05value\x18\x01 \x01(\x0b\x32&.ozonmp.ise_apartment_api.v1.ApartmentR\x05value\"b\n\x18\x43reateApartmentV1Request\x12\x46\n\x05value\x18\x01 \x01(\x0b\x32&.ozonmp.ise_apartment_api.v1.ApartmentB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01R\x05value\">\n\x19\x43reateApartmentV1Response\x12!\n\x0c\x61partment_id\x18\x01 \x01(\x04R\x0b\x61partmentId\"\xad\x02\n\x17ListApartmentsV1Request\x12\x61\n\x06params\x18\x01 \x01(\x0b\x32I.ozonmp.ise_apartment_api.v1.ListApartmentsV1Request.ListApartmentsParamsR\x06params\x1a\xae\x01\n\x14ListApartmentsParams\x12\x10\n\x03ids\x18\x01 \x03(\x04R\x03ids\x12\x1f\n\x06offset\x18\x02 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x06offset\x12\x1d\n\x05limit\x18\x03 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x05limit\x12\"\n\x06object\x18\x04 \x01(\tB\n\xfa\x42\x07r\x05\x10\x01\x18\xe8\x07R\x06object\x12 \n\x05owner\x18\x05 \x01(\tB\n\xfa\x42\x07r\x05\x10\x01\x18\xe8\x07R\x05owner\"X\n\x18ListApartmentsV1Response\x12<\n\x05items\x18\x01 \x03(\x0b\x32&.ozonmp.ise_apartment_api.v1.ApartmentR\x05items\"F\n\x18RemoveApartmentV1Request\x12*\n\x0c\x61partment_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x0b\x61partmentId\"1\n\x19RemoveApartmentV1Response\x12\x14\n\x05\x66ound\x18\x01 \x01(\x08R\x05\x66ound2\xc8\x05\n\x16IseApartmentApiService\x12\xaf\x01\n\x13\x44\x65scribeApartmentV1\x12\x37.ozonmp.ise_apartment_api.v1.DescribeApartmentV1Request\x1a\x38.ozonmp.ise_apartment_api.v1.DescribeApartmentV1Response\"%\x82\xd3\xe4\x93\x02\x1f\x12\x1d/v1/apartments/{apartment_id}\x12\xa8\x01\n\x11\x43reateApartmentV1\x12\x35.ozonmp.ise_apartment_api.v1.CreateApartmentV1Request\x1a\x36.ozonmp.ise_apartment_api.v1.CreateApartmentV1Response\"$\x82\xd3\xe4\x93\x02\x1e\"\x15/v1/apartments/create:\x05value\x12\xa4\x01\n\x10ListApartmentsV1\x12\x34.ozonmp.ise_apartment_api.v1.ListApartmentsV1Request\x1a\x35.ozonmp.ise_apartment_api.v1.ListApartmentsV1Response\"#\x82\xd3\xe4\x93\x02\x1d\"\x13/v1/apartments/list:\x06params\x12\xa9\x01\n\x11RemoveApartmentV1\x12\x35.ozonmp.ise_apartment_api.v1.RemoveApartmentV1Request\x1a\x36.ozonmp.ise_apartment_api.v1.RemoveApartmentV1Response\"%\x82\xd3\xe4\x93\x02\x1f*\x1d/v1/apartments/{apartment_id}BMZKgithub.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api;ise_apartment_apib\x06proto3')
  ,
  dependencies=[validate_dot_validate__pb2.DESCRIPTOR,google_dot_api_dot_annotations__pb2.DESCRIPTOR,])




_APARTMENT = _descriptor.Descriptor(
  name='Apartment',
  full_name='ozonmp.ise_apartment_api.v1.Apartment',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='ozonmp.ise_apartment_api.v1.Apartment.id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\0042\002 \000'), json_name='id', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='object', full_name='ozonmp.ise_apartment_api.v1.Apartment.object', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\007r\005\020\001\030\350\007'), json_name='object', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='owner', full_name='ozonmp.ise_apartment_api.v1.Apartment.owner', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\007r\005\020\001\030\350\007'), json_name='owner', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=139,
  serialized_end=245,
)


_DESCRIBEAPARTMENTV1REQUEST = _descriptor.Descriptor(
  name='DescribeApartmentV1Request',
  full_name='ozonmp.ise_apartment_api.v1.DescribeApartmentV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='apartment_id', full_name='ozonmp.ise_apartment_api.v1.DescribeApartmentV1Request.apartment_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\0042\002 \000'), json_name='apartmentId', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=247,
  serialized_end=319,
)


_DESCRIBEAPARTMENTV1RESPONSE = _descriptor.Descriptor(
  name='DescribeApartmentV1Response',
  full_name='ozonmp.ise_apartment_api.v1.DescribeApartmentV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='ozonmp.ise_apartment_api.v1.DescribeApartmentV1Response.value', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='value', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=321,
  serialized_end=412,
)


_CREATEAPARTMENTV1REQUEST = _descriptor.Descriptor(
  name='CreateApartmentV1Request',
  full_name='ozonmp.ise_apartment_api.v1.CreateApartmentV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='ozonmp.ise_apartment_api.v1.CreateApartmentV1Request.value', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\005\212\001\002\020\001'), json_name='value', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=414,
  serialized_end=512,
)


_CREATEAPARTMENTV1RESPONSE = _descriptor.Descriptor(
  name='CreateApartmentV1Response',
  full_name='ozonmp.ise_apartment_api.v1.CreateApartmentV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='apartment_id', full_name='ozonmp.ise_apartment_api.v1.CreateApartmentV1Response.apartment_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='apartmentId', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=514,
  serialized_end=576,
)


_LISTAPARTMENTSV1REQUEST_LISTAPARTMENTSPARAMS = _descriptor.Descriptor(
  name='ListApartmentsParams',
  full_name='ozonmp.ise_apartment_api.v1.ListApartmentsV1Request.ListApartmentsParams',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ids', full_name='ozonmp.ise_apartment_api.v1.ListApartmentsV1Request.ListApartmentsParams.ids', index=0,
      number=1, type=4, cpp_type=4, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='ids', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='offset', full_name='ozonmp.ise_apartment_api.v1.ListApartmentsV1Request.ListApartmentsParams.offset', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\0042\002 \000'), json_name='offset', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='limit', full_name='ozonmp.ise_apartment_api.v1.ListApartmentsV1Request.ListApartmentsParams.limit', index=2,
      number=3, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\0042\002 \000'), json_name='limit', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='object', full_name='ozonmp.ise_apartment_api.v1.ListApartmentsV1Request.ListApartmentsParams.object', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\007r\005\020\001\030\350\007'), json_name='object', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='owner', full_name='ozonmp.ise_apartment_api.v1.ListApartmentsV1Request.ListApartmentsParams.owner', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\007r\005\020\001\030\350\007'), json_name='owner', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=706,
  serialized_end=880,
)

_LISTAPARTMENTSV1REQUEST = _descriptor.Descriptor(
  name='ListApartmentsV1Request',
  full_name='ozonmp.ise_apartment_api.v1.ListApartmentsV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='params', full_name='ozonmp.ise_apartment_api.v1.ListApartmentsV1Request.params', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='params', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_LISTAPARTMENTSV1REQUEST_LISTAPARTMENTSPARAMS, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=579,
  serialized_end=880,
)


_LISTAPARTMENTSV1RESPONSE = _descriptor.Descriptor(
  name='ListApartmentsV1Response',
  full_name='ozonmp.ise_apartment_api.v1.ListApartmentsV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='items', full_name='ozonmp.ise_apartment_api.v1.ListApartmentsV1Response.items', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='items', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=882,
  serialized_end=970,
)


_REMOVEAPARTMENTV1REQUEST = _descriptor.Descriptor(
  name='RemoveApartmentV1Request',
  full_name='ozonmp.ise_apartment_api.v1.RemoveApartmentV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='apartment_id', full_name='ozonmp.ise_apartment_api.v1.RemoveApartmentV1Request.apartment_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\0042\002 \000'), json_name='apartmentId', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=972,
  serialized_end=1042,
)


_REMOVEAPARTMENTV1RESPONSE = _descriptor.Descriptor(
  name='RemoveApartmentV1Response',
  full_name='ozonmp.ise_apartment_api.v1.RemoveApartmentV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='found', full_name='ozonmp.ise_apartment_api.v1.RemoveApartmentV1Response.found', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='found', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1044,
  serialized_end=1093,
)

_DESCRIBEAPARTMENTV1RESPONSE.fields_by_name['value'].message_type = _APARTMENT
_CREATEAPARTMENTV1REQUEST.fields_by_name['value'].message_type = _APARTMENT
_LISTAPARTMENTSV1REQUEST_LISTAPARTMENTSPARAMS.containing_type = _LISTAPARTMENTSV1REQUEST
_LISTAPARTMENTSV1REQUEST.fields_by_name['params'].message_type = _LISTAPARTMENTSV1REQUEST_LISTAPARTMENTSPARAMS
_LISTAPARTMENTSV1RESPONSE.fields_by_name['items'].message_type = _APARTMENT
DESCRIPTOR.message_types_by_name['Apartment'] = _APARTMENT
DESCRIPTOR.message_types_by_name['DescribeApartmentV1Request'] = _DESCRIBEAPARTMENTV1REQUEST
DESCRIPTOR.message_types_by_name['DescribeApartmentV1Response'] = _DESCRIBEAPARTMENTV1RESPONSE
DESCRIPTOR.message_types_by_name['CreateApartmentV1Request'] = _CREATEAPARTMENTV1REQUEST
DESCRIPTOR.message_types_by_name['CreateApartmentV1Response'] = _CREATEAPARTMENTV1RESPONSE
DESCRIPTOR.message_types_by_name['ListApartmentsV1Request'] = _LISTAPARTMENTSV1REQUEST
DESCRIPTOR.message_types_by_name['ListApartmentsV1Response'] = _LISTAPARTMENTSV1RESPONSE
DESCRIPTOR.message_types_by_name['RemoveApartmentV1Request'] = _REMOVEAPARTMENTV1REQUEST
DESCRIPTOR.message_types_by_name['RemoveApartmentV1Response'] = _REMOVEAPARTMENTV1RESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Apartment = _reflection.GeneratedProtocolMessageType('Apartment', (_message.Message,), dict(
  DESCRIPTOR = _APARTMENT,
  __module__ = 'ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_apartment_api.v1.Apartment)
  ))
_sym_db.RegisterMessage(Apartment)

DescribeApartmentV1Request = _reflection.GeneratedProtocolMessageType('DescribeApartmentV1Request', (_message.Message,), dict(
  DESCRIPTOR = _DESCRIBEAPARTMENTV1REQUEST,
  __module__ = 'ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_apartment_api.v1.DescribeApartmentV1Request)
  ))
_sym_db.RegisterMessage(DescribeApartmentV1Request)

DescribeApartmentV1Response = _reflection.GeneratedProtocolMessageType('DescribeApartmentV1Response', (_message.Message,), dict(
  DESCRIPTOR = _DESCRIBEAPARTMENTV1RESPONSE,
  __module__ = 'ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_apartment_api.v1.DescribeApartmentV1Response)
  ))
_sym_db.RegisterMessage(DescribeApartmentV1Response)

CreateApartmentV1Request = _reflection.GeneratedProtocolMessageType('CreateApartmentV1Request', (_message.Message,), dict(
  DESCRIPTOR = _CREATEAPARTMENTV1REQUEST,
  __module__ = 'ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_apartment_api.v1.CreateApartmentV1Request)
  ))
_sym_db.RegisterMessage(CreateApartmentV1Request)

CreateApartmentV1Response = _reflection.GeneratedProtocolMessageType('CreateApartmentV1Response', (_message.Message,), dict(
  DESCRIPTOR = _CREATEAPARTMENTV1RESPONSE,
  __module__ = 'ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_apartment_api.v1.CreateApartmentV1Response)
  ))
_sym_db.RegisterMessage(CreateApartmentV1Response)

ListApartmentsV1Request = _reflection.GeneratedProtocolMessageType('ListApartmentsV1Request', (_message.Message,), dict(

  ListApartmentsParams = _reflection.GeneratedProtocolMessageType('ListApartmentsParams', (_message.Message,), dict(
    DESCRIPTOR = _LISTAPARTMENTSV1REQUEST_LISTAPARTMENTSPARAMS,
    __module__ = 'ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2'
    # @@protoc_insertion_point(class_scope:ozonmp.ise_apartment_api.v1.ListApartmentsV1Request.ListApartmentsParams)
    ))
  ,
  DESCRIPTOR = _LISTAPARTMENTSV1REQUEST,
  __module__ = 'ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_apartment_api.v1.ListApartmentsV1Request)
  ))
_sym_db.RegisterMessage(ListApartmentsV1Request)
_sym_db.RegisterMessage(ListApartmentsV1Request.ListApartmentsParams)

ListApartmentsV1Response = _reflection.GeneratedProtocolMessageType('ListApartmentsV1Response', (_message.Message,), dict(
  DESCRIPTOR = _LISTAPARTMENTSV1RESPONSE,
  __module__ = 'ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_apartment_api.v1.ListApartmentsV1Response)
  ))
_sym_db.RegisterMessage(ListApartmentsV1Response)

RemoveApartmentV1Request = _reflection.GeneratedProtocolMessageType('RemoveApartmentV1Request', (_message.Message,), dict(
  DESCRIPTOR = _REMOVEAPARTMENTV1REQUEST,
  __module__ = 'ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_apartment_api.v1.RemoveApartmentV1Request)
  ))
_sym_db.RegisterMessage(RemoveApartmentV1Request)

RemoveApartmentV1Response = _reflection.GeneratedProtocolMessageType('RemoveApartmentV1Response', (_message.Message,), dict(
  DESCRIPTOR = _REMOVEAPARTMENTV1RESPONSE,
  __module__ = 'ozonmp.ise_apartment_api.v1.ise_apartment_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_apartment_api.v1.RemoveApartmentV1Response)
  ))
_sym_db.RegisterMessage(RemoveApartmentV1Response)


DESCRIPTOR._options = None
_APARTMENT.fields_by_name['id']._options = None
_APARTMENT.fields_by_name['object']._options = None
_APARTMENT.fields_by_name['owner']._options = None
_DESCRIBEAPARTMENTV1REQUEST.fields_by_name['apartment_id']._options = None
_CREATEAPARTMENTV1REQUEST.fields_by_name['value']._options = None
_LISTAPARTMENTSV1REQUEST_LISTAPARTMENTSPARAMS.fields_by_name['offset']._options = None
_LISTAPARTMENTSV1REQUEST_LISTAPARTMENTSPARAMS.fields_by_name['limit']._options = None
_LISTAPARTMENTSV1REQUEST_LISTAPARTMENTSPARAMS.fields_by_name['object']._options = None
_LISTAPARTMENTSV1REQUEST_LISTAPARTMENTSPARAMS.fields_by_name['owner']._options = None
_REMOVEAPARTMENTV1REQUEST.fields_by_name['apartment_id']._options = None

_ISEAPARTMENTAPISERVICE = _descriptor.ServiceDescriptor(
  name='IseApartmentApiService',
  full_name='ozonmp.ise_apartment_api.v1.IseApartmentApiService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1096,
  serialized_end=1808,
  methods=[
  _descriptor.MethodDescriptor(
    name='DescribeApartmentV1',
    full_name='ozonmp.ise_apartment_api.v1.IseApartmentApiService.DescribeApartmentV1',
    index=0,
    containing_service=None,
    input_type=_DESCRIBEAPARTMENTV1REQUEST,
    output_type=_DESCRIBEAPARTMENTV1RESPONSE,
    serialized_options=_b('\202\323\344\223\002\037\022\035/v1/apartments/{apartment_id}'),
  ),
  _descriptor.MethodDescriptor(
    name='CreateApartmentV1',
    full_name='ozonmp.ise_apartment_api.v1.IseApartmentApiService.CreateApartmentV1',
    index=1,
    containing_service=None,
    input_type=_CREATEAPARTMENTV1REQUEST,
    output_type=_CREATEAPARTMENTV1RESPONSE,
    serialized_options=_b('\202\323\344\223\002\036\"\025/v1/apartments/create:\005value'),
  ),
  _descriptor.MethodDescriptor(
    name='ListApartmentsV1',
    full_name='ozonmp.ise_apartment_api.v1.IseApartmentApiService.ListApartmentsV1',
    index=2,
    containing_service=None,
    input_type=_LISTAPARTMENTSV1REQUEST,
    output_type=_LISTAPARTMENTSV1RESPONSE,
    serialized_options=_b('\202\323\344\223\002\035\"\023/v1/apartments/list:\006params'),
  ),
  _descriptor.MethodDescriptor(
    name='RemoveApartmentV1',
    full_name='ozonmp.ise_apartment_api.v1.IseApartmentApiService.RemoveApartmentV1',
    index=3,
    containing_service=None,
    input_type=_REMOVEAPARTMENTV1REQUEST,
    output_type=_REMOVEAPARTMENTV1RESPONSE,
    serialized_options=_b('\202\323\344\223\002\037*\035/v1/apartments/{apartment_id}'),
  ),
])
_sym_db.RegisterServiceDescriptor(_ISEAPARTMENTAPISERVICE)

DESCRIPTOR.services_by_name['IseApartmentApiService'] = _ISEAPARTMENTAPISERVICE

# @@protoc_insertion_point(module_scope)
