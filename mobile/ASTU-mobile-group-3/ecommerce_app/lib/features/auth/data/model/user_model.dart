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
  @override
  // ignore: overridden_fields
  final String id;

  @override
  // ignore: overridden_fields
  final int v;
  const UserModel({
    required this.name,
    required this.email,
    required this.password,
    required this.id,
    required this.v,
  }) : super(name: name, email: email, password: password, id: id, v: v);

  factory UserModel.fromStrings(String name, String email, String password) {
    return UserModel(
        name: name, email: email, password: password, id: '', v: 0);
  }

  factory UserModel.fromJson(Map<String, dynamic> map) {
    return UserModel(
        name: map['name'],
        email: map['email'],
        id: map['id'],
        password: '',
        v: 0);
  }
  factory UserModel.fromSellerJson(Map<String, dynamic> map) {
    return UserModel(
        name: map['name'],
        email: map['email'],
        id: map['_id'],
        v: map['__v'].toInt(),
        password: '');
  }
  UserEntity toEntity() {
    return UserEntity(
        name: name, email: email, password: password, id: id, v: v);
  }

  factory UserModel.fromEntity(UserEntity user) {
    return UserModel(
        name: user.name,
        email: user.email,
        id: user.id,
        password: user.password,
        v: user.v);
  }

  /// The purpose of these method is that because we cannot use user model in ProductModel because it wiill extends from UserEntity
  /// That is it will inherit as entity even though we pass model
  static UserEntity toEntityParam(UserEntity user) {
    return UserEntity(
        name: user.name,
        email: user.email,
        id: user.id,
        password: user.password,
        v: user.v);
  }
}
