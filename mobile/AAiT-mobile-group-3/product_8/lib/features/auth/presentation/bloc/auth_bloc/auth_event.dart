part of 'auth_bloc.dart';

abstract class AuthEvent extends Equatable {
  const AuthEvent();

  @override
  List<Object> get props => [];
}

class SigninEvent extends AuthEvent {
  final SignInUserEntitiy signInUserEntitiy;

  SigninEvent({required this.signInUserEntitiy});
}

class SignupEvent extends AuthEvent {
  final SignUpUserEntitiy signUpUserEntitiy;

  SignupEvent({required this.signUpUserEntitiy});
}

class LogoutEvent extends AuthEvent {}

class GetUserEvent extends AuthEvent {}