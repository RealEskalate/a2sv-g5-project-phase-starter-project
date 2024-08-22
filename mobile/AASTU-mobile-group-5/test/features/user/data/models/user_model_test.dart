import 'package:flutter_test/flutter_test.dart';
import 'package:task_9/features/user/data/models/user_model.dart';
import 'package:task_9/features/user/domain/entities/user.dart';


void main() {
  const tUserModel = UserModel(
    id: '1',
    name: 'Test User',
    email: 'test@example.com',
    password: 'password123',
  );

  group('UserModel', () {
    test('should be a subclass of User entity', () async {
      // Assert
      expect(tUserModel, isA<User>());
    });

    test('fromJson should return a valid model', () async {
      // Arrange
      final Map<String, dynamic> jsonMap = {
        'id': '1',
        'name': 'Test User',
        'email': 'test@example.com',
        'password': 'password123',
      };

      // Act
      final result = UserModel.fromJson(jsonMap);

      // Assert
      expect(result, tUserModel);
    });

    test('toJson should return a JSON map containing the proper data', () async {
      // Act
      final result = tUserModel.toJson();

      // Assert
      final expectedMap = {
        'id': '1',
        'name': 'Test User',
        'email': 'test@example.com',
        'password': 'password123',
      };
      expect(result, expectedMap);
    });
  });
}
