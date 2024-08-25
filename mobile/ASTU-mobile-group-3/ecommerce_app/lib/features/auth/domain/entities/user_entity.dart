import 'package:equatable/equatable.dart';

class UserEntity extends Equatable {
  final String name;
  final String email;
  final String password;
  final String id;
  final int v;

  const UserEntity({
    required this.name,
    required this.email,
    required this.password,
    required this.id,
    required this.v,
  });

  @override
  List<Object?> get props => [name, password, email, id, v];
}
