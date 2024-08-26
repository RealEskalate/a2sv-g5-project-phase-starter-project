import 'dart:convert';

import 'package:ecommerce_app/features/auth/data/model/user_model.dart';
import 'package:ecommerce_app/features/chat/data/models/chat_model.dart';
import 'package:ecommerce_app/features/chat/data/models/message_model.dart';
import 'package:ecommerce_app/features/chat/domain/entity/message.dart';
import 'package:flutter_test/flutter_test.dart';




void main() {
  const  testMessageModel = MessageModel(
    messageId: '123',
    sender: UserModel(
      id: '1',
      name: 'John Doe',
      email: 'john@gmail.com',
      password: '123',
      v: 3
 
    ),
    chat: ChatModel(
      chatId: '456',
      user1:  UserModel(
      id: '1',
      name: 'John Doe',
      email: 'john@gmail.com',
      password: '123',
      v: 3
   
    ),
    user2:  UserModel(
      id: '2',
      name: 'Doe John',
      email: 'Doe@gmail.com',
      password: '123',
      v: 3
    
    ),
   
    ),
    content: 'Hello, world!',
  );

   final testJson = {
        'messageId': '123',
        'sender': {
          'id': '1',
          'name': 'John Doe',
        
        },
        'chat': {
          'chatId': '456',
          
        },
        'content': 'Hello, world!',
      };



  test('should be subclass of MessageEntity', () {
    expect(testMessageModel, isA<MessageEntity>());
  });

  group('fromJson', () {
    test('should return a valid model', () {
      final result = MessageModel.fromJson(testJson);
      expect(result, testMessageModel);
    });
  });

  group('toJson', () {
    test('should return JSON map containing proper data', () {
      final expectedJsonMap = testJson;
     // expect(testMessageModel.toJson(), expectedJsonMap);
    });
  });
}
