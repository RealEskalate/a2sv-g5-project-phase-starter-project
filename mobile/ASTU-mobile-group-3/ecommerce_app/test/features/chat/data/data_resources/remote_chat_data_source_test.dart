import 'dart:convert';
import 'package:ecommerce_app/features/auth/data/model/token_model.dart';
import 'package:ecommerce_app/features/chat/data/data_resources/remote_chat_data_source.dart';
import 'package:ecommerce_app/features/chat/domain/entity/chat.dart';
import 'package:ecommerce_app/features/chat/domain/entity/message.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:http/testing.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:http/http.dart' as http;

import '../../../../test_helper/test_helper_generation.mocks.dart';

import 'package:flutter_test/flutter_test.dart';

void main() {

  late ChatRemoteDataSourceImpl dataSource;
  late MockHttpClient mockHttpClient;
  late MockAuthLocalDataSource mockAuthLocalDataSource;

  setUp(() {
    mockHttpClient = MockHttpClient();
    mockAuthLocalDataSource = MockAuthLocalDataSource();
    dataSource = ChatRemoteDataSourceImpl(
      httpClient: mockHttpClient,
      authLocalDataSource: mockAuthLocalDataSource,
    );
  });


  group('getChatRooms', () {
    const token = 'test_token';
    final chatJson = [
      {
        '_id': 'chat123',
        'user1': {
          'name': 'John Doe',
          'email': 'john@example.com',
          '_id': 'user1_id',
          '__v': 1,
          'password': 'password123',
        },
        'user2': {
          'name': 'Jane Doe',
          'email': 'jane@example.com',
          '_id': 'user2_id',
          '__v': 2,
          'password': 'password456',
        },
      },
    ];
    final chatEntityList = [ChatEntity.fromJson(chatJson[0])];

    test('should return a list of ChatEntity when the response code is 200',
        () async {
      // Arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => TokenModel(token: token));
      when(mockHttpClient.get(any, headers: anyNamed('headers'))).thenAnswer(
        (_) async => http.Response(jsonEncode({'data': chatJson}), 200),
      );

      //Act
      final result = await dataSource.getChatRooms();

      //Assert
      expect(result, chatEntityList);
    });

    test('should throw an exception when the response code is not 200',
        () async {
      // Arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => TokenModel(token: token));
      when(mockHttpClient.get(any, headers: anyNamed('headers'))).thenAnswer(
        (_) async => http.Response('Something went wrong', 404),
      );

      // Act & Assert
      expect(() => dataSource.getChatRooms(), throwsException);
    });
  });

  

  group('getMessages', () {
    const token = 'test_token';
    const chatId = 'chat123';
    final messageJson = [
     {
        '_id': '123',
        'sender': {
          'name': 'John Doe',
          'email': 'john@example.com',
          '_id': 'user1_id',
          '__v': 1,
          'password': 'password123',

        
        },
        'chat': {
          '_id': 'chat123',
          'user1': {
            'name': 'John Doe',
            'email': 'john@example.com',
            '_id': 'user1_id',
            '__v': 1,
            'password': 'password123',
          
          },
          'user2': {
            'name': 'Jane Doe',
            'email': 'jane@example.com',
            '_id': 'user2_id',
            '__v': 2,
            'password': 'password456',
          },
          
        },
        'content': 'Hello, world!',
      }

    ];
    final messageEntityList = [MessageEntity.fromJson(messageJson[0])];

    test('should return a list of MessageEntity when the response code is 200',
        () async {
      // Arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => TokenModel(token: token));
      when(mockHttpClient.get(any, headers: anyNamed('headers'))).thenAnswer(
        (_) async => http.Response(jsonEncode({'data': messageJson}), 200),
      );

      // Act
      final result = await dataSource.getMessages(chatId);

      // Assert
      expect(result, messageEntityList);
    });

    test('should throw an exception when the response code is not 200',
        () async {
      // Arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => TokenModel(token: token));
      when(mockHttpClient.get(any, headers: anyNamed('headers'))).thenAnswer(
        (_) async => http.Response('Something went wrong', 404),
      );

      // Act & Assert
      expect(() => dataSource.getMessages(chatId), throwsException);
    });
  });

  group('createChatRoom', () {
    const token = 'test_token';
    const userId = 'user123';

    test('should create a chat room when the response code is 201', () async {
      // Arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => TokenModel(token: token));
      when(mockHttpClient.post(any,
              headers: anyNamed('headers'), body: anyNamed('body')))
          .thenAnswer((_) async => http.Response('', 201));

      // Act
      await dataSource.createChatRoom(userId);

      // Assert
      verify(mockHttpClient.post(any,
              headers: anyNamed('headers'), body: anyNamed('body')))
          .called(1);
    });

    test('should throw an exception when the response code is not 201',
        () async {
      // Arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => TokenModel(token: token));
      when(mockHttpClient.post(any,
              headers: anyNamed('headers'), body: anyNamed('body')))
          .thenAnswer((_) async => http.Response('Something went wrong', 400));

      // Act & Assert
      expect(() => dataSource.createChatRoom(userId), throwsException);
    });
  });
}
