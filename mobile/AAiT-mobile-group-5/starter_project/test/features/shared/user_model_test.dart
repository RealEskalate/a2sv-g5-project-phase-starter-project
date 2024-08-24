import 'package:flutter_test/flutter_test.dart';
import 'package:starter_project/features/shared/data/user_model.dart';
import 'package:starter_project/features/shared/entities/user.dart';

void main() {
  const tUserModel = UserModel(
    name: 'Test User',
    email: 'test@example.com',
    password: 'password123',
  );

  const tUserEntity = User(
    name: 'Test User',
    email: 'test@example.com',
    password: 'password123',
  );

  final tUserModelList = [
    tUserModel,
    const UserModel(
      name: 'Another User',
      email: 'another@example.com',
      password: 'password456',
    ),
  ];

  final tUserEntityList = [
    tUserEntity,
    const User(
      name: 'Another User',
      email: 'another@example.com',
      password: 'password456',
    ),
  ];

  final tUserJson = {
    'name': 'Test User',
    'email': 'test@example.com',
    'password': 'password123',
  };

  final tUserJsonList = [
    tUserJson,
    {
      'name': 'Another User',
      'email': 'another@example.com',
      'password': 'password456',
    },
  ];

  group('UserModel', () {
    test('should convert UserModel to User entity', () {
      final result = tUserModel.toEntity();
      expect(result, tUserEntity);
    });

    test('should convert User entity to UserModel', () {
      final result = UserModel.toModel(tUserEntity);
      expect(result, tUserModel);
    });

    test('should convert JSON to UserModel', () {
      final result = UserModel.fromJson(tUserJson);
      expect(result, tUserModel);
    });

    test('should convert UserModel to JSON', () {
      final result = tUserModel.toJson();
      expect(result, tUserJson);
    });

    test('should convert list of JSON to list of UserModels', () {
      final result = UserModel.fromJsonList(tUserJsonList);
      expect(result, tUserModelList);
    });

    test('should convert list of UserModels to list of JSON', () {
      final result = UserModel.toJsonList(tUserModelList);
      expect(result, tUserJsonList);
    });

    test('should convert list of UserModels to list of User entities', () {
      final result = tUserModel.toEntityList(tUserModelList);
      expect(result, tUserEntityList);
    });

    test('should convert list of User entities to list of UserModels', () {
      final result = tUserModel.toModelList(tUserEntityList);
      expect(result, tUserModelList);
    });
  });
}
