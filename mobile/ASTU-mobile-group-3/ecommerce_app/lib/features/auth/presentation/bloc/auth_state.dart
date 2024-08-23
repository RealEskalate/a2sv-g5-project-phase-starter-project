part of 'auth_bloc.dart';

// ignore: must_be_immutable
sealed class AuthState extends Equatable {
  List<Object> objects;
  AuthState([this.objects = const []]);

  @override
  List<Object> get props => [objects];
}

// ignore: must_be_immutable
final class AuthInitial extends AuthState {}

// ignore: must_be_immutable
final class AuthLoadingstate extends AuthState {}

// ignore: must_be_immutable
final class LogInSuccessState extends AuthState {}

// ignore: must_be_immutable
final class RegisterSuccessState extends AuthState {}

// ignore: must_be_immutable
final class LoginErrorState extends AuthState {
  final String message;
  LoginErrorState({required this.message}) : super([message]);
}

// ignore: must_be_immutable
final class SignupErrorState extends AuthState {
  final String message;
  SignupErrorState({required this.message}) : super([message]);
}

// ignore: must_be_immutable
final class LogoutSuccess extends AuthState {}

// ignore: must_be_immutable
final class LogoutFailedState extends AuthState {
  final String message;

  LogoutFailedState({required this.message});
}
