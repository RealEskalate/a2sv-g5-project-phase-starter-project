part of 'auth_bloc.dart';

sealed class AuthEvent extends Equatable {
  const AuthEvent();

  @override
  List<Object> get props => [];
}

class GetCurrentUserEvent extends AuthEvent {}

class SignUpEvent extends AuthEvent {
  final SignUpEntity signUpEntity;
  const SignUpEvent({required this.signUpEntity});

  @override
  List<Object> get props => [signUpEntity];
}

class LogInEvent extends AuthEvent {
  
  final LogInEntity logInEntity;

  const LogInEvent({required this.logInEntity});
  @override
  List<Object> get props => [logInEntity];
}

class LogOutEvent extends AuthEvent {}
