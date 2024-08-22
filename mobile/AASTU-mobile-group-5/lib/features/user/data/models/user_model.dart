import '../../domain/entities/user.dart';

class UserModel extends User {
  @override
  // ignore: overridden_fields
  final String? id;
  @override
  // ignore: overridden_fields
  final String? name;
  @override
  // ignore: overridden_fields
  final String email;
  @override
  // ignore: overridden_fields
  final String? password;

  const UserModel({
    this.id,
    this.name,
    required this.email,
    this.password,
  }) : super(
          email: '',
          password: '',
          name: '',
          id: '',
        );

  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
      id: json['id'],
      name: json['name'],
      email: json['email'],
      password: json['password'],
    );
  }
  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'email': email,
      'password': password,
    };
  }

  // factory UserModel.fromEntity(User user) {
  //   return UserModel(
  //     id: user.id,
  //     name: user.name,
  //     email: user.email,
  //     password: user.password,
  //   );
  // }

  // User toEntity() {
  //   return User(
  //     id: id,
  //     name: name,
  //     email: email,
  //     password: password,
  //   );
  // }

  // List<Object?> get props => [id, name, email, password];
}