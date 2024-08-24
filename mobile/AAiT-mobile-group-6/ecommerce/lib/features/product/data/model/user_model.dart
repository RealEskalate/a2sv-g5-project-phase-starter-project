import '../../domain/entitity/user.dart';

class UserModel extends User {
  const UserModel({
    required super.id,
    required super.name,
    required super.email,
    required super.password,
  });

  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
        id: json['_id'],
        name: json['name'],
        email: json['email'],
        password: json['password']);
  }

  Map<String, dynamic> toJson() {
    return {
      'id': super.id,
      'name': super.name,
      'email': super.email,
      'password': super.password,
    };
  }

  factory UserModel.fromUser(User user){
    return UserModel(id: user.id, name: user.name, email: user.email, password: user.password);
  }
}