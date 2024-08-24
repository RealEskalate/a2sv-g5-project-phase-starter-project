import 'package:equatable/equatable.dart';

class SignUpEntity extends Equatable {
  final String id;
  final String name;
  final String email;
  final String password;

  const SignUpEntity({
    required this.id,
    required this.name,
    required this.email,
    required this.password,
  });

  @override
  List<Object> get props => [id, name, email, password];
}
