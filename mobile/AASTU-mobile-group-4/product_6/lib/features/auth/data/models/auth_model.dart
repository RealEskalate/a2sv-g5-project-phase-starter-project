import '../../domain/entity/auth_entity.dart';

class AuthResponseModel extends AuthResponseEntity {
  AuthResponseModel({
    required String accessToken,
  }) : super(accessToken: accessToken);

  factory AuthResponseModel.fromJson(Map<String, dynamic> json) {
    return AuthResponseModel(
      accessToken: json['access_token'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'access_token': accessToken,
    };
  }
}

class AuthModel extends AuthEntity {
  AuthModel({
    required String email,
    required String password,
  }) : super(email: email, password: password);

  factory AuthModel.fromJson(Map<String, dynamic> json) {
    return AuthModel(
      email: json['email'],
      password: json['password'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'email': email,
      'password': password,
    };
  }
}

class UserModel extends UserEntity {
  UserModel({
    required String id,
    required String name,
    required String email,
  }) : super(id: id, name: name, email: email);

  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
      id: json['_id'] ?? json['id'],
      name: json['name'],
      email: json['email'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      '_id': id,
      'name': name,
      'email': email,
    };
  }
}

class SentUserModel {
  final String name;
  final String email;
  final String password;

  SentUserModel({
    required this.name,
    required this.email,
    required this.password,
  });

  // Factory constructor to create a SentUserModel from a map (e.g., from JSON)
  factory SentUserModel.fromJson(Map<String, dynamic> json) {
    return SentUserModel(
      name: json['name'],
      email: json['email'],
      password: json['password'],
    );
  }

  // Method to convert a SentUserModel instance into a map (e.g., for JSON serialization)
  Map<String, dynamic> toJson() {
    return {
      'name': name,
      'email': email,
      'password': password,
    };
  }
}
