part of 'auth_bloc.dart';

sealed class AuthState extends Equatable {
  const AuthState();

  @override
  List<Object> get props => [];
}

final class AuthInitial extends AuthState {}

final class AuthLoadingState extends AuthState {}

final class AuthErrorState extends AuthState {
  final String message;

  const AuthErrorState({required this.message});
}

final class AuthSignedUpState extends AuthState {}

final class AuthUserLoaded extends AuthState {
  final UserEntity userEntity;

  const AuthUserLoaded({required this.userEntity});
  @override
  List<Object> get props => [userEntity];
}

final class AuthSignedInState extends AuthState {}

final class AuthLogOutState extends AuthState {}
