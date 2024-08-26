import 'dart:math';

import 'package:ecommerce_app/features/auth/data/model/user_model.dart';
import 'package:ecommerce_app/features/chat/data/models/chat_model.dart';
import 'package:ecommerce_app/features/chat/domain/entity/chat.dart';
import 'package:flutter_test/flutter_test.dart';


void main() {
  // Test data
  const testUser1 = UserModel(
    name: 'John Doe',
    email: 'john@example.com',
    password: '',
    id: 'user1_id',
    v: 0,
  );

  const testUser2 = UserModel(
    name: 'Jane Doe',
    email: 'jane@example.com',
    password: '',
    id: 'user2_id',
    v: 0,
  );

  const testChatModel = ChatModel(
    chatId: 'chat123',
    user1: testUser1,
    user2: testUser2,
  );

  final testJson = {
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
  };

  test('should be a subclass of ChatEntity', () {
    expect(testChatModel, isA<ChatEntity>());
  });

  group('fromJson', () {
    test('should return a valid ChatModel from JSON', () {
      final result = ChatModel.fromJson(testJson);
      
      expect(result, testChatModel);
    });
  });

 
  
}
