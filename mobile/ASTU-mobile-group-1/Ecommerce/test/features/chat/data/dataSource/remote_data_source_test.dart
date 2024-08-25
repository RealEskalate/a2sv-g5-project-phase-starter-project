// import 'dart:convert';

// import 'package:mockito/annotations.dart';
// import 'package:mockito/mockito.dart';
// import 'package:http/http.dart' as http;

// import 'package:flutter_test/flutter_test.dart';
// import 'package:http/testing.dart';
// import 'package:mockito/mockito.dart';
// import 'package:http/http.dart' as http;
// import 'package:product_6/core/constants/constants.dart';
// import 'package:product_6/core/error/exception.dart';
// import 'package:product_6/features/auth/domain/entities/user_entity.dart';
// import 'package:product_6/features/chat/data/data_source/remote_data_source/remote_data_source_impl.dart';
// import 'package:product_6/features/chat/data/model/chat_model.dart';
// import 'package:product_6/features/chat/data/model/message_model.dart';

// import './remote_data_source_test.mocks.dart';
// import '../../../../helpers/test_helper.mocks.dart';

// // Generate a MockClient using the Mockito package.
// // Create a new file called 'remote_data_source_impl_test.mocks.dart' by running:
// // flutter pub run build_runner build
// @GenerateMocks([http.Client])
// void main() {
//   late RemoteDataSourceImpl dataSource;
//   late MockClient mockHttpClient;

//   const String accessToken = 'test_token';
//   const String chatId = '1';
//   const String recieverId = '2';
//   const UserEntity tempuser1 = UserEntity(id: '1', name: '1', email: '1');
//   const UserEntity tempuser2 = UserEntity(id: '2', name: '3', email: '2');
//   const  ChatModel tempchatModel = ChatModel(chatId: '1', user1: tempuser1, user2: tempuser2);
//   MessageModel tempmessageModel = MessageModel(messageId: '1',chatEntity: tempchatModel,content:'hello',type:'twxt',sender: tempuser1);


//   setUp(() {
//     mockHttpClient = MockClient();
//     dataSource = RemoteDataSourceImpl(accessToken, client: mockHttpClient);
//   });

//   group('getAllChats', () {
//     // final tChatModels = [
//     //   ChatModel(chatId: '1', tempuser1: 'tempuser1', user2: 'user2'),
//     // ];

//     test('should return a list of ChatModels when the response code is 200', () async {
//       // Arrange
//       when(mockHttpClient.get(Uri.parse(Urls.baseChat), headers: anyNamed('headers')))
//           .thenAnswer((_) async => http.Response(json.encode({'data': tempchatModel}), 200));

//       // Act
//       final result = await dataSource.getAllChats();

//       // Assert
//       expect(result, isA<List<ChatModel>>());
//     });

//     test('should throw ServerException when the response code is not 200', () async {
//       // Arrange
//       when(mockHttpClient.get(Uri.parse(Urls.baseChat), headers: anyNamed('headers')))
//           .thenAnswer((_) async => http.Response('Something went wrong', 404));

//       // Act
//       final call = dataSource.getAllChats;

//       // Assert
//       expect(() => call(), throwsA(isA<ServerException>()));
//     });
//   });

//   group('getChatById', () {
//     // final tChatModel = ChatModel(chatId: '1', tempuser1: 'tempuser1', user2: 'user2');

//     test('should return ChatModel when the response code is 200', () async {
//       // Arrange
//       when(mockHttpClient.post(Uri.parse(Urls.getChatById(chatId)), headers: anyNamed('headers')))
//           .thenAnswer((_) async => http.Response(json.encode({'data': tempchatModel}), 200));

//       // Act
//       final result = await dataSource.getChatById(chatId);

//       // Assert
//       expect(result, isA<ChatModel>());
//     });

//     test('should throw Exception when the response code is not 200', () async {
//       // Arrange
//       when(mockHttpClient.post(Uri.parse(Urls.getChatById(chatId)), headers: anyNamed('headers')))
//           .thenAnswer((_) async => http.Response('Something went wrong', 404));

//       // Act
//       final call = dataSource.getChatById;

//       // Assert
//       expect(() => call(chatId), throwsA(isA<Exception>()));
//     });
//   });

//   group('deleteChat', () {
//     test('should return true when the response code is 204', () async {
//       // Arrange
//       when(mockHttpClient.delete(Uri.parse(Urls.getChatById(chatId)), headers: anyNamed('headers')))
//           .thenAnswer((_) async => http.Response('', 204));

//       // Act
//       final result = await dataSource.deleteChat(chatId);

//       // Assert
//       expect(result, true);
//     });

//     test('should return false when the response code is not 204', () async {
//       // Arrange
//       when(mockHttpClient.delete(Uri.parse(Urls.getChatById(chatId)), headers: anyNamed('headers')))
//           .thenAnswer((_) async => http.Response('Something went wrong', 404));

//       // Act
//       final result = await dataSource.deleteChat(chatId);

//       // Assert
//       expect(result, false);
//     });

//     test('should throw Exception on error', () async {
//       // Arrange
//       when(mockHttpClient.delete(Uri.parse(Urls.getChatById(chatId)), headers: anyNamed('headers')))
//           .thenThrow(Exception('Failed to delete chat'));

//       // Act
//       final call = dataSource.deleteChat;

//       // Assert
//       expect(() => call(chatId), throwsA(isA<Exception>()));
//     });
//   });

//   group('initiateChat', () {
//     // final tChatModel = ChatModel(chatId: '1', tempuser1: 'tempuser1', user2: 'user2');

//     test('should return ChatModel when the response code is 201', () async {
//       // Arrange
//       when(mockHttpClient.post(Uri.parse(Urls.baseChat), headers: anyNamed('headers'), body: anyNamed('body')))
//           .thenAnswer((_) async => http.Response(json.encode({'data': tempchatModel}), 201));

//       // Act
//       final result = await dataSource.initiateChat(recieverId);

//       // Assert
//       expect(result, isA<ChatModel>());
//     });

//     test('should throw Exception when the response code is not 201', () async {
//       // Arrange
//       when(mockHttpClient.post(Uri.parse(Urls.getChatById(chatId)), headers: anyNamed('headers'), body: anyNamed('body')))
//           .thenAnswer((_) async => http.Response('Something went wrong', 400));

//       // Act
//       final call = dataSource.initiateChat;

//       // Assert
//       expect(() => call(recieverId), throwsA(isA<Exception>()));
//     });
//   });

//   group('getChatMessages', () {
//     final tMessageModels = [
//       // MessageModel(messageId: '1', chatId: '1', senderId: 'tempuser1', text: 'Hello'),
//       tempmessageModel
//     ];

//     test('should return a list of MessageModels when the response code is 200', () async {
//       // Arrange
//       when(mockHttpClient.delete(Uri.parse(Urls.getChatById(chatId)), headers: anyNamed('headers')))
//           .thenAnswer((_) async => http.Response(json.encode({'data': tMessageModels}), 200));

//       // Act
//       final result = await dataSource.getChatMessages(chatId);

//       // Assert
//       expect(result, isA<List<MessageModel>>());
//     });

//     test('should throw Exception when the response code is not 200', () async {
//       // Arrange
//       when(mockHttpClient.delete(Uri.parse(Urls.getChatById(chatId)), headers: anyNamed('headers')))
//           .thenAnswer((_) async => http.Response('Something went wrong', 404));

//       // Act
//       final call = dataSource.getChatMessages;

//       // Assert
//       expect(() => call(chatId), throwsA(isA<Exception>()));
//     });
//   });
// }
