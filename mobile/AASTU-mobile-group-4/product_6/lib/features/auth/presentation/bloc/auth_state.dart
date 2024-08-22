import 'package:equatable/equatable.dart';

import '../../domain/entity/auth_entity.dart';

abstract class AuthState extends Equatable {
  const AuthState();

  @override
  List<Object> get props => [];
}

class AuthInitial extends AuthState {}

class AuthLoading extends AuthState {}

class AuthLoaded extends AuthState {
  final AuthResponseEntity authResponse;

  const AuthLoaded({required this.authResponse});

  @override
  List<Object> get props => [authResponse];
}

class UserProfileLoaded extends AuthState {
  final UserEntity user;

  const UserProfileLoaded({required this.user});

  @override
  List<Object> get props => [user];
}

class UserSuccessRegister extends AuthState {}

class AuthError extends AuthState {
  final String message;

  const AuthError({required this.message});

  @override
  List<Object> get props => [message];
}

class LogoutSuccess extends AuthState {}