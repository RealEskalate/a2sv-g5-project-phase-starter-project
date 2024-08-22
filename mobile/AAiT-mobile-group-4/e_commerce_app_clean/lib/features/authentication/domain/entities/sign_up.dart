import 'package:equatable/equatable.dart';

class SignUpEntity extends Equatable{
  final String email;
  final String password;
  final String username;

  const SignUpEntity({
    required this.email,
    required this.password,
    required this.username,
  });
  @override
  List<Object?> get props =>[
    email,
    password,
    username,
  ] ;
}