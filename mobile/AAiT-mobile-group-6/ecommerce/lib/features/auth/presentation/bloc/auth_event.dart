part of 'auth_bloc.dart';

@immutable
sealed class AuthEvent {
  const AuthEvent();
}

final class SingUpEvent extends AuthEvent {
  final SignUpEntity signUpEntity;
  const SingUpEvent({required this.signUpEntity});
}

final class LogInEvent extends AuthEvent {
  final LogInEntity logInEntity;
  const LogInEvent({required this.logInEntity});
}

final class LogOutEvent extends AuthEvent {

}