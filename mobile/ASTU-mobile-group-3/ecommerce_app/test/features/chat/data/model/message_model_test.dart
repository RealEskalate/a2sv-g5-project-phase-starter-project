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
      id: 'user1_id',
      name: 'John Doe',
      email: 'john@example.com',
      password: '',
      v: 1
 
    ),
    chat: ChatModel(
      chatId: 'chat123',
      user1:  UserModel(
      id: 'user1_id',
      name: 'John Doe',
      email: 'john@example.com',
      password: '',
      v: 0
   
    ),
    user2:  UserModel(
      id: 'user2_id',
      name: 'Jane Doe',
      email: 'jane@example.com',
      password: '',
      v: 0
    
    ),
   
    ),
    content: 'Hello, world!',
  );

   final testJson = {
        'messageId': '123',
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
            'id': 'user1_id',
            '__v': 1,
            'password': 'password123',
          
          },
          'user2': {
            'name': 'Jane Doe',
            'email': 'jane@example.com',
            'id': 'user2_id',
            '__v': 2,
            'password': 'password456',
          },
          
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

 
}
