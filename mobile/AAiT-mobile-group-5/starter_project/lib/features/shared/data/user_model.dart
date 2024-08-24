import 'package:starter_project/features/shared/entities/user.dart';

class UserModel extends User {
  const UserModel({
    required super.name,
    required super.email,
    required super.password,
  });

  // Factory method to create a UserModel from JSON
  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
      name: json['name'],
      email: json['email'],
      password: json['password'],
    );
  }

  // Static method to create a list of UserModels from a list of JSON objects
  static List<UserModel> fromJsonList(List<dynamic> jsonList) {
    return jsonList
        .map((json) => UserModel.fromJson(json as Map<String, dynamic>))
        .toList();
  }

  // Method to convert a UserModel to JSON
  Map<String, dynamic> toJson() {
    return {
      'name': name,
      'email': email,
      'password': password,
    };
  }

  // Static method to convert a list of UserModels to a list of JSON objects
  static List<Map<String, dynamic>> toJsonList(List<UserModel> users) {
    return users.map((user) => user.toJson()).toList();
  }

  // Method to convert UserModel to User entity
  User toEntity() {
    return User(
      name: name,
      email: email,
      password: password,
    );
  }

  // Method to convert a list of UserModels to a list of User entities
  List<User> toEntityList(List<UserModel> userModels) {
    return userModels.map((userModel) => userModel.toEntity()).toList();
  }

  // Method to convert a User entity to UserModel
  static UserModel toModel(User user) {
    return UserModel(
      name: user.name,
      email: user.email,
      password: user.password,
    );
  }

  // Method to convert a list of User entities to a list of UserModels
  List<UserModel> toModelList(List<User> users) {
    return users.map((user) => toModel(user)).toList();
  }
}
