// import 'dart:convert';

// import 'package:flutter_test/flutter_test.dart';
// import 'package:product_8/features/auth/data/models/sign_in_user_model.dart';

// import '../../../../helpers/jason_reader.dart';

// void main() {
//   const testSignInModel = SignInUserModel(
//     email: 'ex@gmail.com',
//     password:'123456',
//   );

//   test('should be a subclass of SignInUserModel', () async {
//     //assert
//     expect(testSignInModel, isA<SignInUserModel>());
//   });

//   test('should return a valid model from json', () async {
//     // arrange
//     final Map<String, dynamic> jsonMap = json
//         .decode(readJson('helpers/dummy_data/dummy_sign_in_response.json'));

//     // act
//     final result = SignInUserModel.fromJson(jsonMap);

//     // assert
//     expect(result, testSignInModel);
//   });

//   test('should return a valid json map when toJson is called', () async {

//     final Map<String, dynamic> jsonMap = json
//         .decode(readJson('helpers/dummy_data/dummy_sign_in_response.json'));
//     // Act
//     final result = testSignInModel.toJson();

//     // Assert
//     expect(result, jsonMap);
//   });
// }
