import '../../domain/entity/user.dart';

class UserModel extends User {
  final String id;
  final String email;
  final String username;
  final String password;

  UserModel({
    required this.id,
    required this.email,
    required this.username,
    required this.password,
  }) : super(id: id, email: email, username: username, password: password);
  

  @override
  List<Object?> get props => [id, email, username, password];

  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
      id: json['id'],
      email: json['email'],
      username: json['name'],
      password: json['password'] ?? '',
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'email': email,
      'username': username,
      'password': password ,
    };
  }


  User toEntity() => User(id: id, username: username, password: password, email: email);
}