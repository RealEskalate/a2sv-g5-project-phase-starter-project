// import 'package:flutter_test/flutter_test.dart';
// import 'package:product_8/features/auth/data/models/user_data_model.dart';
// import 'package:product_8/features/auth/domain/entities/user_data_entity.dart';

// void main() {
//   const dataModel = DataModel(
//     name: 'John Doe',
//     email: 'test@example.com',
//   );

//   const userDataModel = UserDataModel(
//     data: dataModel,
//     token: 'someToken123',
//   );

//   const jsonMap = {
//     'data': {
//       'name': 'John Doe',
//       'email': 'test@example.com',
//     },
//     'token': 'someToken123',
//   };

//   test('should be a subclass of UserDataEntity', () async {
//     // Assert
//     expect(userDataModel, isA<UserDataEntity>());
//   });

//   test('should return a valid model when fromJson is called', () async {
//     // Act
//     final result = UserDataModel.fromJson(jsonMap);

//     // Assert
//     expect(result, userDataModel);
//   });

//   test('should return a valid json map when toJson is called', () async {
//     // Act
//     final result = userDataModel.toJson();

//     // Assert
//     expect(result, jsonMap);
//   });

//   test('should correctly parse data model when fromJson is called', () async {
//     // Act
//     final result = DataModel.fromJson(jsonMap['data'] as Map<String, dynamic>);

//     // Assert
//     expect(result, dataModel);
//   });
// }
