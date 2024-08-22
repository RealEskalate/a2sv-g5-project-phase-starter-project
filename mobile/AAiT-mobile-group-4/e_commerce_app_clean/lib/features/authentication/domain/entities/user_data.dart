import 'package:equatable/equatable.dart';

class UserEntity extends Equatable {
  final String email;
  final String name;

  const UserEntity({
    required this.email,
    required this.name,
  });

  @override
  List<Object?> get props => [email, name];
}
