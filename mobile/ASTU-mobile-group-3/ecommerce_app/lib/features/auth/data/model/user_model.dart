import '../../domain/entities/user_entity.dart';

class UserModel extends UserEntity {
  @override
  // ignore: overridden_fields
  final String name;
  @override
  // ignore: overridden_fields
  final String email;
  @override
  // ignore: overridden_fields
  final String password;

  const UserModel(
      {required this.name, required this.email, required this.password})
      : super(email: email, name: name, password: password);

  // Factory method to create UserModel from JSON
  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
      name: json['name'],
      email: json['email'],
      password: json['password'],
    );
  }

  // Method to convert UserModel to JSON
  Map<String, dynamic> toJson() {
    return {
      'name': name,
      'email': email,
      'password': password,
    };
  }

  factory UserModel.getModel(String name, String email, String password) {
    return UserModel(name: name, email: email, password: password);
  }

  UserEntity toEntity() {
    return UserEntity(name: name, email: email, password: password);
  }
}
