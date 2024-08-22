import 'package:equatable/equatable.dart';

class SignInUserEntitiy extends Equatable {
  final String email;
  final String password;

  const SignInUserEntitiy({
    required this.email,
    required this.password,
  });

  @override
  List<Object?> get props => [email, password];
}