part of 'auth_bloc.dart';

abstract class AuthState extends Equatable {
  const AuthState();

  @override
  List<Object> get props => [];
}

class AuthInitial extends AuthState {}

class AuthLoading extends AuthState {}

class AuthAuthenticated extends AuthState {
  final UserDataEntity userDataEntity;

  AuthAuthenticated({required this.userDataEntity});

  @override
  List<Object> get props => [userDataEntity];
}

class AuthError extends AuthState {
  final String message;

  AuthError({required this.message});

  @override
  List<Object> get props => [message];
}

class AuthLoggedOut extends AuthState {}

class AuthSuccess extends AuthState {}

class AuthRegisterSuccess extends AuthState {}