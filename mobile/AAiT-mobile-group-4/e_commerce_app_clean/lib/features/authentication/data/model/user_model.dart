import '../../domain/entities/user_data.dart';

class UserModel extends UserEntity {
  const UserModel({
    required super.email,
    required super.name,
  });

  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
      email: json['email'],
      name: json['name'],
    );
  }
  UserEntity toUserEntity() {
    return UserEntity(
      email: email,
      name: name,
    );
  }
}
