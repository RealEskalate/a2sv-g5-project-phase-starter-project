import 'package:e_commerce_app/features/auth/domain/entities/user.dart';

class UserModel extends User {
  final String id;
  final String email;
  final String name;

  UserModel({
    required this.id,
    required this.email,
    required this.name,
  }) : super(id: id, email: email, name: name);

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'email': email,
      'name': name,
    };
  }
  Map<String, dynamic> toJson2() {
    return {
      '_id': id,
      'email': email,
      'name': name,
    };
  }

  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
      id: json['id'],
      email: json['email'],
      name: json['name'],
    );
  }
  factory UserModel.fromJson2(Map<String, dynamic> json) {
    return UserModel(
      id: json['_id'],
      email: json['email'],
      name: json['name'],
    );
  }

  User toUser() => User(id: id, email: email, name: name);

  @override
  List<Object?> get props => [id, email, name];
}
