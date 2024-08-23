import 'package:equatable/equatable.dart';

import '../../data/models/user_model.dart';

class UserEntity extends Equatable {
  final String id;
  final String name;
  final String email;
  final String accessToken;

  const UserEntity(
      {required this.id,
      required this.name,
      required this.email,
      required this.accessToken});
  factory UserEntity.fromModel(UserModel model) {
    return UserEntity(
        id: model.id,
        name: model.name,
        email: model.email,
        accessToken: model.accessToken);
  }
  @override
  List<Object?> get props => [id, name, email, accessToken];
}
