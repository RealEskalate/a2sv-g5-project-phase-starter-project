import 'package:equatable/equatable.dart';

class SignInEntity extends Equatable {
  final String email;
  final String password;

  const SignInEntity({required this.email, required this.password});

  @override
  List<Object?> get props => [email, password];
}
