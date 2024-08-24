import 'dart:convert';

import 'package:ecom_app/core/error/exception.dart';
import 'package:ecom_app/features/chat/data/datasources/chat_remote_data_source.dart';
import 'package:ecom_app/features/chat/data/models/user_chat_model.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late MockCustomHttpClient mockHttpClient;
  late ChatRemoteDataSourceImpl dataSource;

  setUp(() {
    mockHttpClient = MockCustomHttpClient();
    dataSource = ChatRemoteDataSourceImpl(client: mockHttpClient);
  });

  group('getChats', () {
    final tEmail = 'user@gmail.com';
    final tUri = '';

    final tChatList = [
      {
        '_id': '66c837f2b7068ee15142f66a',
        'user1': {
          '_id': '66c44fd86198f150e643c827',
          'name': 'User1',
          'email': 'user1@gmail.com'
        },
        'user2': {
          '_id': '66bde36e9bbe07fc39034cdd',
          'name': 'Mr. User',
          'email': 'user@gmail.com'
        },
      },
      {
        '_id': '66c83a67b7068ee15142f6c8',
        'user1': {
          '_id': '66c8386ab7068ee15142f684',
          'name': 'User2',
          'email': 'user2@gmail.com'
        },
        'user2': {
          '_id': '66bde36e9bbe07fc39034cdd',
          'name': 'Mr. User',
          'email': 'user@gmail.com'
        },
      }
    ];

    test('should return a list of UserChatModels when the call is successful',
        () async {
      // Arrange
      when(mockHttpClient.get(any)).thenAnswer(
        (_) async => http.Response(
          json.encode({'statusCode': 200, 'data': tChatList}),
          200,
        ),
      );

      // Act
      final result = await dataSource.getChats(tEmail);

      // Assert
      expect(result, isA<List<UserChatModel>>());
      expect(result.length, 2);
      expect(result[0].name, 'User1');
      expect(result[1].name, 'User2');
    });

    // test('should throw a ServerException when the response status is not 200',
    //     () async {
    //   // Arrange
    //   when(mockHttpClient.get(any)).thenAnswer(
    //     (_) async => http.Response('Something went wrong', 404),
    //   );

      

    //   // Assert
    //   expect(() => dataSource.getChats(tEmail), throwsA(isA<ServerException>()));
    // });

    test('should throw an Exception when an error occurs', () async {
      // Arrange
      when(mockHttpClient.get(any)).thenThrow(Exception('Failed to load'));

      // Act
      final call = dataSource.getChats;

      // Assert
      expect(() => call(tEmail), throwsA(isA<Exception>()));
    });
  });
}
