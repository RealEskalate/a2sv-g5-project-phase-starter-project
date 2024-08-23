// import 'package:flutter_test/flutter_test.dart';
// import 'package:product_8/features/auth/data/models/sign_up_user_model.dart';
// import 'package:product_8/features/auth/domain/entities/sign_up_user_entitiy.dart';

// void main() {
//   const signUpUserModel = SignUpUserModel(
//     email: 'test@example.com',
//     password: 'password123',
//     name: 'John Doe',
//   );

//   const jsonMap = {
//     'email': 'test@example.com',
//     'password': 'password123',
//     'name': 'John Doe',
//   };

//   test('should be a subclass of SignUpUserEntitiy', () async {
//     // Assert
//     expect(signUpUserModel, isA<SignUpUserEntitiy>());
//   });

//   test('should return a valid model when fromJson is called', () async {
//     // Act
//     final result = SignUpUserModel.fromJson(jsonMap);

//     // Assert
//     expect(result, signUpUserModel);
//   });

//   test('should return a valid json map when toJson is called', () async {
//     // Act
//     final result = signUpUserModel.toJson();

//     // Assert
//     expect(result, jsonMap);
//   });
// }
