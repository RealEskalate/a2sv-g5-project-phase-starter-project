import 'package:equatable/equatable.dart';

class AuthState extends Equatable {
  const AuthState();
  @override
  List<Object> get props => [];
}

class InitalState extends AuthState {}

class LoadingState extends AuthState {}

class SignedInState extends AuthState {}

class SignInState extends AuthState {}

class SignedUpState extends AuthState {}

class SignUpState extends AuthState {}

class UserIsReady extends AuthState {}

class ErrorState extends AuthState {
  final String message;

  const ErrorState({required this.message});

  @override
  List<Object> get props => [message];
}
