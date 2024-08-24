part of 'auth_bloc.dart';

@immutable
sealed class AuthState {
  const AuthState();
}

final class AuthInitial extends AuthState {}

final class AuthLoadingState extends AuthState{}

final class AuthErrorState extends AuthState{
  final String message;
  const AuthErrorState({required this.message});
}

final class AuthSuccessState extends AuthState{
  final String message;
  const AuthSuccessState({required this.message});
}


enum AuthStatus { loaded, loading, error }
class UserLogoutState extends AuthState {
  final String message;
  final AuthStatus status;

  const UserLogoutState({required this.message, required this.status});
  @override
  List<Object> get props => [status];
}