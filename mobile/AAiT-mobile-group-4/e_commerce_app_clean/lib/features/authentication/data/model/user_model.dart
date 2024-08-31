import '../../domain/entities/user_data.dart';

class UserModel extends UserEntity {
  const UserModel({
    required super.id,
    required super.email,
    required super.name,
  });

  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
      id: json['id'] ?? json['_id'],
      email: json['email'],
      name: json['name'],
    );
  }
  UserEntity toUserEntity() {
    return UserEntity(
      id: id,
      email: email,
      name: name,
    );
  }
}
