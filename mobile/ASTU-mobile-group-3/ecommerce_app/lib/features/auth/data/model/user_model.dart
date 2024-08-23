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

  factory UserModel.getModel(String name, String email, String password) {
    return UserModel(name: name, email: email, password: password);
  }

  UserEntity toEntity() {
    return UserEntity(name: name, email: email, password: password);
  }
}
