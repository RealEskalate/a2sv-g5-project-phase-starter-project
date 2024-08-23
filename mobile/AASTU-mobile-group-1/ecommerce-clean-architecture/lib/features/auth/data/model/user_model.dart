import '../../domain/entities/user.dart';

class UserModel extends UserEntity {
  UserModel({
    id,
    required String name,
    required String email,
    String? password,
  }) : super(
          id: id,
          name: name,
          email: email,
          password: password ?? '',
        );

  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
      id: json['id'],
      name: json['name'],
      email: json['email'],
    );
  }

  factory UserModel.forSeller(Map<String, dynamic> json) {
    return UserModel(
      id: json['_id'],
      name: json['name'],
      email: json['email'],
    );
  }

   Map<String, dynamic> toJsonForSeller() => {
        '_id': id,
        'name': name,
        'email': email,
      };
  Map<String, dynamic> toJson() => {
        'id': id,
        'name': name,
        'email': email,
        'password': password,
      };

  factory UserModel.fromEntity(UserEntity entity) {
    return UserModel(
      id: entity.id,
      name: entity.name,
      email: entity.email,
      password: entity.password,
    );
  }
}
