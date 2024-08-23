import 'package:equatable/equatable.dart';

class AuthEvent extends Equatable {
  @override
  List<Object?> get props => [];
}

class SignInEvent extends AuthEvent {
  final String email;
  final String password;

  SignInEvent({required this.email, required this.password});
}

class SignUpEvent extends AuthEvent {
  final String name;
  final String email;
  final String password;
  final String repeatedPassword;

  SignUpEvent(
      {required this.name,
      required this.email,
      required this.password,
      required this.repeatedPassword});
}

class LogOutEvent extends AuthEvent {}

class CheckSignedInEvent extends AuthEvent {}

class GetUserEvent extends AuthEvent {}
